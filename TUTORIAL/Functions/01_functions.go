package main 
import(
   "fmt"     
)
func message(msg string) (string) {
  fmt.Println(msg)
   return (msg + " viv-bhai")
}
func main(){
  newMessage := message("Hello")
  fmt.Println(newMessage)
}
/*Output:
Hello
Hello viv-bhai
*/