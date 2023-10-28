package funciones

import (
	"fmt"
	"rerepolez/errores"
	"rerepolez/votos"
	"strings"
	"tdas/cola"
)

func obtenerMaxVotante(slice []votos.Votante) int {
	max := slice[0].LeerDNI()
	for i := 1; i < len(slice); i++ {
		if slice[i].LeerDNI() > max {
			max = slice[i].LeerDNI()
		}
	}
	return max
}

// Ingresa el dni válido a la cola.
func Ingresar(dni int, sliceVotantes []votos.Votante, colaVotantes *cola.Cola[votos.Votante]) {
	if dni <= 0 {
		err := errores.DNIError{}
		fmt.Println(err.Error())
		return
	}
	votante, existe := verificarDni(dni, sliceVotantes)
	if !existe {
		err := errores.DNIFueraPadron{}
		fmt.Println(err.Error())
		return
	}
	(*colaVotantes).Encolar(votante)
	fmt.Println("OK")

}

// Funcion de votar
func Votar(numeroLista int, tipoVoto string, colaVotantes *cola.Cola[votos.Votante], contadorVotos *int) {
	tipo := transformarTipoVoto(strings.ToLower(tipoVoto))
	persona := (*colaVotantes).VerPrimero()
	err := persona.Votar(tipo, numeroLista)
	if err != nil {
		(*colaVotantes).Desencolar()
		fmt.Println(err.Error())
		return
	}
	(*contadorVotos)++
	fmt.Println("OK")
}

func Deshacer(colaVotantes *cola.Cola[votos.Votante]) {
	persona := (*colaVotantes).VerPrimero()
	err := persona.Deshacer()
	errorVotoFraude := errores.ErrorVotanteFraudulento{persona.LeerDNI()}
	if err != nil {
		if err == errorVotoFraude {
			(*colaVotantes).Desencolar()
		}
		fmt.Println(err.Error())
		return
	}
	fmt.Println("OK")
}

func ImprimirResltador(arregloDePartidos []votos.Partido, partidoEnBlanco votos.Partido, VotosImpugnados int) {
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
	if VotosImpugnados == 1 {
		fmt.Printf("Votos Impugnados: %d voto\n", VotosImpugnados)
		return
	}
	fmt.Printf("Votos Impugnados: %d votos\n", VotosImpugnados)
}

func FinVoto(votante votos.Votante, arregloDePartidos *[]votos.Partido, partidoEnBlanco *votos.Partido, contadorVotos int, VotosImpugnados int) int {
	voto, posibleError := votante.FinVoto()
	if voto.Impugnado {
		VotosImpugnados++
	}
	todoVotoEnBlanco := [3]int{0, 0, 0}
	cantidadCandidatos := 3
	if posibleError == nil && !voto.Impugnado && contadorVotos > 0 && voto.VotoPorTipo != todoVotoEnBlanco {

		for i := 0; i < cantidadCandidatos; i++ {

			if voto.VotoPorTipo[i] != 0 {
				(*arregloDePartidos)[voto.VotoPorTipo[i]].VotadoPara(votos.TipoVoto(i))
			} else {
				(*partidoEnBlanco).VotadoPara(votos.TipoVoto(i))
			}
		}
	} else if posibleError == nil && !voto.Impugnado && voto.VotoPorTipo == todoVotoEnBlanco {

		for i := 0; i < cantidadCandidatos; i++ {
			(*partidoEnBlanco).VotadoPara(votos.TipoVoto(i))
		}
	}
	return VotosImpugnados

}
