package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	test := os.Args[2:]
	if len(test) == 0 {
		fmt.Println("No arguments provided or not enough arguments provided")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, stat.Size())
	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}
	url := strings.Replace(os.Args[2], "http://", "", -1)
	yoururl := b64.StdEncoding.EncodeToString([]byte("http://" + url))
	var EZ string = string(buffer)
	EZ = strings.Replace(EZ, "www.boomlings.com/database", url, -1)
	EZ = strings.Replace(EZ, "aHR0cDovL3d3dy5ib29tbGluZ3MuY29tL2RhdGFiYXNl", yoururl, -1)

	file.Close()
	err = os.Remove(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	file, err = os.Create(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	file.WriteString(EZ)
	file.Close()
	println("Made GDPS successfully!")
	println("Press enter to exit...")
	fmt.Scanln()
	os.Exit(0)
}
