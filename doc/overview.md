# Tri Overview

## Parsing process

1. Compiler performs checks on declaration that are automatic and built into compile-time checking.
   
   Go compiler ensures declaration has correct bracketing, that arrays elements are explicitly or implicitly correctly typed, and so on.

3. Runtime splits CLI invocation string by spaces and provides this in os.Args array

4. Runtime first executes Go var declarations, which creates the in-memory form of the declaration

5. Validate application's Tri declaration, this should be inside an `init()` or first within the `main()`.
   
   > In this step all of the other-than-zero defaults on Vars will be determined to be correctly specified in as far as presence, and in most cases, type checking. 
   > 
   > **TODO: Validators need to type-check between Default and Slot type in Vars.**
   
   There is no need to create a resultant struct, as the use of Slot elements connects the destination to the source definition, with its defaults (or implied zeroes) specified. The declaration provides paths and types and content elements for default, and is used to point to the final destination for each of the variables.

6. Read JSON configuration file, that should contain persistent values for Var and Trigger items.

7. Configuration loader then places overridden values into their respective Slot, after checking type is correct.

8. Parse os.Args one by one until a valid trigger or variable definition is found, then the name is sought within the Tri, and if found, value is parsed to expected type, and value assigned to dereferenced Slot pointer value. (each of these can error and halt execution)
   
   As well as those specified in the Tri itself, there is several top-level Var and Trigger types that can also be valid both in config and CLI. Triggers are processed immediately as they are found.

9. Parsing is now complete, all values are placed implicitly indirectly to their correct destination in this process and non-terminating trigger handlers are run (init, save), and then application main is launched with the composed configuration ready.

## Types for Vars

In the target application configuration structures for the intended purpose for writing this library, the destination configuration structures have a set of variable types that we must correctly validate and parse.

### Types

These are the types that the initial implementation of Tri will target, their in-configuration storage type, their actual type which they must parse correctly to, and thus be validated both in Default and Slot lists, elaborated as a tree from the configuration store type to the content constraint subtype, if applicable (url, addresses, durations, etc)

- bool

   true/false values, default is false unless Default is present.

- int

   Mostly these are actually scalars. Qualifiers: Some have valid ranges, some have special meaning to -1 (genthreads meaning all available threads).

- uint32

   these are all scalars, in most cases zero is not a default and is invalid. Some refer to sizes in bytes - parser should understand KMG (kilo mega giga) - for this case for simplicity KiB MiB GiB, the base 1024. User likely would not use such multipliers unless it is a byte size, so one parser can handle all of these, as their outputs generally can not be over 2^32 (4 billion, 4GiB).

- float64

   In pod these are all amounts of currency, maximum precision of 8 decimal places (satoshi). Excess decimal places are truncated before parsing string to float

- string

   Strings are a mix of quite different types. One is a port number spec, others are filesystem paths, some are URLs and some are network addresses. This especially hints towards creating validator handlers (I think maybe this is the whole solution)

- []string

   Most of these are either URLs or addresses that one or more instances of the variable name may be present and each item appends to the slice if valid

- time.Duration

   Time has a simple parser for this. A wrapper is needed for it.

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