# Tri Configuration File Format

In keeping with the goals of the declaration format designed for Tri, the configuration format has been designed to the same ends - to be easy to read for humans, as well as easy to process by the machine.

It is not intended that editing the file be necessary, as far as possible, by humans, but to not make it difficult for a human to comprehend and compose them.

Tri's configuration does not have a great deal of depth to it, the tree has a maximum length of 3 nodes from the root and the root does not need to to be named

- Trigger and Var fields at root level are just name, space, value.

- Each command section starts with the name of the command, and then every element related to it has a prefix tab, then name and the remainder of the line is the value.

- Config will write only fields that have different than default. A human can add such a setting if they want, but the next time it persists it (like when the flag 'save' is given), this redundant field will be omitted, as the output is generated only when the content of the Slot differs from the Default.

- Incorrect lines inside a group have the same effect as an empty line or nothing at all - no effect. Thus, one can use anything at all aside from an alphabet character or tab on the first position in the line and it will be ignored as a comment, but the only use for this is such as to share a configuration on a help forum or so.

- Comment lines will be lost when Tri is asked to write the configuration to disk - because preserving them requires complex parsing and there really is no good reason for this. The user can set several values in one shot by using the 'save' root level built in trigger, and the non-default values after full parse and composition will be written, and the help built-in trigger will list every possible item anyway, at the root specified after 'help'.

- The config writer will always put every subcommand's header line in, regardless of whether any variables in the group are different from default, so that a human can see where to insert the tab prefixed, and so that the parser only needs to know what the last root level name was, until it encounters another one.

### Example configuration items

```
datadir /path/to/nondefault/data/directory
loglevel debug
profile /path/to/profile (default ./)
cpuprofile http://address:port (default http://localhost:1100)
testnet
simnet
ctl
   wallet
   rpcserver http://address:port
   username username
   password pa55word
gui
node
   listener http://address:port
   useragentcomments crypto widget miners
   dropcfindex
   generate
   genthreads 4
   algo sha256d
shell
wallet
   rpcserver http://address:port
   username username
   password pa55word
   file /path/to/wallet
```

## Configuration Parsing Procedure

Firstly, before reading the configuration, the CLI args must be parsed. In these args, the data directory root can be specified.

It is not necessary, really, to provide any more customisable paths than this, so it is suggested to users of the library if the target config has these, just derive them off the datadir. 

The default tree inside the datadir is simple and unambiguous - each subcommand's configured root will be a folder matching the command's name, the default log file location, for example, should just be ./commandname.log, and in the case this library is developed for, under node would be testnet/mainnet/simnet and inside there where the databases will be stored. This is really outside the scope of this library, really just a recommendation.

With the datadir in hand, the filename 'config' is looked for, if it does not exist, the empty file will be created.

If the file has contents, all letters-only first field represent nodes at the base of the Tri

Command names always appear in the configuration, and mark the beginning of triggers and variables related to a subcommand

All subcommand values start with a tab

Any line that does not have either a tab or a letters-only name as the first thing, is skipped over immediately

If the name appearing first on any line (prefix tab for command groups) is not found where it is expected, the parser will error and the context and position will be shown along with a suggestion to use the 'save' builtin trigger, which also omits all non-parsing lines.

Configuration is one element per line, so always first word is a node name, and everything after that and a space is the value