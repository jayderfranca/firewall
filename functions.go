package main

import (
	"os"
	"bufio"
	//"fmt"
	"strings"
	"io/ioutil"
	"path/filepath"
)

// isFdDirectory indica se o caminho informado eh um diretorio
func isFDDirectory(path string) bool {

	// retorna as informacoes do caminho
	info, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}

	// valida se eh um diretorio
	return info.IsDir()
}

// isFDExecutable indica se o caminho informado eh um executavel
func isFDExecutable(path string) bool {

	// eh um diretorio
	if isFDDirectory(path) {
		return false
	}

	// obtem as informacoes do caminho
	info, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}

	// valida se o arquivo eh executavel
	// isto so eh valido para SO basado *nix, Windows validar como verificar se eh executavel
	if info.Mode()&0111 != 0 {
		return true
	}

	// retorno padrao
	return false
}

// IsFDRegular indica se o caminho informado eh um arquivo normal
func IsFDRegular(path string) bool {

	// eh um diretorio
	if isFDDirectory(path) {
		return false
	}

	// obtem as informacoes do caminho
	info, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}

	// retorno indicando se o arquivo eh normal
	return info.Mode().IsRegular()
}

// ReadLines realiza a leitura de todas as linhas de um arquivo
func ReadLines(file string) []string {

	// linhas do arquivo
	lines := []string{}

	// verifica se eh um arquivo valido
	if !IsFDRegular(file) {
		return lines
	}

	// abre o arquivo para leitura
	handle, err := os.Open(file)
	if err != nil {
		return lines
	}

	// fecha o arquivo no final da funcao
	defer handle.Close()

	// leitura das linhas
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// retorno das linhas
	return lines
}

// FindExecutable procura no sistema operacional oexecutavel informado
func FindExecutable(name string) string {

	// retorna a lista de caminhos atribuidos a variavel PATH
	paths := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))

	// loop dos caminhos encontrados
	for _, path := range paths {

		// eh um diretorio ?
		if !isFDDirectory(path) {
			continue
		}

		// lista todos descritores de arquivos no diretorio
		files, err := ioutil.ReadDir(path)
		if err != nil {
			continue
		}

		// loop da lista encontrada
		for _, file := range files {

			// monta o caminho completo
			full := filepath.Join(path, file.Name())

			// valida se fd encontrado eh um executavel e
			// se o executavel que sendo procurado
			if isFDExecutable(full) && file.Name() == name {
				return full
			}
		}
	}

	// caso nao encontre, retorna vazio
	return ""
}