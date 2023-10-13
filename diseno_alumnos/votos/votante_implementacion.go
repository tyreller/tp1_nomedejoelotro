package votos

import (
	"rerepolez/errores"
)


type votanteImplementacion struct {
	dni int
	votoFinalizado bool
	votos[] int 
}


func CrearVotante(dni int) Votante {
	return &votanteImplementacion{dni, false,[]int{} }
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
		if  votante.votoFinalizado{
			dni:= votante.dni
			err := errores.ErrorVotanteFraudulento{dni}
	
			return err
		}
		votante.votos = append(votante.votos,alternativa)
		return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if len(votante.votos) == 0{
		err := errores.ErrorNoHayVotosAnteriores{}
		return err
	}else if  votante.votoFinalizado{
		dni:= votante.dni
		err := errores.ErrorVotanteFraudulento{dni}
		return err
	}
	votante.votos = votante.votos[:len(votante.votos)-1]
		return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	return Voto{}, nil
}

