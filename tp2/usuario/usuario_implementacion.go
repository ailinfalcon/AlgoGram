package usuario

import (
	TDAHeap "tdas/cola_prioridad"
)

type usuario struct {
	nombre   string
	feed     TDAHeap.ColaPrioridad[*postFeed]
	afinidad int
}

func CrearUsuario(nombreUsuario string, afinidadUsuario int) *usuario {
	return &usuario{
		nombre: nombreUsuario,
		feed:,
		afinidad: afinidadUsuario,
	}
}

func (usuario *usuario) ObtenerAfinidad() int {
	return usuario.afinidad
}

func (usuario *usuario) ObtenerNombre() string {
	return usuario.nombre
}