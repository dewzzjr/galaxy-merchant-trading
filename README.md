# Galaxy Merchant Trading Guide

## Problem Description

You decided to give up on earth after the latest financial collapse left 99.99% of the earth's population with 0.01% of the wealth. Luckily, with the scant sum of money that is left in your account, you are able to afford to rent a spaceship, leave earth, and fly all over the galaxy to sell common metals and dirt (which apparently is worth a lot). Buying and selling over the galaxy requires you to convert numbers and units, and you decided to write a program to help you.The numbers used for intergalactic transactions follows similar convention to the roman numerals and you have painstakingly collected the appropriate translation between them. Roman numerals are based on seven symbols:

**Symbol Value**
 - `I 1`
 - `V 5`
 - `X 10`
 - `L 50`
 - `C 100`
 - `D 500`
 - `M 1,000`

Numbers are formed by combining symbols together and adding the values. For example, MMVI is 1000 + 1000 + 5 + 1 = 2006. Generally, symbols are placed in order of value, starting with the largest values. When smaller values precede larger values, the smaller values are subtracted from the larger values, and the result is added to the total. For example MCMXLIV = 1000 + (1000 − 100) + (50 − 10) + (5 − 1) = 1944.

The symbols "I", "X", "C", and "M" can be repeated three times in succession, but no more. (They may appear four times if the third and fourth are separated by a smaller value, such as XXXIX.) "D", "L", and "V" can never be repeated.

"I" can be subtracted from "V" and "X" only. "X" can be subtracted from "L" and "C" only. "C" can be subtracted from "D" and "M" only. "V", "L", and "D" can never be subtracted.

Only one small-value symbol may be subtracted from any large-value symbol.

A number written in Arabic numerals can be broken into digits. For example, 1903 is composed of 1, 9, 0, and 3. To write the Roman numeral, each of the non-zero digits should be treated separately. In the above example, 1,000 = M, 900 = CM, and 3 = III. Therefore, 1903 = MCMIII.

**Source:** [Wikipedia](https://en.wikipedia.org/wiki/Roman_numerals)

## System Design Solution
- Program implemented as simple Command Line Interface Application
- There are two usecases:
    - **Query**: to process input text into action and input data (Roman numerals symbols, Foreign words, Credits, and Unit types of common metals).
    - **Translator**: there are several functionalities:
        - create dictionary from foreign words to roman numerals symbols, then translate them
        - create price list of common metals with credits
        - calculate the total credits of a given unit of metal
- Package library to calculate roman numerals to numeric decimal

## Requirements

- GoLang 1.18
- make (recommended)
- Visual Studio Code, with GoLang and Delve plugins (recommended)

## How to Run
- run `make all` on your terminal / command line
- if you don't have `make`, you can run
```
go mod vendor -v
go build -o galaxy-merchant-trading cmd/main.go
```
- if you don't have `Visual Studio Code`, you can run `./galaxy-merchant-trading run`
- to use input from file you can try using `./galaxy-merchant-trading run -f {filename}`
- to get more information you can run `./galaxy-merchant-trading help run`

## Testing
We have add complete unit test that can be run on github actions when there is changes on main branch. These test cases can also run on your local machine.

### How to test on local
- run `make all` on your terminal / command line
- if you don't have `make`, you can run
```
go mod vendor -v
go vet ./...
go test -race -cover -v ./...

```