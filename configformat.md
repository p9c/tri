# Tri Configuration File Format

In keeping with the goals of the declaration format designed for Tri, the configuration format has been designed to the same ends - to be easy to read for humans, as well as easy to process by the machine.

## Rules for structure of config

1. As in the declaration, all names are letters only, case insensitive and normalised when output automatically to lower case
2. Names starting at the beginning of the line refer to root level Var and Trigger items
3. Command names always appear at the first position of the line, in a group, after all of the non-default values from the root, and all of the command names are always present even if they have no non-default values stored after them.
4. Items belonging to commands are prefixed by a tab at the beginning of the line, and the group is delimited by the next command name at the start or the end of file
5. Items that represent arrays, are likewise grouped under their parent name, with two tabs as prefix, and group ends at the first line with less than two tabs at the start.
6. All content after the name and maybe prefix tabs, after one space after the name, is one whole string that is the value, thus one can have space- and tab-containing content, the only thing a value cannot have is a carriage return, because that is the end marker

By keeping the rules simple, the programmer's task is made simpler, letting them instead spend time on the complicated things that are necessary. Yes, maybe it might be easier to use json or other structured variable type parser/formatter, however, this syntax is so simple the parser is barely more wordy than using the libraries as commonly used.