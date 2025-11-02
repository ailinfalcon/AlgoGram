package diccionario_test

import (
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func igualdadIntsAbb(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func TestDiccionarioVacioAbb(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](igualdadIntsAbb)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(3))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(3) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(3) })
}

func TestUnElementoAbb(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardarAbb(t *testing.T) {
	claves := []int{1, 2, 3, 4}
	valoresPares := []int{2, 4, 6, 8}

	dic := TDADiccionario.CrearABB[int, int](igualdadIntsAbb)

	for i := range claves {
		require.False(t, dic.Pertenece(claves[i]))
		dic.Guardar(claves[i], valoresPares[i])
		require.EqualValues(t, i+1, dic.Cantidad())
	}

	for i := range claves {
		require.EqualValues(t, valoresPares[i], dic.Obtener(claves[i]))
	}
}

func TestReemplazoDatoAbb(t *testing.T) {
	claves := []string{"Alan", "Barbara", "Pedro"}

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	for i := range claves {
		dic.Guardar(claves[i], i+2)
		require.True(t, dic.Pertenece(claves[i]))
		require.EqualValues(t, i+2, dic.Obtener(claves[i]))
		require.EqualValues(t, i+1, dic.Cantidad())
	}

	dic.Guardar(claves[0], 1)
	dic.Guardar(claves[2], 2)

	require.EqualValues(t, 1, dic.Obtener(claves[0]))
	require.EqualValues(t, 2, dic.Obtener(claves[2]))
}

func TestDiccionarioBorrarAbb(t *testing.T) {
	clave := []int{2, 6, 7, 9, 11}
	valor := []int{3, 7, 8, 10, 12}
	dic := TDADiccionario.CrearABB[int, int](igualdadIntsAbb)

	for i := range clave {
		require.False(t, dic.Pertenece(clave[i]))
		dic.Guardar(clave[i], valor[i])
	}

	require.True(t, dic.Pertenece(clave[1]))
	require.EqualValues(t, valor[1], dic.Borrar(clave[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(clave[1]) })
	require.EqualValues(t, 4, dic.Cantidad())
	require.False(t, dic.Pertenece(clave[1]))

	require.True(t, dic.Pertenece(clave[0]))
	require.EqualValues(t, valor[0], dic.Borrar(clave[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(clave[0]) })
	require.EqualValues(t, 3, dic.Cantidad())
	require.False(t, dic.Pertenece(clave[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(clave[0]) })
}

func TestReutlizacionDeBorradosAbb(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	clave := "hola"
	dic.Guardar(clave, 2)
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, 4)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, 4, dic.Obtener(clave))
}

func TestBorrarRaizAbb(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](igualdadIntsAbb)
	dic.Guardar(4, "Gato")
	dic.Guardar(5, "Perro")
	dic.Guardar(10, "Tortuga")
	dic.Guardar(1, "Mapache")

	require.EqualValues(t, "Gato", dic.Borrar(4))
	require.EqualValues(t, 3, dic.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(4) })
	require.EqualValues(t, "Perro", dic.Obtener(5))
	require.EqualValues(t, "Tortuga", dic.Obtener(10))
	require.EqualValues(t, "Mapache", dic.Obtener(1))
}

func TestGuardarYBorrarRepetidasVecesAbb(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](igualdadIntsAbb)
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}

func TestVolumenAbb(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](igualdadIntsAbb)

	claves := make([]int, 1000)
	valores := make([]int, 1000)

	for i := range claves {
		valores[i] = (i * 2) + 2
		claves[i] = i
	}

	for i := range claves {
		j := rand.Intn(1 + i)
		claves[j], claves[i] = claves[i], claves[j]
	}

	for i, clave := range claves {
		dic.Guardar(clave, valores[i])
		require.EqualValues(t, valores[i], dic.Obtener(clave))
	}

	require.EqualValues(t, 1000, dic.Cantidad())

	for i, clave := range claves {
		require.EqualValues(t, valores[i], dic.Borrar(clave))
		require.False(t, dic.Pertenece(clave))
	}

}

/* PRUEBAS ITERADOR INTERNO */
func TestIteradorAbbSinCorte(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](igualdadIntsAbb)

	for i := range 10 {
		dic.Guardar(i, i)
	}

	var suma int
	sumaEsperada := 45

	dic.Iterar(func(clave, dato int) bool {
		suma += clave
		return true
	})

	require.Equal(t, sumaEsperada, suma)
}

func TestIteradorAbbConCorte(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](igualdadIntsAbb)
	slice := []int{2, 4, 6, 7, 8}

	for i, elem := range slice {
		dic.Guardar(i, elem)
	}

	var cantPar int
	cantEsperada := 3

	dic.Iterar(func(clave, dato int) bool {
		if dato%2 != 0 {
			return false
		}

		cantPar += 1

		return true
	})

	require.Equal(t, cantEsperada, cantPar)
}

func TestIteradorPorRangoAbbSinCorte(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](igualdadIntsAbb)

	for i := range 15 {
		dic.Guardar(i, i)
	}

	var suma int
	sumaEsperada := 88

	desde := 3
	hasta := 13
	dic.IterarRango(&desde, &hasta, func(clave, dato int) bool {
		suma += clave
		return true
	})

	require.Equal(t, sumaEsperada, suma)
}

func TestIteradorPorRangoAbbConCorte(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](igualdadIntsAbb)
	slice := []int{1, 3, 5, 7, 10, 12}

	for i, elem := range slice {
		dic.Guardar(i, elem)
	}

	var cantImpar int
	cantEsperada := 4

	dic.IterarRango(nil, nil, func(clave, dato int) bool {
		if dato%2 == 0 {
			return false
		}

		cantImpar += 1

		return true
	})

	require.Equal(t, cantEsperada, cantImpar)
}

func TestIteradorRangoAbbDesdeNil(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	claves := []string{"h", "a", "b", "k", "w"}

	for i, elem := range claves {
		dic.Guardar(elem, i+1)
	}

	ordenElem := []string{"a", "b", "h", "k", "w"}
	var res []string
	hasta := "w"

	dic.IterarRango(nil, &hasta, func(clave string, dato int) bool {
		res = append(res, clave)
		return true
	})

	require.Equal(t, ordenElem, res)
}

/* PRUEBAS ITERADOR EXTERNO */
func TestIterarDiccionarioVacioAbb(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterarDiccionarioInOrderAbb(t *testing.T) {
	clave1 := "Burro"
	clave2 := "Aguila"
	clave3 := "Gato"
	clave4 := "Perro"
	clave5 := "Vaca"

	claves := []string{clave1, clave2, clave3, clave4, clave5}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	for _, elem := range claves {
		dic.Guardar(elem, "animal")
	}

	clavesInOrder := []string{clave2, clave1, clave3, clave4, clave5}
	res := make([]string, len(clavesInOrder))

	i := 0
	for iter := dic.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		claveIter, _ := iter.VerActual()
		res[i] = claveIter
		i++
	}

	for i := range clavesInOrder {
		require.Equal(t, clavesInOrder[i], res[i])
	}

}

func TestIterarPorRangosDiccionarioVacioAbb(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](igualdadIntsAbb)
	desde := 10
	hasta := 20
	iter := dic.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterarPorRangosNilAbb(t *testing.T) {
	t.Log("Valida que se recorra todo el diccionario con el iterador externo si desde y hasta son nil")

	clave1 := 20
	clave2 := 30
	clave3 := 10
	clave4 := 5
	clave5 := 40

	claves := []int{clave1, clave2, clave3, clave4, clave5}
	dic := TDADiccionario.CrearABB[int, int](igualdadIntsAbb)

	for i, clave := range claves {
		dic.Guardar(clave, i)
	}

	clavesInOrder := []int{clave4, clave3, clave1, clave2, clave5}
	res := make([]int, len(clavesInOrder))

	i := 0
	for iter := dic.IteradorRango(nil, nil); iter.HaySiguiente(); iter.Siguiente() {
		claveIter, _ := iter.VerActual()
		res[i] = claveIter
		i++
	}

	require.Equal(t, clavesInOrder, res)
}

func TestIterarPorRangosAbb(t *testing.T) {
	t.Log("Verifica que se recorra todo el diccionario con el iterador externo en el rango indicado")

	clave1 := 20
	clave2 := 30
	clave3 := 10
	clave4 := 5
	clave5 := 40

	claves := []int{clave1, clave2, clave3, clave4, clave5}
	dic := TDADiccionario.CrearABB[int, int](igualdadIntsAbb)

	for i, clave := range claves {
		dic.Guardar(clave, i)
	}

	resEsperado := []int{clave3, clave1, clave2, clave5}
	res := make([]int, len(resEsperado))

	i := 0
	desde := 10
	hasta := 45
	for iter := dic.IteradorRango(&desde, &hasta); iter.HaySiguiente(); iter.Siguiente() {
		claveIter, _ := iter.VerActual()
		res[i] = claveIter
		i++
	}

	require.Equal(t, resEsperado, res)
}

func TestIterarFueraDeRangoAbb(t *testing.T) {
	t.Log("Verifica que si las claves del diccionario no pertenecen al rango, no haya nada que iterar")

	clave1 := "Suricata"
	clave2 := "Tortuga"
	clave3 := "Elefante"
	clave4 := "Jirafa"

	claves := []string{clave1, clave2, clave3, clave4}
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	for i, clave := range claves {
		dic.Guardar(clave, i)
	}

	desde := "Vaca"
	hasta := "Vibora"

	for iter := dic.IteradorRango(&desde, &hasta); iter.HaySiguiente(); iter.Siguiente() {
		require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	}
}
