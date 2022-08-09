package main

import "fmt"

const (
	StatusOk                   = 200
	StatusCreated              = 201
	StatusAccepted             = 202
	StatusNonAuthoritativeInfo = 203
	StatusNoContent            = 204
	StatusResetContent         = 205
)

func main() {
	const StatusPartialContent = 206
	fmt.Println(StatusAccepted, StatusOk, StatusPartialContent)
}
