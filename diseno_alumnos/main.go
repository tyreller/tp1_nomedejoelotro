package main

import (
	"rerepolez/lista"
	"rerepolez/errores"
	"rerepolez/votos"
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
)
//Implementacion de 2 listas enlazadas para guardar los partidos y el padron.
var listaVotantes lista.Lista[votos.Votante]
var listaPartidos lista.Lista[votos.Partido]

//Implementacion de 1 cola para el orden en que se ingresan los votantes.

//Detecta un error si falta parametros al comenzar
func detectarErrorParametro(parametros []string,cantidadParametros int)bool{
	if len(parametros) != cantidadParametros {
		error := errores.ErrorParametros{}
		
		fmt.Println(error)
		return true
	}
	return false
}

//Detecta un error al abrir el archivo
func detectarErrorArchivo(archivo string) (*os.File,error){
	archivoAbierto, err := os.Open(archivo)
	if err != nil{
		error:= errores.ErrorLeerArchivo{}
		fmt.Println(error)
		return nil,err
	
	}
	return archivoAbierto,nil
}

//Funcion para leer el archivo del padron
func lecturaDePadron(archivo string) bool {
	archivoAbierto,errorArchivo:= detectarErrorArchivo(archivo)
	defer archivoAbierto.Close()
	if errorArchivo != nil{
		return false
	}
	lector:=bufio.NewReader(archivoAbierto)
	for {
		linea, err := lector.ReadString('\n')
		if err != nil {
			break
		}
		linea = strings.TrimSuffix(linea, "\n")
		
		dni,_:= strconv.Atoi(linea)
		votante:= votos.CrearVotante(dni)
		
		listaVotantes.InsertarUltimo(votante)
	}
	return true
}	

//Funcion para leer el archivo de las boletas. Ademas crea cada Partido y los guarda en una lista enlazada.
func lecturaDeBoletas(archivo string) bool{
	archivoAbierto,errorArchivo :=detectarErrorArchivo(archivo)
	defer archivoAbierto.Close()
	if errorArchivo != nil{
		return false
	}
	lector:=bufio.NewReader(archivoAbierto)
	contador := 0
	for {
		
		linea, err := lector.ReadString('\n')
		if err != nil {
			break
		}
		partidoArreglo :=  strings.Split(linea, ",")
		nombrePartido := partidoArreglo[0]
		var candidatos  [3]string
		candidatos[0]=partidoArreglo[1]
		candidatos[1]=partidoArreglo[2]
		candidatos[2]=partidoArreglo[3]
		partido:= votos.CrearPartido(nombrePartido,contador,candidatos)
		contador ++
		listaPartidos.InsertarUltimo(partido)
		
	}
	return true
}

func main(){
	params:= os.Args[1:]
	listaVotantes = lista.CrearListaEnlazada[votos.Votante]()
	listaPartidos = lista.CrearListaEnlazada[votos.Partido]()

	if detectarErrorParametro(params,2){
		return
	}
	
	if !lecturaDeBoletas(params[0]){
		return
	}
	if !lecturaDePadron(params[1]){
		return
	}
	
	parametro:=""
	fmt.Println(listaVotantes.VerPrimero())
	fmt.Println(listaVotantes.VerUltimo())
	fmt.Println(listaPartidos.VerPrimero())
	fmt.Println(listaPartidos.VerUltimo())
	for parametro!="fin-votar"{
		fmt.Scanf("%s", &parametro);

	}
	
}

/*
FALTA:

	FUNCION DE INGRESAR
	FUNCION DE VOTAR
	FUNCION DE DESHACER
	IMPRIMIR RESULTADOS

	Terminar las implementaciones de votantes/partidos.
*/