package main

import (
	"fmt"
	"time"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	for _, port := range []Int{
	8010,
	8020,
	8030,
} {
		args := os.Args[0:]

	 
	    com := make(chan io.Reader)
	    defer close(com)
	  
	    for{
	    		conn, err := net.Dial("tcp", port)
				if err != nil {
					log.Fatal(err)	
				} com <- conn
				time <- com
				loc, err2 := io.Copy(os.Stdout, time)
			    if err2 != nil {
			        log.Fatal(err2)
			    }
		}
	}
}
