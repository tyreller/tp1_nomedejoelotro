package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos encolados, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento a le lista en el primer lugar.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento a le lista en el último lugar.
	InsertarUltimo(T)

	// BorrarPrimero borra el primer elemento de la lista y devuelve ese valor.  Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero devuelve el valor del primer elemento de la lista. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo devuelve el valor del ultimo elemento de la lista.Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista.
	Largo() int

	// Iterar crea un iterador de la lista interno.
	Iterar(visitar func(T) bool)

	// Iterador crea un iterador de la lista externo.Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento en donde este posicionado el iterador.
	VerActual() T

	// HaySiguiente devuelve verdadero si todavía hay algun elemento de la lista por ver, en caso contrario devuelve falso.
	HaySiguiente() bool

	// Siguiente modifica la posicion del iterador al siguiente elemento.
	Siguiente()

	// Insertar inserta un elemento a la lista en donde este posicionado el iterador.
	Insertar(T)

	// Borrar borra el elemento de la lista en donde este posicionado el iterador, y ademas devuelve el valor de ese elemento.
	Borrar() T
}
