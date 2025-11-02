package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	cola.Encolar(11)
	require.Equal(t, 11, cola.VerPrimero())
	require.False(t, cola.EstaVacia())

	require.Equal(t, 11, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestComportamientoFIFO(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	slice := []int{2, 4, 6, 8, 10, 12, 14, 18, 20, 22, 24, 26}

	for _, valor := range slice {
		cola.Encolar(valor)
	}

	require.False(t, cola.EstaVacia())

	for i := 0; i < len(slice); i++ {
		require.Equal(t, slice[i], cola.VerPrimero())
		cola.Desencolar()
	}

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}

func TestEncolarElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}

	for _, valor := range slice {
		cola.Encolar(valor)
	}

	for i := 0; i < len(slice); i++ {
		require.Equal(t, slice[i], cola.VerPrimero())
		cola.Desencolar()
	}

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.True(t, cola.EstaVacia())

	slice2 := []string{"j", "k", "l", "m"}
	for _, valor := range slice2 {
		cola.Encolar(valor)
	}

	for i := 0; i < len(slice2); i++ {
		require.Equal(t, slice2[i], cola.VerPrimero())
		cola.Desencolar()
	}
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	for i := 0; i < 10000; i++ {
		cola.Encolar(i)
	}

	for i := 0; i < 10000; i++ {
		require.Equal(t, i, cola.VerPrimero())
		cola.Desencolar()
	}

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}
