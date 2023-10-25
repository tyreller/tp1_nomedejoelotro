package funciones

import (
	"bufio"
	"rerepolez/votos"
	"strconv"
	"strings"
)


// Funcion para leer el archivo del padron
func LecturaDePadron(archivo string, sliceVotantes *[]votos.Votante) bool {
	archivoAbierto, errorArchivo := DetectarErrorArchivo(archivo)
	if errorArchivo != nil {
		return false
	}
	
	defer archivoAbierto.Close()
	lector := bufio.NewReader(archivoAbierto)
	for {
		linea, err := lector.ReadString('\n')
		if err != nil {
			break
		}
		linea = strings.TrimSuffix(linea, "\n")
		dni, _ := strconv.Atoi(linea)
		votante := votos.CrearVotante(dni)
		*sliceVotantes = append(*sliceVotantes, votante) //Almacena uno a uno todos los DNI's
	}

	*sliceVotantes = mergeSortVotantes(*sliceVotantes) //Ordena el slice de votantes

	return true
}



// Funcion para leer el archivo de las boletas. Ademas crea cada Partido y los guarda en una lista enlazada.
func LecturaDeBoletas(archivo string, arregloDePartidos *[]votos.Partido) bool {
	archivoAbierto, errorArchivo := DetectarErrorArchivo(archivo)
	defer archivoAbierto.Close()
	if errorArchivo != nil {
		return false
	}

	lector := bufio.NewReader(archivoAbierto)
	partidoNulo := votos.CrearPartido("", votos.LISTA_IMPUGNA, [votos.CANT_VOTACION]string{"", "", ""}, [votos.CANT_VOTACION]int{votos.LISTA_IMPUGNA, votos.LISTA_IMPUGNA, votos.LISTA_IMPUGNA})
	*arregloDePartidos = append(*arregloDePartidos, partidoNulo)
	contador := 1
	for {
		linea, err := lector.ReadString('\n')
		if err != nil {
			break
		}
		partidoArreglo := strings.Split(linea, ",")
		nombrePartido := partidoArreglo[0]
		var candidatos [votos.CANT_VOTACION]string

		candidatos[votos.PRESIDENTE] = strings.TrimSuffix(partidoArreglo[1], "\n")
		candidatos[votos.GOBERNADOR] = partidoArreglo[2]
		candidatos[votos.INTENDENTE] = strings.TrimSuffix(partidoArreglo[3], "\n")
		partido := votos.CrearPartido(nombrePartido, contador, candidatos, [votos.CANT_VOTACION]int{0, 0, 0})
		contador++
		*arregloDePartidos = append(*arregloDePartidos, partido)
	}
	return true
}
