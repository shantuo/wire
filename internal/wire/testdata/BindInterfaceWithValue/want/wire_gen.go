// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"io"
	"os"
)

// Injectors from wire.go:

func injectReader() io.Reader {
	file := _wireFileValue
	return file
}

var (
	_wireFileValue = os.Stdin
)
