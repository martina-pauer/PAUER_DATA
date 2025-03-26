const ARCHIVO : &str = "lista.txt";

const CARACTER : &str = "_";

const CARACTERES : u8 = 42;

const LINEAS : usize = 3;

use std :: fs :: File;

use std :: io :: {Read, BufReader, BufRead, Result};

fn separar(separadores : u8, caracter : &str)
{
    // Imprime de 1 a n caracteres para hacer separador visual de lineas
    print!("\t");
    for _separador in 1 ..= separadores
    {
        print!("{}", caracter);
    }
    println!("");    
}

fn categorias(ruta_archivo : &str) -> Result<[[String; 2]; LINEAS]>
{
    /* Relaciona lineas contenidas en un archivo y en una lista dice
       que categorias son y cuantas. */
       
       // Junta columna de categorias y de cuenta categoria
       let mut repes : [[String; 2]; LINEAS] = [["-".to_string(), "0".to_string()], ["-".to_string(), "0".to_string()], ["-".to_string(), "0".to_string()]];
       
       let mut indice = 0;
       
       let lector = BufReader :: new (File :: open(ruta_archivo)?);
       
       let mut suma : u32 = 1;
       
       for linea in lector.lines()
       {
            // Agrego si la categoria es nueva
            if repes[indice][0] == "-".to_string()
            {
                // Agrego desde archivo
                repes[indice][0] = linea?;
                // Cuento que existe una de esa categoria en su respectivo lugar
                repes[indice][1] = suma.to_string();
            }
            // Sino cuento uno mas en la categoria
            else
            {
                suma += 1;
            }
            
            if indice < LINEAS
            {
                indice += 1;
            }
       }
       
       return Ok(repes);
       
}

fn main() -> Result <()>
{
    /* Algoritmo:
    
            1) Leer lineas de un archivo y relacionarlas a una categoría.
            
            2)  En la lista con categorías contar cuanto se repite de cada una.
            
            3)  Mostrar resultado en una tabla
    */
    separar(CARACTERES, CARACTER);
    println!("\n\t\tCATEGORÍA\t|\tVECES\t");
    separar(CARACTERES, CARACTER);
    // Digo cuantas veces se repite cada categoría 
    for dato in &categorias(ARCHIVO)
    {
        println!("\n\t{:?}\t\t|\t{:?}\t", dato[0], dato[1]);
        separar(CARACTERES, CARACTER);
    }
    Ok(())
}       
