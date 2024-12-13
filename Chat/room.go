package main

type room struct {
    forward chan []byte
    join chan *client
    leave chan *client
    client map[*client]bool
}
