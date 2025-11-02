package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	cola.primero = nil
	cola.ultimo = nil

	return cola
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	return cola.primero.dato
}

func crearNuevoNodo[T any](elem T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = elem
	nodo.prox = nil

	return nodo
}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.
func (cola *colaEnlazada[T]) Encolar(elem T) {
	nuevoNodo := crearNuevoNodo(elem)

	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.prox = nuevoNodo
	}

	cola.ultimo = nuevoNodo
}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	primerElem := cola.primero
	cola.primero = cola.primero.prox

	return primerElem.dato
}
