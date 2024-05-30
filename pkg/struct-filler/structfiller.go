package sego

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"

	es "github.com/munaja/exam-deals-yc-w22/pkg/error-structure"
)

// Fill struct with form-data content
func HttpFormData(container any, r *http.Request) error {
	// identiy value and loop if its pointer until reaches non pointer
	cv := reflect.ValueOf(container)

	// loop until we get what kind lays behind the input
	for cv.Kind() == reflect.Pointer || cv.Kind() == reflect.Interface {
		cv = cv.Elem()
	}

	// non struct cant be filled
	if cv.Kind() != reflect.Struct {
		panic("input requires struct type")
	}

	// check each field
	errList := es.XErrors{}
	ct := cv.Type()
	for i := 0; i < cv.NumField(); i++ {
		// identify field type and value of the field
		ft := ct.Field(i)
		fv := cv.Field(i)
		if ft.Anonymous {
			if err := setEmbedValueFormData(container, r, []string{ft.Name}); err != nil {
				return nil
			}
			continue
		}

		for fv.Kind() == reflect.Ptr {
			fv = fv.Elem()
		}
		if fv.Kind() == reflect.Struct || !fv.CanSet() {
			continue
		}

		key := keyOrJsonTag(ft.Name, ft.Tag.Get("json"))
		rv := r.PostFormValue(key)
		if rv == "" {
			// try once more if fail, mostly not called tho
			r.ParseForm()
			rv = r.FormValue(key)
		}

		fvKind := fv.Kind()
		ftName := ft.Name
		err := reflectValueFiller(fv, fvKind, ftName, rv)
		if err != nil {
			errList[key] = err.(es.XError)
		}
	}

	if len(errList) > 0 {
		return errList
	}
	return nil
}

// Fill struct with url encoded content
func UrlQueryParam(container any, url url.URL) error {
	// identiy value and loop if its pointer until reaches non pointer
	cv := reflect.ValueOf(container)

	// loop until we get what kind lays behind the input
	for cv.Kind() == reflect.Pointer || cv.Kind() == reflect.Interface {
		cv = cv.Elem()
	}

	// non struct cant be filled
	if cv.Kind() != reflect.Struct {
		panic("input requires struct type")
	}

	errList := es.XErrors{}
	ct := cv.Type()
	values := url.Query()
	for i := 0; i < cv.NumField(); i++ {
		// identify field type and value of the field
		ft := ct.Field(i)
		fv := cv.Field(i)
		if ft.Anonymous {
			if err := setEmbedValueUrlQuery(container, url, []string{ft.Name}); err != nil {
				return nil
			}
			continue
		}

		for fv.Kind() == reflect.Ptr {
			fv = fv.Elem()
		}
		if !fv.CanSet() {
			continue
		}

		key := keyOrJsonTag(ft.Name, ft.Tag.Get("json"))
		vals, ok := values[key]
		if !ok {
			continue
		}

		fvKind := fv.Kind()
		ftName := ft.Name
		err := reflectValueFiller(fv, fvKind, ftName, vals[0])
		if err != nil {
			errList[key] = err.(es.XError)
		}
	}

	if len(errList) > 0 {
		return errList
	}
	return nil
}

// Fill struct json content from io.reader
func IOReaderJson(container any, input io.Reader) error {
	decoder := json.NewDecoder(input)
	err := decoder.Decode(&container)
	if err != nil {
		cv := reflect.ValueOf(container)
		for cv.Kind() == reflect.Pointer || cv.Kind() == reflect.Interface {
			cv = cv.Elem()
		}
		structName := cv.Type().Name()
		return es.XError{
			Code:        "parse-fail",
			Message:     "failed to parse input, error: " + err.Error(),
			ExpectedVal: "value of " + structName,
		}
	}

	return nil
}

// embeded helper for form data
func setEmbedValueFormData(container any, r *http.Request, keyPath []string) error {
	cv := reflect.ValueOf(container)
	for cv.Kind() == reflect.Pointer || cv.Kind() == reflect.Interface {
		cv = cv.Elem()
	}

	for idx := range keyPath {
		cv = cv.FieldByName(keyPath[idx])
		for cv.Kind() == reflect.Pointer || cv.Kind() == reflect.Interface {
			cv = cv.Elem()
		}
	}

	// check each field
	errList := es.XErrors{}
	ct := cv.Type()
	for i := 0; i < cv.NumField(); i++ {
		// identify field type and value of the field
		ft := ct.Field(i)
		fv := cv.Field(i)
		if ft.Anonymous {
			setEmbedValueFormData(container, r, append(keyPath, ft.Name))
			continue
		}

		for fv.Kind() == reflect.Ptr {
			fv = fv.Elem()
		}
		if !fv.CanSet() {
			continue
		}

		key := keyOrJsonTag(ft.Name, ft.Tag.Get("json"))
		rv := r.PostFormValue(key)
		if rv == "" {
			// try once more if fail, mostly not called tho
			r.ParseForm()
			rv = r.FormValue(key)
		}

		fvKind := fv.Kind()
		ftName := ft.Name
		err := reflectValueFiller(fv, fvKind, ftName, rv)
		if err != nil {
			errList[key] = err.(es.XError)
		}
	}

	if len(errList) > 0 {
		return errList
	}

	return nil
}

// embeded helper for form data
func setEmbedValueUrlQuery(container any, url url.URL, keyPath []string) error {
	cv := reflect.ValueOf(container)
	for cv.Kind() == reflect.Pointer || cv.Kind() == reflect.Interface {
		cv = cv.Elem()
	}

	for idx := range keyPath {
		cv = cv.FieldByName(keyPath[idx])
		for cv.Kind() == reflect.Pointer || cv.Kind() == reflect.Interface {
			cv = cv.Elem()
		}
	}

	// check each field
	errList := es.XErrors{}
	ct := cv.Type()
	values := url.Query()
	for i := 0; i < cv.NumField(); i++ {
		// identify field type and value of the field
		ft := ct.Field(i)
		fv := cv.Field(i)
		if ft.Anonymous {
			setEmbedValueUrlQuery(container, url, append(keyPath, ft.Name))
			continue
		}

		for fv.Kind() == reflect.Ptr {
			fv = fv.Elem()
		}
		if !fv.CanSet() {
			continue
		}

		key := keyOrJsonTag(ft.Name, ft.Tag.Get("json"))
		vals, ok := values[key]
		if !ok {
			continue
		}

		fvKind := fv.Kind()
		ftName := ft.Name
		err := reflectValueFiller(fv, fvKind, ftName, vals[0])
		if err != nil {
			errList[key] = err.(es.XError)
		}
	}

	if len(errList) > 0 {
		return errList
	}

	return nil
}
