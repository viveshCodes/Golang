package main

import (
    "fmt"    
)

func main(){
	var cityPopulation =make(map[string] int )  // cityPopulation := make(make[string]int)
    cityPopulation =map[string]int{
        "Birgunj": 7781,
        "Delhi": 8871,
        "Chennai":9980,
    }
    fmt.Println("Population of each city is :", cityPopulation)
}
/*
Output:
_______
Population of each city is : map[Birgunj:7781 Chennai:9980 Delhi:8871]

*/