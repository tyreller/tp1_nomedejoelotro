package votos

type partidoImplementacion struct {
	nombre string
	NumeroLista int
	candidatos [] string
	tipoVoto TipoVoto
	cantidadVotos [3]TipoVoto
}

type partidoEnBlanco struct {
	tipoVoto TipoVoto
	cantidadVotos [3]TipoVoto
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	return &partidoImplementacion{nombre,0, candidatos,0,0,0}
}

func CrearVotosEnBlanco() Partido {
	return &partidoEnBlanco{0, 0}
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	*partido.tipo++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	stringReturn := "Votos del Partido " + partido.nombre + ".\n"
	stringReturn += stringReturn + "Resultados Presidente: " + partido.cantidadVotos[PRESIDENTE] + " votos.\n"
	stringReturn += stringReturn + "Resultados Gobernador: " + partido.cantidadVotos[GOBERNADOR] + " votos.\n"
	stringReturn += stringReturn + "Resultados Intendente: " + partido.cantidadVotos[INTENDENTE] + " votos.\n"
	return stringReturn
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	*blanco.cantidadVotos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	stringReturn := "Votos en Blanco " + partido.nombre + ".\n"
	stringReturn += stringReturn + "Resultados Presidente: " + blanco.cantidadVotos[PRESIDENTE] + " votos.\n"
	stringReturn += stringReturn + "Resultados Gobernador: " + blanco.cantidadVotos[GOBERNADOR] + " votos.\n"
	stringReturn += stringReturn + "Resultados Intendente: " + blanco.cantidadVotos[INTENDENTE] + " votos.\n"
	return ""
}
