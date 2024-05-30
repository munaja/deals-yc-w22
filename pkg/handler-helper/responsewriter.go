package handlerhelper

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/google/uuid"

	sv "github.com/munaja/exam-deals-yc-w22/pkg/struct-validator"

	lo "github.com/munaja/exam-deals-yc-w22/pkg/api-core/lang-own"
	ds "github.com/munaja/exam-deals-yc-w22/pkg/data-structure"
	es "github.com/munaja/exam-deals-yc-w22/pkg/error-structure"
	l "github.com/munaja/exam-deals-yc-w22/pkg/language"
)

func UseLang(src *l.LangData) {
	lo.I = *src
}

// The primary function that writes json output through http.ResponseWriter
func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) {
	js, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte("{ \"message\": \"error converting data or result to json\"}"))
		w.WriteHeader(500)
		return
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

// Writes an error response, will try to match list of recognized error to decide
// the http status code
func WriteError(w http.ResponseWriter, err es.XError) {
	httpCode := http.StatusUnprocessableEntity
	for idx, code := range ErrorCodes {
		if err.Code == idx {
			httpCode = code
			break
		}
	}

	WriteJSON(w, httpCode, err, nil)
}

// Generate response based on the condition of data and error
// Note that it should be called for things related to data processing and
// with the possible error only for unprocessable entity that related to
// data field which can have more than one error. Any non-field error, which
// normally a single error will be disposed to WriteError funcion
func DataResponse(w http.ResponseWriter, data, err any) {
	v := reflect.ValueOf(data)
	vKind := v.Kind()
	if (data == nil || (vKind == reflect.Pointer && v.IsNil())) && err == nil {
		WriteJSON(w, http.StatusNotFound, es.XError{
			Code:    "data-notFound",
			Message: lo.I.Msg("data-notFound"),
		}, nil)
	} else if err != nil {
		if stringErr, ok := err.(string); ok {
			WriteJSON(w, http.StatusUnprocessableEntity, ds.IS{"Message": stringErr}, nil)
		} else {
			if castedErr, ok := err.(es.XError); ok {
				// this is the only error that can have non unprocessable entity error
				WriteError(w, castedErr)
			} else if castedErr, ok := err.(es.XErrors); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, ds.II{
					"meta":   ds.IS{"count": strconv.Itoa(len(castedErr))},
					"errors": err,
				}, nil)
			} else if castedErr, ok := err.(map[string]any); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, ds.II{
					"meta":   ds.IS{"count": strconv.Itoa(len(castedErr))},
					"errors": castedErr}, nil)
			} else if castedErr, ok := err.(error); ok {
				WriteJSON(w, http.StatusUnprocessableEntity, es.XError{
					Code:    "unknown",
					Message: castedErr.Error(),
				}, nil)
			} else {
				// worst case unknown error
				WriteJSON(w, http.StatusUnprocessableEntity, ds.II{"errors": err}, nil)
			}
		}
	} else {
		if dataVal, ok := data.(ds.Data); ok {
			WriteJSON(w, http.StatusOK, dataVal, nil)
		} else if message, ok := data.(string); ok {
			WriteJSON(w, http.StatusOK, ds.IS{"message": message}, nil)
		} else {
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			vKind = v.Kind()
			if vKind != reflect.Struct && vKind != reflect.Map {
				WriteJSON(w, http.StatusOK, ds.II{"value": data}, nil)
			} else {
				WriteJSON(w, http.StatusOK, data, nil)
			}
		}
	}
	// if (vKind != reflect.Pointer && data != nil) || (vKind == reflect.Pointer && !v.IsNil())) && err == nil {
	// } else if err != nil {
	// } else {
	// }
}

// Validates a string assuming the field is required.
// Error occured will be categorized as data-field
func ValidateString(w http.ResponseWriter, fieldName, input string) string {
	if !requiredString(w, fieldName, input) {
		return ""
	}
	return input
}

// Validates an int value from string, assuming the field is required.
// Error occured will be categorized as data-field
func ValidateInt(w http.ResponseWriter, fieldName, input string) int {
	// val := chi.URLParam(r, input)
	if !requiredString(w, fieldName, input) {
		return 0
	}
	output, err := strconv.Atoi(input)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ds.II{"errors": es.XErrors{
			fieldName: es.XError{
				Code:    "val-int",
				Message: lo.I.Msg("val-int"),
			},
		}}, nil)
		return 0
	}
	return output
}

// Validates a UUID from string assuming the field is required
// Error occured will be categorized as data-field
func ValidateIdUuid(w http.ResponseWriter, fieldName, input string) uuid.UUID {
	if !requiredString(w, fieldName, input) {
		return uuid.Nil
	}
	output, err := uuid.Parse(input)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ds.II{"errors": es.XErrors{
			fieldName: es.XError{
				Code:    "val-uuid",
				Message: lo.I.Msg("val-uuid"),
			},
		}}, nil)
		return uuid.Nil
	}
	return output
}

// Validates struct
func ValidateStruct(w http.ResponseWriter, data any) bool {
	err := sv.Validate(data)
	if err != nil {
		DataResponse(w, nil, err)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return bool true on success
func ValidateStructByIOR(w http.ResponseWriter, body io.Reader, data any) bool {
	err := sv.ValidateIoReader(&data, body)
	if err != nil {
		DataResponse(w, nil, err)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return bool true on success
func ValidateStructByURL(w http.ResponseWriter, url url.URL, data any) bool {
	err := sv.ValidateURL(&data, url)
	if err != nil {
		DataResponse(w, nil, err)
		return false
	}

	return true
}

// by form-data version of ValidateStruct, to cover form-data, return bool true on success
func ValidateStructByFD(w http.ResponseWriter, r *http.Request, data any) bool {
	err := sv.ValidateFormData(&data, r)
	if err != nil {
		DataResponse(w, nil, err)
		return false
	}

	return true
}
