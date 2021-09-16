package main

//https://medium.com/@greenraccoon23/multi-thread-for-loops-easily-and-safely-in-go-a2e915302f8b
import (
	"fmt"
	"os/exec"
	"time"
)

func PingHost(url string, ch chan string) {
	if url == "" {
		ch <- ""
	}
	//add error handling
	cmd, _ := exec.Command("ping", "-c3", url).Output()
	ch <- string(cmd)
}

func main() {
	//runtime.GOMAXPROCS(4)
	ch := make(chan string)
	urls := [5]string{
		"google.com", "cisco.com", "ya.ru", "mail.ru", "golang.org",
	}
	start := time.Now()

	//for loop sequentially runs the routines
	//figure out the mistake here: how to prevent waiting in the for loop
	for i := range urls {
		go PingHost(urls[i], ch)
		fmt.Println(<-ch)
	}
	duration := time.Since(start)
	fmt.Println("time", duration)

}
