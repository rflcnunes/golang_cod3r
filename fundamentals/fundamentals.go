package fundamentals

import (
	"fmt"

	"github.com/rflcnunes/golang_cod3r/fundamentals/commands"
	"github.com/rflcnunes/golang_cod3r/fundamentals/constvar"
	"github.com/rflcnunes/golang_cod3r/fundamentals/first"
	"github.com/rflcnunes/golang_cod3r/fundamentals/prints"
)

func Setup() {
	fmt.Printf("Other program!")

	first.Setup()
	commands.Setup()
	constvar.Constvar()
	prints.Setup()
}
