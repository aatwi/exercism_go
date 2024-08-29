# Instructions
The Collatz Conjecture or 3x+1 problem can be summarized as follows:

Take any positive integer n. If n is even, divide n by 2 to get n / 2. If n is odd, multiply n by 3 and add 1 to get 3n + 1. Repeat the process indefinitely. The conjecture states that no matter which number you start with, you will always reach 1 eventually.

Given a number n, return the number of steps required to reach 1.

## Examples
Starting with n = 12, the steps would be as follows:

1. 12
1. 6
1. 3
1. 10
1. 5
1. 16
1. 8
1. 4
1. 2
1. 1

Resulting in 9 steps. So for input n = 12, the return value would be 9.

# How to debug
When a test fails, a message is displayed describing what went wrong and for which input. You can also use the fact that console output will be shown too. You can write to the console using:

```go
import "fmt"
fmt.Println("Debug message")
````

**Note:** When debugging with the in-browser editor, using e.g. `fmt.Print` will not add a newline after the output which can provoke a bug in `go test --json` and result in messed up test output.