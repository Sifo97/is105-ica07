//Orginal kode hentet fra:
//https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/
package main
 
import (
    "fmt"
    "net"
    "os"
	"../crypt"
)
 
/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}
 
func main() {
    /* Lets prepare a address at any address at port 10001*/   
	
    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:27015")
    CheckError(err)
 
    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()
 
    buf := make([]byte, 1024)
 
    for {
		key,addr,err := ServerConn.ReadFromUDP(buf) // server key
		chipmsg,addr,err := ServerConn.ReadFromUDP(buf) // chip message
        //n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Received ",string(buf[0:key]), " from ",addr)
		fmt.Println("Received ",string(buf[0:chipmsg]), " from ",addr)
		plaintext, err := crypt.Decrypt(key1,chipmsg)
		fmt.Println(key)
        if err != nil {
            fmt.Println("Error: ",err)
        } 
    }
}