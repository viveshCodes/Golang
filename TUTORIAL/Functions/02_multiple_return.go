package main 
import(
   "fmt"
)
func swapNumbers(num1 , num2 int) (int , int){
   var temp int
   temp= num1 
   num1 = num2
   num2= temp
   return (num1) , (num2)
}
func main(){
  num1 , num2 := 7 ,8
  fmt.Printf("num1 = %v  and num2 = %v \n",num1 , num2)
  num1 , num2 = swapNumbers( num1 , num2)
  fmt.Printf("num1 = %v  and num2 = %v", num1 , num2)
}