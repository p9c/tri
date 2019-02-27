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

      - [ ] string in index 0
      - [ ] string is valid name (letters only)
      - [ ] more than one brief not allowed
      - [ ] more than one handler not allowed
      - [ ] handler not nil
      - [ ] Brief field present
      - [ ] Handler present

   - [ ] `Commands.Validate()`

      This array alias has no errors of its own

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

   - [ ] `Examples.Validate()`

      - [ ] must not be empty
      - [ ] must have pairs of elements
      - [ ] elements must be strings
      - [ ] even numbered (first in pair) elements have no control characters

   - [ ] `Group.Validate()`

      - [ ] contains only one element
      - [ ] element is a string
      - [ ] string is a ValidName

   - [ ] `Help.Validate()`

      - [ ] contains only one element
      - [ ] element is a string

   - [ ] `RunAfter.Validate()`

      - [ ] may not contain anything

   - [ ] `Short.Validate()`

      - [ ] contains only one element
      - [ ] element is a rune (single character/unicode point)

   - [ ] `Slot.Validate()`

      - [ ] slots are all the same type (pointer to said type)

   - [ ] `Terminates.Validate()`

      - [ ] contains no elements

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

   - [ ] `Usage.Validate()`

      - [ ] only one element
      - [ ] element is string
      - [ ] string is no more than 80 chars long
      - [ ] string contains no control characters

   - [ ] `Var.Validate()`

      - [ ] contains at least 3 elements
      - [ ] first is string
      - [ ] name is ValidName
      - [ ] only one Brief
      - [ ] only one handler
      - [ ] has one each of Brief and Handler

   - [ ] `Version.Validate()`

      - [ ] has no more than 4 fields
      - [ ] has at least 3 fields
      - [ ] first three are integers
      - [ ] integers are under 100
      - [ ] 4th field is a string
      - [ ] string contains only letters and numbers

   - [ ] `tri.ValidName()`

      - [ ] contains only letters


## Documentation

   - [x] Markdown declaration syntax description
   - [x] godoc comments properly formatted and reasonably complete
   - [x] nice checklist (this!)

## Commandline Parser

   - [ ] Scanner