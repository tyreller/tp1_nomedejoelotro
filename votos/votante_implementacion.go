package votos

import (
	"rerepolez/errores"
	"tdas/pila"
)

// Implementacion de pila para guardar los votos del votante.
type votanteImplementacion struct {
	dni                  int
	votoFinalizado       bool
	iteraciones          int
	impugnado            bool
	pilaVotosTipo        pila.Pila[TipoVoto]
	pilaVotosAlternativa pila.Pila[int]
}

func CrearVotante(dni int) Votante {
	pilaVotosTipo := pila.CrearPilaDinamica[TipoVoto]()
	pilaVotosAlternativa := pila.CrearPilaDinamica[int]()
	return &votanteImplementacion{dni, false, 0, false, pilaVotosTipo, pilaVotosAlternativa}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {

	if votante.votoFinalizado {
		dni := votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return err
	}
	if alternativa == 0 {
		votante.impugnado = true
	}
	votante.pilaVotosTipo.Apilar(tipo)
	votante.pilaVotosAlternativa.Apilar(alternativa)
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
	alternativa := votante.pilaVotosAlternativa.Desapilar()
	if alternativa == 0 {
		votante.impugnado = false
	}
	votante.pilaVotosTipo.Desapilar()
	votante.iteraciones--

	return nil
}

func VaciarPilas(pilaVotosTipo pila.Pila[TipoVoto], pilaVotosAlternativa pila.Pila[int]) (pila.Pila[TipoVoto], pila.Pila[int]) {
	for !pilaVotosTipo.EstaVacia() {
		pilaVotosTipo.Desapilar()
	}
	for !pilaVotosAlternativa.EstaVacia() {
		pilaVotosAlternativa.Desapilar()
	}
	return pilaVotosTipo, pilaVotosAlternativa
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	voto := [CANT_VOTACION]int{0, 0, 0}
	if votante.votoFinalizado {
		dni := votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return Voto{[CANT_VOTACION]int{0, 0, 0}, votante.impugnado}, err
	}
	if votante.impugnado {
		votante.pilaVotosTipo, votante.pilaVotosAlternativa = VaciarPilas(votante.pilaVotosTipo, votante.pilaVotosAlternativa)
		return Voto{[CANT_VOTACION]int{0, 0, 0}, votante.impugnado}, nil
	}
	contadorPresidente := 0
	contadorGobernador := 0
	contadorIntendente := 0
	for !votante.pilaVotosTipo.EstaVacia() && (contadorGobernador != 1 || contadorPresidente != 1 || contadorIntendente != 1) {
		tipo := votante.pilaVotosTipo.Desapilar()
		alternativa := votante.pilaVotosAlternativa.Desapilar()
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
	votante.pilaVotosTipo, votante.pilaVotosAlternativa = VaciarPilas(votante.pilaVotosTipo, votante.pilaVotosAlternativa)
	return Voto{voto, votante.impugnado}, nil
}
