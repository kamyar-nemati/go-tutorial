package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
)

type client struct {
	conn net.Conn        //the connection to the client
	msg  bytes.Buffer    //the message client sends
	wg   *sync.WaitGroup //the reference to the goRoutine collection
}

func handleIntSig(ln *net.Listener, _running *bool) {
	/*
	 * Create a channel so that it can
	 * deliver the thrown signal.
	 */
	sigCh := make(chan os.Signal, 1)

	/*
	 * Register the channel to deliver
	 * the Interrupt signal.
	 */
	signal.Notify(sigCh, os.Interrupt)

	/*
	 * As the channel has capacity of one(1),
	 * therefore it is a blocking approach.
	 *
	 * Hence, it'll wait until a signal is caught.
	 */
	for sig := range sigCh {
		//print out the caught signal
		fmt.Printf("Signal caught: %s\n", sig)

		//stop the server
		*_running = false

		//cancel the socket that is waiting for a connection
		(*ln).Close()
	}
}

func checkError(err *error) bool {
	//check there's an error
	if (*err) != nil {
		//print out the error
		fmt.Printf("Error: %s\n", (*err).Error())
		return true
	}
	return false
}

func readPackets(cl *client, ch chan string) {
	//the tag that ends the data transfer
	EOT := "<~EOT~>"

	/*
	 * Create a buffer to hold incoming
	 * data from the client.
	 */
	buff := make([]byte, 1024)

	//keep reading
	for {
		n, err := cl.conn.Read(buff)

		//ignore if nothing is read
		if n == 0 {
			continue
		}

		//ignore if failed to read
		if checkError(&err) {
			continue
		}

		/*
		 * Transform the byte-data into string
		 * and send it to the channel.
		 */
		packet := string(buff[:n])
		ch <- packet

		//stop reading if the tag received
		if strings.Contains(packet, EOT) {
			break
		}
	}

	//close the channel
	close(ch)
}

func handleClient(cl *client) {
	fmt.Printf("Connection established: %s\n",
		cl.conn.RemoteAddr().String())

	/*
	 * Create a channel so that this function
	 * can communicate to the function that is
	 * responsible to read packets from client.
	 */
	ch := make(chan string, 1)

	//read incoming packets
	go readPackets(cl, ch)

	//observe the channel
	for {
		message, open := <-ch

		//check if the channel is still open
		if !open {
			break
		}

		//accumulate the read message
		cl.msg.WriteString(message)
	}

	//prepare some response for the client
	response := []byte(cl.msg.String())

	//write the response to the client
	_, err := cl.conn.Write(response)

	checkError(&err)

	//terminate the client's connection
	err = cl.conn.Close()

	checkError(&err)

	fmt.Printf("Connection closed: %s\n",
		cl.conn.RemoteAddr().String())

	//set the goRoutine as finished
	cl.wg.Done()
}

func main() {
	fmt.Println("Starting the socket server...")

	/*
	 * The flag that stops the socket server from
	 * accepting any other incoming connections.
	 */
	_running := true

	//server's configuration
	protocol := "tcp"
	endPoint := "127.0.0.1:9050"

	//listen to the defined endpoint
	ln, err := net.Listen(protocol, endPoint)

	//die if failed to listen
	if checkError(&err) {
		return
	}

	/*
	 * Launch a goRoutine to handle the Interrupt
	 * signal.
	 *
	 * The server would refuse to accept further
	 * connections once the signal is caught.
	 *
	 * However, the server will wait until all
	 * ongoing connections are handled before
	 * it terminates itself.
	 */
	go handleIntSig(&ln, &_running)

	fmt.Println("Bound on:", ln.Addr().String())
	fmt.Println("Awaiting connections...")

	//keep track of running goRoutines
	var wg sync.WaitGroup

	//begin accepting connections
	for {
		conn, err := ln.Accept()

		//check for failure on connection acception
		if !checkError(&err) {
			//make a client object
			cl := new(client)
			cl.conn = conn
			cl.wg = &wg

			//launch a goRoutine to handle the connected client
			go handleClient(cl)

			//add one into the collection of goRoutines
			wg.Add(1)
		}

		//check if the server is still running
		if !_running {
			fmt.Println("Socket server stopped")
			break
		}
	}

	fmt.Println("Shutting down socket server...")

	//wait for all clients to be handled
	wg.Wait()

	fmt.Println("Socket server terminated")
}
