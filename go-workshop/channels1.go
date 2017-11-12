package main

import (
    "fmt"
    "time"
)
// START OMIT
type Computer struct {
    Model string
    Price int
}

func (c *Computer) StartTimer(channel chan string, t time.Duration) {
    fmt.Println("Starting timer...")
    time.Sleep(t)
    channel <- "Time up!"
}

func main() {
    computer := Computer{Model: "Macbook",Price: 1000,}
    channel := make(chan string)
    go computer.StartTimer(channel, 3 * time.Second)
    
    select {
        case msg := <-channel:
            fmt.Println(msg)
    }
    fmt.Println("Exited")
}

// END OMIT