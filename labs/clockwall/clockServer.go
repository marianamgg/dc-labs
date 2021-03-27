// Clock Server is a concurrent TCP server that periodically writes the time.
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

	var location string
	location = os.Getenv("TZ")

	listener, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(port)
    for {
        conn, err := listener.Accept()
        log.Print(conn)
        if err != nil {
            log.Print(err) // e.g., connection aborted
            continue
        }
        go handleConn(conn, location) // handle one connection at a time
    }
}

}

func handleConn(c net.Conn, zone string) {
	loc, err := time.LoadLocation(zone)
    defer c.Close()

    for {
    	 t, err := TimeIn(time.Now(), loc)

    	 if err == nil {
			fmt.Println(t.Location(), t.Format("15:04\n"))
		} else {
			fmt.Println(name, "<time unknown>")
			}

        time.Sleep(1 * time.Second)
    }
}
