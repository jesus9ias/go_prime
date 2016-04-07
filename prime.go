package main

import (
    "fmt"
    "math"
)

func main() {
    var i int
    fmt.Print("Enter a number: ")
    fmt.Scanf("%d\n", &i)
    num := int(i)
    isPrime := true
    switch {
      case num <= 0:
        fmt.Println(num, "not natural number")
      case num <= 3:
      	fmt.Println("Prime number")
      default:
        if math.Mod(float64(num), 2) == 0 {
  				fmt.Println(num, "not a Prime. There is not pair prime numbers (except 2)")
  				isPrime = false
  			} else {
          limite := math.Sqrt(float64(num))
  				for a := float64(3); a <= limite; a += 2 {
  					if math.Mod(float64(num), a) == 0 {
  						fmt.Println("Was found factor on", int(a))
  						isPrime = false
  					}
  				}
  				if isPrime {
  					fmt.Println("Number", num , "is a Prime")
  				} else {
  					fmt.Println("Number", num ,"is not a Prime")
  				}
  			}
    }
}
