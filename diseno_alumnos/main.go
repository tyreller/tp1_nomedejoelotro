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

// Implementacion de 2 listas enlazadas para guardar los partidos y el padron.
var listaVotantes lista.Lista[votos.Votante]
var listaPartidos lista.Lista[votos.Partido]

//Implementacion de 1 cola para el orden en que se ingresan los votantes.

var ColaVotantes cola.Cola[votos.Votante]

// Detecta un error si falta parametros al comenzar
func detectarErrorParametro(parametros []string, cantidadParametros int) bool {
	if len(parametros) != cantidadParametros {
		error := errores.ErrorParametros{}

		fmt.Println(error)
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
	contador := 1
	votosEnBlanco:= votos.CrearVotosEnBlanco()
	listaPartidos.InsertarPrimero(votosEnBlanco)
	for {

		linea, err := lector.ReadString('\n')
		if err != nil {
			break
		}
		partidoArreglo := strings.Split(linea, ",")
		nombrePartido := partidoArreglo[0]
		var candidatos [3]string
		candidatos[0] = partidoArreglo[1]
		candidatos[1] = partidoArreglo[2]
		candidatos[2] = partidoArreglo[3]
		partido := votos.CrearPartido(nombrePartido, contador, candidatos)
		contador++
		listaPartidos.InsertarUltimo(partido)

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
	ColaVotantes.Encolar(votante)
	fmt.Println("OK")
	return
}

//Verificacion de Error cuando se realiza la funcion dee Votar
func verificarErrores(tipoVoto string, numeroLista int){
	cantidadPartidos:= listaPartidos.Largo()-1
	if ColaVotantes.EstaVacia(){
		err:= errores.FilaVacia{}
		fmt.Println(err.Error())
	}else if strings.ToLower(tipoVoto) != "presidente" && strings.ToLower(tipoVoto) != "gobernador" && strings.ToLower(tipoVoto) != "intendente"{
		err:= errores.ErrorTipoVoto{}
		fmt.Println(err.Error())
	}else if cantidadPartidos<numeroLista{
		err:=errores. ErrorAlternativaInvalida{}
		fmt.Println(err.Error())
	}
}


func main() {
	params := os.Args[1:]
	listaVotantes = lista.CrearListaEnlazada[votos.Votante]()
	listaPartidos = lista.CrearListaEnlazada[votos.Partido]()

	if detectarErrorParametro(params, 2) {
		return
	}

	if !lecturaDeBoletas(params[0]) {
		return
	}
	if !lecturaDePadron(params[1]) {
		return
	}

	ColaVotantes = cola.CrearColaEnlazada[votos.Votante]()
	//Usamos esta forma, ya que es la que encontramos por internet. El fmt.Scan() nos estaba generando problemas con separar por ejemplo el ingresar <dni> en 2 .
	escanerInput := bufio.NewScanner(os.Stdin)
	var parametro string

	for strings.ToLower(parametro) != "fin-votar" {
		escanerInput.Scan()
		parametro = escanerInput.Text()
		parametroSeparado := strings.Fields(parametro)
		comando:=parametroSeparado[0]

		switch strings.ToLower(comando) {
		case "ingresar":
			dniIngresado, _ := strconv.Atoi(parametroSeparado[1])
			ingresar(dniIngresado)

		case "votar":
			tipoVoto:= parametroSeparado[1]
			numeroLista,_ := strconv.Atoi(parametroSeparado[2])
			verificarErrores(tipoVoto,numeroLista)
		}
	}
}

/*
FALTA:
	FUNCION DE VOTAR
	FUNCION DE DESHACER
	IMPRIMIR RESULTADOS

	Terminar las implementaciones de votantes/partidos.
*/
