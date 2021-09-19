package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var pings map[int]map[string]string
var times map[int]string

func UserInput() []string {
	var webs string
	fmt.Println("Input the urls that you would like to check, followed by spaces. To run the default hit enter.")
	fmt.Scanln(&webs)
	websArr := strings.Split(webs, " ")
	if(websArr[0] == "") {
		urls := []string{
			"google.com", "bc.edu", "facebook.com", "instagram.com", "amazon.com",
		}
		return urls
	}
	return websArr
}

func PingHost(out chan string, used chan string, urls []string) {
	urlsLength := len(urls)

	for i := 0; i < urlsLength; i++ {
		go func(i int) {
			n := urls[i]
			cmd, ok := exec.Command("ping", "-c3", n).Output()
			if ok != nil {
				used <- n
				out <- n + " not found"
			} else {
				used <- n
				out <- string(cmd)
			}
		}(i)
	}
}

func main() {
	pings = make(map[int]map[string]string)
	times = make(map[int]string)
	urls := UserInput()

	for i := 1; i < 8; i++ {
		runtime.GOMAXPROCS(i)
		pings[i] = make(map[string]string)
		out := make(chan string, len(urls))
		used := make(chan string, len(urls))
		start := time.Now()
		PingHost(out, used, urls)

		for range urls {

			v := <-out
			u := <-used

			average := regexp.MustCompile("min/avg/max/stddev = ([0-9./]+)")
			result := average.FindStringSubmatch(v)
			if len(result) > 0 {
				parts := strings.Split(result[1], "/")
				pings[i][u] = parts[1]
			} else {
				pings[i][u] = "Failed"
			}
		}
		duration := time.Since(start)
		times[i] = fmt.Sprintf("%f", duration.Seconds())
	}
	for key := 1; key < 8; key++ {
		element := times[key]
		fmt.Println("For", key, "processes, it took ", element, "Seconds")
		tmp := pings[key]
		for subkey, subelem := range tmp {
			fmt.Println("Average time for", subkey, "was", subelem, "ms")
		}
		fmt.Println("")
	}

}

/*
{
		"google.com", "google.co.jp", "google.co.uk", "google.es", "google.ca", "google.de", "google.it", "google.fr",
		"google.com.au", "google.com.tw", "google.nl", "google.com.br", "google.com.tr", "google.be", "google.com.gr",
		"google.co.in", "google.com.mx", "google.dk", "google.com.ar", "google.ch", "google.cl", "google.at",
		"google.co.kr", "google.ie", "google.com.co", "google.pl", "google.pt", "google.com.pk",
	}
*/
