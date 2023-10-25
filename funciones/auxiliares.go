package funciones

import (
	"rerepolez/votos"
)
// MergeSort para ordenar de menor a mayor el padron
// Parte el slice en mitades recursivamente hasta que quede solo 1 elemento
func mergeSortVotantes(slice []votos.Votante) []votos.Votante {
	if len(slice) <= 1 {
		return slice
	}

	mit := len(slice) / 2
	izq := mergeSortVotantes(slice[:mit])
	der := mergeSortVotantes(slice[mit:])

	return merge(izq, der)
}

// Junta las mitades en un slice pero de forma ordenada
func merge(izq, der []votos.Votante) []votos.Votante {
	sliceFinal := make([]votos.Votante, 0, len(izq)+len(der))
	//Le fijamos la cap para que no tenga que redimensionar

	i, d := 0, 0 //izquierda y derecha

	//Mientras no se acabe ninguna de las mitades
	for i < len(izq) && d < len(der) {
		if izq[i].LeerDNI() < der[d].LeerDNI() {
			sliceFinal = append(sliceFinal, izq[i])
			i++
		} else {
			sliceFinal = append(sliceFinal, der[d])
			d++
		}
	}

	//Inserta todos los elementos que faltaran, si es que faltaba alguno
	for j := i; j < len(izq); j++ {
		sliceFinal = append(sliceFinal, izq[j])
	}
	for j := d; j < len(der); j++ {
		sliceFinal = append(sliceFinal, der[j])
	}

	return sliceFinal
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