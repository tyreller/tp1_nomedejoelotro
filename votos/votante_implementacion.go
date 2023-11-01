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
	pilaVotos      pila.Pila[[]int]
	//La Lista aQuienVoto esta pensada para cuando deshace, saber de que pila borrar el elemento.
	listaAQuienVoto         []TipoVoto
	cantidadDeImpugnaciones int
}

func CrearVotante(dni int) Votante {

	pilaVotos := pila.CrearPilaDinamica[[]int]()
	listaInicial := []int{-1, -1, -1}
	pilaVotos.Apilar(listaInicial)
	listaAQuienVoto := []TipoVoto{}
	return &votanteImplementacion{dni, false, false, pilaVotos, listaAQuienVoto, 0}
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
		votante.cantidadDeImpugnaciones++
	}
	listaDeVotos := make([]int, CANT_VOTACION)
	copy(listaDeVotos, votante.pilaVotos.VerTope())
	listaDeVotos[tipo] = alternativa
	votante.listaAQuienVoto = append(votante.listaAQuienVoto, tipo)
	votante.pilaVotos.Apilar(listaDeVotos)
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

	votosDeshechos := votante.pilaVotos.Desapilar()

	votosActuales := votante.pilaVotos.VerTope()
	tipo := int(votante.listaAQuienVoto[len(votante.listaAQuienVoto)-1])
	votante.listaAQuienVoto = votante.listaAQuienVoto[:len(votante.listaAQuienVoto)-1]
	alternativa := votosDeshechos[tipo]
	if alternativa == 0 && votosDeshechos[tipo] != votosActuales[tipo] {
		votante.cantidadDeImpugnaciones--
	}
	if votante.cantidadDeImpugnaciones == 0 {
		votante.impugnado = false
	}

	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	voto := [CANT_VOTACION]int{0, 0, 0}
	votosDelVontante := votante.pilaVotos.VerTope()

	if votante.votoFinalizado {
		dni := votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return Voto{[CANT_VOTACION]int{0, 0, 0}, votante.impugnado}, err
	}
	if votante.impugnado {
		return Voto{[CANT_VOTACION]int{0, 0, 0}, votante.impugnado}, nil
	}
	voto[PRESIDENTE] = votosDelVontante[PRESIDENTE]
	voto[GOBERNADOR] = votosDelVontante[GOBERNADOR]
	voto[INTENDENTE] = votosDelVontante[INTENDENTE]
	votante.votoFinalizado = true

	return Voto{voto, votante.impugnado}, nil
}
