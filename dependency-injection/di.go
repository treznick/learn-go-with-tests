package main

import (
	"fmt"
	"io"
	"os"
)

type Greet struct {
	Writer io.Writer
}

type Option func(*Greet)

func Options(opts ...Option) Option {
	return func(g *Greet) {
		for _, opt := range opts {
			opt(g)
		}
	}
}

func Output(writer io.Writer) Option {
	return func(g *Greet) {
		g.Writer = writer
	}
}

func NewGreet(opts ...Option) *Greet {
	greeter := &Greet{Writer: os.Stdout}

	for _, opt := range opts {
		opt(greeter)
	}

	return greeter
}

func (g Greet) Hello(name string) {
	fmt.Fprintf(g.Writer, "Hello, %s", name)
}

func main() {
	NewGreet().Hello("elodie")
}
