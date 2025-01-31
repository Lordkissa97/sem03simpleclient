package main

import (
	"github.com/Lordkissa97/is105sem03/mycrypt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.3:8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("os.Args[1] = ", os.Args[1])

	_, err = conn.Write([]byte(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	response := string(buf[:n])
	log.Printf("reply from proxy: %s", response)

	kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.AlfSem03, 4)
	log.Println("Kryptert melding: ", string(kryptertMelding))
	_, err = conn.Write([]byte(string(kryptertMelding)))
}
