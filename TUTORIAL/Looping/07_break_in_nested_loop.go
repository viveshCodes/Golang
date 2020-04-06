package main
import(
	"fmt"
)
func main(){
	for i :=1;i<=5;i++{
	       for j:=1;j<=5;j++{
	           if i*j==20{
	               break // closest loop is broken
	           } else{
	               fmt.Println(i*j)
	           }
	       }
	}
}