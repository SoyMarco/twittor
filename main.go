package main

import (
	"log"

	"github.com/SoyMarco/twittor/bd"
	"github.com/SoyMarco/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
