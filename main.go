package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

var pings map[int]map[string]string
var times map[int]string

/*
UserInput Take input urls, to be used in Main(). If no input uses default URL's
*/
func UserInput() []string {
	var webs string
	fmt.Println("Input the urls that you would like to check, followed by spaces. To run the default hit enter.")
	val, _ := fmt.Scanln(&webs)
	fmt.Println(val, "inputs")
	// Split string input into separate string links in an arr
	websArr := strings.Split(webs, " ")
	if websArr[0] == "" {
		urls := []string{
			"google.com", "google.co.jp", "google.co.uk", "google.es", "google.ca", "google.de", "google.it", "google.fr",
			"google.com.au", "google.com.tw", "google.nl", "google.com.br", "google.com.tr", "google.be", "google.com.gr",
			"google.co.in", "google.com.mx", "google.dk", "google.com.ar", "google.ch", "google.cl", "google.at",
			"google.co.kr", "google.ie", "google.com.co", "google.pl", "google.pt", "google.com.pk",
		}
		return urls
	}
	return websArr
}

/*
PingHost uses Go Routines in loop to Ping urls in parallel. Each loop uses GO channels
'Out' and 'Used' to store values from each ping
*/
func PingHost(out chan string, used chan string, urls []string) {
	urlsLength := len(urls)

	var wg sync.WaitGroup //waitgroups allow for goroutines to interact easier, and to ensure all goroutines in PingHost terminate

	for i := 0; i < urlsLength; i++ {
		wg.Add(1)
		go func(i int) {
			//defer wg.Done()
			n := urls[i]
			//Executes Ping with five threads ('c5'), error handles, and assigns output to out channel
			cmd, ok := exec.Command("ping", "-c5", n).Output()
			if ok != nil {
				used <- n
				out <- n + " not found"
			} else {
				used <- n
				out <- string(cmd)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
main Handles time of execution, and calls
*/
func main() {
	pings = make(map[int]map[string]string)
	times = make(map[int]string)
	urls := UserInput()

	//Run processes. Starting with 1 and ending at 7.
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

			//Turns the cmd line results from Ping into strings in result.
			average := regexp.MustCompile("min/avg/max/stddev = ([0-9./]+)")
			result := average.FindStringSubmatch(v)
			if len(result) > 0 {
				//Splits up result into a splice, placing the average time for the ping in
				//position 1 of the splice.
				parts := strings.Split(result[1], "/")
				//Takes the average value from parts[1] and puts it in the Pings Map for this specific
				//process run i
				pings[i][u] = parts[1]
			} else {
				pings[i][u] = "Failed"
			}
		}
		duration := time.Since(start)
		times[i] = fmt.Sprintf("%f", duration.Seconds())
	}
	//Formats value's stored in Ping and Time Maps
	for key := 1; key < 8; key++ {
		element := times[key]
		fmt.Println("For", key, "processes, it took ", element, "Seconds")
		tmp := pings[key]
		for subkey, subelem := range tmp {
			fmt.Println("Average time for", subkey, "was", subelem, "ms")
		}
		fmt.Println("")
	}
	fmt.Println(times)
}
