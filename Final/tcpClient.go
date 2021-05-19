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
	// checks to make sure that the user povided a server to connect to
        if len(os.Args) == 1 {
                fmt.Println("Please provide a valid host and port as two separate arguments")
                return
        }
	
	// estabilishes a tcp connection to the specified server and catches any errors
        connection, error := net.Dial("tcp", os.Args[1] + ":" + os.Args[2])
        if error != nil {
                fmt.Println(error)
                return
        }

	// main infinite loop to handle data transfer
        for {
		// use bufio to get user input properly buffered
                reader := bufio.NewReader(os.Stdin)
                fmt.Print(">> ")
                data, _ := reader.ReadString('\n')

                // send data to the connection using standard input/output methods
                fmt.Fprintf(connection, data+"\n")

                // receive data from the connection and print it out
                message, _ := bufio.NewReader(connection).ReadString('\n')
                fmt.Print("Server: " + message)

                // check for exit message to stop the script
                if strings.TrimSpace(string(data)) == "exit" || strings.TrimSpace(string(data)) == "quit" {
                        fmt.Println("Bye!")
                        return
                }
        }
}
