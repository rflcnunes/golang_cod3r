package fundamentals

import (
	"fmt"

	"github.com/rflcnunes/golang_cod3r/fundamentals/arithmeticoperators"
	"github.com/rflcnunes/golang_cod3r/fundamentals/assignment"
	"github.com/rflcnunes/golang_cod3r/fundamentals/commands"
	"github.com/rflcnunes/golang_cod3r/fundamentals/constvar"
	"github.com/rflcnunes/golang_cod3r/fundamentals/conversions"
	"github.com/rflcnunes/golang_cod3r/fundamentals/first"
	"github.com/rflcnunes/golang_cod3r/fundamentals/functions"
	"github.com/rflcnunes/golang_cod3r/fundamentals/prints"
	"github.com/rflcnunes/golang_cod3r/fundamentals/relationaloperators"
	"github.com/rflcnunes/golang_cod3r/fundamentals/types"
)

func Setup() {
	fmt.Printf("Other program!")

	first.Setup()
	commands.Setup()
	constvar.Setup()
	prints.Setup()
	types.Setup()
	conversions.Setup()
	functions.Setup()
	arithmeticoperators.Setup()
	assignment.Setup()
	relationaloperators.Setup()
}
