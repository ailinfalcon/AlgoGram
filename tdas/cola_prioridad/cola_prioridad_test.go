package cola_prioridad_test

import (
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func igualdadInts(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func cmpStringsDesc(a, b string) int {
	return strings.Compare(b, a)
}

func igualdadIntsHeapDeMin(a, b int) int {
	return igualdadInts(b, a)
}

func TestHeapVacio(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](igualdadInts)
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestCumplePrioridad(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](igualdadInts)

	require.True(t, heap.EstaVacia())
	heap.Encolar(18)
	heap.Encolar(99)
	heap.Encolar(2)
	heap.Encolar(23)
	heap.Encolar(101)

	arrEsperado := []int{101, 99, 23, 18, 2}

	for i := 0; i < len(arrEsperado); i++ {
		require.Equal(t, arrEsperado[i], heap.VerMax())
		heap.Desencolar()
	}

	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}
func TestVerMaxTrasEncolar(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[string](strings.Compare)
	heap.Encolar("Gato")
	require.EqualValues(t, "Gato", heap.VerMax())
	heap.Encolar("Perro")
	require.EqualValues(t, "Perro", heap.VerMax())
	heap.Encolar("Tortuga")
	require.EqualValues(t, "Tortuga", heap.VerMax())
}

func TestDesencolar(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](igualdadInts)
	heap.Encolar(30)
	heap.Encolar(100)
	heap.Encolar(15)

	require.EqualValues(t, 100, heap.Desencolar())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 30, heap.VerMax())

	require.EqualValues(t, 30, heap.Desencolar())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 15, heap.VerMax())

	heap.Encolar(5)
	heap.Encolar(200)
	require.EqualValues(t, 200, heap.VerMax())

	require.EqualValues(t, 200, heap.Desencolar())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 15, heap.VerMax())
}

func TestEncolarDesencolarRepetidamente(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](igualdadInts)
	for i := range 100 {
		heap.Encolar(i)
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i, heap.Desencolar())
		require.EqualValues(t, 0, heap.Cantidad())
	}
}

func TestVolumen(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](igualdadInts)
	for i := range 10000 {
		heap.Encolar(i)
		require.EqualValues(t, i, heap.VerMax())
	}
	for i := range 9999 {
		require.EqualValues(t, 9999-i, heap.Desencolar())
		require.EqualValues(t, 9999-1-i, heap.VerMax())
	}
	require.EqualValues(t, 0, heap.Desencolar())
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestVerMaxHeapDeMinimos(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](igualdadIntsHeapDeMin)
	heap.Encolar(10)
	require.EqualValues(t, 10, heap.VerMax())
	heap.Encolar(20)
	require.EqualValues(t, 10, heap.VerMax())
	heap.Encolar(2)
	require.EqualValues(t, 2, heap.VerMax())
	require.EqualValues(t, 3, heap.Cantidad())
}

func TestCrearHeapArr(t *testing.T) {
	arr := []int{5, 2, 9, 1}
	heap := TDAColaPrioridad.CrearHeapArr[int](arr, igualdadInts)
	require.EqualValues(t, 4, heap.Cantidad())

	require.EqualValues(t, 9, heap.VerMax())
	require.EqualValues(t, 9, heap.Desencolar())

	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, 5, heap.Desencolar())

	require.EqualValues(t, 2, heap.VerMax())
	require.EqualValues(t, 2, heap.Desencolar())

	require.EqualValues(t, 1, heap.VerMax())
	require.EqualValues(t, 1, heap.Desencolar())

	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSortAsc(t *testing.T) {
	arr := []int{4, 2, 7, 10, 4, 3, 8, 5}
	arrOrdenado := []int{2, 3, 4, 4, 5, 7, 8, 10}

	TDAColaPrioridad.HeapSort(arr, igualdadInts)

	for i := 0; i < len(arr); i++ {
		require.Equal(t, arrOrdenado[i], arr[i])
	}
}

func TestHeapSortDesc(t *testing.T) {
	arr := []string{"z", "a", "x", "y", "c", "m"}
	arrOrdenado := []string{"z", "y", "x", "m", "c", "a"}

	TDAColaPrioridad.HeapSort(arr, cmpStringsDesc)

	for i := 0; i < len(arr); i++ {
		require.Equal(t, arrOrdenado[i], arr[i])
	}
}
