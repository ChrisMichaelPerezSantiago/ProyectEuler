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