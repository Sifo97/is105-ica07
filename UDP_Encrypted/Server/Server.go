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
	key := []byte("a very very very very secret key") // must be 32 bytes
    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()
 
   var buf []byte = make([]byte, 1500)
 
    for {
		chipmsg,addr,err := ServerConn.ReadFromUDP(buf) // chip message
		fmt.Println("Received(ENCRYPTED):",string(buf[0:chipmsg]), " from ",addr)
		fmt.Println("Attempting to decrypt...")
		plaintext, err := crypt.Decrypt(key, buf[0:chipmsg])
		fmt.Println("Decrypted message: ", string(plaintext))
        if err != nil {
            fmt.Println("Error: ",err)
        } 
    }
}