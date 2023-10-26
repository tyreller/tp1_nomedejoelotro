package funciones

import (
	"bufio"
	"os"
	"strings"
)

func CrearEscaner() *bufio.Scanner {
	escanerInput := bufio.NewScanner(os.Stdin)
	return escanerInput
}
func ObtenerParametroComando(escaner *bufio.Scanner) ([]string, string){
	parametro := escaner.Text()
	parametroSeparado := strings.Fields(parametro)
	comando := parametroSeparado[0]
	return parametroSeparado,comando
}