// Package Tri is a library for defining the configuration parameters for a CLI application, managing the configuration files, application data directory, defining subcommands, parameters, default values and automatically loading configuration structures with the processed configuration.
//
// The primary construct used in Go to implement generic types is the interface, and to implement any kind of structure that involves lists, the slice-of-interface []interface{} is a structure that can hold zero or more items of any type whatsoever.
//
// By deriving from this primary primitive using names, the type aliases can become metadata that can be used to specify how its contents are to be interpreted.
//
// Implementation Notes
//
// In this implementation you can see I take advantage of the possibility to put any type into the list to place string labels at the heads of the lists as identifiers, which then can form a tagged tree structure that can be addressed by providing a list of the identifier strings at each node.
//
// Go's strict static typing means that such hierarchies cannot be written without a complete set of types already pre-defined, so in each case the implementation has to be written specifically for the types of data used.
//
// In spite of it being perhaps less logical, all elements of a Tri are a derivative of a Tri. This way one never has to type assert the container, only the contents, and in most cases, the possible types depend on the type of a branch.
//
// Declaration Syntax Validation
//
// Validators for each type in this package automatically trigger the validators for their (valid set of) constituent elements, so only one Valid() invocation is needed, on the Tri type. If Valid functions return an error, it passes back through the cascade to the root where it is printed to tell (the proogrammer) that their declaration was erroneous.
//
// Without validity checks, in a strict static typing language, dynamic variables can cause errors when assuming the input is correct, and such errors can potentially escape notice for a long time, so by adding these runtime bounds and set membership tests, such errors are caught before the application they configure even starts execution of the main() - the Tri type's Valid() method should be invoked in a init() so it runs after the calls in any var blocks and before the main() tries to parse the CLI flags.
//
// NOTE: all types use the Tri []interface{} type because, while it is possible to omit fields by using field tags, if the object is another composite object, its type must be specified unless it is an array and the members are implicitly typed. By using interface slice instead, the declaration syntax has minimum redundant type specifications.
package tri

import (
	"fmt"
)

// Branch is an interface that all Tri nodes implement to do runtime checks on the contents of a Tri derived type, basically something like Assert for the constraints required of the subtype.
//
// The validation method is used at runtime and basically is a declaration syntax check, and only the root Valid() function must be called, it does the rest and returns as soon as it finds an error when walking the tree, or proceeds to next step of reading CLI args and compositing defaults, config file, env together filling the configuration structure.
//
// This validation has no benefit if the Tri declaration is valid, but if it is not checked, the parse/configure process will almost certainly panic when it finds the wrong type of object in the wrong position, so it may seem overly verbose but this ensures that no further checking is required before parsing the several inputs that result in a filled out configuration struct.
//
// Constraints are not the most visible feature of generic types, but if constraints aren't applied to generic types they can be very hard bugs to spot, and simply do not occur if you first check the parts of the structures are correct.
type Branch interface {
	Validate() error
}

// TODO: write the english version of what structure each of these has

// Brief is a short description up to 80 characters long containing one string with no control characters, that is intended to describe the item it is embedded in.
type Brief Tri

// Command is the specification for an individual subcommand. Shown below is the full set of allowable items, the metadata items may only appear once, there must be a Brief, the name at the start, and a Handler function.
/*
	{"name",
		Short{"c"}, // single character shortcut for full length name
		Brief{"brief"},
		Usage{"usage"},
		Help{"help"},
		Examples{
			"example 1", "explaining text",
			...
		},
		Var{...
		},
		Trigger{...
		},
		func(Tri) int {
			...
			return 0
		},
	}
*/
type Command Tri

// Commands is just an array of Command, providing a symbol-free and human-friendly name for the array of commands in an application declaration.
type Commands []Command

// Default is specifies the default value for a Variable, it must contain only one variable inside its first element.
type Default Tri

// DefaultCommand specifies the Command that should run when no subcommand is specified on the commandline.
type DefaultCommand Tri

// DefaultOn specifies that the trigger it is inside is disabled by its name appearing in the invocation.
type DefaultOn Tri

// Examples is is a list of pairs of strings containing a snippet of an example invocation and a short description of the effect of this example.
type Examples Tri

// Group is a single string tag with the same format as name fields that functions as a tag to gather related items in the help output.
type Group Tri

// Handler is the function signature of a subcommand, and is modeled after the parameterisation of a CLI command. The return value is passed through by the real main function of che handler back to the commandline as a return value. Zero indicates success, nonzero is error, with the possibility of arbitrary attribution of meaning to the number, and is returned to the shell application that launches the command.
type Handler func(Tri) int

// Help is a free-form text that is interpreted as markdown syntax and may optionally be formatted using ANSI codes by a preprocessor to represent the structured text that a markdown parser will produce, by default all markdown annotations will be removed.
type Help Tri

// RunAfter is a flag indicating that a Trigger element of a Command should be run during shutdown instead of before startup.
type RunAfter Tri

// Short is a single character symbol that can be used instead of the name at the top of the Tri-derived type in invocation.
type Short Tri

// Slot can contain pointers to one or more items of the same type and is intended to allow the parser to directly populate the value in a possibly external struct.
type Slot Tri

// Terminates is a flag for Trigger types that indicates that the function will terminate execution of the application once it completes its work.
type Terminates Tri

// Tri is the root type where the base of an application parameter definition starts.
/*
	var exampleTri = Tri{
		"appname", // only letters in Tri tags
		Brief{"up to 80 char string, no control characters, not nil"},
		Usage{"up to 80 char string, no control characters, not nil"},
		Version{0, 1, 1, "alpha"},
		DefaultCommand{"help"},
		Var{"datadir",
			Short{"d"},
			Brief{"brief"},
			Usage{"usage"},
			Help{"help"},
			Default{"~/.pod"},
			Slot{""},
		},
		Trigger{"trigger",
		...
		},
		Commands{},
	}
*/
// Note that this base specification may have variables and triggers associated with it that can be used to set configuration values common to all (or most) of the Commands specified in the declaration.
type Tri []interface{}

// Trigger is for initiating the execution of one-shot functions that often terminate execution, or rewrite, sort, re-index, and this kind of thing.
type Trigger Tri

// Usage is is an example showing the invocation of a Tri CLI flag.
type Usage Tri

// Var is defines a configuration variable and the means to populate this variable in an optionally separate configuration structure.
type Var Tri

// Version is a short specification implementing a semver version string.
type Version Tri

var exampleTri = Tri{
	"appname",
	Brief{"brief"},
	Version{0, 1, 1, "alpha"},
	DefaultCommand{"help"},
	Var{"datadir",
		Short{"d"},
		Brief{"brief"},
		Usage{"usage"},
		Help{"help"},
		Default{"~/.pod"},
		Slot{""},
	},
	Commands{
		{"ctl",
			Short{"c"},
			Brief{"brief"},
			Usage{"usage"},
			Help{"help"},
			Group{"groupname"},
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
				DefaultOn{},
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

// selftest is just a self test to make golint not tell me about unused things
func selftest() {
	brief := exampleTri[1].(Brief)
	fmt.Println(brief[0])
}
