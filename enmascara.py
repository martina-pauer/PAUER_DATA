#!/usr/bin/python3

import sys

class Mascara:
    '''
        Convierte palabras de un lado a otro de
        un texto en otras.

        Propiedades:

            - Mascara.contigua:str, palabra que divide un lado y otro del texto

            - Mascara.primero:str, texto antes de la palabra contigua

            - Mascara.segundo:str, texto después de la palabra contgua

            - Mascara.anterior:str, texto nuevo con que reemplazar Mascara.primero

            - Mascara.posterior:str, texto nuevo con que reemplazar Mascara.segundo

        Métodos:

            - Mascara.reemplazo(antes:str, luego:str), reemplaza por un texto u otro segun posicion respecto a palabra contigua

        Ejemplo: 

            "hola como adios", alterar palabra "adios" al lado
            del "como" a "¿como están?", al otro cambiar a misma
            palabra.
    '''
    def __init__(self, palabra : str, texto : str):

        self.contigua : str = palabra

        self.primero : str = ''

        self.segundo : str = ''

        texto = texto.split(' ')

        for palabra in texto:
            # Voy sumando palabras hasta encontrar palabra que divide al texto
            if (palabra != self.contigua) and (texto.index(palabra) < texto.index(self.contigua)):
                self.primero += (palabra + ' ')
            # Si sige después de la contigua es la otra parte del texto
            elif (palabra == self.contigua):
                # Agrego palabra contigua a una de las partes del texto, para tener copia del original
                self.segundo += (' ' + palabra + ' ')
            elif (palabra != self.contigua):
                # Sumo palabras a la otra parte
                self.segundo += (' ' + palabra)

        self.anterior : str = ''

        self.posterior : str = ''

    def reemplazo(self, antes : str, luego : str):
        '''
            Reemplaza por diferente texto Mascara.anterior y
            Mascara.posterior
        '''

        self.anterior = antes.__str__()

        self.posterior = luego.__str__()

# Prueba de que todo ande bien, completada el 9 de abril de 2025 a las 10:45 AM (hora Argentina)
#probar = Mascara('como', 'hola como estan ustedes')
#probar.reemplazo('una', 'bien')
#print(f'antes "{probar.primero}" después "{probar.segundo}", reemplazo ({probar.anterior}, {probar.posterior})')

# Guarda lineas en lectura y escribe el texto de reemplzo Mascara.anterior + Mascara.contigua + Mascara.posterior

lectura = open(sys.argv[1].__str__(), 'r')

texto = ''

for linea in lectura.readlines():
    # Voy convirtiendo en texto que luego reemplazo
    texto += linea
# Cierro y limpio lo innecesario de memoria
lectura.close()
del lectura
# Reemplaza con una lista compuesta de palabra contigua[2], anterior[3] y posterior[4]
adaptar = Mascara(sys.argv[2].__str__(), texto)
# Hago esto por seguridad para evitar inyecciones de código en el script
adaptar.reemplazo(sys.argv[3].__str__(), sys.argv[4].__str__())
# Actualizo el texto y elimino objeto que no lo necesito antes que lo parasiten
texto = adaptar.anterior.__str__() + adaptar.posterior.__str__()
# Muestro texto
print(f'\n\toriginal: \n\n{adaptar.primero.__str__() + adaptar.segundo.__str__()}\n\n\tmodficado : \n\n{texto.__str__()}')
del adaptar
# 12:13 AM me aseguré que todo esté bien ahora reemplazo texto del archivo
escritura = open(sys.argv[1].__str__(), 'w')
escritura.write(texto + '\n')
escritura.close()
del escritura
