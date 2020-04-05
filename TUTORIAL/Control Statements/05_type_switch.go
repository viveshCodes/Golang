package main
import(
    "fmt"
)
 func main(){
     var i interface{}= 7
     switch i.(type) {
         case int:
         fmt.Println("It's integer.")
     
         case float64 :
         fmt.Println("It's float.")
                      
         case string :
         fmt.Println("It's string.")

         default:
         fmt.Println("Other Data Type.")
     }
       
     
 }