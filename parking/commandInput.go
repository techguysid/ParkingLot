package parking
import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)
var cars = []Car{}
var current_empty_slot int = 1
var  total_slots int
func StartCommandOperation()  {
  fmt.Println("Enter commands-->")
  reader := bufio.NewReader(os.Stdin)
  for{
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)
    //fmt.Println(text);
    args := strings.Split(text, " ")
    switch args[0] {
    case "create_parking_lot":
      total_slots,_= strconv.Atoi(args[1])
      fmt.Println(total_slots)
    case "park":
      if current_empty_slot>total_slots{
        fmt.Println("Sorry, parking lot is full")
      }else{
      c := Car{
          Slot_No:current_empty_slot,
          Car_No:  args[1],
          Colour: args[2],
      }
      cars = append(cars,c)
      fmt.Println("Allocated slot number:",current_empty_slot);
      current_empty_slot++;
    }
    case "leave":
      i,_:=strconv.Atoi(args[1])
      j:=0;
      for _, c := range cars {
          if c.Slot_No==i {
            cars = append(cars[:j], cars[j+1:]...)
          }
          j++;
      }

    case "status":
      fmt.Println("Slot No.\t Registration No \t Colour")
      for _, c := range cars {
          fmt.Println(c.Slot_No,"\t\t", c.Car_No,"\t\t",c.Colour)
      }
    case "registration_numbers_for_cars_with_Colour":
      for _, c := range cars {
          if(c.Colour==args[1]){
            fmt.Printf(c.Car_No,"\t")
          }
      }
    default:
      fmt.Println("Incorrect Choice")
    }
  }
}
