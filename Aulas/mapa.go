package main

import "fmt"

func main(){
	/*x:= make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])*/
	elemento:= make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Helio"
	elemento["Ne"] = "Neon"
	fmt.Println(elemento)
	// quase um tradutor
}

