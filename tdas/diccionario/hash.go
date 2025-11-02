package diccionario

import "fmt"

type estadoClave int

const (
	tamanioInicial         = 11
	factorDeCargaAgrandar  = 0.7
	factorDeCargaAchicar   = 0.3
	coeficienteRedimension = 2
)

const (
	VACIO = estadoClave(iota)
	BORRADO
	OCUPADO
)

type celdaHash[K any, V any] struct {
	clave  K
	dato   V
	estado estadoClave
}

type hashCerrado[K any, V any] struct {
	tabla         []celdaHash[K, V]
	cantidad      int
	tam           int
	borrados      int
	clavesIguales func(K, K) bool
}

type iterHash[K any, V any] struct {
	hash   *hashCerrado[K, V]
	actual int
}

/*  FUNCIÓN DE HASH */

const (
	uint64Offset uint64 = 0xcbf29ce484222325
	uint64Prime  uint64 = 0x00000100000001b3
)

func fvnHash(data []byte) (hash uint64) {
	hash = uint64Offset

	for _, b := range data {
		hash ^= uint64(b)
		hash *= uint64Prime
	}

	return
}

/* La función hash es de: https://golangprojectstructure.com/hash-functions-go-code/ */

func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func crearTabla[K any, V any](tam int) []celdaHash[K, V] {
	return make([]celdaHash[K, V], tam)
}

func CrearHash[K any, V any](comparar func(K, K) bool) Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.tabla = crearTabla[K, V](tamanioInicial)
	hash.cantidad = 0
	hash.tam = tamanioInicial
	hash.clavesIguales = comparar
	return hash
}

func abs(ind int) int {
	if ind < 0 {
		return -ind
	}

	return ind
}

func (hash *hashCerrado[K, V]) calcularIndice(clave K, tam int) int {
	return abs(int(fvnHash(convertirABytes(clave))) % tam)
}

func (hash *hashCerrado[K, V]) buscar(tabla []celdaHash[K, V], clave K) (int, bool) {
	tam := len(tabla)
	indice := hash.calcularIndice(clave, tam)

	for tabla[indice].estado != VACIO {
		if tabla[indice].estado == OCUPADO && hash.clavesIguales(tabla[indice].clave, clave) {
			return indice, true
		}

		indice = (indice + 1) % tam
	}

	return indice, false
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if float32(hash.cantidad+hash.borrados)/float32(hash.tam) > factorDeCargaAgrandar {
		hash.redimensionar(coeficienteRedimension * hash.tam)
	}

	ind, _ := hash.buscar(hash.tabla, clave)
	hash.tabla[ind].clave = clave
	hash.tabla[ind].dato = dato

	if hash.tabla[ind].estado == VACIO {
		hash.tabla[ind].estado = OCUPADO
		hash.cantidad++
	}
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	if len(hash.tabla) == 0 {
		return false
	}

	_, existe := hash.buscar(hash.tabla, clave)

	return existe
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	ind, existe := hash.buscar(hash.tabla, clave)

	if !existe {
		panic("La clave no pertenece al diccionario")
	}

	return hash.tabla[ind].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	ind, existe := hash.buscar(hash.tabla, clave)

	if !existe {
		panic("La clave no pertenece al diccionario")
	}

	dato := hash.tabla[ind].dato
	hash.tabla[ind].estado = BORRADO
	hash.cantidad--
	hash.borrados++

	if float32(hash.cantidad)/float32(hash.tam) < factorDeCargaAchicar {
		nuevoTam := hash.tam / coeficienteRedimension
		if nuevoTam < tamanioInicial {
			nuevoTam = tamanioInicial
		}

		hash.redimensionar(nuevoTam)
	}

	return dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) redimensionar(tam int) {
	tabla := hash.tabla

	hash.tabla = crearTabla[K, V](tam)
	hash.tam = tam
	hash.cantidad = 0
	hash.borrados = 0

	for i := range tabla {
		if tabla[i].estado == OCUPADO {
			nuevoIndice, _ := hash.buscar(hash.tabla, tabla[i].clave)
			hash.tabla[nuevoIndice].clave = tabla[i].clave
			hash.tabla[nuevoIndice].dato = tabla[i].dato
			hash.tabla[nuevoIndice].estado = OCUPADO
			hash.cantidad++
		}
	}
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(K, V) bool) {
	for i := range hash.tabla {
		celda := hash.tabla[i]
		if celda.estado == OCUPADO && !visitar(celda.clave, celda.dato) {
			return
		}
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterHash[K, V])
	iter.hash = hash
	iter.actual = 0

	for iter.actual < len(iter.hash.tabla) && iter.hash.tabla[iter.actual].estado != OCUPADO {
		iter.actual++
	}

	return iter
}

func (iter *iterHash[K, V]) HaySiguiente() bool {
	return iter.actual != iter.hash.tam
}

func (iter *iterHash[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	celda := iter.hash.tabla[iter.actual]

	return celda.clave, celda.dato
}

func (iter *iterHash[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	iter.actual++

	for iter.actual < len(iter.hash.tabla) && iter.hash.tabla[iter.actual].estado != OCUPADO {
		iter.actual++
	}
}
