package main

import "fmt"

// Variáveis publicas ou privadas:
//Se coloco a variável com a primeira letra maiúscula, quer dizer que ela será pública, se for minúscula será privada
//ex.
var Cidade, bairro string // variável Cidade(pública), bairro(privada)

//var <nome da variável><tipo>
// As variáveis criadas sem atribuição de valor não viram nulas e sim com um valor defaul, ex. endereco =""
// posso definir mais de uma variavel na mesma linha
//var endereco string
//var telefone = "9999-9999"
var endereco, telefone string

// O int em go é uma forma default, mas temos outros tipos de int, são eles: int, int8, int32, int16 e int64
// se tivermos uma máquina 32 bits, o int vai ser para 32, se for 64 será para 64, ele vai se adaptar
var quantidade int // por default vira com o valor 0
var comprou bool   // por default virá como false
var valor float64  // por default virá como 0.00

// posso também definir variáveis dentro de parenteses
var (
	numero int
	estado string
	cep    int
)

// caracteres especiais
var palavras rune

func main() {
	// Outras formas de definir uma variável
	// os ":" substituem o var
	teste := "valor de teste"

	fmt.Printf("endereco: %s\r\n", endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("teste: %v\r\n", teste)

}
