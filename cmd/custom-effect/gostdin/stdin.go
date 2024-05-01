// Package gostdin provides helpers for working with stdin
package gostdin

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// OnSignal calls handle when the given signal is received
func OnSignal(sig syscall.Signal, handle func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, sig)
	go func() {
		<-c
		handle()
	}()
}

// OnInput calls handle when any input is received
func OnInput(prompt string, handle func(string), closeCh chan bool) {
	ch := make(chan string)
	go func(ch chan string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(prompt + " ")
			s, err := reader.ReadString('\n')
			if err != nil {
				close(ch)
				return
			}
			ch <- s
		}
	}(ch)

stdinloop:
	for {
		select {
		case stdin, ok := <-ch:
			if !ok {
				break stdinloop
			} else {
				handle(strings.TrimRight(stdin, "\n"))
			}
		case <-closeCh:
			close(ch)
			return
		}
	}
}
