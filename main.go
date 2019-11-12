// Copyright (C) 2019 PeerNova, Inc.
//
// All rights reserved.
//
// PeerNova and Cuneiform are trademarks of PeerNova, Inc. References to
// third-party marks or brands are the property of their respective owners.
// No rights or licenses are granted, express or implied, unless set forth in
// a written agreement signed by PeerNova, Inc. You may not distribute,
// disseminate, copy, record, modify, enhance, supplement, create derivative
// works from, adapt, or translate any content contained herein except as
// otherwise expressly permitted pursuant to a written agreement signed by
// PeerNova, Inc.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

const (
	HOST = "H"
	PORT = "P"
	TYPE = "T"

	DefaultConnHost = "0.0.0.0"
	DefaultConnPort = "9090"
	DefaultConnType = "tcp"
)

func main() {
	connHost := flag.String(HOST, DefaultConnHost, "Connection Host")
	connPort := flag.String(PORT, DefaultConnPort, "Connection Port")
	connType := flag.String(TYPE, DefaultConnType, "Connection Type")

	flag.Parse()

	// Listen for incoming connections
	ln, err := net.Listen(*connType, *connHost+":"+*connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes
	defer ln.Close()

	fmt.Println("listening on " + *connHost + ":" + *connPort)

	// Listen for an incoming connection
	conn, _ := ln.Accept()

	// run loop forever (or until Ctrl-C)

	for {
		// listen for message to process ending in new line (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// convert to string
		logs := string(message)

		// output message received
		fmt.Print(logs)

		// Send a response back to person contacting us
		conn.Write([]byte(logs + "\n"))
	}
}
