package firewall

import (
	"testing"
	//"fmt"
	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	result, err := ReadLines("./_examples/filter/input")
	assert.Exactly(t, nil, err)
	assert.Exactly(t, 63, len(result))
	assert.Exactly(t, "# libera o protocolo icmp", result[19])
}

func TestReadLineErrorFileNotExists(t  *testing.T) {
	result, err := ReadLines("./_examples/filter/eu")
	assert.Error(t, err, "open ./_examples/filter/eu: no such file or directory")
	assert.Exactly(t, []string(nil), result)
}

func TestReadLineErrorDirectory(t *testing.T) {
	result, err := ReadLines("./_examples/filter")
	assert.Error(t, err, "ReadLines ./_examples/filter: invalid file type")
	assert.Exactly(t, []string(nil), result)
}

func TestReadLineEmptyFile(t *testing.T) {
	result, err := ReadLines("./_examples/filter/output")
	assert.Exactly(t, nil, err)
	assert.Exactly(t, []string{}, result)
}

func TestIsComment(t *testing.T) {
	result := IsComment("#este e um comentario")
	assert.Exactly(t, true, result)
}

func TestIsCommentNotComment(t *testing.T) {
	result := IsComment("este nap e um comentario")
	assert.Exactly(t, false, result)
}

func TestIsCommentEmpty(t *testing.T) {
	result := IsComment("")
	assert.Exactly(t, false, result)
}

func TestIsEmpty(t *testing.T) {
	result := IsEmpty("")
	assert.Exactly(t, true, result)
}

func TestIsEmptyNotEmpty(t *testing.T) {
	result := IsEmpty("nao esta vazio")
	assert.Exactly(t, false, result)
}

func TestIsEmptyWithSpaces(t *testing.T) {
	result := IsEmpty("    ")
	assert.Exactly(t, true, result)
}