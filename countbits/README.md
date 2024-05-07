## Kata - Bit Counting

This module provides a `CountBits` function that calculates the number of bits set to '1' in the binary representation of a non-negative integer. This function is particularly useful in understanding binary operations and data analysis involving bit manipulation.

### Function Description

`CountBits` takes a single non-negative integer (`uint`) as input and returns the count of bits that are set to '1' in its binary representation.

### Example

The binary representation of 1234 is `10011010010`, so the function should return `5` because there are five '1's in the binary representation.

### Usage

To use the CountBits function, you need to include it in your Go project. Below is a simple example of how you can call this function from another Go program:

```go

package main

import (
    "fmt"
    "yourproject/kata" // Replace 'yourproject' with your actual project name
)

func main() {
    number := uint(1234)
    result := kata.CountBits(number)
    fmt.Printf("Number of 1s in the binary representation of %d: %d\n", number, result)
}
```
### Running the Example

To run the example above, ensure you have the Go environment setup on your system. Save the code in a file (e.g., main.go), navigate to the directory containing the file via your terminal or command prompt, and run the following command:

```bash

go run main.go
```
This command compiles and runs your Go program, outputting the count of '1's in the binary representation of the given number.

### Testing

To test the correctness of the CountBits function, refer to the provided test suite in countbits_test.go. To execute the tests, run:

```bash

go test -v
```
Ensure that your working directory is where the test file is located, or specify the path to the test file if needed.

### Conclusion

The CountBits function is a simple yet powerful tool for analyzing the bit-level representation of integers. It is helpful for educational purposes, debugging, or any application requiring bit manipulation.


## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Authors
- [Wambita](https://github.com/Wambita)