# Taxonomy of a Tri

## `name`

List elements of a Tri, being Tri, Command, Var and Trigger all require an initial string name field that may contain only letter characters. Case should be lower but the parser normalises the strings before using the definition.

## `Brief`

Brief fields are a single maximum 80 character string intended to contain a brief description of a list's contents and purpose.

## `Usage`

Usage fields are a representation of the flag as it should be typed in a command line invocation. If they are absent, they will be constructed using the Default value as an example.

## `Help`

Help fields are a plain text field that should approximately follow the conventions used by Godoc - single lines starting with a capital letter of no more than 10 words and no final full-stop are headings, paragraphs are separated by a double carriage return. The text will be rendered to fit within an 80 column terminal matrix.

## `Version`

The version element contains 3 integer numbers no larger than 99, followed by an optional letter-only string representing a build label, usually 'alpha' or 'beta' or similar.

## `DefaultCommand`

DefaultCommand specifies a subcommand to run, it must be found in the Commands list. If it is omitted, an invocation without a Command will print the top level help text.

## `Examples`

Examples are pairs of strings in a list that are shown at the bottom of the help text for the Tri, Command, Variable or Trigger. If the first field is an empty string ("") the help printer generates an example based on the Default field (see below)

## `Group`

Group is a single tag string with the same constraints as a `name` field

## `Short`

Short is a single character (case sensitive) that can be substituted for the `name` field in invocations for convenience.

## `Slot`

Slot is intended to store a pointer to another variable which usually will be a configuration field of an external configuration variable, and will have the final value parsed out of the configuration composition loaded into it using dereferencing.

## `Handler`

Handler is a function that accepts a Tri as its sole parameter and returns an int with the same semantics as a shell command return, 0 means no error, nonzero means error, any specific numbers can be used to indicate a specific error type.

## `Terminates`

Terminates is a flag that indicates that a trigger ends execution.

## `RunAfter`

RunAfter indicates that the command will run on shutdown instead of before startup.

## `Trigger`

Trigger is a one-shot function that will be used for things like resetting configurations to default, running reindexing or replay or other similar one-off processes that may sometimes be needed for the application.

## `Default`

The Default field is found in Var containers and is intended to hold the default value that will be assigned to the Slot if no other configuration setting has a value provided.

## `DefaultOn`

DefaultOn is for Triggers and indicates the presence of the Trigger flag means to disable the one-shot function associated with the trigger.

## `Var`

Var is a Tri containing a variable that sets a value for configuration. There is a set of permissible types in Vars that is based on the conventions for JSON values: integer, floating point, string, network address, URL, boolean and lists (separated by commas).

A Var must contain only one each of `name`, `Brief` and `Slot`. Unspecified Default will make zero or empty the resultant value in the absence of any setting.

## `Command`

Command is a Tri containing the definition of a subcommand. `name`, `Brief` and `Handler` are mandatory and singular values that must appear, and optionally one of `Short`, `Usage`, `Group`, `Var`, `Trigger` and `Examples` also may be found here.

## `Commands`

Commands is just an array of Command, containing zero or more Command items.

## `Tri`

Tri is the top-level definition for the application, it reqires the `name`, `Brief` and `Version` fields, and optionally a Commands item and zero or more Var items. 

The `Var` fields define values that are common to all or most of the `Command` fields.

## Syntax Definition Pseudocode

Below is all of the elements within their valid positions (only `name` fields must have specific position), * indicates mandatory and 1 indicates only one is permitted:

      Tri{
         "name", *1
         Brief{""}, *1
         Version{0, 1, 1, "alpha"}, *1
         DefaultCommand{""}, 1
         Var{"name", *1
            Short{"d"}, 1
            Brief{"brief"}, *1
            Usage{"usage"}, 1
            Help{"help"}, 1
            Default{"~/.pod"}, 1
            Slot{""}, 1
         },
         Commands{
            {"ctl", *1
               Short{"c"}, 1
               Brief{"brief"}, *1
               Usage{"usage"}, 1
               Help{"help"}, 1
               Group{"groupname"}, 1
               Examples{ 1
                  "example 1", "explaining text", (pairs of strings)
               },
               Var{}, *1
               Trigger{"init", *1
                  Short{"I"}, 1
                  Brief{"brief"}, *1
                  Usage{"usage"}, 1
                  Help{"help"}, 1
                  DefaultOn{}, 1
                  Terminates{}, 1
                  RunAfter{}, 1
                  func(Tri) int { *1
                     return 0
                  },
               },
               func(Tri) int { *1
               },
            },
         },
      }