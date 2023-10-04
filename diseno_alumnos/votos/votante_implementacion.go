package votos

type votanteImplementacion struct {
	dni int
}



func CrearVotante(dni int) Votante {
	return &votanteImplementacion{dni,/*FUNCION PARA OBTENER EL SIGUIENTE LUGAR*/}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if *NO ERROR*
		return nil

	*VOTAR*
}

func (votante *votanteImplementacion) Deshacer() error {
	if *HAY ALGO QUE DESHACER*
		*DESHACER*
		return nil
	
	return *HUBO ERROR*
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	return Voto{}, nil
}
