package main

import "fmt"

type Processor struct {
	op  func(in float64) float64
	in  chan float64
	out chan float64
}

func (processor *Processor) process() {

	for inValue := range processor.in {
		processor.out <- processor.op(inValue)
	}
	close(processor.out) // Done writing all values
}

func makeChain(processors []*Processor, generator func() (float64, bool), consumer func(float64)) {

	in := make(chan float64)
	out := make(chan float64)

	processors[0].in = in
	processors[len(processors)-1].out = out

	for i := 0; i < len(processors)-1; i++ {
		channel := make(chan float64) // A channel between ith and i+1 th processors
		processors[i].out = channel   // out is already set for last processor
		processors[i+1].in = channel  // in is already set for first [0] processor
	}

	for i := 0; i < len(processors); i++ {
		go processors[i].process() // start each processor in a separate go routine
	}

	go func() {
		for value, ok := generator(); ok; value, ok = generator() {
			in <- value
		}
		close(in)
	}()

	func() {
		for outValue := range out {
			consumer(outValue)
		}
	}()

}

// Sample operations

func add(value float64) func(in float64) float64 {
	return func(in float64) float64 {
		return in + value
	}
}

func mul(value float64) func(in float64) float64 {
	return func(in float64) float64 {
		return in * value
	}
}

func div(value float64) func(in float64) float64 {
	return func(in float64) float64 {
		return in / value
	}
}

func square(in float64) float64 {
	return in * in
}

func main() {

	processors := []*Processor{
		{op: add(2)},
		{op: square},
		{op: div(6)},
	}

	genValue := 0.0

	makeChain(processors, func() (float64, bool) {
		if genValue < 11 {
			value := genValue
			genValue++
			return value, true
		}
		return 0.0, false
	}, func(output float64) {
		fmt.Println(output)
	})

}
