package tri

import (
	"time"
)

// In this file are defined the permissible types found in Slots and the Default specified must concur

// Following are local scoped variables for each type, implicitly in their zeroed state. These are the base primitive types
var vBool bool
var vString string
var vStringSlice []string
var vInt int
var vFloat float64
var vDuration time.Duration

// ValidTypes is a slice of interface{} containing a zero value variable for each of the valid types permitted in Var declarations, which are listed directly above as local scoped vars
var ValidTypes = []interface{} {
	
}

// ValidateDefault is intended to be used with the Default and Slot fields of Var elements to ensure that the former is a valid value to store in the latter
func ValidateDefault(def, slot interface{}) {
	// First check that slot type is a pointer to a valid type of value
	switch slot.(type) {
	case *int:
	default:
	}
}