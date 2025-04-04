package main

import (
	"fmt"
	"log"
	"os"
)

func lineas(nombre_de_archivo string, longitud_de_caracteres uint8) []string {
	/*
		Devuelve una lista con lineas del archivo pasado como parametro
		y considerando que por linea solo puede haber un entero fijo de
		caracteres
	*/

	caracteres_de_archivo, mensaje_de_log := os.ReadFile(nombre_de_archivo)

	if mensaje_de_log != nil {
		// Si problemas con leer el archivo, mostrar mensage generado por el sistema
		log.Fatal(mensaje_de_log)
	}

	lista := make([]string, len(caracteres_de_archivo))

	for linea := 0; linea < len(lista); linea += 1 {

		if (caracteres_de_archivo[linea] < 65 || caracteres_de_archivo[linea] > 90) && (caracteres_de_archivo[linea] < 97 || caracteres_de_archivo[linea] > 122) && (caracteres_de_archivo[linea] < 37 || caracteres_de_archivo[linea] > 58) {
			// Si el caracter no es alfanumerico, analizo los siguientes
			continue
		}

		if (linea + int(longitud_de_caracteres)) < len(lista) {
			// Asigno caracter de linea en la lista
			lista[linea] = string(caracteres_de_archivo[linea:(linea + int(longitud_de_caracteres))])
			// Salto 16 caracteres asignados a la linea
			linea += int(longitud_de_caracteres)
		} else {
			// Paro ejecucion si no puedo asignar mas
			continue
		}
	} // Fin de iteracion

	return lista
} // Fin de funcion lineas

func contar(elemento string, lista []string) uint {
	// Busca cuantas veces se repite un elemento en una lista

	var apariciones uint = 0

	for conteo := 0; conteo < len(lista); conteo++ {
		// Recorro lista, debe ser sin duplicados
		if elemento == string(lista[conteo]) {
			// Comparo con elemento de lista
			apariciones += 1
		} // Solo cuento una aparicion mas si es lo que busco
	}

	return apariciones

}

func categorias(lista []string) []string {
	// Devuelve una lista sin duplicados de hasta longitud de lista
	clasificacion := make([]string, len(lista))

	for indice := 0; indice < len(lista); indice += 1 {
		// Por cada elemento de la lista, agrego los que no esten en clasificacion
		if contar(string(lista[indice]), clasificacion) == 0 {
			clasificacion[indice] = string(lista[indice])
		}
	} // Cerrar recorrida entre elementos sin dupliicar

	return clasificacion
} // fin funcion categorias

func separar(n int) {
	for caracteres := 0; caracteres < n; caracteres++ {
		fmt.Print("_")
	}
}

const caracteres_linea int = 49

func main() {
	// Muestra tabla con repeticiones de cada categoria en el archivo
	fmt.Println("|\tCategorías\t|\tRepeticiones\t|")
	// Hago así para que se ejecute una vez la funcion lineas
	recorrido := lineas("categorias.lista", 16)
	clasificaciones := categorias(recorrido)
	// Recorro cada una de las categorias y veo cuanto de cada una tiene el archivo
	for indice := 0; indice < len(clasificaciones); indice++ {
		if contar(clasificaciones[indice], recorrido) != 368 {
			separar(caracteres_linea)
			fmt.Println("\n\t", clasificaciones[indice], "\t\t|\t", contar(clasificaciones[indice], recorrido), "\t")
		}
	}
}
