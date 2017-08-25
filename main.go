package main

import (
  parking "./parking"
)

func main() {
  parking.Init()
  //Code for Command Shell Input
  if parking.Type_Of_Input==1{
      parking.StartCommandOperation()
    } else if parking.Type_Of_Input==2 {   //Code for file input

  }else{


  }

}
