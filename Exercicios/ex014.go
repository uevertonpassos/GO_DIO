//Calcule a soma dos dígitos de um número inteiro.

package main

import "fmt"

func main() {
    numero := 12 // Substitua pelo número que você deseja calcular a soma dos dígitos
    soma := 0

    for numero > 0 {
        digito := numero % 10
        soma += digito
        numero /= 10
    }

    fmt.Println("A soma dos dígitos é:", soma)
}
