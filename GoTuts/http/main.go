package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type console struct{}

func main() {
	resp, err := http.Get("http://www.google.com/")

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	// fmt.Println(resp) //this gives response (ststus, ststus code, body, and many)
	// fmt.Println()

	//this is one way of doing things
	// bs := make([]byte, 100000)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//this is another way of doing things
	//io.Copy(os.Stdout, resp.Body)

	//our own type that implements Writer interface

	c := console{}
	io.Copy(c, resp.Body)
}

func (console) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes written:", len(bs))
	return len(bs), nil
}
