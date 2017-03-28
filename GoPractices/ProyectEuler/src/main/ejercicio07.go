/*
	10001st prime
	Problem 7

	By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
	we can see that the 6th prime is 13. What is the 10001st prime number?
	
	@Author Chris M. Perez
	@Date   3/28/2017
 */
package main
import(
   "fmt"
   "math"
)


func main(){
   MAX := 10001
   fmt.Println(PrimeNumbers(MAX))
}

func PrimeNumbers(MAX int) int{
   var x int
   var counter int
   
   
   for counter != MAX {
      x++
      if IsPrime(x) { counter++ }
   }
   return x
}

func IsPrime(value int) bool {
   if value < 2 { return false }
   if value == 2 { return true }
   
   x := int(math.Sqrt(float64(value)))
   
   for i:=2;i<=x;i++{
      if value % i == 0 {return false}
   }
   return true
}