package votos

import (
	"rerepolez/errores"
	"tdas/pila"
)

var VotosImpugnados = 0

// Implementacion de pila para guardar los votos del votante.
type votanteImplementacion struct {
	dni            int
	votoFinalizado bool
	iteraciones    int
	impugnado      bool
}

var pilaVotosTipo pila.Pila[TipoVoto]
var pilaVotosAlternativa pila.Pila[int]

func inicializarPilas() {
	pilaVotosTipo = pila.CrearPilaDinamica[TipoVoto]()
	pilaVotosAlternativa = pila.CrearPilaDinamica[int]()
}

func CrearVotante(dni int) Votante {
	inicializarPilas()
	return &votanteImplementacion{dni, false, 0, false}
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
		VotosImpugnados++
	}
	pilaVotosTipo.Apilar(tipo)
	pilaVotosAlternativa.Apilar(alternativa)
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
	alternativa := pilaVotosAlternativa.Desapilar()
	if alternativa == 0 {
		votante.impugnado = false
		VotosImpugnados--
	}
	pilaVotosTipo.Desapilar()
	votante.iteraciones--

	return nil
}

func VaciarPilas() {
	for !pilaVotosTipo.EstaVacia() {
		pilaVotosTipo.Desapilar()
	}
	for !pilaVotosAlternativa.EstaVacia() {
		pilaVotosAlternativa.Desapilar()
	}
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	voto := [CANT_VOTACION]int{0, 0, 0}
	if votante.votoFinalizado {
		dni := votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return Voto{[CANT_VOTACION]int{0, 0, 0}, votante.impugnado}, err
	}
	if votante.impugnado {
		VaciarPilas()
		return Voto{[CANT_VOTACION]int{0, 0, 0}, votante.impugnado}, nil
	}
	contadorPresidente := 0
	contadorGobernador := 0
	contadorIntendente := 0
	for !pilaVotosTipo.EstaVacia() && (contadorGobernador != 1 || contadorPresidente != 1 || contadorIntendente != 1) {
		tipo := pilaVotosTipo.Desapilar()
		alternativa := pilaVotosAlternativa.Desapilar()
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
	VaciarPilas()
	return Voto{voto, votante.impugnado}, nil
}
