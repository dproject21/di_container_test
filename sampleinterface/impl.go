package sampleinterface

import (
	"fmt"
)

type SampleInterface interface {
	Hoge()
	Fuga()
}

type SampleFoo struct {
}

func NewSampleFoo() SampleInterface {
	return new(SampleFoo)
}

type SampleBar struct {
}

func NewSampleBar() SampleInterface {
	return new(SampleBar)
}

type SamplePrinter interface {
	Print()
}

type DefaultPrinter struct {
	printer SampleInterface
}

func NewDefaultPrinter(s SampleInterface) SamplePrinter {
	d := new(DefaultPrinter)
	d.printer = s
	return d
}

func (p *DefaultPrinter) Print() {
	p.printer.Hoge()
	p.printer.Fuga()
}

func (s *SampleFoo) Hoge() {
	fmt.Println("hoge...")
}

func (s *SampleFoo) Fuga() {
	fmt.Println("fuga...")
}

func (s *SampleBar) Hoge() {
	fmt.Println("HOGE!!!")
}

func (s *SampleBar) Fuga() {
	fmt.Println("FUGA!!!")
}
