package lista_test

import (
	TDAlista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	RANGO = 100
	INT_1 = 1
	INT_2 = 2
)

func TestListaVacia(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[string]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())

	lista.InsertarPrimero("11")
	require.Equal(t, "11", lista.VerPrimero())
	lista.InsertarUltimo("23")
	require.Equal(t, "23", lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())

	lista.BorrarPrimero()
	require.Equal(t, "23", lista.VerPrimero())
	lista.BorrarPrimero()

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestInsertarYBorrarHastaVaciar(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := range RANGO {
		lista.InsertarUltimo(i)
	}
	for range RANGO {
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
}

func TestInsertarTrasBorrar(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := range RANGO {
		lista.InsertarUltimo(i)
	}
	for range RANGO {
		lista.BorrarPrimero()
	}
	for i := range RANGO {
		lista.InsertarPrimero(i)
	}
	require.Equal(t, RANGO-1, lista.VerPrimero())
}

func TestInsertarPrimero(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := range RANGO {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
	}
}

func TestInsertarUltimo(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := range RANGO {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.VerUltimo())
	}
}

func TestInsertarYBorrarRepetidamente(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := range RANGO {
		require.Equal(t, 0, lista.Largo())
		lista.InsertarUltimo(i)
		lista.BorrarPrimero()
		lista.InsertarPrimero(i)
		lista.BorrarPrimero()
		require.Equal(t, 0, lista.Largo())
	}
}

func TestUnElemento(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(INT_1)
	require.Equal(t, INT_1, lista.VerPrimero())
	require.Equal(t, INT_1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
}

// --- Test Iterador Externo
func TestIteradorListaVacia(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(INT_1)
	require.Equal(t, INT_1, lista.VerPrimero())

	iter.Borrar()
	require.Equal(t, 0, lista.Largo())
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIteradorInsertarPrimero(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("abcd")
	iter := lista.Iterador()

	iter.Insertar("efgh")
	require.Equal(t, "efgh", lista.VerPrimero())
}

func TestIteradorInsertaEnMedio(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(0)
	lista.InsertarUltimo(2)

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(1)
	require.Equal(t, 1, iter.VerActual())

	i := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		require.Equal(t, i, iter.VerActual())
		i++
	}
}

func TestIteradorInsertarUltimo(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[string]()
	slice := []string{"2", "4", "6", "8"}

	iter := lista.Iterador()
	for _, elem := range slice {
		iter.Insertar(elem)
	}

	for iter.HaySiguiente() {
		iter.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

	iter.Insertar("10")
	require.Equal(t, "10", lista.VerUltimo())
	require.Equal(t, 5, lista.Largo())
}

func TestIteradorBorrarPrimero(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	slice := []int{9, 11, 10}

	for _, elem := range slice {
		lista.InsertarPrimero(elem)
		require.Equal(t, elem, lista.VerPrimero())
	}

	iter := lista.Iterador()
	iter.Borrar()
	require.Equal(t, 11, lista.VerPrimero())
}

func TestIteradorBorrarMedio(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	slice := []int{1, 2, 4, 3, 4}

	for _, elem := range slice {
		lista.InsertarUltimo(elem)
	}

	iter := lista.Iterador()
	for i := 0; i < 2; i++ {
		iter.Siguiente()
	}

	iter.Borrar()

	i := 1
	for it := lista.Iterador(); it.HaySiguiente(); it.Siguiente() {
		require.Equal(t, i, it.VerActual())
		i++
	}
}

func TestIteradorBorraUltimo(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(INT_1)
	lista.InsertarUltimo(INT_2)
	iter := lista.Iterador()

	for iter.HaySiguiente() && iter.VerActual() != INT_2 {
		iter.Siguiente()
	}

	require.Equal(t, INT_2, iter.Borrar())
	require.Equal(t, INT_1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
}

func TestIteradorInsertarYBorrarMismaPosicion(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	for i := range 10 {
		iter.Insertar(i)
		require.Equal(t, i, iter.VerActual())
		iter.Borrar()
		require.True(t, lista.EstaVacia())
		require.False(t, iter.HaySiguiente())
	}
}

// Test iterador interno

func TestIteradorSinCorte(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	slice := []int{1, 2, 3, 4}

	for _, elem := range slice {
		lista.InsertarPrimero(elem)
	}

	var suma int
	sumaEsperada := 10

	lista.Iterar(func(elem int) bool {
		suma += elem
		return true
	})

	require.Equal(t, sumaEsperada, suma)
}

func TestIteradorConCorte(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	slice := []int{2, 4, 6, 7, 8}

	for _, elem := range slice {
		lista.InsertarUltimo(elem)
	}

	var cantPar int
	lista.Iterar(func(elem int) bool {
		if elem%2 != 0 {
			return false
		}

		cantPar += 1

		return true
	})

	require.Equal(t, 3, cantPar)
}
