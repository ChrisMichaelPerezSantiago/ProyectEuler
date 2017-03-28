/*
	Largest palindrome product
	Problem 4

	A palindromic number reads the same both ways. The largest palindrome made from the product
	of two 2-digit numbers is 9009 = 91 × 99.Find the largest palindrome made from the product
	of two 3-digit numbers.

	@Author Chris M. Perez
	@Date   3/27/2017
*/

package main
import (
   "fmt"
   "strconv"
)


func main(){
   largest , actual , product := 0 , 0 , 0
   
   for i:=100;i <= 999;i++ {
      for j := 100; j <= 999; j++ {
	 product = i * j
	 
	 if isPalindrome(strconv.Itoa(product)){
	    actual = product
	    
	    if isLargest(actual , largest){
	       largest = actual
	    }
	 }
      }
   }
   fmt.Println(largest)
}

func isPalindrome(value string) bool{
   array:= [] rune(value)
   
   for i:=range array{
      if array[i] != array[len(array)-i-1] {
	 return false
      }
   }
   return true
}

func isLargest(a , b int) bool {
   return a > b
}
