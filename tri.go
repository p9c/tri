// Package Tri is an implementation of a simple abstract container type that uses type aliasing to concisely designate different types of elements that can be captured using type switches.
/*
	The primary construct used in Go to implement generic types is the interface, and to implement any kind of structure that involves lists, the slice-of-interface []interface{} is a structure that can hold zero or more items of any type whatsoever.

	By aliasing this primary primitive using names, the type aliases can become metadata that can be used to specify how its contents are to be interpreted.

	In this implementation you can see I take advantage of the possibility to put any type into the list to place string labels at the heads of the lists as identifiers, which then can form a tagged tree structure that can be addressed by providing a list of the identifier strings at each node.

	Go's strict static typing means that such hierarchies cannot be written without a complete set of types already pre-defined, so in each case the implementation has to be written specifically for the types of data used.

	In spite of it being perhaps less logical, all elements of a Tri are a derivative of a Tri. This way one never has to type assert the container, only the contents, and in most cases, the possible types depend on the type of the parent of a branch.

	Most generics systems used in other languages, and default parameters have the ability to set constraints on the contents of the node. Instead, Tri has an interface called Branch, which has a single method Valid(). This function returns true if the mandatory items are present, and all other items have values within the acceptable range required by the implementation.

	Validators for each type in this package automatically trigger the validators for their (valid set of) constituent elements, so only one Valid() invocation is needed, on the Tri type. If Valid functions return an error, it passes back through the cascade to the root where it is printed to tell the (proogrammer) their declaration was erroneous.

	Without validity checks, in a strict static typing language, dynamic variables can cause errors when assuming the input is correct, and such errors can potentially escape notice for a long time, so by adding these runtime bounds and set membership tests, such errors are caught before the application they configure even starts execution of the main() - the Tri type's Valid() method should be invoked in a init() so it runs after the calls in any var blocks and before the main() tries to parse the CLI flags.
*/
package main

import (
	"errors"
	"fmt"
	"unicode"
)

// Branch is an interface that all Tri nodes implement to do runtime checks on the contents of a Tri derived type, basically something like Assert for the constraints required of the subtype.
/*
	Since in Go every bit of unknown-at-runtime data type must be specifically implemented, it is advised that when implementing this type one not flag extraneous elements as erroneous, but only that some set of items are specified, and for all others that are valid, that the values are inside the necessary bounds, or match in their type or abstract structure for composite types.

	The validation method is used at runtime and basically is a declaration syntax check, and only the root Valid() function must be called, it does the rest and returns as soon as it finds an error when walking the tree.

	This validation has no benefit if the Tri declaration is valid, but if it is not checked, the parse/configure process will almost certainly panic when it finds the wrong type of object in the wrong position, so it may seem overly verbose but this ensures that no further checking is required before parsing the several inputs that result in a filled out configuration struct.

	Constraints are not the most visible feature of generic types, but if constraints aren't applied to generic types they can be very hard bugs to spot, and simply do not occur if you first check the parts of the struct fit the mandatory minimum contents specification, which in Go means a function with validation.

	All types must check their members contain only valid element types in the set defined for each, and when possible, trigger the validators on their constituent elements. When a child validator finds an error, it is passed back through the chain of parents back to the root so the exact position in the specification (relative to the root and identified by the names) where the incorrect item (or lack of mandatory item) was found. Parsing terminates at this point - it is assumed that the programmer will not frequently misconstruct a Tri and in any case collating errors and defining a valid point of exit thereafter is arbitrary, and in the case of multiple mandatory elements, the list of valid elements is provided in the error text.
*/
type Branch interface {
	Valid() error
}

type Brief Tri
type Command Tri
type Commands []Command
type Default Tri
type DefaultCommand Tri
type Examples Tri
type Group Tri
type Handler func(Tri) int
type Help Tri
type RunAfter Tri
type Short Tri
type Slot Tri
type Terminates Tri
type Tri []interface{}
type Trigger Tri
type Usage Tri
type Var Tri
type Version struct {
	Major int
	Minor int
	Rev   int
	Build string
}

var exampleTri = Tri{
	"appname",
	Brief{"brief"},
	Usage{"usage"},
	Version{0, 1, 1, "alpha"},
	DefaultCommand{"help"},
	Commands{
		{"ctl",
			Short{"c"},
			Brief{"brief"},
			Usage{"usage"},
			Help{"help"},
			Examples{
				"example 1", "explaining text",
				"example 2", "explaining text",
			},
			Var{"datadir",
				Short{"d"},
				Brief{"brief"},
				Usage{"usage"},
				Help{"help"},
				Default{"~/.pod"},
				Slot{""},
			},
			Trigger{"init",
				Short{"I"},
				Brief{"brief"},
				Usage{"usage"},
				Help{"help"},
				Default{},
				Terminates{},
				RunAfter{},
			},
			func(Tri) int {
				return 0
			},
		},
		{"node",
			Short{"n"},
			Brief{"brief"},
			Usage{"usage"},
			Help{"help"},
			Examples{"example1", "example2"},
			func(Tri) int { return 0 },
		},
	},
}

// Valid checks to ensure the contents of this node type satisfy constraints
func (
	r *Brief,
) Valid() error {

	// Brief only contains one thing, so we make sure it has it - one string. This string may not contain any type of control characters, and is limited to 80 characters in length
	R := (*r)
	if len(R) < 1 {
		return errors.New("Brief field must have at least one item")
	}
	s, ok := R[0].(string)
	if !ok {
		return errors.New("Brief's mandatory first field is not a string")
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
func (
	r *Commands,
) Valid() error {

	// This validator only has to check the elements of the slice contain the mandatory elements: first element is the name string, and that there is a Brief, so there is something to print out in the top level help.
	R := (*r)
	for i, x := range R {
		if len(x) < 3 {
			return fmt.Errorf(
				"Commands item %d has less than minimum elements, Brief and Handler are required", i)
		}
		y, ok := x[0].(string)
		if !ok {
			return fmt.Errorf("first element of Commands element %d is not a string", i)
		}
		if e := ValidName(y); e != nil {
			return fmt.Errorf(
				"Commands element %d's first element is not a string, %s: %s",
				i, y, e.Error(),
			)
		}
		// Next, check that all present members are in the valid set for this container
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
func (
	r *Default,
) Valid() error {

	// The only constraint on the Default subtype is that it contains at only one element, the value is checked for correct typing by the Commands validator
	R := (*r)
	if len(R) != 1 {
		return errors.New("the Default container must only contain one element")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
func (
	r *DefaultCommand,
) Valid() error {

	// The constraint of DefaultCommand is that it has at least one element, and that the 0 element is a string. The check for the command name's presence in the Commands set is in the Tri validator
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
func (
	r *Examples,
) Valid() error {

	// The constraints of examples are minimum 1 element and all elements are strings
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
func (
	r *Group,
) Valid() error {

	// A group must contain one string, anything else is invalid
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
func (
	r *Help,
) Valid() error {

	// Help may only contain one string
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
func (
	r *RunAfter,
) Valid() error {

	// RunAfter is a simple flag that indicates by existence of an empty value, so error if it has anything inside it
	R := *r
	if len(R) > 0 {
		return errors.New(
			"RunAfter may not contain anything, empty declaration only")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
func (
	r *Short,
) Valid() error {

	// Short names contain only a single Rune variable
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
func (
	r *Slot,
) Valid() error {

	// Slot may only contain one element. The type check is in the Var validator
	R := *r
	if len(R) != 1 {
		return errors.New("Slot type must contain (only) one element")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
func (
	r *Terminates,
) Valid() error {

	// Terminates is a flag value, and may not contain anything
	R := *r
	if len(R) > 0 {
		return errors.New("Terminates type may not contain anything")
	}
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
func (
	r *Tri,
) Valid() error {

	// A Tri, the base type, in a declaration must contain a name as first element, a Brief, Version and a Commands item, and only one of each. Also, this and several other subtypes of Tri
	R := *r
	if len(R) < 4 {
		return errors.New("a Tri must contain 4 elements: name, Brief, Version and Commands")
	}
	// validSet is an array of 4 elements that represent the presence of the 4 mandatory parts.
	var validSet [4]bool
	brief, version, commands := 1, 2, 3
	_, ok := R[0].(string)
	if !ok {
		return errors.New("first element of a Tri must be the application name")
	}
	for i, x := range R {
		switch y := x.(type) {
		case Brief:
			if validSet[brief] {
				return fmt.Errorf(
					"Tri contains more than one Brief, second found at index %d", i)
			}
			validSet[brief] = true
			if e := y.Valid(); e != nil {
				return e
			}
		case Version:
			if validSet[brief] {
				return fmt.Errorf(
					"Tri contains more than one Version, second found at index %d", i)
			}
			if e := y.Valid(); e != nil {
				return e
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
				return e
			}
			// (mostly) Empty conditions only to filter out element types that are not valid in a Tri (default case will trigger for any type not in the set)
		case Usage:
			e := y.Valid()
			if e != nil {
				return e
			}
		case DefaultCommand:
			e := y.Valid()
			if e != nil {
				return e
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
func (
	r *Trigger,
) Valid() error {

	// Trigger must contain (one) name, Brief and Handler, and nothing other than these and Short, Usage, Help, Default, Terminates, RunAfter
	R := *r
	if len(R) < 3 {
		return errors.New(
			"Trigger must contain a name, Brief and Handler at minimum")
	}
	name, ok := R[0].(string)
	if !ok {
		return errors.New("first element of Trigger must be the name")
	} else if e := ValidName(name); e != nil {
		return e
	}
	// check for presence of all mandatory and non-presence of impermissible element types
	for i, x := range R[1:] {
		switch y := x.(type) {
		case Brief:
			if e := y.Valid(); e != nil {
				return fmt.Errorf(
					"Trigger contains invalid element at %d :%s", i, e)
			}
		case Handler:
			if y == nil {
				return fmt.Errorf("Handler at index %d may not be nil", i)
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
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
func (
	r *Usage,
) Valid() error {

	//
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
func (
	r *Var,
) Valid() error {

	//
	return nil
}

// Valid checks to ensure the contents of this node type satisfy constraints
func (
	r *Version,
) Valid() error {

	//
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

func main() {
	brief := exampleTri[1].(Brief)
	fmt.Println(brief[0])
}
