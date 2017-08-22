package main

import (
  "github.com/go-vgo/robotgo"
)

func main() {
  booleanFlag := false
  //Switch back to RS client
  robotgo.KeyTap("h", "command")

  //Hold down button to get rid of item

  for booleanFlag == false {
    robotgo.KeyToggle("1", "down", "shift")
    robotgo.SetKeyDelay(300)
    robotgo.KeyTap("left")
    robotgo.KeyTap("right")
  }
} 