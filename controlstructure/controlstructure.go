package controlstructure

import (
	"fmt"

	"github.com/rflcnunes/golang_cod3r/controlstructure/ifelse"
	"github.com/rflcnunes/golang_cod3r/controlstructure/ifelseif"
)

func Setup() {
	fmt.Println("Control Structure")

	ifelse.Setup()
	ifelseif.Setup()
}
