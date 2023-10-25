package funciones

import (
	"rerepolez/errores"
	"rerepolez/votos"
	"tdas/cola"
	"strings"
	"fmt"
	"os"
)

// Detecta un error si falta parametros al comenzar
func DetectarErrorParametro(parametros []string, cantidadParametros int) bool {
	if len(parametros) != cantidadParametros {
		err := errores.ErrorParametros{}

		fmt.Println(err.Error())
		return true
	}
	return false
}

// Detecta un error al abrir el archivo
func DetectarErrorArchivo(archivo string) (*os.File, error) {
	archivoAbierto, err := os.Open(archivo)
	if err != nil {
		error := errores.ErrorLeerArchivo{}
		fmt.Println(error)
		return nil, err

	}
	return archivoAbierto, nil
}

func verificarDni(dni int, sliceVotantes []votos.Votante) (votos.Votante, bool) {
	indiceSlice := busquedaBinaria(sliceVotantes, dni)
	if indiceSlice >= 0 { //Si el indice es no negativo, significa que el votante fue encontrado
		return sliceVotantes[indiceSlice], true
	}
	return nil, false
}

// Verificacion de Error de ingreso del comando votar
func VerificarErroresVotacion(tipoVoto string, numeroLista int, arregloDePartidos []votos.Partido, colaVotantes cola.Cola[votos.Votante]) bool {
	cantidadPartidos := len(arregloDePartidos) - 1
	if colaVotantes.EstaVacia() {
		err := errores.FilaVacia{}
		fmt.Println(err.Error())
		return false
	} else if strings.ToLower(tipoVoto) != "presidente" && strings.ToLower(tipoVoto) != "gobernador" && strings.ToLower(tipoVoto) != "intendente" {
		err := errores.ErrorTipoVoto{}
		fmt.Println(err.Error())
		return false
	} else if cantidadPartidos < numeroLista {
		err := errores.ErrorAlternativaInvalida{}
		fmt.Println(err.Error())
		return false
	}
	return true
}


func DetectarVotantesFaltantes(colaVotantes cola.Cola[votos.Votante]) {
	if !colaVotantes.EstaVacia() {
		err := errores.ErrorCiudadanosSinVotar{}
		fmt.Println(err.Error())
	}
}
func DetectarErrorFin(colaVotantes cola.Cola[votos.Votante]) bool {
	if colaVotantes.EstaVacia() {
		err := errores.FilaVacia{}
		fmt.Println(err.Error())
		return true
	}
	return false
}