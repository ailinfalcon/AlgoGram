package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K any, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

type abb[K any, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iterAbb[K any, V any] struct {
	ab    *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func crearNodo[K any, V any](clave K, dato V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = dato
	return nodo
}

func CrearABB[K any, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.cmp = funcionCmp
	abb.cantidad = 0
	return abb
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	if ab.raiz == nil {
		ab.raiz = crearNodo(clave, dato)
		ab.cantidad++
	}

	nodoActual, nodoPadre := ab.buscar(ab.raiz, clave, nil)

	if nodoActual != nil {
		nodoActual.dato = dato
		return
	}

	nuevoNodo := crearNodo(clave, dato)

	if ab.cmp(clave, nodoPadre.clave) < 0 {
		nodoPadre.izq = nuevoNodo
	} else {
		nodoPadre.der = nuevoNodo
	}

	ab.cantidad++
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	nodoObtenido, _ := ab.buscar(ab.raiz, clave, nil)
	return nodoObtenido != nil
}

func (ab *abb[K, V]) Obtener(clave K) V {
	nodoObtenido, _ := ab.buscar(ab.raiz, clave, nil)

	if nodoObtenido == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodoObtenido.dato
}

func (ab *abb[K, V]) Borrar(clave K) V {
	nodoABorrar, padre := ab.buscar(ab.raiz, clave, nil)

	if nodoABorrar == nil {
		panic("La clave no pertenece al diccionario")
	}

	dato := nodoABorrar.dato

	if nodoABorrar.izq == nil || nodoABorrar.der == nil {
		var hijo *nodoAbb[K, V]
		if nodoABorrar.izq != nil {
			hijo = nodoABorrar.izq
		} else {
			hijo = nodoABorrar.der
		}
		ab.reemplazarNodo(padre, nodoABorrar, hijo)
		ab.cantidad--
		return dato
	}

	nodoReemplazo, padreReemplazo := ab.buscarReemplazo(nodoABorrar.izq, nodoABorrar)

	ab.reemplazarNodo(padreReemplazo, nodoReemplazo, nodoReemplazo.izq)
	nodoABorrar.clave = nodoReemplazo.clave
	nodoABorrar.dato = nodoReemplazo.dato
	ab.cantidad--

	return dato
}

func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	ab.raiz.iterar(visitar)
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return ab.IteradorRango(nil, nil)
}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if ab.raiz == nil {
		return
	}

	if desde == nil {
		primerNodo := ab.obtenerPrimero(ab.raiz)
		desde = &primerNodo.clave
	}

	if hasta == nil {
		ultimoNodo := ab.obtenerUltimo(ab.raiz)
		hasta = &ultimoNodo.clave
	}

	if !ab.iterarPorRango(ab.raiz, desde, hasta, visitar) {
		return
	}
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.ab = ab
	iter.desde = desde
	iter.hasta = hasta

	ab.apilarPorRango(iter, ab.raiz)

	return iter
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	nodoAnterior := iter.pila.Desapilar()
	iter.ab.apilarPorRango(iter, nodoAnterior.der)
}

func (ab *abb[K, V]) buscarReemplazo(nodo, padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo.der == nil {
		return nodo, padre
	}
	return ab.buscarReemplazo(nodo.der, nodo)
}

func (ab *abb[K, V]) buscar(nodoActual *nodoAbb[K, V], clave K, nodoPadre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodoActual == nil {
		return nil, nodoPadre
	}

	cmp := ab.cmp(clave, nodoActual.clave)
	if cmp == 0 {
		return nodoActual, nodoPadre
	} else if cmp < 0 {
		return ab.buscar(nodoActual.izq, clave, nodoActual)
	} else {
		return ab.buscar(nodoActual.der, clave, nodoActual)
	}
}

func (ab *abb[K, V]) reemplazarNodo(padre, nodoABorrar, hijo *nodoAbb[K, V]) {
	if padre == nil {
		ab.raiz = hijo
	} else if padre.izq == nodoABorrar {
		padre.izq = hijo
	} else {
		padre.der = hijo
	}
}

func (nodo *nodoAbb[K, V]) iterar(visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}

	if !nodo.izq.iterar(visitar) {
		return false
	}

	if !visitar(nodo.clave, nodo.dato) {
		return false
	}

	if !nodo.der.iterar(visitar) {
		return false
	}

	return true
}

func (ab *abb[K, V]) obtenerPrimero(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.izq == nil {
		return nodo
	}

	return ab.obtenerPrimero(nodo.izq)
}

func (ab *abb[K, V]) obtenerUltimo(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.der == nil {
		return nodo
	}

	return ab.obtenerUltimo(nodo.der)
}

func (ab *abb[K, V]) iterarPorRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}

	if ab.cmp(nodo.clave, *desde) > 0 {
		if !ab.iterarPorRango(nodo.izq, desde, hasta, visitar) {
			return false
		}
	}

	if ab.cmp(nodo.clave, *desde) >= 0 && (ab.cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}

	if ab.cmp(nodo.clave, *hasta) < 0 {
		if !ab.iterarPorRango(nodo.der, desde, hasta, visitar) {
			return false
		}
	}

	return true
}

func (ab *abb[K, V]) apilarPorRango(iter *iterAbb[K, V], nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if (iter.desde == nil || iter.ab.cmp(nodo.clave, *iter.desde) >= 0) &&
			(iter.hasta == nil || iter.ab.cmp(nodo.clave, *iter.hasta) <= 0) {
			iter.pila.Apilar(nodo)
			nodo = nodo.izq
		} else if iter.desde != nil && iter.ab.cmp(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.der
		} else {
			nodo = nodo.izq
		}
	}
}
