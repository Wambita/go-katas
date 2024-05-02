## Kata - Multiples of 3 and 5

This Go program efficiently calculates the sum of all numbers that are multiples of 3 or 5 and are less than a specified input number.
Problem Statement

The challenge is to find the sum of all numbers below a given number which are multiples of either 3 or 5.

### Getting Started

These instructions will guide you through the setup process and get you up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, ensure you have Go installed on your computer. The recommended version is Go 1.15 or later.

### Installation

Follow these steps to set up the project on your local machine:

-Clone the repository:

   ``` bash

git clone https://go-katas.com
```
Navigate to the project directory:

```bash

    cd multiplesof3and5
```
### Running the Tests

You can run the automated tests for this system by executing:

```bash

go test -v
```
### Algorithm

The program features a function named Multiple3And5 which encapsulates the core logic. Here's how it works:

- Initialize the Sum: Start with a sum of zero.
- Iterate through Numbers: Loop from the given number minus one down to 1.
- Check for Multiples: During each iteration, check if the current number is divisible by 3 or 5.
- Sum Multiples: If a number is a multiple of either, add it to the sum.
- Return the Sum: After the loop completes, print and return the total sum.

This algorithm ensures that all relevant multiples are added precisely once, providing an efficient solution to the problem.

### Example Usage

Here's how you can use the function in a simple Go program:

```go

package main

import (
    "fmt"
)

func main() {
    result := Multiple3And5(30)
    fmt.Println("The sum of multiples of 3 and 5 below 30 is:", result)
}
```
### Contributing

We welcome contributions to this project. Please fork the repository and submit a pull request with your improvements.

### License

This project is licensed under the MIT License - see the LICENSE.md file for details.

### Authors
This Project was created by ***@Wambita***