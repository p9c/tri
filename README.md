# tri
Tri is a CLI parameter parsing and configuration library designed to populate app configurations automatically based on defaults, env, config files and CLI args. The name Tri relates to the way the declarations are tree structured, as well as having an intrinsic tripartite structure, tag, container and parameter set.

Virtually all CLI args/configuration libraries written in Go have several irritating deficiencies:

- Boilerplate-laden declaration syntax... or
- Verbose system of parameter setting methods (a bit of the previous plus really ugly syntax)
- Configuring defaults is often completely manual and thus error prone
- Even when the configuration composition is automated, it can't easily be parlayed into an existing config structure

I have constrained the scope of this library to only cover configuration.

It is not quite complete in that it lacks log rotation and placeholder implementation of logging to file, but https://github.com/parallelcointeam/pod/tree/master/pkg/util/clog is a channel-based logging library I have written that aims to reduce runtime overhead of allocation and function context switching by making the logging path pass through channels, and is designed such that one simple, very short log.go file can be added and entire packages are covered.

The declaration types used by Tri are designed to combine readability with type-enforced structuring at compile time as far as possible, and then a full set of validators for every possible are defined and every container subtype checks for mandatory, impermissible and malformed elements, which executes in a depth-first descent and when it errors, returns the value through a series of error returns that halts at the first parse/validation error and informs the programmer exactly where the declarations are incorrect. 

The validators are very strict, and implement the checking that is not possible to enforce at runtime either due to Go's design or the nature of the specification. It is intended also that this library serves as an example of how 'generic' types are correctly implemented in Go. There is no such thing as a trivial parsing error, something all Go programmers eventually fully understand, and so, it naturally follows therefore that a Go programmer should adhere to these principles. There is no such thing as a trivial or negligable error. Edge cases can skulk around in code for years before someone finds it, and really, it would be better if that was a good guy rather than someone intent on stealing or destroying data they have no moral right to.

Within each subcommand, each configurable parameter has a default, sane value, which is overridden by the configuration file, environment variables and CLI flags, in that order, the last one in the line being the one that prevails. The values are stored into a slot that is filled during the declaration by the use of a pointer to a configuration struct/variable used by the application this library is used to configure, so once Tri finishes validation and gathers and composes all of the values, the configuration variables for the application are fully validated and filled and can immediately then be executed. 

## Motivation

This library was written in order to ease the implementation of multi-command concurrent servers in the root of the repository referred to above, which contains a bitcoin-based full node and wallet server, separated for security reasons, to be able to launch them at the same time automatically configured to connect to each other using loopback, and specifically for the one case not considered by the original designers of these servers, being impractical for bitcoin - a full node wallet GUI without any SPV shortcuts (as the chain is currently very small in data size, under 150Mb at the time of writing.

The climax library was what I initially used, but its lack of facilities to set defaults or pass the composed set of parameters to the final configuration structure meant that I ended up writing extremely bulky, very long, and thus error prone configuration parsing functions, which proved to be a showstopper on the path to release. Of course, careful and irritating double, triple, and quadruple checking of this fluff might have solved the issue as quickly as I could write this library, but whoever was up for the job of maintenance, which is hopefully me, will be visited again by this abomination, so the exact, error-resistant and robust method I had in mind is the basis of every aspect of the design of this library.

> At this point I am not 100% certain this will not simply extend climax, but I think it probably will not. Climax features several parts that seem to me to be pretty much beyond the scope that almost anyone needs for a configuration/flags library, wikis, 'topics' and the like, and it has no configuration file handling or profile directory handling, and the existing code that does this all needs to be totally rewritten to protect me from my wandering attention, since the gravity of an exploitable bug in this kind of application can cost its victims literally millions of dollars, something that I want to assure myself I am not allowing the possibility of, before it gets out of my hands and out of my control.
