package main

import ("fmt")

func main() {
  var empty_string string // That's an empty string.
  var full_string = "Woah, I'm full!"
  var integer1, integer2 = 10, 50
  no_var_here := "I am an string too!"
  
  // Well, there are 3 types of printing.
  fmt.Println("Text.") // If you know C#, that's the WriteLine. If you don't, it printed the thing and skipped to the next line.
  fmt.Print("Text.\n")   // If you know C#, that's the Writel. If you don't, it's actually the normal print.
  fmt.Printf("Text. But i can show you the type of things. %T", no_var_here)  // 
  
  fmt.Println("\n------" + empty_string + "\n" + full_string + "\n", integer1, "\n", integer2, "\n" + no_var_here)
}
