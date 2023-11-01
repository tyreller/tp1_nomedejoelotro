package main

import (
	"os"
	"rerepolez/funciones"
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

	colaVotantes, partidoEnBlanco, votosImpugnados := funciones.BucleDelPrograma(arregloDePartidos, sliceVotantes)

	funciones.DetectarVotantesFaltantes(colaVotantes)
	funciones.ImprimirResltador(arregloDePartidos, partidoEnBlanco, votosImpugnados)
}
