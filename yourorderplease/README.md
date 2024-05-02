# Kata - Your Order, Please

## Overview

This kata project a function named `Order` which takes a string and reorders each word in it based on a number embedded in the word itself. It also includes a test file that has a struct that ensures the output matches the expected output. The project demonstrates the use of string manipulation, sorting algorithms, and Go testing practices.

## Algorithm
      
   #### Order Function:
        **Input/Output:** Accepts a sentence string, returns the reordered sentence based on embedded numbers.

        **Splitting:**  Uses strings.Split to break the sentence into words.
        
        **Mapping Numbers to Words:**  Uses a loop to examine each word, and a nested loop to examine each character. If a character is a digit (excluding '0'), it maps the word to its numeric position using a map where the key is the extracted number.

        **Compiling Results:** Allocates a slice to store words in their correct order. It then fills this slice using the map, ensuring words are placed according to their numerical position as indicated by their embedded number.

        **Constructing the Ordered Sentence: **  Merge the ordered words back into a single string.

   #### Main Function:
       - The program defines a sentence string that includes words embedded with numbers.
       - It calls the Order function which is designed to parse these numbers and reorder the words accordingly.
        The result is printed using println.


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