package main

import ("fmt")

func main() {
    var days = 0
    
    fmt.Print("How much days do you want to hire the car? <> ")
    fmt.Scanln(&days)
    
    if !(days <= 0) {
        if (days >= 1) && (days <= 5) {
            fmt.Print(days, "days price is: ", days * 70)
        } else if (days >= 6) && (days <= 10) {
            fmt.Print(days, "days price is: ", days * 60)
        } else if (days >= 11) && (days <= 20) {
            fmt.Print(days, "days price is: ", days * 50)
        } else {
            fmt.Print(days, "days price is: ", days * 30)
        }
    } else {
        fmt.Print("You must input a positive number.")
    }
}
