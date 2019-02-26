package tri

import (
	"reflect"
	"errors"
	"fmt"
	"unicode"
)

// Valid checks to ensure the contents of this node type satisfy constraints
// Brief only contains one thing, so we make sure it has it - one string. This string may not contain any type of control characters, and is limited to 80 characters in length
func (
	r *Brief,
) Valid() error {

	R := (*r)
	if len(R) < 1 {
		return errors.New("Brief field must have (only) one item")
	}
	s, ok := R[0].(string)
	if !ok {
		return errors.New("Brief's mandatory field is not a string")
	}
	if len(s) > 80 {
		return errors.New("Brief's text may not be over 80 characters in length")
	}
	for _, x := range s {
		if unicode.IsControl(x) {
			return errors.New("Brief text may not contain any control characters")
		}
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// This validator only has to check the elements of the slice are zero or more Command items
func (
	r *Command,
) Valid() error {

	R := *r
	s, ok := R[0].(string)
	if !ok {
		return fmt.Errorf("first element of Command must be a string")
	}
	if e := ValidName(s); e != nil {
		return fmt.Errorf("error in name of Command: %v", e)
	}
	// validSet is an array of 4 elements that represent the presence of the 4 mandatory parts.
	var validSet [2]bool
	brief, handler := 0, 1
	for i, x := range R[1:] {

	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// This validator only has to check the elements of the slice are zero or more Command items
func (
	r *Commands,
) Valid() error {

	R := (*r)
	for i, x := range R {
		e := x.Valid()
		if e != nil {
			return fmt.Errorf("error in element %d of Commands list: %v", i, e)
		}
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// The only constraint on the Default subtype is that it contains at only one element, the value is checked for correct typing by the Commands validator
func (
	r *Default,
) Valid() error {

	R := (*r)
	if len(R) != 1 {
		return errors.New("the Default container must only contain one element")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// The constraint of DefaultCommand is that it has at least one element, and that the 0 element is a string. The check for the command name's presence in the Commands set is in the Tri validator
func (
	r *DefaultCommand,
) Valid() error {

	R := (*r)
	if len(R) < 1 {
		return errors.New(
			"the DefaultCommand element must contain only one element")
	}
	_, ok := R[0].(string)
	if !ok {
		return errors.New("element 0 of DefaultCommand must be a string")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// The constraints of examples are minimum 1 element and all elements are strings
func (
	r *Examples,
) Valid() error {

	R := *r
	if len(R) < 1 {
		return errors.New("Examples field may not be empty")
	}
	for i, x := range R {
		_, ok := x.(string)
		if !ok {
			return fmt.Errorf("Examples elements may only be strings, element %d is not a string", i)
		}
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// A group must contain one string, anything else is invalid
func (
	r *Group,
) Valid() error {

	R := *r
	if len(R) != 1 {
		return errors.New("Group must (only) contain one element")
	}
	_, ok := R[0].(string)
	if !ok {
		return errors.New("Group element must be a string")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// Help may only contain one string
func (
	r *Help,
) Valid() error {

	R := *r
	if len(R) != 1 {
		return errors.New("Help field must contain (only) one item")
	}
	_, ok := R[0].(string)
	if !ok {
		return errors.New("Help field element is not a string")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// RunAfter is a simple flag that indicates by existence of an empty value, so error if it has anything inside it
func (
	r *RunAfter,
) Valid() error {

	R := *r
	if len(R) > 0 {
		return errors.New(
			"RunAfter may not contain anything, empty declaration only")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// Short names contain only a single Rune variable
func (
	r *Short,
) Valid() error {

	R := *r
	if len(R) != 1 {
		return errors.New("Short name item must contain (only) one item")
	}
	_, ok := R[0].(rune)
	if !ok {
		return errors.New("Short name element must be a rune (enclose in '')")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// Slot may only contain one type of element. The type check is in the Var, here we only ensure the slots contain pointers to the same type, the parser will put the final parsed value in all of them. Multiple variables are permitted here to enable the configuration of more than one application
func (
	r *Slot,
) Valid() error {

	R := *r
	var slotTypes []reflect.Type
	for _, x := range R {
		slotTypes = append(slotTypes, reflect.TypeOf(x))
	}
	for i, x := range slotTypes {
		if i > 0 {
			if slotTypes[i] != slotTypes[i-1] {
				return fmt.Errorf("slot contains more than one type of variable, found %v at index %d", x, i)
			}
		}
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// Terminates is a flag value, and may not contain anything
func (
	r *Terminates,
) Valid() error {

	R := *r
	if len(R) > 0 {
		return errors.New("Terminates type may not contain anything")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// A Tri, the base type, in a declaration must contain a name as first element, a Brief, Version and a Commands item, and only one of each. Also, this and several other subtypes of Tri
func (
	r *Tri,
) Valid() error {
	R := *r
	if len(R) < 4 {
		return errors.New("a Tri must contain at least 4 elements: name, Brief, Version and Commands")
	}
	// validSet is an array of 4 elements that represent the presence of the 4 mandatory parts.
	var validSet [4]bool
	brief, version, commands := 1, 2, 3
	_, ok := R[0].(string)
	if !ok {
		return errors.New("first element of a Tri must be the application name")
	}
	// The mandatory elements also may not be repeated:
	for i, x := range R[1:] {
		switch y := x.(type) {
		case Brief:
			if validSet[brief] {
				return fmt.Errorf(
					"Tri contains more than one Brief, second found at index %d", i)
			}
			validSet[brief] = true
			if e := y.Valid(); e != nil {
				return fmt.Errorf("Tri field %d: %s", i, e)
			}
		case Version:
			if validSet[brief] {
				return fmt.Errorf(
					"Tri contains more than one Version, second found at index %d", i)
			}
			if e := y.Valid(); e != nil {
				return fmt.Errorf("Tri field %d: %s", i, e)
			}
			validSet[version] = true
		case Commands:
			if validSet[commands] {
				return fmt.Errorf(
					"Tri contains more than one Commands, second found at index %d", i)
			}
			validSet[commands] = true
			e := y.Valid()
			if e != nil {
				return fmt.Errorf("Tri field %d: %s", i, e)
			}
			// (mostly) Empty conditions only to filter out element types that are not valid in a Tri (default case will trigger for any type not in the set)
		case Usage:
			e := y.Valid()
			if e != nil {
				return fmt.Errorf("Tri field %d: %s", i, e)
			}
		case DefaultCommand:
			e := y.Valid()
			if e != nil {
				return fmt.Errorf("Tri field %d: %s", i, e)
			}
		default:
			return fmt.Errorf(
				"Tri contains an element type it may not contain, at index %d", i)
		}
	}
	switch {
	case !validSet[brief]:
		return errors.New("Tri is missing its Brief field")

	case !validSet[version]:
		return errors.New("Tri is missing its Version field")

	case !validSet[commands]:
		return errors.New("Tri is missing its Commands field")

	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// Trigger must contain (one) name, Brief and Handler, and nothing other than these and Short, Usage, Help, Default, Terminates, RunAfter
func (
	r *Trigger,
) Valid() error {

	R := *r
	if len(R) < 3 {
		return errors.New(
			"Trigger must contain a name, Brief and Handler at minimum")
	}
	name, ok := R[0].(string)
	if !ok {
		return errors.New("first element of Trigger must be the name")
	} else if e := ValidName(name); e != nil {
		return fmt.Errorf("Invalid Name in Trigger at index 0: %v", e)
	}
	// validSet is an array of 4 elements that represent the presence of the 4 mandatory parts.
	var validSet [2]bool
	brief, handler := 0, 1
	// check for presence of all mandatory and non-presence of impermissible element types
	for i, x := range R[1:] {

		switch y := x.(type) {

		case Brief:
			if e := y.Valid(); e != nil {
				return fmt.Errorf(
					"Trigger contains invalid element at %d :%s", i, e)
			}
			if validSet[brief] {
				return fmt.Errorf("Trigger may must (only) contain one Brief, second found at index %d", i)
			} else {
				validSet[brief] = true
			}

		case Handler:
			if y == nil {
				return fmt.Errorf("Handler at index %d may not be nil", i)
			}
			if validSet[handler] {
				return fmt.Errorf(
					"Trigger may must (only) contain one Brief, second found at index %d", i)
			} else {
				validSet[handler] = true
			}

		case Short:
			if e := y.Valid(); e != nil {
				return fmt.Errorf(
					"Trigger contains invalid element at %d :%s", i, e)
			}

		case Usage:
			if e := y.Valid(); e != nil {
				return fmt.Errorf(
					"Trigger contains invalid element at %d :%s", i, e)
			}

		case Help:
			if e := y.Valid(); e != nil {
				return fmt.Errorf(
					"Trigger contains invalid element at %d :%s", i, e)
			}

		case Default:
			if e := y.Valid(); e != nil {
				return fmt.Errorf(
					"Trigger contains invalid element at %d :%s", i, e)
			}

		case Terminates:
			if e := y.Valid(); e != nil {
				return fmt.Errorf(
					"Trigger contains invalid element at %d :%s", i, e)
			}

		case RunAfter:
			if e := y.Valid(); e != nil {
				return fmt.Errorf(
					"Trigger contains invalid element at %d :%s", i, e)
			}

		default:
			return fmt.Errorf(
				"found invalid item type at element %d in a Trigger", i)
		}
	}
	if !(validSet[brief] && validSet[handler]) {
		return errors.New("Trigger must contain one each of Brief and Handler")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// Usage fields contain only one string of no more than 80 characters and no control characters
func (
	r *Usage,
) Valid() error {
	R := *r
	if len(R) > 1 {
		return errors.New("Usage field must contain (only) one string")
	}
	s, ok := R[0].(string)
	if !ok {
		return errors.New("Usage field element is not a string")
	}
	if ll := len(s); ll > 80 {
		return fmt.Errorf("Usage string is %d chars long, may not be longer than 80", ll)
	}
	for i, x := range s {
		if unicode.IsControl(x) {
			return fmt.Errorf(
				"Usage field string may not contain control characters, one found at index %d", i)
		}
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// Var must contain name, Brief and Slot, and optionally, Short, Usage, Help and Default. The type in the Slot and the Default must be the same.
func (
	r *Var,
) Valid() error {

	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
// A version item contains three integers and an optional (less than 16 character) string, and the numbers may not be more than 99
func (
	r *Version,
) Valid() error {

	R := *r
	if len(R) > 4 {
		return errors.New("Version field may not contain more than 4 fields")
	}
	if len(R) < 3 {
		return errors.New("Version field must contain at least 3 fields")
	}
	for i, x := range R[:4] {
		n, ok := x.(int)
		if !ok {
			return fmt.Errorf("Version field %d is not an integer: %d", i, n)
		}
		if n > 99 {
			return fmt.Errorf("Version field %d value is over 99: %d", i, n)
		}
	}
	if len(R) == 4 {
		s, ok := R[3].(string)
		if !ok {
			return fmt.Errorf("optional field 4 of Version is not a string")
		}
		for i, x := range s {
			if !(unicode.IsLetter(x) || unicode.IsDigit(x)) {
				return fmt.Errorf(
					"optional field 4 of Version contains other than letters and numbers at position %d: '%v,", i, x)
			}
		}
	}
	return nil
}

// ValidName checks that a Tri name element that should be a name only contains letters
func ValidName(s string) error {

	for i, x := range s {
		if !unicode.IsLetter(x) {
			return fmt.Errorf(
				"element %d, '%v' of name is not a letter", i, x)
		}
	}
	return nil
}
