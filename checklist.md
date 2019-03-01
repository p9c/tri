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

### Tests

Each validator has several error conditions, so each of them are elaborated under the validation function's heading.

- [x] Test stubs written

   - [x] `Brief.Validate()`

      - [x] one item only
      - [x] string typed item
      - [x] string length < 80
      - [x] no control characters
      - [x] no error!

   - [ ] `Command.Validate()`

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

      This array alias has no errors of its own, Go enforces that its contents are typed correctly.

   - [x] `Default.Validate()`

      - [x] only one item
      - [x] item is string
      - [x] item is a ValidName
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

   - [ ] `Tri.Validate()`

      - [ ] contains at least 4 elements
      - [ ] first element is a string
      - [ ] string is a ValidName
      - [ ] contains (only) one Brief
      - [ ] contains (only) one Version
      - [ ] contains (only) one Commands
      - [ ] contains no more than one DefaultCommand
      - [ ] only contains element from set of possible elements
      - [ ] Brief is missing
      - [ ] Version is missing
      - [ ] Commands is missing

   - [ ] `Trigger.Validate()`

      - [ ] must contain name, Brief and Handler
      - [ ] first element is name
      - [ ] name is a ValidName
      - [ ] only one Brief
      - [ ] only one Handler
      - [ ] rest of items only (one of) Short, Usage, Help, DefaultOn, Terminates and RunAfter

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

## Commandline Parser

   - [ ] Scanner