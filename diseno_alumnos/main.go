package main

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/cola"
	"rerepolez/errores"
	"rerepolez/lista"
	"rerepolez/votos"
	"strconv"
	"strings"
)

var (
	// Implementacion de una lista enlazada para guardar el padron.
	listaVotantes lista.Lista[votos.Votante]

	// Implementacion de un arreglo  para guardar los partidos
	arregloDePartidos []votos.Partido

	//Implementacion de 1 cola para el orden en que se ingresan los votantes.
	colaVotantes    cola.Cola[votos.Votante]
	partidoEnBlanco votos.Partido
)

// Detecta un error si falta parametros al comenzar
func detectarErrorParametro(parametros []string, cantidadParametros int) bool {
	if len(parametros) != cantidadParametros {
		error := errores.ErrorParametros{}

		fmt.Println(error.Error())
		return true
	}
	return false
}

// Detecta un error al abrir el archivo
func detectarErrorArchivo(archivo string) (*os.File, error) {
	archivoAbierto, err := os.Open(archivo)
	if err != nil {
		error := errores.ErrorLeerArchivo{}
		fmt.Println(error)
		return nil, err

	}
	return archivoAbierto, nil
}

// Funcion para leer el archivo del padron
func lecturaDePadron(archivo string) bool {
	archivoAbierto, errorArchivo := detectarErrorArchivo(archivo)
	defer archivoAbierto.Close()
	if errorArchivo != nil {
		return false
	}
	lector := bufio.NewReader(archivoAbierto)
	for {
		linea, err := lector.ReadString('\n')
		if err != nil {
			break
		}
		linea = strings.TrimSuffix(linea, "\n")
		dni, _ := strconv.Atoi(linea)
		votante := votos.CrearVotante(dni)
		listaVotantes.InsertarUltimo(votante)
	}
	return true
}

// Funcion para leer el archivo de las boletas. Ademas crea cada Partido y los guarda en una lista enlazada.
func lecturaDeBoletas(archivo string) bool {
	archivoAbierto, errorArchivo := detectarErrorArchivo(archivo)
	defer archivoAbierto.Close()
	if errorArchivo != nil {
		return false
	}
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
	return true
}

// Esta funcion verifica si el DNI pertenece al padron cargado previamente.
func verificarDni(dni int) (votos.Votante, bool) {
	iterador := listaVotantes.Iterador()
	for iterador.HaySiguiente() {
		if dni == (iterador.VerActual().LeerDNI()) {
			return iterador.VerActual(), true
		}
		iterador.Siguiente()
	}
	return nil, false
}

// Ingresa el dni válido a la cola.
func ingresar(dni int) {
	if dni <= 0 {
		err := errores.DNIError{}
		fmt.Println(err.Error())
		return
	}
	votante, existe := verificarDni(dni)
	if !existe {
		err := errores.DNIFueraPadron{}
		fmt.Println(err.Error())
		return
	}
	colaVotantes.Encolar(votante)
	fmt.Println("OK")

}

// Verificacion de Error de ingreso del comando votar
func verificarErroresVotacion(tipoVoto string, numeroLista int) bool {
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

// Tranasfoma la palabra del tipoVoto a un numero.
func transformarTipoVoto(tipoVoto string) votos.TipoVoto {
	if tipoVoto == "presidente" {
		return votos.PRESIDENTE //0
	} else if tipoVoto == "gobernador" {
		return votos.GOBERNADOR //1
	} else {
		return votos.INTENDENTE //2
	}
}

// Funcion de votar
func votar(numeroLista int, tipoVoto string) {
	tipo := transformarTipoVoto(tipoVoto)
	persona := colaVotantes.VerPrimero()
	err := persona.Votar(tipo, numeroLista)
	if err != nil {
		//El unico error que puede haber es que el votante ya haya votado (voto fraudulento)
		//En ese caso, tambien hay que sacarlo de la cola
		colaVotantes.Desencolar()
		fmt.Println(err.Error())
		return
	}
	fmt.Println("OK")
}

func deshacer() {
	persona := colaVotantes.VerPrimero()
	err := persona.Deshacer()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("OK")
	return
}

func imprimirResltador() {
	fmt.Println("\nPresidente:")
	stringPresidente := partidoEnBlanco.ObtenerResultado(votos.PRESIDENTE)
	for _, partido := range arregloDePartidos {
		stringPresidente += partido.ObtenerResultado(votos.PRESIDENTE)
		stringPresidente += "\n"
	}
	fmt.Println(stringPresidente)
	fmt.Println("Gobernador:")
	stringGobernador := partidoEnBlanco.ObtenerResultado(votos.GOBERNADOR)
	for _, partido := range arregloDePartidos {
		stringGobernador += partido.ObtenerResultado(votos.GOBERNADOR)
		stringGobernador += "\n"
	}
	fmt.Println(stringGobernador)
	fmt.Println("\nIntendente:")
	stringIntendente := partidoEnBlanco.ObtenerResultado(votos.INTENDENTE)
	for _, partido := range arregloDePartidos {
		stringIntendente += partido.ObtenerResultado(votos.INTENDENTE)
		stringIntendente += "\n"
	}
	fmt.Println(stringIntendente)
	fmt.Printf("\nVotos impugnados: %d\n", votos.VotosImpugnados)
}

func finVoto(votante votos.Votante) {
	voto, posibleError := votante.FinVoto()
	if posibleError == nil && !voto.Impugnado {
		cantidadCandidatos := 3
		for i := 0; i < cantidadCandidatos; i++ {
			if voto.VotoPorTipo[i] != 0 {
				arregloDePartidos[voto.VotoPorTipo[i]].VotadoPara(votos.TipoVoto(i))
			}
		}
	}
}
func votarEnBlanco() {
	cantidadCandidatos := 3
	for !colaVotantes.EstaVacia() {
		for i := 0; i < cantidadCandidatos; i++ {
			partidoEnBlanco.VotadoPara(votos.TipoVoto(i))
		}
		colaVotantes.Desencolar()
	}
}

func detectarVotantesFaltantes() {
	if !colaVotantes.EstaVacia() {
		err := errores.ErrorCiudadanosSinVotar{}
		fmt.Println(err.Error())
		votarEnBlanco()
	}
}
func detectarErrorFin() bool {
	if colaVotantes.EstaVacia() {
		err := errores.FilaVacia{}
		fmt.Println(err.Error())
		return true
	}
	return false
}

func main() {
	params := os.Args[1:]
	listaVotantes = lista.CrearListaEnlazada[votos.Votante]()
	if detectarErrorParametro(params, 2) {
		return
	}
	if !lecturaDeBoletas(params[0]) {
		return
	}
	if !lecturaDePadron(params[1]) {
		return
	}
	partidoEnBlanco = votos.CrearVotosEnBlanco()

	colaVotantes = cola.CrearColaEnlazada[votos.Votante]()
	//Usamos esta forma, ya que es la que encontramos por internet. El fmt.Scan() nos estaba generando problemas con separar por ejemplo el ingresar <dni> en 2 .
	escanerInput := bufio.NewScanner(os.Stdin)
	var parametro string

	for escanerInput.Scan() {

		parametro = escanerInput.Text()
		parametroSeparado := strings.Fields(parametro)
		comando := parametroSeparado[0]

		switch strings.ToLower(comando) {
		case "ingresar":
			dniIngresado, _ := strconv.Atoi(parametroSeparado[1])
			ingresar(dniIngresado)

		case "votar":
			tipoVoto := parametroSeparado[1]
			numeroLista, _ := strconv.Atoi(parametroSeparado[2])
			if verificarErroresVotacion(tipoVoto, numeroLista) {
				votar(numeroLista, tipoVoto)
			}
		case "deshacer":
			deshacer()
		case "fin-votar":
			if !detectarErrorFin() {
				votante := colaVotantes.VerPrimero()
				finVoto(votante)
				colaVotantes.Desencolar()
			}
		}
	}
	detectarVotantesFaltantes()
	imprimirResltador()

}

/*
FALTA:

	Aclaracion: Hay que crear partido en blanco

	FUNCION DE Fin Votar
	IMPRIMIR RESULTADOS (Falta el como leer cuando se finaliza el programa asi ejecutar la funcion)

	Terminar las implementaciones de votantes/partidos.
*/
