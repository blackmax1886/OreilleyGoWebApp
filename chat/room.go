package main

type room struct {
	// forward is channel to keep message that is sent to other client
	forward chan []byte
}
