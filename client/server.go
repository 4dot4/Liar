// main.go

package main

// import the package we need to use
import (
  "fmt"
  "log"
  "net"
  "net/http"
)

func server() {

  // set a HTTP request handle function for path /greeting and registrate it
    http.HandleFunc("/greeting", func (w http.ResponseWriter, 
        r *http.Request) {
  
        // when receive the request, print the greeting meassage
        fmt.Fprint(w, "Hello World")
  
      })

  // print out the server is going to start and show the time
  log.Println("Starting server....")

  // create server at localhost:8080 and using tcp as the network
  listener, err := net.Listen("tcp", ":8080")

  // if recieve error, record it and exit the program
  if err != nil {
    log.Fatal(err)
  }

  // setup HTTP connection for the listener of the server
  http.Serve(listener, nil)

}