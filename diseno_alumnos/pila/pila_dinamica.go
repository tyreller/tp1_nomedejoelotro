package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

// Variables de la pila:
const (
	tamInicial        = 2
	achicarCapacidad  = 0.5
	agrandarCapicidad = 2.0
	cantidadMinima    = 4.0
)

func CrearPilaDinamica[T any]() Pila[T] {
	cantidad := 0 /*Puse 0, porque cuando se crea la pila no tiene nada. Antes habia hecho que la funcion tenga como parametros a datos pero leyendo las faq, me di cuenta que no tengo que dar información a Barbara del struct*/
	datos := make([]T, tamInicial)

	pila := pilaDinamica[T]{datos, cantidad}
	return &pila
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (p *pilaDinamica[T]) EstaVacia() bool {
	if p.cantidad == 0 {
		return true
	}
	return false
}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) redimensionar(tam_nuevo float64) {
	tamañoEntero := int(tam_nuevo)
	nuevoTamDatos := make([]T, tamañoEntero)
	copy(nuevoTamDatos, p.datos)
	p.datos = nuevoTamDatos

}

// Apilar agrega un nuevo elemento a la pila.
func (p *pilaDinamica[T]) Apilar(elem T) {
	if p.cantidad == cap(p.datos) {
		capacidad := float64(cap(p.datos)) * agrandarCapicidad
		p.redimensionar(capacidad)
	}
	p.datos[p.cantidad] = elem
	p.cantidad++
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (p *pilaDinamica[T]) Desapilar() T {
	capacidad := float64(cap(p.datos))
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	if float64(p.cantidad)*cantidadMinima <= capacidad {
		capacidad = capacidad * achicarCapacidad
		p.redimensionar(capacidad)
	}
	ultimoDato := p.datos[p.cantidad-1]
	p.cantidad--
	return ultimoDato
}