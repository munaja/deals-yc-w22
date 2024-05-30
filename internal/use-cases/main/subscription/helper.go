package subscription

import "strconv"

// generating redis key is managed here in care there is some need in
// hash, encryoting, or something else
func generateRedisKey(User_Id int, theType string) string {
	return "subscription_" + strconv.Itoa(User_Id) + "_" + string(theType)
}
