package tri

import (
	"time"
)

// LoadAllDefaults walks a Tri and calls LoadDefaults on each one for the first step in composition of configuration
func LoadAllDefaults(t *Tri) {
	T := *t
	var counter, varsfound int
	for _, x := range T {
		if v, ok := x.(Var); ok {
			varsfound++
			if LoadDefaults(&v) {
				counter++
			}
		}
		if v, ok := x.(Commands); ok {
			for _, x := range v {
				for _, y := range x {
					if v, ok := y.(Var); ok {
						varsfound++
						if LoadDefaults(&v) {
							counter++
						}
					}
				}
			}
		}
	}
}

// LoadDefaults reads the Default (if any) in a Var, and copies the value into the Slot, returns true if there was a Default and it was filled
func LoadDefaults(v *Var) (found bool) {
	// First find if there is a default
	var def Default
	V := *v
	for _, x := range V {
		if j, ok := x.(Default); ok {
			def = j
			found = true
		}
	}
	if !found {
		return false
	}
	found = false
	var slot Slot
	for _, x := range V {
		if j, ok := x.(Slot); ok {
			slot = j
			found = true
		}
	}
	if !found {
		return false
	}
	for _, x := range slot {
		switch S := x.(type) {
		case *string:
			s := def[0].(string)
			*S = s
		case *int:
			s := def[0].(int)
			*S = s
		case *uint32:
			s := def[0].(uint32)
			*S = s
		case *float64:
			s := def[0].(float64)
			*S = s
		case *[]string:
			s := def[0].([]string)
			*S = s
		case *time.Duration:
			s := def[0].(time.Duration)
			*S = s
		default:
			panic("unrecognised type found in slot")
		}
	}
	return true
}
