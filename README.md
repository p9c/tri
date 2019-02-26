# tri
Tri is a CLI parameter parsing and configuration library designed to populate app configurations automatically based on defaults, env, config files and CLI args. The name Tri relates to the way the declarations are tree structured, as well as having an intrinsic tripartite structure, tag, container and parameter set.

Virtually all CLI args/configuration libraries written in Go have several irritating deficiencies:

- Boilerplate-laden declaration syntax... or
- Verbose system of parameter setting methods (a bit of the previous plus really ugly syntax)
- Configuring defaults is often completely manual and thus error prone
- Even when the configuration composition is automated, it can't easily be parlayed into an existing config structure

I have constrained the scope of this library to only cover configuration.

It is not quite complete in that it lacks log rotation and placeholder implementation of logging to file, but https://github.com/parallelcointeam/pod/tree/master/pkg/util/clog is a channel-based logging library I have written that aims to reduce runtime overhead of allocation and function context switching by making the logging path pass through channels, and is designed such that one simple, very short log.go file can be added and entire packages are covered.

## Readable declaration syntax in pure Go

The declaration types used by Tri are designed to combine readability with type-enforced structuring at compile time as far as possible, and then a full set of validators for every possible are defined and every container subtype checks for mandatory, impermissible and malformed elements, which executes in a depth-first descent and when it errors, returns the value through a series of error returns that halts at the first parse/validation error and informs the programmer exactly where the declarations are incorrect.

By creating a readable syntax for declarations, parsing is handled by the Go compiler and reduces the need for some arbitrary object notation syntax (json, toml, csv, xml, do we really need more of them???), the other thing is that it is much more complicated to bind go executable code into such structures anyway.

## Comprehensive input sanitisers for app declarations

The validators are very strict, and implement the checking that is not possible to enforce at runtime either due to Go's design or the nature of the specification. It is intended also that this library serves as an example of how 'generic' types are correctly implemented in Go, in that in essence the declarations are a subset of Go.

Edge cases can skulk around in code for years before someone finds it, and really, it would be better if that was a good guy rather than someone intent on stealing or destroying data they have no moral right to.

## Automatic loading of external configuration structures for upgrades and multi-function executables

Within each subcommand, each configurable parameter has a default, sane value, which is overridden by the configuration file, environment variables and CLI flags, in that order, the last one in the line being the one that prevails. The values are stored into a slot that is filled during the declaration by the use of a pointer to a configuration struct/variable used by the application this library is used to configure.

Once Tri finishes validation and gathers and composes all of the values, the configuration variables for the application are fully validated and filled and can immediately then be executed, eliminating a major cause of errors in such code, where values are not properly loaded, loaded into the wrong place, or not effectively exported to the application at all.

## Motivation

This library was written in order to ease the implementation of multi-command concurrent servers in the root of the repository referred to above about the logger (the parallelcoin pod), which contains a bitcoin-based full node and wallet server, separated for security reasons, but a complete mismatch with our needs, and with the light weight of the chain, the most direct route to a nice GUI.

The climax library was what I initially used, but its lack of facilities to set defaults or pass the composed set of parameters to the final configuration structure meant that I ended up writing extremely bulky, very long, and thus error prone configuration parsing functions, which proved to be a showstopper on the path to release.

Of course, careful and irritating double, triple, and quadruple checking of this fluff might have solved the issue as quickly as I could write this library, but whoever was up for the job of maintenance, which is hopefully me, will be visited again by this abomination, which I have already suffered from enough.
