package main

import (
    "fmt"
    "math"
)

func main() {
    var i int
    fmt.Print("Enter a number: ")
    num, _ := fmt.Scanf("%d\n", &i)

    isPrime := true

    switch true {
      case num <= 0:
        fmt.Println("No es un numero natural")
      case num <= 3:
      	fmt.Println("Numero primo")
      default:
        if math.Mod(float64(num), 2) == 0 {
  				fmt.Println("No es numero primo. No hay numeros pares primos (excepto el '2')")
  				isPrime = false
  			} else {
          limite := math.Sqrt(float64(num))

  				for a := float64(3); a <= limite; a += 2 {
  					if math.Mod(float64(num), a) == 0 {
  						fmt.Println("Se encontro un factor: %d", int(a))
  						isPrime = false
  					}
  				}

  				if isPrime {
  					fmt.Println("Conclusion: el numero %d es primo", num)
  				} else {
  					fmt.Println("Conclusion: el numero %d no es primo", num)
  				}
  			}
    }
}
