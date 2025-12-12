package io_test

import (
	io "golang_learn/go_basic/io"
	"testing"
)

func TestWriteFile(t *testing.T) {
	io.WriteFile()

}

func TestReadFile(t *testing.T) {
	io.ReadFile()
}

func TestReadFileWithBuffer(t *testing.T) {
	io.ReadFileWithBuffer()
}

func TestWriteFileWithBuffer(t *testing.T) {
	io.WriteFileWithBuffer()
}

func TestJsonSerialize(t *testing.T) {
	io.JsonSerialize()
}

func TestSlog(t *testing.T) {
	io.Slog(io.NewSlogger())
}
