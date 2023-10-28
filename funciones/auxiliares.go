package funciones

import (
	"rerepolez/votos"
)

func obtenerDigito(persona votos.Votante, div int) int {
	return (persona.LeerDNI() / div) % 10
}

func busquedaBinaria(sliceVotantes []votos.Votante, dniBuscado int) int {
	izq, der := 0, len(sliceVotantes)-1

	for izq <= der {
		mit := (izq + der) / 2

		if sliceVotantes[mit].LeerDNI() < dniBuscado {
			izq = mit + 1
		}
		if sliceVotantes[mit].LeerDNI() > dniBuscado {
			der = mit - 1
		}
		if sliceVotantes[mit].LeerDNI() == dniBuscado {
			return mit
		}
	}

	return -1
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
