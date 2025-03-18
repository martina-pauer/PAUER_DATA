#!/usr/bin/python3

def categorias(archivo:str, valor_inicial:int = 0) -> dict:
    '''
        Devuelve un diccionario con cada una de las clasificaciones
        escritas en un archivo y un valor inicial desde donde ir
        contando cuantas veces aparece algo dentro de esa categoría.
    '''
    diccionario = {}
    
    lectura = open(archivo, 'r')
    
    for linea in lectura.readlines():
    # Agrega solo categorías nuevas y cuenta desde un valor inicial
        if not diccionario.keys().__contains__(linea):
            diccionario[linea] = valor_inicial
            
    lectura.close()
      
    return diccionario
    
def apariciones(archivo:str, categoria:str, conteo:int = 0) -> int:
    '''
        Cuenta desde un valor inicial cuantas veces aparece una categoría en un archivo.
    '''
    analizado = open(archivo, 'r')
    
    for linea in analizado.readlines():
        if linea == categoria:
            # Solo cuenta si aparece la categoría buscada
            conteo += 1
            
    analizado.close()
    
    return conteo
