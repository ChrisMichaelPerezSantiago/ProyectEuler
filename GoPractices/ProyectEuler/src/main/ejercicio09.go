/*
	Special Pythagorean triplet
	Problem 9

	A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
	a2 + b2 = c2

	For example, 3_2 + 4_2 = 9 + 16 = 25 = 52.

	There exists exactly one Pythagorean triplet for which a + b + c = 1000.
	Find the product abc.
	
	@Author Chris M. Perez
	@Date   3/28/2017
 */
package main
import (
   "fmt"
)


func main(){
   MAX := 1000
   fmt.Println(PythagoreanTriplet(MAX))
}

func PythagoreanTriplet(MAX int) int {
   correctPythagoreanTriplet := false
   k := 0
   product := 1
   
   for i:=1;i<=MAX/3;i++{
      for j:=i;j<=MAX/2;j++{
	 k = MAX - i - j
	 
	 if PythagoreanTheorem(i , j , k) {
	    correctPythagoreanTriplet = true
	    product = Product(i , j , k)
	    break
	 }
      }
      if correctPythagoreanTriplet {
	 break
      }
   }
   return product
}

func PythagoreanTheorem(a , b , c int) bool {
   return (a * a) + (b * b) == (c * c)
}

func Product(a , b , c int) int {
   return a * b * c
}