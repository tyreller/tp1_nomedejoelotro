package votos

import (
	"rerepolez/errores"
)


type votanteImplementacion struct {
	dni int
	votoFinalizado bool
	iteraciones int	
	votosLista[] int
	votosTipo[] TipoVoto
	impugnado bool
}


func CrearVotante(dni int) Votante {
	return &votanteImplementacion{dni, false,0,[]int{},[]TipoVoto{},false}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
		if  votante.votoFinalizado{
			dni := votante.dni
			err := errores.ErrorVotanteFraudulento{dni}
			return err
		}
		if alternativa == 0{
			votante.impugnado = true
		}
		votante.votosLista= append(votante.votosLista,alternativa)
		votante.votosTipo= append(votante.votosTipo,tipo)
		votante.iteraciones++
		return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.iteraciones == 0{
		err := errores.ErrorNoHayVotosAnteriores{}
		return err
	}else if  votante.votoFinalizado{
		dni:= votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return err
	}
	votante.votosLista = votante.votosLista[:len(votante.votosLista)-1]
	votante.votosTipo = votante.votosTipo[:len(votante.votosTipo)-1]

	votante.iteraciones --
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.votoFinalizado{
		dni:= votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return Voto{[CANT_VOTACION]int{0,0,0},votante.impugnado},err
	}
	votante.votoFinalizado = true
	
	return Voto{[CANT_VOTACION]int{0,0,0},votante.impugnado}, nil
}

