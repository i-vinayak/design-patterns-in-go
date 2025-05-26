package main

import (
	"fmt"
)

// Context holds the state function
type Context struct {
	state StateFunc
	name  string
}

type StateFunc func(ctx *Context) StateFunc

func (ctx *Context) NextState() StateFunc {
	if ctx.state != nil {
		ctx.state = ctx.state(ctx)
	}
	return nil
}

func idleState(ctx *Context) StateFunc {
	fmt.Println(ctx.name, "is in idle state")
	return processingState
}

func processingState(ctx *Context) StateFunc {
	fmt.Println(ctx.name, "is processing")
	return completeState
}

func completeState(ctx *Context) StateFunc {
	fmt.Println(ctx.name, "has completed")
	return nil
}

// --- Main ---

func main() {
	ctx := &Context{
		name:  "Job123",
		state: idleState,
	}
	fmt.Println("Changing state")
	fmt.Print("\n\n")
	ctx.NextState()

	fmt.Println("Changing state - again")
	fmt.Print("\n\n")
	ctx.NextState()

	fmt.Println("Changing state - again part 2")
	fmt.Print("\n\n")
	ctx.NextState()

}
