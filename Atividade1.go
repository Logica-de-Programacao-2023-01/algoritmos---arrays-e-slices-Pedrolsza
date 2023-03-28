package main

import "fmt"

func main() {
	nomes := []string{"Pedro", "Zveiter", "Sanderson", "Klaus", "Cayo"}
	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	}

}
