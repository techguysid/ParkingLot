package parking
import (
  "fmt"
)
var Type_Of_Input int
func Init()  {
  fmt.Println("*****************************Welcome to Parking Lot System*****************************")
  fmt.Println("--------------Developed By : Siddharth Udeniya--------------")
  fmt.Println("-------------------------------------------------------------------------")
  fmt.Println("How do you want to pass the input ? Type 1 for command shell | 2 for file")
  fmt.Scanf("%d", &Type_Of_Input)
}
