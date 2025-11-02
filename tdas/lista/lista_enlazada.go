package lista

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	lista.primero = nil
	lista.ultimo = nil
	lista.largo = 0

	return lista
}

type iterListaEnlazada[T any] struct {
	lista  *listaEnlazada[T]
	actual *nodoLista[T]
	ant    *nodoLista[T]
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func crearNuevoNodo[T any](elem T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = elem
	nodo.prox = nil

	return nodo
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevoNodo := crearNuevoNodo(elem)

	if lista.EstaVacia() {
		lista.primero = nuevoNodo
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.prox = lista.primero
		lista.primero = nuevoNodo
	}

	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	nuevoNodo := crearNuevoNodo(elem)

	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.prox = nuevoNodo
	}

	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	primerElem := lista.primero
	lista.primero = primerElem.prox

	lista.largo--

	if lista.primero == nil {
		lista.ultimo = nil
	}

	return primerElem.dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero

	for actual != nil && visitar(actual.dato) {
		actual = actual.prox
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iterListaEnlazada[T])
	iter.lista = lista
	iter.actual = lista.primero
	iter.ant = nil

	return iter
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	iter.ant = iter.actual
	iter.actual = iter.actual.prox
}

func (iter *iterListaEnlazada[T]) Insertar(elem T) {
	nuevoNodo := crearNuevoNodo(elem)

	if !iter.HaySiguiente() {
		iter.lista.ultimo = nuevoNodo
	}

	if iter.ant == nil {
		iter.lista.primero = nuevoNodo
	} else {
		iter.ant.prox = nuevoNodo
	}

	nuevoNodo.prox = iter.actual
	iter.actual = nuevoNodo
	iter.lista.largo++
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	actual := iter.actual
	elem := actual.dato

	if iter.ant == nil {
		iter.lista.primero = actual.prox
	} else {
		iter.ant.prox = actual.prox
	}

	if actual == iter.lista.ultimo {
		iter.lista.ultimo = iter.ant
	}

	iter.actual = actual.prox
	iter.lista.largo--

	return elem
}
