package extension

import (
	"fmt"
	"testing"
)

type BasicOp interface {
	RunBasicOp()
}

type BasicMode struct {
	BasicParams int
}

type BasicMode1 struct {
	BasicParams int
}

func (bm *BasicMode) RunBasicOp() {
	fmt.Println("Run basic mode with params", bm.BasicParams)
}

func (bm *BasicMode1) RunBasicOp() {
	fmt.Println("Run basic mode1 with params", bm.BasicParams)
}

type ExtendedMode struct {
	ExtParams int
	bm        *BasicMode
}

type ExtendedMode1 struct {
	BasicMode
	ExtParams int
	BasicOp
}

func (em *ExtendedMode) RunExtendedOp() {
	fmt.Println("Run ext mode with params", em.ExtParams)
}

func (em *ExtendedMode) RunBasicOp() {
	em.bm.RunBasicOp()
}

// func (em *ExtendedMode1) RunBasicOp() {
// 	fmt.Println("Run ext Basic mode with params", em.ExtParams)
// }

func runBasicMode(bm BasicOp) {
	bm.RunBasicOp()
}

func TestExtendedMode(t *testing.T) {
	em := &ExtendedMode{10, &BasicMode{1}}
	em.RunExtendedOp()
	//runBasicMode(em)
}

func TestExtendedMode1(t *testing.T) {
	em := &ExtendedMode1{}
	//	em.BasicParams = 1
	em.ExtParams = 2
	//em.RunBasicOp()
	em.BasicMode.RunBasicOp()
	runBasicMode(em)
}
