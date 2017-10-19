package firewall

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

// ReadLines realiza a leitura de todas as linhas de um arquivo
func ReadLines(file string) ([]string, error) {

	// abre o arquivo para leitura
	handle, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	// pbtem as informacoes do arquivo
	info, err := handle.Stat()
	if err != nil {
		return nil, err
	}

	// verifica se eh um arquivo valido
	if !info.Mode().IsRegular() {
		return nil, fmt.Errorf("ReadLines %s: invalid file type", file)
	}

	// fecha o arquivo no final da funcao
	defer handle.Close()

	// linhas do arquivo
	lines := []string{}

	// leitura das linhas
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// retorno das linhas
	return  lines, nil
}

// IsComment indica que o texto informado eh um comentario (iniciado com #)
func IsComment(text string) bool {
	if text == "" || text[:1] != "#" {
		return false
	}
	return true
}

// IsEmpty identifica que o valor informado eh vazio
func IsEmpty(text string) bool {
	return strings.Trim(text, " ") == ""
}