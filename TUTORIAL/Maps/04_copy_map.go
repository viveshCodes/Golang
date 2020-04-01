package main
import(
	"fmt"
)
func main(){
	cityPopulation := make(map[string] int)
	cityPopulation = map[string] int {
		"Birgunj": 7781,
		"Ranchi": 1983,
		"Chennai":2005,
	}
	
	sp :=cityPopulation  // We're not simply copying cityPopulation into sp .sp is reference to cityPopulation
	delete(sp,"Chennai")
	fmt.Println(sp)
	fmt.Println(cityPopulation)
	
	cityPopulation["Patna"] = 1000
	fmt.Println(sp)
	fmt.Println(cityPopulation)
	
	
	
	
}
/*Output:
map[Birgunj:7781 Ranchi:1983]
map[Birgunj:7781 Ranchi:1983]
map[Birgunj:7781 Patna:1000 Ranchi:1983]
map[Birgunj:7781 Patna:1000 Ranchi:1983]

*/
