/*_____________fallthrough_______________*/
package main
import(
    "fmt"
)
 func main(){
     number :=7
     switch {
         case number <=10:
                         fmt.Println("The number is less than or equal to 10 .")
                         fallthrough
         case number >=10 :
                         fmt.Println("The number is greater than or equal to 10 .")
                         fallthrough
         case number >5 && number < 10 :
                         fmt.Println("The number lies between 5 and 10 .")
         default:
                         fmt.Println("Number out of range.")
     }
       
     
 }

 /* OUTPUT :
The number is less than or equal to 10 .
The number lies between 5 and 10 .
 
 */