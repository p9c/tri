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

	Note that all types use the Tri []interface{} type because, while it is possible to omit fields by using field tags, if the object is another composite object, its type must be specified unless it is an array and the members are implicitly typed. By using interface slice instead, the declaration syntax has minimum redundant type specifications.
*/
package tri

import (
	"fmt"
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
type Version Tri

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
				func(Tri) int {
					return 0
				},
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
			func(Tri) int {
				return 0
			},
		},
	},
}

func Selftest() {
	brief := exampleTri[1].(Brief)
	fmt.Println(brief[0])
}
