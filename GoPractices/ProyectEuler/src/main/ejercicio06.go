/*
      Sum square difference
      Problem 6
      
      The sum of the squares of the first ten natural numbers is,
      12 + 22 + ... + 102 = 385
      
      The square of the sum of the first ten natural numbers is,
      (1 + 2 + ... + 10)2 = 552 = 3025
      
      Hence the difference between the sum of the squares of the first ten natural
      numbers and the square of the sum is 3025 − 385 = 2640.
      
      Find the difference between the sum of the squares of the first one hundred
      natural numbers and the square of the sum.

      @Author Chris M. Perez
      @Date   3/28/2017
*/
package main
import(
   "fmt"
)


func main(){
   MAX := 100
   fmt.Println(SumSquareDifference(MAX))

}

func SumSquareDifference(MAX int) int{
   sumSquare := 0
   squareSum := 0
   
   for i:=1;i<=MAX;i++{
      sumSquare += i*i
   }
   
   for j:=1;j<=MAX;j++{
      squareSum += j
   }
   return (squareSum * squareSum) - sumSquare
}