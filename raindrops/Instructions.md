# Introduction
Raindrops is a slightly more complex version of the FizzBuzz challenge, a classic interview question.

# Instructions
Your task is to convert a number into its corresponding raindrop sounds.

If a given number:

- is divisible by 3, add "Pling" to the result.
- divisible by 5, add "Plang" to the result.
- is divisible by 7, add "Plong" to the result.
- is not divisible by 3, 5, or 7, the result should be the number as a string.

## Examples
- 28 is divisible by 7, but not 3 or 5, so the result would be "Plong".
- 30 is divisible by 3 and 5, but not 7, so the result would be "PlingPlang".
- 34 is not divisible by 3, 5, or 7, so the result would be "34".

### Note
A common way to test if one number is evenly divisible by another is to compare the remainder or modulus to zero. Most languages provide operators or functions for one (or both) of these.

# How to debug
When a test fails, a message is displayed describing what went wrong and for which input. You can also use the fact that console output will be shown too. You can write to the console using:

```go
import "fmt"
fmt.Println("Debug message")
```
**Note:** When debugging with the in-browser editor, using e.g. fmt.Print will not add a newline after the output which can provoke a bug in go test --json and result in messed up test output.