package cola_prioridad

const (
	_CAPACIDAD_INICIAL       = 10
	_COEFICIENTE_REDIMENSION = 2
	_COEFICIENTE_REDUCCION   = 4
	_DIVISOR_REDUCCION       = 2
)

type heap[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcionCmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = make([]T, _CAPACIDAD_INICIAL)
	heap.cmp = funcionCmp
	return heap
}

func CrearHeapArr[T any](arr []T, funcionCmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])

	if len(arr) < _CAPACIDAD_INICIAL {
		heap.datos = make([]T, _CAPACIDAD_INICIAL)
		copy(heap.datos, arr)
	} else {
		heap.datos = make([]T, len(arr))
		copy(heap.datos, arr)
	}

	heap.cant = len(arr)
	heap.cmp = funcionCmp

	heapify(heap.datos, heap.cant, heap.cmp)

	return heap
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *heap[T]) Encolar(elem T) {
	if heap.cant == len(heap.datos) {
		heap.redimensionar(len(heap.datos) * _COEFICIENTE_REDIMENSION)
	}

	heap.datos[heap.cant] = elem
	heap.cant++
	upheap(heap.datos, heap.cant-1, heap.cmp)
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}

	ultimo := heap.cant - 1

	swap(&heap.datos[ultimo], &heap.datos[0])
	heap.cant--

	downheap(heap.datos, 0, heap.cant, heap.cmp)

	if (heap.cant*_COEFICIENTE_REDUCCION) <= len(heap.datos) && len(heap.datos) > _CAPACIDAD_INICIAL {
		heap.redimensionar(len(heap.datos) / _DIVISOR_REDUCCION)
	}

	return heap.datos[ultimo]
}

func (heap *heap[T]) Cantidad() int {
	return heap.cant
}

func HeapSort[T any](datos []T, cmp func(T, T) int) {
	cant := len(datos)
	heapify(datos, cant, cmp)

	for i := cant - 1; i >= 0; i-- {
		swap(&datos[0], &datos[i])
		downheap(datos, 0, i, cmp)
	}
}

func heapify[T any](datos []T, cant int, cmp func(T, T) int) {
	for i := cant - 1; i >= 0; i-- {
		downheap(datos, i, cant, cmp)
	}
}

func (heap *heap[T]) redimensionar(tam int) {
	nuevosDatos := make([]T, tam)
	copy(nuevosDatos, heap.datos)
	heap.datos = nuevosDatos
}

func upheap[T any](datos []T, pos int, cmp func(T, T) int) {
	if pos == 0 {
		return
	}

	posPadre := (pos - 1) / 2

	if cmp(datos[pos], datos[posPadre]) > 0 {
		swap(&datos[pos], &datos[posPadre])
		upheap(datos, posPadre, cmp)
	}
}

func obtenerHijoMayor[T any](datos []T, izq, der int, cantidad int, cmp func(T, T) int) int {
	if izq >= cantidad {
		return -1
	}

	if der >= cantidad || cmp(datos[izq], datos[der]) > 0 {
		return izq
	}

	return der
}

func downheap[T any](datos []T, pos int, cant int, cmp func(T, T) int) {
	hijoIzq := (2 * pos) + 1
	hijoDer := (2 * pos) + 2

	posMayor := obtenerHijoMayor(datos, hijoIzq, hijoDer, cant, cmp)
	if posMayor == -1 {
		return
	}

	if cmp(datos[pos], datos[posMayor]) < 0 {
		swap(&datos[pos], &datos[posMayor])
		downheap(datos, posMayor, cant, cmp)
	}
}

func swap[T any](a *T, b *T) {
	*a, *b = *b, *a
}
