A simple go library to make taking input/using scanners easier.

How to use:
NewScanner() will return a new scanner.
scanner.ReadLine() will read a line of user input from the console.
### Example:
```
package main

import (
    "fmt"
    "github.com/junglehornet/goscan"
)

func main() {
s := goscan.NewScanner()
input := s.ReadLine() // The program will wait for user input then store it in the variable 'input'.
fmt.Println(input) // The program will then print the user's input in the console:
}
```
