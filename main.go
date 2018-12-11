package main

import (
	"log"

	"github.com/alecthomas/repr"
	"github.com/paulopotter/go-course-project/anime"
)

func main() {
	anime, err := anime.SearchAnime("No%20game")
	if err == nil {
		repr.Print(anime)
	} else {
		log.Println("Falha ao buscar usuário: ", err)
	}
}
