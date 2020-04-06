package main 
import(
	"fmt"
)
func main(){
	for i ,j := 7,8 ; i<=7; i,j = i+1 ,j+1{
		fmt.Println(i,j)
	}
}
/*The following will throw an error :
for i ,j := 7,8 ; i<=7;  i++ ,j++{
		fmt.Println(i,j)
	}

*/