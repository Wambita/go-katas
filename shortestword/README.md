## Kata - Shortest Word 

This program includes a function `FindShort` that takes a sentence as input and returns the length of the shortest word in that sentence. 

## Function Description

`FindShort` is a Go function that operates by:
1. Splitting the input string into a slice of words based on spaces.
2. Sorting this slice by the length of each word in ascending order.
3. Returning the length of the shortest word in the input string.

### Function Signature

```go
func FindShort(s string) int
```
### Usage

To use this function in a Go project, you can include it directly in your program. Below is a basic example of how to use the FindShort function:

```go

package main

import "fmt"

func main() {
    sentence := "Let's travel abroad shall we"
    shortestWordLength := FindShort(sentence)
    fmt.Printf("The shortest word length is: %d\n", shortestWordLength)
}
```
### How to Run the Program

To run the example program, follow these steps:

- Ensure you have Go installed on your computer. If not, you can download and install it from Go's official site.
- Save the above example code in a file named main.go.
- Open a terminal or command prompt.
- Navigate to the directory containing main.go.
- Run the program by typing go run main.go.

### Testing

This program comes with a test file named shortestword_test.go that includes several test cases to ensure the functionality works correctly under various conditions. To run the tests, follow these steps:

- Save the test code in a file named shortestword_test.go in the same directory as main.go.
- Open a terminal or command prompt.
- Navigate to the directory containing your Go files.
- Run the tests by typing go test in the terminal. For verbose output, use go test -v.

### Contributing

Contributions to improve the program or expand the test cases are welcome. Please ensure to run the existing tests and add new ones if you're introducing changes to the function.

### License

This project is open-source and available under the MIT License.


## Authors
- [Wambita](https://github.com/Wambita)