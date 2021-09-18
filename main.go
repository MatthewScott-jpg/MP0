package main

//TODO: Add user input, store performance statistics in map, iterate through gomaxprocs possibilities

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func PingHost(out chan string, urls []string) {
	urlsLength := len(urls)

	fmt.Println("Running for loop...")
	for i := 0; i < urlsLength; i++ {
		go func(i int) {
			n := urls[i]
			cmd, ok := exec.Command("ping", "-c3", n).Output()
			if ok != nil {
				out <- n + " not found"
			} else {
				out <- string(cmd)
			}
		}(i)
	}
}

func main() {
	runtime.GOMAXPROCS(8)
	urls := []string{
		"google.com",
		"google.co.jp",
		"google.co.uk",
		"google.es",
		"google.ca",
		"google.de",
		"google.it",
		"google.fr",
		"google.com.au",
		"google.com.tw",
		"google.nl",
		"google.com.br",
		"google.com.tr",
		"google.be",
		"google.com.gr",
		"google.co.in",
		"google.com.mx",
		"google.dk",
		"google.com.ar",
		"google.ch",
		"google.cl",
		"google.at",
		"google.co.kr",
		"google.ie",
		"google.com.co",
		"google.pl",
		"google.pt",
		"google.com.pk",
	}

	out := make(chan string, len(urls))
	start := time.Now()

	go PingHost(out, urls)

	for range urls {
		fmt.Println(<-out)
	}

	duration := time.Since(start)
	fmt.Println("time", duration)
}
