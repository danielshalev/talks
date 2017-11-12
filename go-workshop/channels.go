// Use make() to create a channel
var channel := make(chan int)

// Push to channel using '<-'
channel <- 5