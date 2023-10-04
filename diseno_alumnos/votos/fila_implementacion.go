package votos

/* Definición del struct pila proporcionado por la cátedra. */

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	votante *votanteImplementacion
	prox *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{nil, nil}
}

func crearNodo[T any](persona votanteImplementacion) *nodoCola[T] {
	return &nodoCola[T]{persona, nil}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.persona.dni
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nodo := crearNodo[T](dato)
	if c.EstaVacia() {
		c.primero = nodo
	} else {
		c.ultimo.prox = nodo
	}
	c.ultimo = nodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	dni := c.VerPrimero()
	c.primero = c.primero.prox
	if c.EstaVacia() {
		c.ultimo = nil
	}
	return dni
}