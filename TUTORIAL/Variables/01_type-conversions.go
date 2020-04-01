package main

import (
	"fmt"
)

func main() {

	var i float32 = 7.7
	fmt.Printf(" Value = %v  and Type = %T \n", i, i)

	var j int
	j = int(i)
	fmt.Printf("Value = %v and Type = %T \n", j, j)

	/*
		 var i int =7
		fmt.Printf(" Value = %v  and Type = %T \n",i,i)

		var j float32
		j = float32(i)
		fmt.Printf("Value = %v and Type = %T \n",j,j)

	*/

}

/*
//CONVERTING INTEGER TO STRING:
_______________________________
			package main

			import (
				"fmt"
				"strconv"
			)

			func main() {

				var i int =7
				fmt.Printf(" Value = %v  and Type = %T \n",i,i)

				var j string
				j = strconv.Itoa(i)
				fmt.Printf("Value = %v and Type = %T \n",j,j)

			}
_______________________________________________________________________

*/
