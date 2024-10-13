package main

import (
	"fmt"

	"com.github/anicolaspp/unionfind"
)

func main() {

	uf := unionfind.New[int]()

	for i := 0; i < 10; i++ {
		uf.Add(i)
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

	g := map[string][]string{
		"A": []string{"B", "C"},
		"C": []string{"D"},
		"M": []string{"N", "X"},
		"O": []string{"X"},
	}

	fmt.Println(graphComponents(g))
}

func graphComponents(g map[string][]string) [][]string {
	uf := unionfind.New[string]()

	for k, adj := range g {
		for _, m := range adj {
			uf.Union(k, m)
		}
	}

	return uf.Sets()
}
