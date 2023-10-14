package cola

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{nil, nil}
}

func crearNodo[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{dato, nil}
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.

func (c *colaEnlazada[T]) Encolar(dato T) {
	nodo := crearNodo[T](dato)
	if c.EstaVacia() {
		c.primero = nodo
	} else {
		c.ultimo.prox = nodo
	}
	c.ultimo = nodo
}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := c.VerPrimero()
	c.primero = c.primero.prox
	if c.EstaVacia() {
		c.ultimo = nil
	}
	return dato
}
