package parking
import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "log"
  "strconv"
)
var Is_Slot_Empty bool=false
func StartFileOperation()  {
  fmt.Println("Enter file path-->")
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  // convert CRLF to LF
  text = strings.Replace(text, "\n", "", -1)
  if file, err := os.Open(text); err == nil {
  // make sure it gets closed
  defer file.Close()
  // create a new scanner and read the file line by line
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      //fmt.Println(scanner.Text())
      args:= strings.Split(scanner.Text(), " ")
      //fmt.Println(args[0]);
      switch args[0] {
      case "create_parking_lot":
        total_slots,_= strconv.Atoi(args[1])
        //fmt.Println(total_slots)
      case "park":
        if len(cars)==total_slots{
          fmt.Println("Sorry, parking lot is full")
        }else{
        i:=1
        if Is_Slot_Empty {

        for i <= total_slots {
            if i!=cars[i-1].Slot_No{
                c := Car{
                    Slot_No:i,
                    Car_No:  args[1],
                    Colour: args[2],
                }
                cars = append(cars[:i-1], append([]Car{c}, cars[i-1:]...)...)
                fmt.Println("Allocated slot number:",i);
                Is_Slot_Empty=!Is_Slot_Empty
                break
            }
            i++;
          }

        }else{
        //a = append(a[:i], append([]T{x}, a[i:]...)...)
        c := Car{
            Slot_No:current_empty_slot,
            Car_No:  args[1],
            Colour: args[2],
        }
        cars = append(cars,c)
        fmt.Println("Allocated slot number:",current_empty_slot);
        current_empty_slot++;
      }
      }
      case "leave":
        i,_:=strconv.Atoi(args[1])
        j:=0;
        for _, c := range cars {
            if c.Slot_No==i {
              cars = append(cars[:j], cars[j+1:]...)
              Is_Slot_Empty=!Is_Slot_Empty
              fmt.Println("Slot No ",i," is free")
            }
            j++;
        }

      case "status":
        fmt.Println("Slot No.\t Registration No \t Colour")
        for _, c := range cars {
            fmt.Println(c.Slot_No,"\t\t", c.Car_No,"\t\t",c.Colour)
        }
      case "registration_numbers_for_cars_with_colour":
        for _, c := range cars {
            if(c.Colour==args[1]){
              fmt.Printf("%s\t",c.Car_No)
            }
        }
        fmt.Printf("\n")
      case "slot_numbers_for_cars_with_colour":
        for _, c := range cars {
            if(c.Colour==args[1]){
              fmt.Printf("%s\t",strconv.Itoa(c.Slot_No))
            }
        }
        fmt.Printf("\n")
      case "slot_number_for_registration_number":
        for _, c := range cars {
            if(c.Car_No==args[1]){
              fmt.Println(c.Slot_No)
            }
        }
      default:
        fmt.Println("Incorrect Choice")
      }

  }

  // check for errors
  if err = scanner.Err(); err != nil {
    log.Fatal(err)
  }

} else {
  log.Fatal(err)
}

}
