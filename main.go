package main

//https://medium.com/@greenraccoon23/multi-thread-for-loops-easily-and-safely-in-go-a2e915302f8b
import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func PingHost(out chan string, urls []string) {
	//runtime.GOMAXPROCS(6)
	urlsLength := len(urls)

	var wg sync.WaitGroup
	wg.Add(urlsLength)

	fmt.Println("Running for loop...")
	for i := 0; i < urlsLength; i++ {
		go func(i int) {
			defer wg.Done()
			n := urls[i]
			cmd, ok := exec.Command("ping", "-c3", n).Output()
			if ok != nil {
				out <- n + " not found"
			} else {
				out <- string(cmd)
			}
		}(i)
	}
	wg.Wait()
}

func main() {
	runtime.GOMAXPROCS(4)
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

	PingHost(out, urls)

	for range urls {
		fmt.Println(<-out)
	}

	duration := time.Since(start)
	fmt.Println("time", duration)

}

/*
package main

//https://medium.com/@greenraccoon23/multi-thread-for-loops-easily-and-safely-in-go-a2e915302f8b
import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

func PingHost(out chan string, urls []string){
	//runtime.GOMAXPROCS(6)
	urlsLength := len(urls)

	var wg sync.WaitGroup
	wg.Add(urlsLength)

	fmt.Println("Running for loop...")
	for i := 0; i < urlsLength; i++ {
		go func(i int) {
			defer wg.Done()
			n := urls[i]
			cmd, ok := exec.Command("ping", "-c3", n).Output()
			if ok != nil {
				out <- n + " not found"
			}else {
				out <- string(cmd)
			}
		}(i)
	}
	wg.Wait()
}

func main() {
	urls := []string{
		"https://www.Google.com",
		"https://www.google.co.jp",
		"https://www.google.co.uk",
		"https://www.google.es",
		"https://www.google.ca",
		"https://www.google.de",
		"https://www.google.it",
		"https://www.google.fr",
		"https://www.google.com.au",
		"https://www.google.com.tw",
		"https://www.google.nl",
		"https://www.google.com.br",
		"https://www.google.com.tr",
		"https://www.google.be",
		"https://www.google.com.gr",
		"https://www.google.co.in",
		"https://www.google.com.mx",
		"https://www.google.dk",
		"https://www.google.com.ar",
		"https://www.google.ch",
		"https://www.google.cl",
		"https://www.google.at",
		"https://www.google.co.kr",
		"https://www.google.ie",
		"https://www.google.com.co",
		"https://www.google.pl",
		"https://www.google.pt",
		"google.com.pk",
	}

	out := make(chan string, len(urls))
	start := time.Now()

	PingHost(out, urls)

	for range urls{
		fmt.Println(<-out)
	}

	duration := time.Since(start)
	fmt.Println("time", duration)

}

*/
