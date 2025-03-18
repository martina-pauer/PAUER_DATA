#!/usr/bin/python3

nombre_de_archivo = 'categorias.txt'

import categoric, random

caso_categorias = random.randint(1, 10)

caso_apariciones = random.randint(0, 100)

print(f'\nCuenta de categorías en archivo "{nombre_de_archivo}" {categoric.categorias(nombre_de_archivo, caso_categorias)}:\n\n')
# Extraigo categoría del diccionario y cuento del archivo cuanto hay de cada una
for clasificacion in categoric.categorias(nombre_de_archivo).keys():
    print(f'\tDe "{clasificacion.replace("\n", "")}" hay {categoric.apariciones(nombre_de_archivo, clasificacion)}\n')
    print(f'\t\t(Con dibujo de {caso_apariciones} "{clasificacion.replace("\n", "")}" hay {categoric.apariciones(nombre_de_archivo, clasificacion, caso_apariciones)})\n')
