//Gere os primeiros 10 números da sequência de Fibonacci.
//F(n) = F(n-1) + F(n-2) para n > 1

package main

import "fmt"

func main(){
	n:= 10 //numero de termos da sequência
	a, b := 0,1

	for i:=0; i< n; i++{
		fmt.Println(a)
		a, b = b, a+b
	}
}