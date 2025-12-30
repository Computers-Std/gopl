// Exercise 8.15: Failure of any client program to read data in a
// timely manner ultimately causes all clients to get stuck. Modify
// the broadcaster to skip a message rather than wait if a client
// writer is not ready to accept it. Alternatively, add buffering to
// each client’s out going message channel so that most messages are
// not dropped; the broadcaster should use a non-blocking send to this
// channel.

package main

// See [file: /home/ukiran/prog/gopl/ch8/exercises/ex-8.14/v3/main.go]
