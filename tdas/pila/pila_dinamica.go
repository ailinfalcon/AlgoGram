package pila

const (
	capacidadInicial         = 10
	coefCapacidadCrecimiento = 2
	coefCapacidadReduccion   = 4
	divisorReduccion         = 2
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, capacidadInicial)
	pila.cantidad = 0

	return pila
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) redimensionar(tam int) {
	nuevosDatos := make([]T, tam)
	copy(nuevosDatos, pila.datos)
	pila.datos = nuevosDatos
}

// Apilar agrega un nuevo elemento a la pila.
func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == len(pila.datos) {
		pila.redimensionar(len(pila.datos) * coefCapacidadCrecimiento)
	}

	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	elemTope := pila.VerTope()
	pila.cantidad--

	if (pila.cantidad*coefCapacidadReduccion) <= len(pila.datos) && len(pila.datos) > capacidadInicial {
		pila.redimensionar(len(pila.datos) / divisorReduccion)
	}

	return elemTope
}
