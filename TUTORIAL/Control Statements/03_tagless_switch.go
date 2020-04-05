package main
import(
    "fmt"
)
 func main(){
     worldCupYear :=2011
     switch{
       case worldCupYear == 2011 :
       fmt.Println("MS Dhoni won ICC trophy for India.")
       case worldCupYear==1983 :
       fmt.Println("Kapil Dev won ICC trophy for India.")
       default :
       fmt.Println("Broken Hearts.")
       
     }
     
 }