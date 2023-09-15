// Ping is a simple pinger using UDP instead of ICMP.
// The client sends a ping message to the provided host, and measures the RTT
// for a response.
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const (
	E_ConnRefused = "connection refused"
	E_Timeout     = "connection timed out"
)

type params struct {
	Count   int
	Timeout int
}

type elements struct {
	Latencies    []float64
	ResolvedHost string
	Failures     []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	return
}

func pingUDP() (err error) {

	// Begin measuring time
	c, err := net.Dial("udp", service)
	checkError(err)

	_, err = c.Write([]byte("PING!"))
	checkError(err)

	// Set a timeout of 1 second
	err = c.SetReadDeadline(time.Now().Add(1 * time.Second))
	checkError(err)

	defer c.Close()

	rb := make([]byte, 1000)

	if _, err := c.Read(rb); err != nil {
		if e := err.(*net.OpError).Timeout(); e {
			return fmt.Errorf(E_Timeout)
		}
		if strings.Contains(err.Error(), "connection refused") {
			return fmt.Errorf(E_ConnRefused)
		}

		return fmt.Errorf("Read error: %w", err.Error())
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
	os.Exit(1)

}
