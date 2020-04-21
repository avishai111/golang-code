package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fileName := "test.txt"
	numberOfData := 2
	var buf chan []byte
	var addr chan *net.UDPAddr
	conn, errlisten := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 5000, Zone: " "})
	if errlisten != nil {
		fmt.Println("eror in listen to adress in udp protocol")
	}
	_, err := os.Create(fileName) // create file
	if err != nil {
		fmt.Println("error in create file")
	}
	file, err := os.Open(fileName) // open file
	if err != nil {
		fmt.Println("error in open file")
	}
	for i := 0; i < numberOfData; i++ { // create go routine
		go readudp(conn, addr, buf) //and write the data that recive from go rouitne to  thefile
		v1, ok1 := <-addr
		close(addr)
		if ok1 != true {
			fmt.Println("error to get the adress to channel v1")
		}
		v2, ok2 := <-buf
		close(buf)
		if ok2 != true {
			fmt.Println("error to get the data to channel v2")
		}
		_, err := fmt.Fprintln(file, string(i)+string(v1.IP)+string(v1.Port)+string(v2))
		fmt.Println(string(i) + string(v1.IP) + string(v1.Port) + string(v2))
		if err != nil {
			fmt.Println("eror write to the file")
		}
		fmt.Println("print into the file")
	}
	file.Close()
}

func readudp(conn *net.UDPConn, addr chan *net.UDPAddr, buf chan []byte) { // goroutine that listen to udp and read from udp
	buf1 := make([]byte, 100)
	n, addr1, erread := conn.ReadFromUDP(buf1)
	if erread != nil {
		fmt.Println("eror in read from udp protocol")
	}
	addr2 := addr1
	buf2 := buf1[:n]
	addr <- addr2
	buf <- buf2
}

// write to file

/*func check_eror1(eror error) {
	if eror != nil {
		fmt.Println("ERROR")
	}
}

func check_eror2(eror bool) {
	if eror != true {
		fmt.Println("ERROR")
	}
}
*/
