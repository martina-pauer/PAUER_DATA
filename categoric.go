package main

import (
	"fmt"
	"log"
	"os"
)

func lineas(nombre_de_archivo string) string {
	// Devuelve un archivo con lineas del archivo

	caracteres_de_archivo, mensaje_de_log := os.ReadFile(nombre_de_archivo)

	if mensaje_de_log != nil {
		// Si problemas con leer el archivo, mostrar mensage generado por el sistema
		log.Fatal(mensaje_de_log)
	}

	texto := ""

	for linea := 0; linea < len(caracteres_de_archivo); linea += 1 {
		if caracteres_de_archivo[linea] != 10 {
			// Sumo caracteres a texto hasta encontrar nueva linea
			texto += string(caracteres_de_archivo[linea])
		} else {
			// Agrego separador de linea para distinguir entre lineas
			texto += ""
		} // Fin condicional iteracion
	} // Fin de iteracion
	return texto
} // Fin de funcion lineas

func contar(elemento string, lista string) uint {
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

func categorias(lista string) string {
	// Devuelve una lista sin duplicados
	clasificacion := string(lista[0])

	for indice := 0; indice < len(lista); indice += 1 {
		// Por cada elemento de la lista, agrego los que no esten en clasificacion
		if contar(string(lista[indice]), clasificacion) == 0 {
			clasificacion += string(lista[indice])
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
	/*
		sobre las funciones...

		- Falta que en vez de string devuelvan una lista de strings

		- Soporte para caracteres Unicode como la o con acento "ó"
	*/
	// Muestra tabla con repeticiones de cada categoria en el archivo
	fmt.Println("|\tCategorias\t|\tRepeticones\t|")
	// Hago así para que se ejecute una vez la funcion lineas
	recorrido := lineas("categorias.lista")
	// Recorro cada una de las categorias y veo cuanto de cada una tiene el archivo
	for indice := 0; indice < len(categorias(recorrido)); indice++ {
		separar(caracteres_linea)
		texto_de_linea := string(categorias(recorrido)[indice])
		fmt.Println("\n\t", texto_de_linea, "\t\t|\t", contar(texto_de_linea, recorrido), "\t")
	}
}
