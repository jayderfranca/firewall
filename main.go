// firewall project main.go
package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

// estrutura com as opcoes
type Options struct {
	Dry   bool `short:"d" long:"dry" description:"Exibe apenas o processamento das regras"`
	Apply bool `short:"a" long:"apply" description:"Aplica as regras na mémoria"`
	Save  bool `short:"s" long:"save" description:"Aplica e salva as regras no formato original do iptables"`
}

// realiza o parse da linha de comando para a estrutura
func ParseOptions() (*Options, error) {

	// armazena os dados dos parametros
	opts := Options{}

	// cria novo parser
	parser := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)

	// parse dos parametros
	_, err := parser.Parse()
	flag, ok := err.(*flags.Error)

	// tratamento de erro no parse
	if err != nil && (!ok || flag.Type != flags.ErrHelp) {
		fmt.Printf("%v: %v\n", os.Args[0], err)
		fmt.Printf("Try '%v --help' for more information.\n", os.Args[0])
		return nil, err
	}

	// exibe o help
	if err != nil && ok && flag.Type == flags.ErrHelp {
		fmt.Fprintln(os.Stdout, err)
	}

	// retorno dos parametros tratados
	return &opts, nil
}

func main() {

	// realiza o tratamento dos parametros informados
	opts, err := ParseOptions()

	// tratamento erro no parse
	if err != nil {
		os.Exit(-1)
	}

	fmt.Println(opts)

	config := Config{}
	files, err := config.FileList()

	for _, file := range files {
		fmt.Println(file)
	}

	// finaliza o programa
	os.Exit(0)
}
