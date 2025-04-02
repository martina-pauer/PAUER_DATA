package main

import (
	"fmt"
	"log"
	"os"
)

func lineas(nombre_de_archivo string) string {
	// Devuelve una lista con lineas del archivo pasado como parametro

	caracteres_de_archivo, mensaje_de_log := os.ReadFile(nombre_de_archivo)

	if mensaje_de_log != nil {
		// Si problemas con leer el archivo, mostrar mensage generado por el sistema
		log.Fatal(mensaje_de_log)
	}

	texto := ""

	for linea := 0; linea < len(caracteres_de_archivo); linea += 1 {
		/*
			- "á" bytes (195, 161) ordinal 225

			- "é" bytes (195, 169) ordinal 233

			- "í" bytes (195, 173) ordinal 237

			- "ó" bytes (195, 179) ordinal 243

			- "ú" bytes (195, 186) ordinal 250
		*/
		if caracteres_de_archivo[linea] == 195 {
			// Avanzo a caracter parametro de este caracter de control
			linea += 1
		}
		// Se debe analizar cada uno de los casos por separado para ser mas flexible
		if caracteres_de_archivo[linea] == 161 {
			// Agrego a con tilde de forma directa para evitar errores
			texto += "á"
			// Voy al siguiente caracter para no mostrar versión defectuosa del caracter
			linea += 1
		}

		if caracteres_de_archivo[linea] == 169 {
			//  Misma lógica de arriba
			texto += "é"
			linea += 1
		}

		if caracteres_de_archivo[linea] == 173 {
			texto += "í"
			linea += 1
		}

		if caracteres_de_archivo[linea] == 179 {
			texto += "ó"
			linea += 1
		}

		if caracteres_de_archivo[linea] == 186 {
			texto += "ú"
			linea += 1
		}

		// Sumo caracteres a texto una vez que me posicione bien
		texto += string(caracteres_de_archivo[linea])
		// Acá debería ir el código para agregar texto a la lista
		// .....
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

			- "" bytes (195, 1 F)lta que en vez de string devuelvan una lista de strings

			- "" bytes (195, 1 S)porte para caracteres Unicode como la o con acento "ó"
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
