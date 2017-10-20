package main

import (
	"testing"
	//"fmt"
	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	result := ReadLines("./_examples/filter/input")
	assert.Exactly(t, 63, len(result))
	assert.Exactly(t, "# libera o protocolo icmp", result[19])
}

func TestReadLineErrorFileNotExists(t  *testing.T) {
	result := ReadLines("./_examples/filter/eu")
	assert.Exactly(t, []string{}, result)
}

func TestReadLineErrorDirectory(t *testing.T) {
	result := ReadLines("./_examples/filter")
	assert.Exactly(t, []string{}, result)
}

func TestReadLineEmptyFile(t *testing.T) {
	result := ReadLines("./_examples/filter/output")
	assert.Exactly(t, []string{}, result)
}

func TestFindExecutable(t *testing.T) {
	result := FindExecutable("sh")
	assert.Exactly(t, "/bin/sh", result)
}

func TestFindExecutableNotFound(t *testing.T) {
	result := FindExecutable("sh1")
	assert.Exactly(t, "", result)
}