// Use select to wait (block) for a single message
select {
case i <- channel:
  fmt.Printf("Got message %d", i)
}

// Watch for multiple messages using range
for i := range channel {
  fmt.Printf("Got message %d", i)
}