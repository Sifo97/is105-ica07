//Orginal kode hentet fra:
//https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/
package main
 
import (
    "fmt"
    "net"
	"../crypt"
	"log"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}
 
func main() {
key := []byte("a very very very very secret key") // must be 32 bytes
plaintext := []byte("Secret message m8")
chipertext, err := crypt.Encrypt(key,plaintext)
fmt.Printf("%0x\n",chipertext)
if err != nil {
log.Fatal(err)
}
    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:27015")
    CheckError(err)
 
    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)
 
    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)
 
    defer Conn.Close()	
		Conn.Write([]byte(chipertext))
        if err != nil {
            fmt.Println(err)
}
}