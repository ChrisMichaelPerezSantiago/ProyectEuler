/*
  	Simple Array Sum

  	Given an array of  integers, can you find the sum of its elements?

	Input Format:
	The first line contains an integer, , denoting the size of the array.
	The second line contains  space-separated integers representing the array's elements.

	Output Format:
	Print the sum of the array's elements as a single integer.
*/

package main
import "fmt"


func main(){
   var array [1000]int
   MAX := 0
   sum := 0
   
   
   fmt.Scan(&MAX)
   
   for i:=0;i<MAX;i++ {
      fmt.Scan(&array[i])
      sum += array[i]
   }
   
   fmt.Println(sum)
}


