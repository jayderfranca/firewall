package main

import (
	"testing"
	//"fmt"
	"github.com/stretchr/testify/assert"
)

func TestIsFDDirectory(t *testing.T) {
	result := isFDDirectory("/usr/bin")
	assert.Exactly(t, true, result)
}

func TestIsFDDirectoryFile(t *testing.T) {
	result := isFDDirectory("/etc/hosts")
	assert.Exactly(t, false, result)
}

func TestIsFDDirectoryEmpty(t *testing.T) {
	result := isFDDirectory("")
	assert.Exactly(t, false, result)
}

func TestIsFDRegular(t *testing.T) {
	result := IsFDRegular("/etc/hosts")
	assert.Exactly(t, true, result)
}

func TestIsFDExecutable(t *testing.T) {
	result := isFDExecutable("/bin/sh")
	assert.Exactly(t, true, result)
}

func TestIsFDExecutableRegular(t *testing.T) {
	result := isFDExecutable("/etc/hosts")
	assert.Exactly(t, false, result)
}

func TestIsFDExecutableDirectory(t *testing.T) {
	result := isFDExecutable("/usr/bin")
	assert.Exactly(t, false, result)
}

func TestIsFDExecutableEmpty(t *testing.T) {
	result := isFDExecutable("")
	assert.Exactly(t, false, result)
}

func TestIsFDRegularNotRegular(t *testing.T) {
	result := IsFDRegular("/dev/null")
	assert.Exactly(t, false, result)
}

func TestIsFDRegularEmpty(t *testing.T) {
	result := IsFDRegular("")
	assert.Exactly(t, false, result)
}

func TestParseLines(t *testing.T) {
	result := ParseLines("./_examples/filter/input")
	assert.Exactly(t, 63, len(result))
	assert.Exactly(t, "# libera o protocolo icmp", result[19])
}

func TestRParseLinesFileNotExists(t  *testing.T) {
	result := ParseLines("./_examples/filter/eu")
	assert.Exactly(t, []string{}, result)
}

func TestParseLinesDirectory(t *testing.T) {
	result := ParseLines("./_examples/filter")
	assert.Exactly(t, []string{}, result)
}

func TestRParseLinesEmptyFile(t *testing.T) {
	result := ParseLines("./_examples/filter/output")
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

func TestGetVariables(t *testing.T) {
	result := GetVariables("./_examples/config")
	assert.Exactly(t, 7, len(result))
	assert.Exactly(t, result["MMulticast"], "224.0.0.1")
}

func TestGetVariablesInvalidFile(t *testing.T) {
	result := GetVariables("/etc/hosts")
	assert.Exactly(t, 0, len(result))
}