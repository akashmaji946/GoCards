package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello, let's check for the sites:")

	links := []string{
		"http://akashmaji.com",
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://golang.org",

		"http://stackoverflow.com",
		"http://rubbish.com",
		"http://garbage.com",
	}

	fmt.Println(links)

	ch := make(chan string)

	for i, link := range links {
		go checkLink(link, ch, i)
	}

	// // may be received in any order
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-ch)
	// 	fmt.Println(i)
	// }
	// //fmt.Println(<-ch) // only one is being shown

	// fmt.Println("Links checked :)")

	j := 0 // just for the syntax
	// for {

	// 	go checkLink(<-ch, ch, j)
	// }

	for link := range ch {
		// putting it here with cause main routine to sleep
		// time.Sleep(time.Second * 2)
		go func(l string) {
			time.Sleep(time.Second * 2)
			checkLink(l, ch, j)
		}(link)
	}
}

func checkLink(link string, ch chan string, i int) {
	//putting sleep here will disturb the use of checkLink()
	//time.Sleep(time.Second * 2)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Link (", link, ") might be down!", i)
		ch <- link
		return
	}
	fmt.Println("Link (", link, ") is working good!", i)
	ch <- link
	return
}
