// config.go
package main

import (
	"os"
	"path/filepath"
)

const (
	path string = "/home/jayderfranca/Sources/firewall/_test/"
)

type parameters []string

type iptables struct {
	table  string
	chain  string
	params parameters
}

type Config struct{}

/*
  FileList lista todos os arquivos de regras
  no diretorio base de configuração
*/
func (c *Config) FileList() ([]string, error) {

	// retorno da lista de arquivos
	files := []string{}

	// realiza a listagem dos arquivos no diretorio base
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	// retorno da listagem
	return files, err
}
