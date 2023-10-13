package votos
import "strconv"
type partidoImplementacion struct {
	nombre string
	NumeroLista int
	candidatos [CANT_VOTACION] string
	votos [CANT_VOTACION]int

}



type partidoEnBlanco struct {
	NumeroLista int
	candidatos [CANT_VOTACION] string
	votos [CANT_VOTACION]int
}

func CrearPartido(nombre string, numero int,candidatos [CANT_VOTACION]string, votos [CANT_VOTACION]int) Partido {
	return &partidoImplementacion{nombre,numero, candidatos,votos}
}

func CrearVotosEnBlanco() Partido {
	var listado = [3]string{"Presidente", "Gobernador", "Intendente"}
	var votos = [3]int{0,0,0}
	return &partidoEnBlanco{0,listado,votos}
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votos[tipo]++
	return
}


func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	stringReturn := partido.nombre + "- Postulante :" + partido.candidatos[tipo] + "-"+ strconv.Itoa(partido.votos[tipo]) + "votos"
	return stringReturn
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votos[tipo]++
	return
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	stringReturn := "Votos en Blanco: " + strconv.Itoa(partido.votos[tipo]) + "votos"
	return stringReturn
}
