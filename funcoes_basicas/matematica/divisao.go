package matematica

// import (
// 	"fmt"
// )

// func main() {

// 	resultadoObtido := divisor(4, 2)
// 	fmt.Printf(" O resultado da divisão é : %d", resultadoObtido)

// 	resultadoObtido, resto := divisaoComResto(4, 2)
// 	fmt.Printf(" O resultado da divisão é : %d resto :%d", resultadoObtido, resto)
// }

func divisor(x int, y int) (variavelDeRetorno int) {

	variavelDeRetorno = x / y

	return
}

func divisaoComResto(x int, y int) (variavelDeRetorno int, resto int) {

	variavelDeRetorno = x / y
	resto = x % y

	return
}
