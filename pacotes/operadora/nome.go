package operadora

import (
	"fmt"
	"strconv"
)

//NomeOperadora variavel publica
var NomeOperadora = "XHXOP Telecon"

//Numero é um código de operadora
var Numero = 67

//PrefixoDaOperadora testando concatenação
var PrefixoDaOperadora = strconv.Itoa(Numero) + " " + NomeOperadora

func main() {
	fmt.Printf("Testando : %s/r/n", PrefixoDaOperadora)
}
