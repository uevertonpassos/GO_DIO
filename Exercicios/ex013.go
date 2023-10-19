//Conte e imprima a quantidade de números pares de 1 a 50.
package main

import "fmt"

func main(){
	contarPares := 0

	for i:=1; i<=50; i++{
		if i % 2 == 0{
			contarPares+=1

		}
		
	}
	fmt.Printf("Existem %d números pares entre 1 e 50\n",contarPares)

}