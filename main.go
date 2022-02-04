package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
)


func main()  {
	leerArchivo()
	leerArchivo2()
	escribirArchivo()
	escribirArchivo2()
}

func leerArchivo()  {
	archivo, err := ioutil.ReadFile("./archivo.txt") //abre el archivo pero no deja manipularlo
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leerArchivo2()  {
	archivo, err := os.Open("./archivo.txt") //abrir el archivo
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo) //Permite capturar mediante un método de entrada, por ejemplo teclado, pero también permite leer ficheros.
		for scanner.Scan() { //leer la linea y guardarla en un buffer
			registro := scanner.Text() //lee del buffer y convierte la línea en una cadena de texto
			fmt.Printf("Linea: " + registro + "\n")
		}
	}
	archivo.Close() //cerrar el archivo
}

func escribirArchivo()  {
	archivo, err := os.Create("./nuevoArchivo.txt") //Crear nuevo archivo, si tengo un archivo con el mismo nombre lo sobreescribe
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Esta es una linea nueva") //Para grabar el texto en un archivo
	archivo.Close() //cerrar el archivo
}

func escribirArchivo2()  {
	fileName := "./nuevoArchivo.txt"
	if Append(fileName, "\n Esta es una segunda linea") == false {
		fmt.Println("Hubo un error")
	} 
}

//Añade nuevas lineas a un archivo que ya existe
func Append(fileName string, texto string) bool  {
	//abrir el archivo
	//con os.O_WRONLY le indicamos que el archivo va a ser de escritura y de lectura
	//con os.O_APPEND le decimos que NO queremos que lo sobreescriba, sino que queremos añadir lineas nuevas
	//0644 son los permisos que le damos al archivo, permisos de lectura y escritura para usuario, para dueño y para grupo
	archivo, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error al abrir el fichero")
		return false
	}

	_, err = archivo.WriteString(texto) //Grabar un String en un archivo
	if err != nil {
		fmt.Println("Hubo un error al escribir el fichero")
		return false
	}

	return true
	
}
