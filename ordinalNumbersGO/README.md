# Ordinal Values GO

## Overview

This program generates and displays a table of numbers from 1 to 100, with their corresponding ordinal suffixes. An ordinal suffix (e.g., "st", "nd", "rd", "th") is used to denote the position of a number in a sequence.

## Features

- Correctly handles special cases like 11th, 12th, and 13th.
- Outputs a formatted table of numbers with their ordinal suffixes.

## Code Explanation
### ordinal Function

The ordinal function takes an integer as input and returns a string representing the appropriate ordinal suffix.
Logic

**Special Cases:**
- The numbers 11, 12, and 13 are exceptions and always have the "th" suffix.

**General Cases:**
- If a number ends in 1 (but is not 11), it has the "st" suffix.
- If a number ends in 2 (but is not 12), it has the "nd" suffix.
- If a number ends in 3 (but is not 13), it has the "rd" suffix.
- All other numbers have the "th" suffix.

### main Function

The main function generates a table of numbers from 1 to 100 and prints them in a formatted way, five numbers per row, each with its ordinal suffix.


Sample Code

```Go
package main

import (
	"fmt"
)

func ordinal(v int) string {
	// exceptions
	if v == 11 || v == 12 || v == 13 {
		return "th"
	}

	v %= 10

	switch v {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

func main() {
	for c := 1; c <= 20; c++ {
		fmt.Printf("%d%s\t%d%s\t%d%s\t%d%s\t%d%s\n",
			c, ordinal(c),
			c+20, ordinal(c+20),
			c+40, ordinal(c+40),
			c+60, ordinal(c+60),
			c+80, ordinal(c+80),
		)
	}
}

```
## Compilation and Execution

To compile and run this program, use the following commands:

```bash

go run .
```
Output

The program will output a formatted table of numbers from 1 to 100, each with its ordinal suffix. For example:
```bash
  1st   21st   41st   61st   81st
  2nd   22nd   42nd   62nd   82nd
  3rd   23rd   43rd   63rd   83rd
  4th   24th   44th   64th   84th
  5th   25th   45th   65th   85th
  6th   26th   46th   66th   86th
  7th   27th   47th   67th   87th
  8th   28th   48th   68th   88th
  9th   29th   49th   69th   89th
 10th   30th   50th   70th   90th
 11th   31st   51st   71st   91st
 12th   32nd   52nd   72nd   92nd
 13th   33rd   53rd   73rd   93rd
 14th   34th   54th   74th   94th
 15th   35th   55th   75th   95th
 16th   36th   56th   76th   96th
 17th   37th   57th   77th   97th
 18th   38th   58th   78th   98th
 19th   39th   59th   79th   99th
 20th   40th   60th   80th  100th
```
## License

This project is licensed under the MIT License. See the LICENSE file for more details.
