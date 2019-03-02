# Tri Overview

## Parsing process

1. Compiler performs checks on declaration that are automatic and built into compile-time checking.
   
   Go compiler ensures declaration has correct bracketing, that arrays elements are explicitly or implicitly correctly typed, and so on.

3. Runtime splits CLI invocation string by spaces and provides this in os.Args array

4. Runtime first executes Go var declarations, which creates the in-memory form of the declaration

5. Validate application's Tri declaration, this should be inside an `init()` or first within the `main()`.
   
   In this step all of the other-than-zero defaults on Vars will be determined to be correctly specified in as far as presence, and in most cases, typing.
   
   There is no need to create a resultant struct, as the use of Slot elements connects the destination to the source definition, with its defaults (or implied zeroes) specified. The declaration provides paths and types and content elements for default, and is used to point to the final destination for each of the variables.

6. Read JSON configuration file, that should contain persistent values for Var and Trigger items.

7. Configuration loader then places overridden values into their respective Slot, after checking type is correct.

8. Parse os.Args one by one until a valid trigger or variable definition is found, then the name is sought within the Tri, and if found, value is parsed to expected type, and value assigned to dereferenced Slot pointer value. (each of these can error and halt execution)
   
   As well as those specified in the Tri itself, there is several top-level Var and Trigger types that can also be valid both in config and CLI. Triggers are processed immediately as they are found.

9. Parsing is now complete, all values are placed implicitly indirectly to their correct destination in this process and non-terminating trigger handlers are run (init, save), and then application main is launched with the composed configuration ready.

## Built in Variables and Triggers

Any application built with Tri implicitly has a data directory (defaulting to app name inside appdata or a unix dot folder), a configuration file (in JSON format), which only contains values different from default, including DefaultOn triggers that have to be explicitly disabled.

Built-in variables:

1. DataDir

   Defaults to ~/`appname`

Built in Triggers:

1. `init`

   Deletes configuration file, then exits. Future runs will then start from defaults

3. `save`
   
   At the end of successful parse of config and CLI args, the new state is persisted into the configuration file. Built-in triggers are the only things that may not be written into the configuration - init is terminating, and save is redundant when no parameters have been specified.