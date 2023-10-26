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
	pilaVotosPresidente      pila.Pila[[]int]
	pilaVotosGobernador      pila.Pila[[]int]
	pilaVotosIntendente      pila.Pila[[]int]
	ultimoQueRealizo  []TipoVoto
}

func CrearVotante(dni int) Votante {
	pilaVotosPresidente := pila.CrearPilaDinamica[[]int]()
	pilaVotosGobernador := pila.CrearPilaDinamica[[]int]()
	pilaVotosIntendente := pila.CrearPilaDinamica[[]int]()
	ultimoQueRealizo :=  []TipoVoto{}
	return &votanteImplementacion{dni, false, 0, false, pilaVotosPresidente,pilaVotosGobernador,pilaVotosIntendente,ultimoQueRealizo}
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
	switch tipo {
	case PRESIDENTE:
		votante.pilaVotosPresidente.Apilar(votos)
		votante.ultimoQueRealizo = append(votante.ultimoQueRealizo, PRESIDENTE)
	case GOBERNADOR:
		votante.pilaVotosGobernador.Apilar(votos)
		votante.ultimoQueRealizo = append(votante.ultimoQueRealizo, GOBERNADOR)
	case INTENDENTE:
		votante.pilaVotosIntendente.Apilar(votos)
		votante.ultimoQueRealizo = append(votante.ultimoQueRealizo, INTENDENTE)
	}
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
	var voto []int
	
	switch votante.ultimoQueRealizo[votante.iteraciones-1] {
		case PRESIDENTE:
			voto = votante.pilaVotosPresidente.Desapilar()
		case GOBERNADOR:
			voto = votante.pilaVotosGobernador.Desapilar()
		case INTENDENTE:
			voto = votante.pilaVotosIntendente.Desapilar()
	}

	votante.ultimoQueRealizo = votante.ultimoQueRealizo[:len(votante.ultimoQueRealizo)-1]
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
		votante.pilaVotosPresidente = VaciarPila(votante.pilaVotosPresidente)
		votante.pilaVotosGobernador = VaciarPila(votante.pilaVotosGobernador)
		votante.pilaVotosIntendente = VaciarPila(votante.pilaVotosIntendente)
		votante.ultimoQueRealizo = 	votante.ultimoQueRealizo[:0]
		return Voto{[CANT_VOTACION]int{0, 0, 0}, votante.impugnado}, nil
	}
	if !votante.pilaVotosPresidente.EstaVacia(){
		voto[PRESIDENTE] = votante.pilaVotosPresidente.Desapilar()[1]
	}
	if !votante.pilaVotosGobernador.EstaVacia(){
		voto[GOBERNADOR] = votante.pilaVotosGobernador.Desapilar()[1]
	}
	if !votante.pilaVotosIntendente.EstaVacia(){
		voto[INTENDENTE] = votante.pilaVotosIntendente.Desapilar()[1]
	}

	votante.votoFinalizado = true
	votante.pilaVotosPresidente = VaciarPila(votante.pilaVotosPresidente)
	votante.pilaVotosGobernador = VaciarPila(votante.pilaVotosGobernador)
	votante.pilaVotosIntendente = VaciarPila(votante.pilaVotosIntendente)
	votante.ultimoQueRealizo = 	votante.ultimoQueRealizo[:0]
	return Voto{voto, votante.impugnado}, nil
}
