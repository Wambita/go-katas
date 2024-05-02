# Kata - Your Order, Please

## Overview

This kata project a function named `Order` which takes a string and reorders each word in it based on a number embedded in the word itself. It also includes a test file that has a struct that ensures the output matches the expected output. The project demonstrates the use of string manipulation, sorting algorithms, and Go testing practices.

## Functionality

The `Order` function processes input strings where each word contains a single digit. This digit indicates the word's position in the output string. For example:

- Input: "is2 Thi1s T4est 3a"
- Output: "Thi1s is2 3a T4est"

## Getting Started

### Prerequisites

- Go (Golang) installed on your machine. [Download Go](https://golang.org/dl/)

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/Wambita/go-katas.git
cd yourorderplease
```

### Running the Tests

To execute the tests and ensure the functionality:

```bash
go test -v
```
This command runs the tests defined in the test file, providing verbose output.

### Usage

Run the program with the main entry file, and the Order function will be executed with a predefined string:

```bash
go run main.go
```

### Example Sentences for Testing

Here are some additional sentences you can use to test the functionality of the Order function:

   - "this4 is2 a3 test1"
   - "example4 of2 sorting6 based3 on5 embedded1 numbers7"
   - "please2 test1 this4 functionality3"

This project is licensed under the MIT License - see the LICENSE.md file for details.

### Authors

   **Wambita**