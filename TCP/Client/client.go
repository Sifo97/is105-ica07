//Orginal kode hentet fra:
// https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "82.164.223.240:27015")
  for { 
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}