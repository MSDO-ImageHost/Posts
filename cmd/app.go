package main

import (
	"context"

	"github.com/asim/nitro-plugins/broker/rabbitmq"
	"github.com/asim/nitro/v3/app/rpc"
)

// Define a request type
type Request struct {
	Name string
}

// Define a response type
type Response struct {
	Message string
}

// Create your public App Handler
type Handler struct{}

// Create a public Handler method which takes request, response and returns an error
func (h *Handler) Call(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Message = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new App
	app := rpc.NewApp()

	// Set the App name
	app.Name("helloworld")

	// Register the Handler
	app.Handle(new(Handler))

	// Run the App (blocking call)
	app.Run()
}
