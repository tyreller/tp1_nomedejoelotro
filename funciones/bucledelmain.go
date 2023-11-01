package funciones

import (
	"fmt"
	"rerepolez/errores"
	"rerepolez/votos"
	"strconv"
	"strings"
	"tdas/cola"
)

func BucleDelPrograma(arregloDePartidos []votos.Partido, sliceVotantes []votos.Votante) (cola.Cola[votos.Votante], votos.Partido, int) {
	contadorVotos := 0
	votosImpugnados := 0
	partidoEnBlanco := votos.CrearVotosEnBlanco()
	colaVotantes := cola.CrearColaEnlazada[votos.Votante]()
	escanerInput := CrearEscaner()

	for escanerInput.Scan() {
		parametroSeparado, comando := ObtenerParametroComando(escanerInput)

		switch strings.ToLower(comando) {

		case "ingresar":
			dniIngresado, _ := strconv.Atoi(parametroSeparado[1])
			Ingresar(dniIngresado, sliceVotantes, &colaVotantes)

		case "votar":
			tipoVoto := parametroSeparado[1]
			numeroLista, err := strconv.Atoi(parametroSeparado[2])

			if err != nil {
				errorAlterniva := errores.ErrorAlternativaInvalida{}
				fmt.Println(errorAlterniva.Error())
			} else if VerificarErroresVotacion(tipoVoto, numeroLista, arregloDePartidos, colaVotantes) {
				Votar(numeroLista, tipoVoto, &colaVotantes, &contadorVotos)
			}

		case "deshacer":
			if !DetectarErrorFin(colaVotantes) {
				Deshacer(&colaVotantes)
			}

		case "fin-votar":
			if !DetectarErrorFin(colaVotantes) {
				votante := colaVotantes.Desencolar()
				votosImpugnados = FinVoto(votante, &arregloDePartidos, &partidoEnBlanco, contadorVotos, votosImpugnados)
				contadorVotos = 0
				fmt.Println("OK")
			}
		}
	}
	return colaVotantes, partidoEnBlanco, votosImpugnados
}
