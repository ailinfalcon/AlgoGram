package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al comienzo de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del último de la lista.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista.
	Largo() int

	// Iterar recorre todos los elementos de la lista, aplicando a cada uno la función visitar
	// hasta que esta devuelva false o se termine la lista.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador externo para recorrer la lista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual obtiene el valor del elemento en la posición en la que se encuentra el iterador.
	// Si ya iteró todos los elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente indica si todavía quedan elementos por iterar.
	HaySiguiente() bool

	// Siguiente avanza el iterador a la siguiente posición de la lista.
	// Si ya iteró todos los elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar agrega un nuevo elemento en la posición donde se encuentra el iterador.
	Insertar(T)

	// Borrar elimina el elemento en la posición donde se encuentra el iterador y devuelve su valor.
	// Si ya iteró todos los elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	Borrar() T
}
