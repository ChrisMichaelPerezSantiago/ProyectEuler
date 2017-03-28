/*
   	Smallest multiple
	Problem 5

	2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
	What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
	
	@Author Chris M. Perez
	@Date   3/28/2017
 */

package main
import(
   "fmt"
)


func main(){
   MAX := 20
   SmallestMultiple(MAX)
   
}

func SmallestMultiple(MAX int){
   actualValue := 20
   
   for ;; {
      flag := true
      for i := 1; i <= MAX; i++ {
	 flag = true
  
	 if actualValue % i != 0 {
	    flag = false
	    break
	 }
      }
      if flag {
	 fmt.Println(actualValue)
	 break
      }
      actualValue++
   }
}