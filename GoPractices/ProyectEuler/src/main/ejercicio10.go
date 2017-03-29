/*
	Summation of primes
	Problem 10

	The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
	Find the sum of all the primes below two million.
	
	@Author Chris M. Perez
	@Date   3/28/2017
 */
package main
import(
   "fmt"
   "math"
)

func main(){
   MAX := 2000000
   fmt.Println(isPrime(MAX))
}

func isPrime(MAX int) [] int{
   array := make([]int,MAX,2000050)
   sum := 0
   
   LOOP:
      for i:=2;i<MAX;i++ {
	 N := int(math.Floor(math.Sqrt(float64(i))))
      
	 for j:=2;j<=N;j++ {
	    if i%j == 0 {
	       continue LOOP
	    }
	 }
	 sum += i
	 array = append(array,sum)
      }
   return array[len(array)-1:]
}
