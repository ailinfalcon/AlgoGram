package pila_test

import (
	"strconv"
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	pila.Apilar(10)
	require.False(t, pila.EstaVacia())

	pila.Desapilar()
	require.True(t, pila.EstaVacia())
}

func TestComportamientoLiFo(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	slice := []int{10, 7, 2, 12, 30}

	for _, valor := range slice {
		pila.Apilar(valor)
		require.Equal(t, valor, pila.VerTope())
	}

	require.False(t, pila.EstaVacia())

	for i := len(slice) - 1; i >= 0; i-- {
		require.Equal(t, slice[i], pila.VerTope())
		pila.Desapilar()
	}

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.True(t, pila.EstaVacia())
}

func TestApilarElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	slice1 := []int{2, 3, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	for _, valor := range slice1 {
		pila.Apilar(valor)
		require.Equal(t, valor, pila.VerTope())
	}

	require.False(t, pila.EstaVacia())

	for !pila.EstaVacia() {
		pila.Desapilar()
	}

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())

	slice2 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}
	for _, valor := range slice2 {
		pila.Apilar(valor)
		require.Equal(t, valor, pila.VerTope())
	}
}

func TestRedimension(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	for i := 1; i <= 10000; i++ {
		pila.Apilar("num: " + strconv.Itoa(i))
		require.Equal(t, "num: "+strconv.Itoa(i), pila.VerTope())
	}

	for i := 10000; i >= 1; i-- {
		require.Equal(t, "num: "+strconv.Itoa(i), pila.VerTope())
		pila.Desapilar()
	}

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}
