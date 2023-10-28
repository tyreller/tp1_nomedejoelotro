package funciones

import (
	"bufio"
	"rerepolez/votos"
	"strconv"
	"strings"
)

func LecturaDePadron(archivo string) ([]votos.Votante, bool) {
	sliceVotantes := make([]votos.Votante, 0)
	sliceDNI := make([]int, 0)
	archivoAbierto, errorArchivo := DetectarErrorArchivo(archivo)
	if errorArchivo != nil {
		return sliceVotantes, true
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
		sliceVotantes = append(sliceVotantes, votante) //Almacena uno a uno todos los DNI's
		sliceDNI = append(sliceDNI, dni)
	}

	radixSortVotantes(sliceVotantes) //Ordena el slice de votantes

	return sliceVotantes, false
}

// Funcion para leer el archivo de las boletas. Ademas crea cada Partido y los guarda en una lista enlazada.
func LecturaDeBoletas(archivo string) ([]votos.Partido, bool) {
	archivoAbierto, errorArchivo := DetectarErrorArchivo(archivo)
	arregloDePartidos := make([]votos.Partido, 0)
	if errorArchivo != nil {
		return arregloDePartidos, true
	}
	defer archivoAbierto.Close()
	lector := bufio.NewReader(archivoAbierto)
	partidoNulo := votos.CrearPartido("", votos.LISTA_IMPUGNA, [votos.CANT_VOTACION]string{"", "", ""}, [votos.CANT_VOTACION]int{votos.LISTA_IMPUGNA, votos.LISTA_IMPUGNA, votos.LISTA_IMPUGNA})
	arregloDePartidos = append(arregloDePartidos, partidoNulo)
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
		arregloDePartidos = append(arregloDePartidos, partido)
	}
	return arregloDePartidos, false
}
