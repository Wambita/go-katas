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
Certainly! Here is the entire README.md content in one complete file. You can copy and paste this Markdown content directly into your project's README.md file:

markdown

# Word Order by Embedded Numbers

## Overview

This Go project includes a function named `Order`, designed to reorder words in a given sentence based on a single digit found within each word. This digit indicates the word's intended positional order in the output sentence. The project showcases practical string manipulation and sorting techniques in Go.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

To run this project, you will need:

- Go (Golang) installed on your machine. [Download and install Go](https://golang.org/dl/).

### Installing

Clone the repository to your local machine using the following commands:

```bash
git clone https://github.com/yourusername/word-order-by-numbers.git
cd word-order-by-numbers

### Running the Tests

To execute the tests and ensure the functionality:

```bash
go test -v

This command runs the tests defined in the test file, providing verbose output.

###Usage

Run the program with the main entry file, and the Order function will be executed with a predefined string:

```bash
go run main.go

You can modify the sentence variable in main.go to test different strings:

### Example Sentences for Testing

Here are some additional sentences you can use to test the functionality of the Order function:

   - "this4 is2 a3 test1"
   - "example4 of2 sorting6 based3 on5 embedded1 numbers7"
   - "please2 test1 this4 functionality3"

Replace the sentence in the main.go file with any of the above and run the program to see the result.
Contributing


This project is licensed under the MIT License - see the LICENSE.md file for details.
Authors

   **Wambita**