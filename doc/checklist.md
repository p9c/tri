# Checklist for completion of components of Tri

## Validators 

### Initial draft

   - [x] `Brief.Validate()`
   - [x] `Command.Validate()`
   - [x] `Commands.Validate()`
   - [x] `Default.Validate()`
   - [x] `DefaultCommand.Validate()`
   - [x] `DefaultOn.Validate()`
   - [x] `Examples.Validate()`
   - [x] `Group.Validate()`
   - [x] `Help.Validate()`
   - [x] `RunAfter.Validate()`
   - [x] `Short.Validate()`
   - [x] `Slot.Validate()`
   - [x] `Terminates.Validate()`
   - [x] `Tri.Validate()`
   - [x] `Trigger.Validate()`
   - [x] `Usage.Validate()`
   - [x] `Var.Validate()`
   - [x] `Version.Validate()`
   - [x] `tri.ValidName()`

### Declaration syntax validator tests

Each validator has several error conditions, so each of them are elaborated under the validation function's heading.

- [x] Test stubs written
- [x] 100% coverage

   - [x] `Brief.Validate()`

      - [x] one item only
      - [x] string typed item
      - [x] string length < 80
      - [x] no control characters
      - [x] no error!

   - [x] `Command.Validate()`

      - [x] not empty
      - [x] string in index 0
      - [x] string is valid name (letters only)
      - [x] more than one brief not allowed
      - [x] brief is invalid
      - [x] more than one handler not allowed
      - [x] handler not nil
      - [x] no more than one Short
      - [x] invalid Short
      - [x] no more than one Usage
      - [x] invalid Usage
      - [x] no more than one Help
      - [x] invalid Help
      - [x] no more than one Examples
      - [x] invalid Examples
      - [x] invalid Var
      - [x] invalid Trigger
      - [x] Brief field present
      - [x] Handler present
      - [x] invalid typed element
      - [x] no errors!

   - [x] `Commands.Validate()`

     - [x] Command elements are all valid

   - [x] `Default.Validate()`

      - [x] only one item
      - [x] no error!

   - [x] `DefaultCommand.Validate()`

      - [x] only one element
      - [x] element is string
      - [x] string is a ValidName
      - [x] no error!

   - [x] `DefaultOn.Validate()`

      - [x] must be empty
      - [x] no error!

   - [x] `Examples.Validate()`

      - [x] must not be empty
      - [x] must have pairs of elements
      - [x] elements must be strings
      - [x] even numbered (first in pair) elements have no control characters
      - [x] first field longer than 40 characters
      - [x] second field longer than 80 characters
      - [x] no error!

   - [x] `Group.Validate()`

      - [x] contains only one element
      - [x] element is a string
      - [x] string is a ValidName
      - [x] no error!

   - [x] `Help.Validate()`

      - [x] contains only one element
      - [x] element is a string
      - [x] no error!

   - [x] `RunAfter.Validate()`

      - [x] may not contain anything
      - [x] no error!

   - [x] `Short.Validate()`

      - [x] contains only one element
      - [x] element is a rune (single character/unicode point)
      - [x] element is a letter or number
      - [x] no error!

   - [x] `Slot.Validate()`

      - [x] slots are all the same type (pointer to said type)
      - [x] slots are all pointers
      - [x] no error!

   - [x] `Terminates.Validate()`

      - [x] contains no elements
      - [x] no error!

   - [x] `Tri.Validate()`

      - [x] contains at least 3 elements
      - [x] first element is a string
      - [x] string is a ValidName
      - [x] contains no more than one Brief
      - [x] contains no more than one Version
      - [x] contains no more than one Commands
      - [x] contains no more than one DefaultCommand
      - [x] DefaultCommand with no Commands array
      - [x] DefaultCommand's name appears in also present Commands array
      - [x] contains invalid Var
      - [x] contains invalid Trigger
      - [x] contains invalid DefaultCommand
      - [x] contains invalid Command in Commands
      - [x] only contains element from set of possible elements
      - [x] contains invalid Brief
      - [x] contains invalid Version
      - [x] Brief is missing
      - [x] Version is missing
      - [x] no error!

   - [x] `Trigger.Validate()`

      - [x] contains at least 3 elements
      - [x] first is string
      - [x] name is ValidName
      - [x] has only one Brief
      - [x] has only one Short
      - [x] has only one Usage
      - [x] has only one Help
      - [x] has only one handler
      - [x] has only one DefaultOn
      - [x] has only one RunAfter
      - [x] has only one Terminates
      - [x] has invalid Brief
      - [x] has invalid Short
      - [x] has invalid Usage
      - [x] has invalid Help
      - [x] has invalid handler
      - [x] has invalid DefaultOn
      - [x] has invalid Terminates
      - [x] has invalid RunAfter
      - [x] has one each of Brief and handler
      - [x] has no other type than those foregoing
      - [x] no error!

   - [x] `Usage.Validate()`

      - [x] only one element
      - [x] element is string
      - [x] string is no more than 80 chars long
      - [x] string contains no control characters
      - [x] no error!

   - [x] `Var.Validate()`

      - [x] contains at least 3 elements
      - [x] first is string
      - [x] name is ValidName
      - [x] has only one Brief
      - [x] has only one Short
      - [x] has only one Usage
      - [x] has only one Help
      - [x] has only one Default
      - [x] has only one Slot
      - [x] has invalid Brief
      - [x] has invalid Short
      - [x] has invalid Usage
      - [x] has invalid Help
      - [x] has invalid Default
      - [x] has invalid Slot
      - [x] has one each of Brief and Slot
      - [x] has no other type than those foregoing
      - [ ] Default value is assignable to dereferenced Slot pointer
      - [x] no error!

   - [x] `Version.Validate()`

      - [x] has no more than 4 fields
      - [x] has at least 3 fields
      - [x] first three are integers
      - [x] integers are under 100
      - [x] 4th field is a string
      - [x] string contains only letters and numbers
      - [x] no error!

   - [x] `tri.ValidName()`

      - [x] 3 or more characters length
      - [x] contains only letters
      - [x] no error!

## Documentation

   - [x] Markdown declaration syntax description
   - [x] godoc comments properly formatted and reasonably complete
   - [x] nice checklist (this!)

## Commandline Scanner

   - [ ] recognise - and -- prefixed var/trigger items
   - [ ] recognise values assigned by --name=value and --name value to be one part
   - [ ] find all of the names in passed Tri declaration that CLI args override and error for those not found
   - [ ] ensure values in Vars are correct type based on Tri declaration
   - [ ] recognise top level Tri builtin trigger version/v, save/S and init/I, being print version, save state after configuration to config file, and revert config to default (ie, empty it) - these triggers should run immediately they are found (this is why arrays were used instead of maps), with the save builtin triggering configuration rewrite
   - [ ] recognise and run custom triggers when and how they are specified, as they are found

## Configuration and triggers

   - [ ] read config and fill fields provided that parse correctly or return error
   - [ ] write only fields that differ from default values
   - [ ] special builtin Tri top-level Var datadir, and library default (based on home dir with dot folder bearing Tri name)

## Configuration Composition

   - [ ] Default base is filled from declaration automatically by Slot fields
   - [ ] Configuration file values replace defaults
   - [ ] Command line parameters load over top of result of previous two steps
   - [ ] When when save/S builtin is found, trigger rewrite of config file prior to launch