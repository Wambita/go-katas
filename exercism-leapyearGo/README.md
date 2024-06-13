# Leap Year implementation in  Go

## Introduction

A leap year (in the Gregorian calendar) occurs:

- In every year that is evenly divisible by 4.
- Unless the year is evenly divisible by 100, in which case it's only a leap year if the year is also evenly divisible by 400.

**Some examples**:

- 1997 was not a leap year as it's not divisible by 4.
- 1900 was not a leap year as it's not divisible by 400.
- 2000 was a leap year!

## Instructions

Your task is to determine whether a given year is a leap year.

## Algorithm
1. Accept a year as input.
2. Check if the year is evenly divisible by 100.
    - If it is, check if it's also divisible by 400. If yes, it's a leap year.
3. If the year is not divisible by 100, check if it's divisible by 4.
    - If it is, it's a leap year.
4. If none of the conditions are met, the year is not a leap year.

## Installation 
Clone the repository and navigate to the project directory

```
git clone https://github.com/Wambita/go-katas.git
 cd exercism-leapyearGo
```
## Usage
To use this program, run it with the go run command followed by the year you want to check. For example:

#### Leap year 
```bash
$ go run . 2024
2024 is a leap year

```
2024 is a leap year as it is divisible only by 4 and it is not divisible by 100

#### Leap year : century year
```bash
$ go run . 2000
2000 is a leap year

```
2000  is a leap year as it is divisible by both 100 and 400

#### Not a Leap year : year not divisible by 4
```bash
$ go run . 1901
1901 is not a leap year

```

1900 is **not** a leap year as it is not divisible by both 100 and 400


#### Not a Leap year : century year
```bash
$ go run . 1900
1900 is not a leap year

```
1900  is **not** a leap year as it is not divisible by both 100 and 400




### Contributing

We welcome contributions to this project. Please fork the repository and submit a pull request with your improvements.

### License

This project is licensed under the MIT License - see the LICENSE.md file for details.

### Authors
This Project was created by [Wambita](https://github.com/Wambita/)

