package votos

import (
	"rerepolez/errores"
	"tdas/pila"
)

// Implementacion de pila para guardar los votos del votante.
type votanteImplementacion struct {
	dni            int
	votoFinalizado bool
	iteraciones    int
	impugnado      bool
	pilaVotos      pila.Pila[[]int]
}

func CrearVotante(dni int) Votante {
	pilaVotos := pila.CrearPilaDinamica[[]int]()
	return &votanteImplementacion{dni, false, 0, false, pilaVotos}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	votos := []int{int(tipo), alternativa}
	if votante.votoFinalizado {
		dni := votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return err
	}
	if alternativa == 0 {
		votante.impugnado = true
	}
	votante.pilaVotos.Apilar(votos)
	votante.iteraciones++
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.votoFinalizado {
		dni := votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return err
	} else if votante.iteraciones == 0 {
		err := errores.ErrorNoHayVotosAnteriores{}
		return err
	}
	voto := votante.pilaVotos.Desapilar()

	alternativa := voto[1]
	if alternativa == 0 {
		votante.impugnado = false
	}
	votante.iteraciones--
	return nil
}

func VaciarPila(pilaVotos pila.Pila[[]int]) pila.Pila[[]int] {
	for !pilaVotos.EstaVacia() {
		pilaVotos.Desapilar()
	}
	return pilaVotos
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	voto := [CANT_VOTACION]int{0, 0, 0}
	if votante.votoFinalizado {
		dni := votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return Voto{[CANT_VOTACION]int{0, 0, 0}, votante.impugnado}, err
	}

	if votante.impugnado {
		votante.pilaVotos = VaciarPila(votante.pilaVotos)
		return Voto{[CANT_VOTACION]int{0, 0, 0}, votante.impugnado}, nil
	}
	contadorPresidente := 0
	contadorGobernador := 0
	contadorIntendente := 0

	for !votante.pilaVotos.EstaVacia() && (contadorGobernador != 1 || contadorPresidente != 1 || contadorIntendente != 1) {
		votos := votante.pilaVotos.Desapilar()
		alternativa := votos[1]
		tipo := TipoVoto(votos[0])

		if contadorPresidente == 0 && tipo == PRESIDENTE {
			contadorPresidente++
			voto[PRESIDENTE] = alternativa
		} else if contadorGobernador == 0 && tipo == GOBERNADOR {
			contadorGobernador++
			voto[GOBERNADOR] = alternativa
		} else if contadorIntendente == 0 && tipo == INTENDENTE {
			contadorIntendente++
			voto[INTENDENTE] = alternativa
		}
	}
	votante.votoFinalizado = true
	votante.pilaVotos = VaciarPila(votante.pilaVotos)
	return Voto{voto, votante.impugnado}, nil
}
