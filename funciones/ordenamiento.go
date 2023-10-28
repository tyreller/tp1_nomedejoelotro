package funciones

import (
	"rerepolez/votos"
)

func radixSortVotantes(slice []votos.Votante) {
	max := obtenerMaxVotante(slice)
	div := 1

	for max/div > 0 {
		countingSortVotantes(slice, div)
		div *= 10
	}
}

func countingSortVotantes(slice []votos.Votante, div int) {
	aux := make([]votos.Votante, len(slice))
	countDigits := make([]int, 10)

	for i := 0; i < len(slice); i++ {
		countDigits[obtenerDigito(slice[i], div)]++
	}

	for i := 1; i < 10; i++ {
		countDigits[i] += countDigits[i-1]
	}

	for i := len(slice) - 1; i >= 0; i-- {
		digit := (slice[i].LeerDNI() / div) % 10
		aux[countDigits[digit]-1] = slice[i]
		countDigits[digit]--
	}

	copy(slice, aux)
}
