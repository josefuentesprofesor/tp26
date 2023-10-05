package main

import (
	"fmt"
	"os"
)

func main() {
	// Crear un mapa para almacenar los apellidos de los jugadores y la descripcion.
	diccionario := map[string]string{
		"Martínez": "Edad: 26 \nNombre completo: Lautaro Martinez \nApodo: El Toro \nFecha de nacimiento: 22 de agosto de 1997 \nEstatura: 1.74 \nPeso: 72 ",
		//TODO agregar un jugador de la seleccion
		"Allister": "Edad: 24 \nNombre completo: Alexis Mac Allister \nApodo: Colo \nFecha de nacimiento: 24 de diciembre de 1998 \nEstatura: 1.76 \nPeso: 72 ",
	}

	// Verificar si se proporciona un argumento (la palabra a buscar).
	if len(os.Args) != 2 {
		fmt.Println("Uso: tp26.exe apellido_jugador")
		return
	}

	// Obtener el apellido proporcionado como argumento.
	apellido := os.Args[1]

	// Utilizar una declaración switch para buscar y mostrar el descripcion.
	switch descripcion := diccionario[apellido]; descripcion {
	case "":
		fmt.Printf("El jugador '%s' no se encuentra.\n", apellido)
	default:
		fmt.Printf("%s:\n%s\n", apellido, descripcion)
	}
}
