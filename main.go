package main

import (
	"fmt"

	"com.github/anicolaspp/unionfind"
)

func main() {

	uf := unionfind.New[int]()

	for i := 0; i < 10; i++ {
		uf.Union(i, i)
	}

	fmt.Println(uf.Sets())

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i%2 == j%2 {
				uf.Union(i, j)
			}
		}
	}

	fmt.Println(uf.Sets())
}
