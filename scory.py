#!/usr/bin/python3

def limite() -> int:
    '''
        Dice cual es el puntaje maximo que
        se puede alcanzar
    '''
    # Hago así porque no existen constantes en Python y es posible que lo reuse
    return 100
    
def puntaje_positivo(texto : str, listado_bueno : str) -> int:
    '''
        Suma de todos los puntos por cada coincidencia
        en el texto del listado en un archivo de texto.
    '''
    puntaje_bueno = 1
    
    archivo_bueno = open(listado_bueno, 'r')
    
    for linea_buena in archivo_bueno.readlines():
        # Si el texto contiene algun termino positivo suma puntos
        if texto.__contains__(linea_buena.replace('\n', '')) and puntaje_bueno <= limite():
            puntaje_bueno *= 2
            
    archivo_bueno.close()
    
    return puntaje_bueno

def puntaje_negativo(texto : str, listado_malo : str) -> int:
    '''
        Suma puntos por cada coincidencia de un termino
        negativo del listado en el texto.
    '''
    puntaje_malo = 1
    
    archivo_malo = open(listado_malo, 'r')
    
    for linea_mala in archivo_malo.readlines():
        if (texto.__contains__(linea_mala.replace('\n', '')) or texto.isupper()) and puntaje_malo >= 1:
            # Por cada termino malo resto puntos
            puntaje_malo //= 3
    
    archivo_malo.close()
    
    return puntaje_malo
# Algortimo principal
def principal():
    print('\n[ANALIZA QUE TAN POSITIVO ES UN TEXTO]\n')             
    # Pido comentario a analizar
    ingresado = input('\tEscriba comentario: ')
    # Calculo promedio (media) entre lo positivo y negativo
    puntos = puntaje_positivo(ingresado, 'texto-positivo.txt')
    
    puntos += puntaje_negativo(ingresado, 'texto-negativo.txt')

    puntos //= 2
    # Muestro resultado
    print(f'\nEl indice de positividad del comentario es {puntos} puntos.')
    # Pregunto si continua
    if input('\n\tDesea continuar S/n: ') != 'n':
        principal()
        
        
principal()    
