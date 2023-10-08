// Esta implementación de cola va a representar la fila para la maquina.

package votos

type Cola[T any] interface {

	// EstaVacia devuelve verdadero si la cola no tiene votantes, false en caso contrario.
	EstaVacia() bool

	// VerPrimero obtiene el valor del primer votante de la cola. Si está vacía, entra en pánico con un mensaje
	// "La fila esta vacia".
	VerPrimero() T

	// Encolar agrega un nuevo votante a la cola, al final de la misma.
	Encolar(T)

	// Desencolar saca el primer votante de la cola. Si la cola tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
	Desencolar() T
}