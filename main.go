package main

import (
	"fmt"
	"os"
	"rerepolez/errores"
	"rerepolez/votos"
	"strconv"
	"strings"
	"tdas/cola"
	"rerepolez/funciones"
)

func main() {
	params := os.Args[1:]
	// Implementacion de un slice para guardar el padron en el.
	sliceVotantes := make([]votos.Votante, 0)

	// Implementacion de un arreglo  para guardar los partidos
	var arregloDePartidos []votos.Partido

	if funciones.DetectarErrorParametro(params, 2) {
		return
	}
	if !funciones.LecturaDeBoletas(params[0], &arregloDePartidos) {
		return
	}
	if !funciones.LecturaDePadron(params[1], &sliceVotantes) {
		return
	}
	VotosImpugnados := 0
	partidoEnBlanco := votos.CrearVotosEnBlanco()

	colaVotantes := cola.CrearColaEnlazada[votos.Votante]()

	//Usamos esta forma, ya que es la que encontramos por internet. El fmt.Scan() nos estaba generando problemas con separar por ejemplo el ingresar <dni> en 2 .
	
	contadorVotos := 0
	escanerInput := funciones.CrearEscaner()
	for escanerInput.Scan() {
		parametroSeparado,comando:=  funciones.ObtenerParametroComando(escanerInput)
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
				VotosImpugnados = funciones.FinVoto(votante, &arregloDePartidos, &partidoEnBlanco, contadorVotos, VotosImpugnados)
				contadorVotos = 0
				fmt.Println("OK")
			}
		}
	}
	funciones.DetectarVotantesFaltantes(colaVotantes)
	funciones.ImprimirResltador(arregloDePartidos, partidoEnBlanco, VotosImpugnados)

}
