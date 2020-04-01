package main
import(
	"fmt"
)
func main(){
	cityPopulation := make(map[string] int)
	cityPopulation = map[string] int {
		"Birgunj": 7781,
		"Ranchi": 1983,
		"Chenaai":2005,
	}
	// Add a new key-value to our map
	cityPopulation["Mumbai"] = 2013
	fmt.Println("The popultaion of Mumbai is"  , cityPopulation["Mumbai"])
	// Let's delte Mumbai from our map
	delete(cityPopulation ,"Mumbai")
	fmt.Println("The popultaion of mumbai is " , cityPopulation["Mumbai"])

	// To be clear if Mumbai really exists
	pop , ok := cityPopulation["Mumbai"]
	fmt.Printf("Pop = %v Ok = %v ",pop,ok)
}
/*OUTPUT
_________
The popultaion of Mumbai is 2013
The popultaion of mumbai is  0
Pop = 0 Ok = false 
*/