# Go Notes

### This document wil hold pointers to key things to remember in go
### This should always be easy to modify and refer back

## Pointers
- A slice in `go` is not the same as slice in `python`
    - A `go` slice will share memory
    - A `python` slice will create a copy
    - To have a similar thing like `python` that is create a copy in `go` we use `copy()` make the destination first, with `make([]int, len(a))`, and then call `copy()`.
- 'a' is a rune (one character) and "a" is a string.
    - Again `python` don't care much but we should and as a habbit character should be `''` and string should be `""`
- Go never mixes int and float64 without an explicit conversion.
    - Well python i guess was made in first place to give this flexiblity ... _wine_ goes with _wine_ and _whisky_ goes with _whisky_