package main

import (
        "fmt"
        "time"
)

func say(s string) {
        for i := 0; i < 5; i++ {
                time.Sleep(100 * time.Millisecond)
                fmt.Println(s)
        }
}

func logo() {
        welcome0 := " ________   ________   ________   _________   ___   ________    _________   ________"
        welcome1:= "|\\   __  \\ |\\   __  \\ |\\   ____\\ |\\___   ___\\|\\  \\ |\\   ___  \\ |\\___   ___\\|\\   __  \\"
        welcome2 := "\\ \\  \\|\\  \\\\ \\  \\|\\  \\\\ \\  \\___|_\\|___ \\  \\_|\\ \\  \\\\ \\  \\\\ \\  \\\\|___ \\  \\_|\\ \\  \\|\\  \\"
        welcome3 := " \\ \\   ____\\\\ \\  \\\\\\  \\\\ \\_____  \\    \\ \\  \\  \\ \\  \\\\ \\  \\\\ \\  \\    \\ \\  \\  \\ \\  \\\\\\  \\"
        welcome4 := "  \\ \\  \\___| \\ \\  \\\\\\  \\\\|____|\\  \\    \\ \\  \\  \\ \\  \\\\ \\  \\\\ \\  \\    \\ \\  \\  \\ \\  \\\\\\  \\"
        welcome5 := "   \\ \\__\\     \\ \\_______\\ ____\\_\\  \\    \\ \\__\\  \\ \\__\\\\ \\__\\\\ \\__\\    \\ \\__\\  \\ \\_______\\"
        welcome6 := "    \\|__|      \\|_______||\\_________\\    \\|__|   \\|__| \\|__| \\|__|     \\|__|   \\|_______|"
        welcome7 := "                         \\|_________|                                                    "
        fmt.Println(welcome0)
        fmt.Println(welcome1)
        fmt.Println(welcome2)
        fmt.Println(welcome3)
        fmt.Println(welcome4)
        fmt.Println(welcome5)
        fmt.Println(welcome6)
        fmt.Println(welcome7)
}

func main() {
        
        
		
}