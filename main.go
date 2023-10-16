package main

import (
	"bufio"
	"fmt"
	"os"
	"tdas/cola"
	
	"rerepolez/errores"
	"rerepolez/votos"
	"strconv"
	"strings"
)

var (
	// Implementacion de una lista enlazada para guardar el padron.
	//listaVotantes lista.Lista[votos.Votante]
	sliceVotantes []votos.Votante
	// Implementacion de un arreglo  para guardar los partidos
	arregloDePartidos []votos.Partido

	//Implementacion de 1 cola para el orden en que se ingresan los votantes.
	colaVotantes    cola.Cola[votos.Votante]
	partidoEnBlanco votos.Partido
	contadorVotos int
)

// Detecta un error si falta parametros al comenzar
func detectarErrorParametro(parametros []string, cantidadParametros int) bool {
	if len(parametros) != cantidadParametros {
		err := errores.ErrorParametros{}

		fmt.Println(err.Error())
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
		sliceVotantes = append(sliceVotantes,votante)
	}

	sliceVotantes = mergeSortVotantes(sliceVotantes) //Ordena el slice de votantes
	
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
	indiceSlice, encontrado := busquedaBinaria(sliceVotantes, dni)
	if encontrado == true {
		return sliceVotantes[indiceSlice],true
	}
	return nil, false
}

func merge(izq, der []votos.Votante) []votos.Votante{
	sliceFinal := make([]votos.Votante, 0, len(izq)+len(der)) 
	//Le fijamos la cap para que no tenga que redimensionar

	i, d := 0,0 //izquierda y derecha

	//Mientras no se acabe ninguna de las mitades
	for i < len(izq) && d < len(der){
		if izq[i].LeerDNI() < der[d].LeerDNI() {
			sliceFinal = append(sliceFinal,izq[i])
			i++
		} else {
			sliceFinal = append(sliceFinal,der[d])
			d++
		}
	}

	//Inserta todos los elementos que faltaran, si es que faltaba alguno
	for j := i; j<len(izq); j++{
		sliceFinal = append(sliceFinal,izq[j])
	}
	for j := d; j<len(der); j++{
		sliceFinal = append(sliceFinal,der[j])
	}

	return sliceFinal
}

func mergeSortVotantes(slice []votos.Votante) []votos.Votante {
	if len(slice) <= 1{
		return slice
	}

	mit := len(slice)/2
	izq := mergeSortVotantes(slice[:mit])
	der := mergeSortVotantes(slice[mit:])

	return merge(izq,der)
}

func busquedaBinaria(sliceVotantes []votos.Votante, dniBuscado int) (int,bool){
	izq, der := 0, len(sliceVotantes)-1

	for izq <= der {
		mit := (izq + der)/2

		if sliceVotantes[mit].LeerDNI() < dniBuscado {
			izq = mit + 1
		}
		if sliceVotantes[mit].LeerDNI() > dniBuscado {
			der = mit - 1
		}
		if sliceVotantes[mit].LeerDNI() == dniBuscado{
			return mit, true
		}
	}

	return -1,false
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
	tipo := transformarTipoVoto(strings.ToLower(tipoVoto))
	persona := colaVotantes.VerPrimero()
	err := persona.Votar(tipo, numeroLista)
	if err != nil {
		colaVotantes.Desencolar()
		fmt.Println(err.Error())
		return
	}
	contadorVotos ++
	fmt.Println("OK")
}

func deshacer() {
	persona := colaVotantes.VerPrimero()
	err := persona.Deshacer()
	errorVotoFraude := errores.ErrorVotanteFraudulento{persona.LeerDNI()}
	if err != nil {
		if err ==  errorVotoFraude{
			colaVotantes.Desencolar()
		}
		fmt.Println(err.Error())
		return
	}
	fmt.Println("OK")
	return
}

func imprimirResltador() {
	fmt.Println("Presidente:")
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
	fmt.Println("Intendente:")
	stringIntendente := partidoEnBlanco.ObtenerResultado(votos.INTENDENTE)
	for _, partido := range arregloDePartidos {
		stringIntendente += partido.ObtenerResultado(votos.INTENDENTE)
		stringIntendente += "\n"
	}
	fmt.Println(stringIntendente)
	if votos.VotosImpugnados == 1{
		fmt.Printf("Votos Impugnados: %d voto\n", votos.VotosImpugnados)
		return	
	}
	fmt.Printf("Votos Impugnados: %d votos\n", votos.VotosImpugnados)
}

func finVoto(votante votos.Votante) {
	voto, posibleError := votante.FinVoto()
	todoVotoEnBlanco := [3]int{0,0,0}
	cantidadCandidatos := 3
	if posibleError == nil && !voto.Impugnado && contadorVotos > 0 && voto.VotoPorTipo != todoVotoEnBlanco{

		for i := 0; i < cantidadCandidatos; i++ {
		
			if voto.VotoPorTipo[i] != 0 {
				arregloDePartidos[voto.VotoPorTipo[i]].VotadoPara(votos.TipoVoto(i))
			}else{
				partidoEnBlanco.VotadoPara(votos.TipoVoto(i))
			}
		}
	} else if posibleError == nil &&  !voto.Impugnado &&voto.VotoPorTipo == todoVotoEnBlanco{

		for i := 0; i < cantidadCandidatos; i++ {
			partidoEnBlanco.VotadoPara(votos.TipoVoto(i))
		}
	}
	
}

func detectarVotantesFaltantes() {
	if !colaVotantes.EstaVacia() {
		err := errores.ErrorCiudadanosSinVotar{}
		fmt.Println(err.Error())
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
	sliceVotantes = make([]votos.Votante, 0)
	
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
	contadorVotos = 0
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
			
			numeroLista, err:= strconv.Atoi(parametroSeparado[2])
			if err != nil{
				errorAlterniva := errores.ErrorAlternativaInvalida{}
				fmt.Println(errorAlterniva.Error())
			}else if verificarErroresVotacion(tipoVoto, numeroLista)  {
				votar(numeroLista, tipoVoto)
			}
		case "deshacer":
			if !detectarErrorFin(){
				deshacer()
			}
		case "fin-votar":
			if !detectarErrorFin() {
				votante := colaVotantes.Desencolar()
				finVoto(votante)
				contadorVotos = 0
				fmt.Println("OK")
			}
		}
	}
	detectarVotantesFaltantes()
	imprimirResltador()

}
