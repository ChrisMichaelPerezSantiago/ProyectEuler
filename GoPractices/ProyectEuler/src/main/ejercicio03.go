/*
	Largest prime factor
	Problem 3

	The prime factors of 13195 are 5, 7, 13 and 29.
	What is the largest prime factor of the number 600851475143 ?
        
        @Author  Chris M. Perez
        @Date    3/27/2017
*/

package main
import(
   "fmt"
)

func main(){
   const MAX = int64(600851475143)
   fmt.Println(LargestPrimeFactorization(MAX))
}

func LargestPrimeFactorization(MAX int64) int64 {
   var i int64
   
   for i = 2;i <= MAX;i++ {
      if MAX % i == 0 {
	 MAX /= i
	 i--
      }
   }
   return i
}
