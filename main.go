package main

import (
"fmt"
"os/exec"
)

func main() {
	for n := range PingHost("google.com", "cisco.com", "ya.ru", "mail.ru", "golang.org") {
		fmt.Println(n)

	}
}

func PingHost(in ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, n := range in {
			cmd, _ := exec.Command("ping", "-c3", n).Output()
			out <- string(cmd)
		}
		close(out)
	}()
	return out
}
