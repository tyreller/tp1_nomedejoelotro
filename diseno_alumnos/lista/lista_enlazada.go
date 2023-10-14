package lista

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterador[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

// CrearListaEnlazada crea una nueva lista enlazada vacia.
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{nil, nil, 0}
}

// crearNodo Crea un nuevo nodo con un dato.
func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato, nil}
}

// Largo devuelve la cantidad de elementos de la lista.
func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

// EstaVacia devuelve verdadero si la lista no tiene elementos encolados, false en caso contrario.
func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.Largo() == 0
}

// InsertarPrimero agrega un nuevo elemento a le lista en el primer lugar.

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := crearNodo(dato)
	if lista.EstaVacia() {
		lista.ultimo = nodo
	} else {
		nodo.prox = lista.primero
	}
	lista.primero = nodo
	lista.largo++
}

// InsertarUltimo agrega un nuevo elemento a le lista en el último lugar.
func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := crearNodo(dato)

	if lista.EstaVacia() {
		lista.primero = nodo
	} else {
		lista.ultimo.prox = nodo
	}
	lista.ultimo = nodo
	lista.largo++
}

// VerPrimero devuelve el valor del primer elemento de la lista. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

// VerUltimo devuelve el valor del ultimo elemento de la lista.Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

// BorrarPrimero borra el primer elemento de la lista y devuelve ese valor.  Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	primero := lista.primero
	if lista.primero.prox == nil {
		lista.ultimo = nil
		lista.primero = nil
	} else {
		lista.primero = primero.prox
	}

	lista.largo--

	return primero.dato
}

// Iterar pasa por cada uno de los elementos de la lista en orden hasta acabarla
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		continuar := visitar(actual.dato)
		if !continuar {
			break
		}
		actual = actual.prox
	}
}

// Iterador crea un iterador de la lista externo.
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterador[T]{lista, lista.primero, nil}
}

// VerActual devuelve el elemento en donde este posicionado el iterador.
func (iterador *iterador[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

// Siguiente modifica la posicion del iterador al siguiente elemento.
func (iterador *iterador[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.prox

}

// HaySiguiente devuelve verdadero si todavía hay algun elemento de la lista por ver, en caso contrario devuelve falso
func (iterador *iterador[T]) HaySiguiente() bool {
	if iterador.actual == nil && iterador.anterior != iterador.lista.ultimo {
		iterador.actual = iterador.lista.ultimo
		return true
	}
	return iterador.actual != nil
}

// Insertar inserta un elemento a la lista en donde este posicionado el iterador.
func (iterador *iterador[T]) Insertar(dato T) {
	nodo := crearNodo(dato)
	nodo.prox = iterador.actual
	if iterador.anterior == nil {
		iterador.lista.primero = nodo
	} else {
		iterador.anterior.prox = nodo
	}
	if !iterador.HaySiguiente() {
		iterador.lista.ultimo = nodo
	}
	iterador.actual = nodo
	iterador.lista.largo++
}

// Borrar borra el elemento de la lista en donde este posicionado el iterador, y ademas devuelve el valor de ese elemento.
func (iterador *iterador[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := iterador.VerActual()

	if iterador.anterior == nil { //Primer Elemento
		iterador.lista.primero = iterador.actual.prox
	}
	if iterador.actual.prox == nil { //Ultimo Elemento
		iterador.lista.ultimo = iterador.anterior
	}

	if iterador.anterior != nil { //Cualquier Elemento excepto el primero
		iterador.anterior.prox = iterador.actual.prox
	}
	iterador.actual = iterador.actual.prox

	iterador.lista.largo--
	return dato
}
