package main

import (
	"fmt"

	"github.com/alecthomas/repr"
	"github.com/paulopotter/go-course-project/anime"
)

func main() {
	anime, err := anime.SearchAnime("No%20game")
	if err == nil {
		repr.Println(anime)
	} else {
		fmt.Println("Falha ao buscar usuário: ", err)
	}
}
