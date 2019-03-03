# Tri Overview

## Parsing process

1. Compiler performs checks on declaration that are automatic and built into compile-time checking.
   
   Go compiler ensures declaration has correct bracketing, that arrays elements are explicitly or implicitly correctly typed, and so on.

3. Runtime splits CLI invocation string by spaces and provides this in os.Args array

4. Runtime first executes Go var declarations, which creates the in-memory form of the declaration, as well as all of the built in and custom handlers for their different types and purposes.

5. Validate application's Tri declaration.
   
   > In this step all of the other-than-zero defaults on Vars will be determined to be correctly specified in as far as presence, and in most cases, type checking. 
   
   There is no need to create a resultant struct, as the use of Slot elements, which are interface{} containing pointer to (optionally multiple) other variables connects the destination to the source definition, with its defaults (or implied zeroes) specified. The declaration provides paths and types and content elements for default, and is used to point to the final destination for each of the variables.

6. Parse os.Args one by one and generate a `map[string]interface{}` containing the names and the type expected is found by locating the name in the Tri and type switch on the Slot that Var types contain.

7. Based on CLI arg specified data directory found in the previous step, or from the default location, JSON configuration file is read and parsed.

8. Configuration loader then places decoded, and validated values into their respective Slot, after checking type is correct, which is performed by the handlers specified in Vars.

9.  Built-in triggers take precedence over custom triggers. Init terminates after clearing the configuration file, save rewrites it after parsing and before initiating the Command handler that is specified.

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

   Time has a simple parser for this. A wrapper is needed for it, and one is implemented in the set of default variable parser handlers that must be present in the Var.

Rather than create an arbitrary set of human readable string type specifications, all of the typing is handled by the Go compiler, through the use of handlers. The handlers determine correct destination type from the Slot, and the default handler is one function with a type switch on the Slot types, in which the input string value attempts to parse, halting if the format of the value is invalid.

If types other than the standard set are needed, the programmer using this library can create their own var handlers to enable more types than the default set.

## Built in Variables and Triggers

Any application built with Tri implicitly has a data directory (defaulting to app name inside appdata or a unix dot folder), a configuration file (in JSON format), which only contains values different from default, including DefaultOn triggers that have to be explicitly disabled.

### Built-in variables:

1. datadir/D

   Defaults to ~/`appname`

   'appname' in this case refers to the first field of the top level Tri structure. In the case of Windows applications this will be in (usually) `C:\users\username\appdata\appname`. If the specified path does not exist, it will be created. The default handler will cover the case of POSIX/Unix style dot folders and windows appdata folders.

### Built in Triggers:

1. `init`

   Deletes configuration file, then exits. Future runs will then start from defaults. Configuration files only store values that are not default.

2. `save`
   
   At the end of successful parse of config and CLI args, the new state is persisted into the configuration file.

3. `defaults`

   Defaults will print the entire set of names as they would appear in the configuration, with their default values afterwards, with one prefix tab grouping command items and two prefix tabs grouping items in lists (if default of this array element *has* more than one item).

Also note that logging configuration is not handled by default, nor is there a parser for it. If the application needs these configurations, the handler must be written by the developer to fit the system they are using. I recommend the logger I wrote, found within the Parallelcoin `pod` repository, [located here](https://github.com/parallelcointeam/pod/tree/master/pkg/util/clog).