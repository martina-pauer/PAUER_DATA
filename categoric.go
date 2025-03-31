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
			texto += fmt.Sprintf("\t[Hasta caracter %d]\n", linea)
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
	/*
		Se repite el ciclo mas externo de 1 a el cuadrado de elementos de la lista
		eso depende de cuantos elementos duplicados tenga la lista y su respectiva longitud
	*/
	for indice := 0; indice < len(lista); indice += 1 {
		// Por cada elemento de la lista, agrego los que no esten en clasificacion
		if contar(string(lista[indice]), clasificacion) == 0 {
			clasificacion += string(lista[indice])
		}
	} // Cerrar recorrida entre elementos sin dupliicar

	return clasificacion
} // fin funcion categorias

func main() {
	/*
		Prueba unitaria hasta el momento de funcion lineas...

		- Funciono, puede imprimir su porpio código fuente

		- Falta que en vez de string devuelva una lista de strings

		- Soporte para caracteres Unicode como la o con acento "ó"
	*/
	// Prueba mostrar lineas
	fmt.Println(lineas("categoric.go"))
	// Prueba unitaria contar repeticiones, falta que sea por lista de lineas y no de caracteres
	fmt.Println("\n\tSe repite {", contar("{", lineas("categoric.go")), "veces en el archivo.")
	// Prueba unitaria de categorias, falta que no se quede en el primer caracter
	fmt.Println("\n\tlos caracteres sin duplicados son", categorias("{{}{}{}{}{}{{}}}"))
}
