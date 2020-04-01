package main

import (
    "fmt"    
)

func main(){
    cityPopulation :=map[string]int{
        "Birgunj": 7781,
        "Delhi": 8871,
        "Chennai":9980,   // Even at thi line we need comma
    }
    fmt.Println("Population of each city is :", cityPopulation)
}
/*
Output:
_______
Population of each city is : map[Birgunj:7781 Chennai:9980 Delhi:8871]

*/