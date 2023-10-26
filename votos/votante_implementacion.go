package votos

import (
	"rerepolez/errores"
	"tdas/pila"
)

// Implementacion de pila para guardar los votos del votante.
type votanteImplementacion struct {
	dni            int
	votoFinalizado bool
	impugnado      bool
	//Implementamos 3 pilas asi la funcion Fin-Votar es de tiempo constante.
	pilaVotosPresidente pila.Pila[[]int]
	pilaVotosGobernador pila.Pila[[]int]
	pilaVotosIntendente pila.Pila[[]int]
	//La Lista aQuienVoto esta pensada para cuando deshace, saber de que pila borrar el elemento.
	listaAQuienVoto []TipoVoto
}

func CrearVotante(dni int) Votante {
	pilaVotosPresidente := pila.CrearPilaDinamica[[]int]()
	pilaVotosGobernador := pila.CrearPilaDinamica[[]int]()
	pilaVotosIntendente := pila.CrearPilaDinamica[[]int]()
	listaAQuienVoto := []TipoVoto{}
	return &votanteImplementacion{dni, false, false, pilaVotosPresidente, pilaVotosGobernador, pilaVotosIntendente, listaAQuienVoto}
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
		votante.listaAQuienVoto = append(votante.listaAQuienVoto, PRESIDENTE)
	case GOBERNADOR:
		votante.pilaVotosGobernador.Apilar(votos)
		votante.listaAQuienVoto = append(votante.listaAQuienVoto, GOBERNADOR)
	case INTENDENTE:
		votante.pilaVotosIntendente.Apilar(votos)
		votante.listaAQuienVoto = append(votante.listaAQuienVoto, INTENDENTE)
	}
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.votoFinalizado {
		dni := votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return err
	} else if len(votante.listaAQuienVoto) == 0 {
		err := errores.ErrorNoHayVotosAnteriores{}
		return err
	}
	var voto []int

	switch votante.listaAQuienVoto[len(votante.listaAQuienVoto)-1] {
	case PRESIDENTE:
		voto = votante.pilaVotosPresidente.Desapilar()
	case GOBERNADOR:
		voto = votante.pilaVotosGobernador.Desapilar()
	case INTENDENTE:
		voto = votante.pilaVotosIntendente.Desapilar()
	}

	votante.listaAQuienVoto = votante.listaAQuienVoto[:len(votante.listaAQuienVoto)-1]
	alternativa := voto[1]
	if alternativa == 0 {
		votante.impugnado = false
	}
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	voto := [CANT_VOTACION]int{0, 0, 0}
	
	if votante.votoFinalizado {
		dni := votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return Voto{[CANT_VOTACION]int{0, 0, 0}, votante.impugnado}, err
	}

	if votante.impugnado {
		return Voto{[CANT_VOTACION]int{0, 0, 0}, votante.impugnado}, nil
	}
	if !votante.pilaVotosPresidente.EstaVacia() {
		voto[PRESIDENTE] = votante.pilaVotosPresidente.Desapilar()[1]
	}
	if !votante.pilaVotosGobernador.EstaVacia() {
		voto[GOBERNADOR] = votante.pilaVotosGobernador.Desapilar()[1]
	}
	if !votante.pilaVotosIntendente.EstaVacia() {
		voto[INTENDENTE] = votante.pilaVotosIntendente.Desapilar()[1]
	}
	votante.votoFinalizado = true
	return Voto{voto, votante.impugnado}, nil
}
