package main

import (
	"fmt"
	"os"
	"rerepolez/errores"
	"rerepolez/funciones"
	"rerepolez/votos"
	"strconv"
	"strings"
	"tdas/cola"
)

func main() {
	params := os.Args[1:]

	// Implementacion de un arreglo  para guardar los partidos
	if funciones.DetectarErrorParametro(params, 2) {
		return
	}
	arregloDePartidos, HayErrorEnBoletas := funciones.LecturaDeBoletas(params[0])
	if HayErrorEnBoletas {
		return
	}
	sliceVotantes, HayErrorEnPadron := funciones.LecturaDePadron(params[1])
	if HayErrorEnPadron {
		return
	}
	contadorVotos := 0
	votosImpugnados := 0
	partidoEnBlanco := votos.CrearVotosEnBlanco()
	colaVotantes := cola.CrearColaEnlazada[votos.Votante]()
	escanerInput := funciones.CrearEscaner()

	for escanerInput.Scan() {
		parametroSeparado, comando := funciones.ObtenerParametroComando(escanerInput)

		switch strings.ToLower(comando) {

		case "ingresar":
			dniIngresado, _ := strconv.Atoi(parametroSeparado[1])
			funciones.Ingresar(dniIngresado, sliceVotantes, &colaVotantes)

		case "votar":
			tipoVoto := parametroSeparado[1]
			numeroLista, err := strconv.Atoi(parametroSeparado[2])

			if err != nil {
				errorAlterniva := errores.ErrorAlternativaInvalida{}
				fmt.Println(errorAlterniva.Error())
			} else if funciones.VerificarErroresVotacion(tipoVoto, numeroLista, arregloDePartidos, colaVotantes) {
				funciones.Votar(numeroLista, tipoVoto, &colaVotantes, &contadorVotos)
			}

		case "deshacer":
			if !funciones.DetectarErrorFin(colaVotantes) {
				funciones.Deshacer(&colaVotantes)
			}

		case "fin-votar":
			if !funciones.DetectarErrorFin(colaVotantes) {
				votante := colaVotantes.Desencolar()
				votosImpugnados = funciones.FinVoto(votante, &arregloDePartidos, &partidoEnBlanco, contadorVotos, votosImpugnados)
				contadorVotos = 0
				fmt.Println("OK")
			}
		}
	}

	funciones.DetectarVotantesFaltantes(colaVotantes)
	funciones.ImprimirResltador(arregloDePartidos, partidoEnBlanco, votosImpugnados)
}

/*
Obligatorio:
El padrón electoral tiene un formato conocido, y puede usare esta información para ordenar mejor que en tiempo
*/
