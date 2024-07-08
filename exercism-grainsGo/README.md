# Grains Go solution  

## Instructions
Calculate the number of grains of wheat on a chessboard given that the number on each square doubles.

## Overview 
There once was a wise servant who saved the life of a prince. The king promised to pay whatever the servant could dream up. Knowing that the king loved chess, the servant told the king he would like to have grains of wheat. One grain on the first square of a chess board, with the number of grains doubling on each successive square.

There are 64 squares on a chessboard (where square 1 has one grain, square 2 has two grains, and so on).

Write code that shows:

    how many grains were on a given square, and
    the total number of grains on the chessboard

## Prerequisites

To compile and run this program, you need to have golang installed on your machine

## Installation

    Clone the repository:

```bash
git  clone https://github.com/Wambita/go-katas.git
cd exercism-grainsGo
```

Compile the program:

```bash
    go run . 4
```

Replace `4` with the desired number for your executable if needed.

## Functions
### square(number float64) int

This function calculates the number of grains on a specific square of a chessboard.

    Parameters: index is a float64 representing the square number on the chessboard (1 to 64).

    Return Value: Returns the number of grains on the index-th square using bitwise left shift operations.

This function calculates the total number of grains on all squares of a chessboard (1 to 64).

    Parameters: This function takes no parameters (void).

    Return Value: Returns the total number of grains on all squares using bitwise left shift operations.

### main 
This function is used to run the program.

## Usage

Run the  program and interact with it as needed. Here's an example usage:

```bash

go run . 3
```
Example

Suppose you want to calculate the number of grains on the 3rd square of the chessboard:

```c
printf("Number of grains on square 3: %ld\n", grains);
```
This will output:

```c

Number of grains on square 3: 8
```
## Contributing

Contributions are welcome! Here's how you can contribute:

   + Fork the repository on GitHub.
   + Clone your forked repository (git clone https://github.com/Wambita/go-katas).
   + Create a new branch (git checkout -b feature-branch).
   + Make your changes and commit them (git commit -am 'Add new feature').
   + Push to the branch (git push origin feature-branch).
   + Create a new Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author
This project is created and maintained by [Wambita](https://github.com/Wambita)