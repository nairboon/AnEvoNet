package main

import (
	flag "github.com/ogier/pflag"
"fmt"
	"log"

)


type APICallChannel chan string

var port int

func main() {
	flag.IntVar(&port, "port", 9000, "the port to start an instance of anevonet")
	flag.Parse()

	log.Printf("staring daemon on %d\n", port)
// declare channels
//api_calls := make(APICallChannel, 5)

// read config file
// open database

// start evolver
// start external listener
// start connection manager
// start internal broker

	fmt.Printf("stopping anevonet daemon\n")
}
