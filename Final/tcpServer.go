package main

//import necessary packages
import (
	"bufio"   //get user input
	"fmt"     // standard i/o
	"net"     // send and receive
	"os"      // access arguments
	"strings" // string manupulation methods
)

func main() {
	// checks to make sure that the user povided a port to listen on
        if len(os.Args) == 1 {
                fmt.Println("Please provide a valid port number")
                return
        }

        // estabilishes a tcp connection listener on the specified port and catches any errors
        listener, err := net.Listen("tcp", ":" + os.Args[1])
        if err != nil {
                fmt.Println(err)
                return
        }

        // when the function returns (due to a client's closed connection or otherwise), stop listening
        defer listener.Close()

        // start accepting new connections to the server and check for errors
        connection, err := listener.Accept()
        if err != nil {
                fmt.Println(err)
                return
        }

        // main infinite loop to handle data transfer
        for {
                // receive data from the connection and check for errors
                data, err := bufio.NewReader(connection).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }

                // check for exit message to stop the script
                if strings.TrimSpace(string(data)) == "exit" || strings.TrimSpace(string(data)) == "quit" {
                        fmt.Println("Bye!")
                        return
                }

                // print out client's message
                fmt.Print("Client: ", string(data))

                // repeat client's message back to them
                connection.Write([]byte(data))
        }
}
