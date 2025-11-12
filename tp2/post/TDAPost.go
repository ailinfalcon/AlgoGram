package TDAPost

import (
	TDADiccionario "tdas/diccionario"
	TDAUsuario "tp2/usuario"
)

type Post interface {

	//
	ObtenerId() int

	//
	ObtenerAutor() TDAUsuario.Usuario

	//
	ObtenerContenido() string

	//
	ObtenerLikes() TDADiccionario.DiccionarioOrdenado[string, TDAUsuario.Usuario]

	//
	ObtenerCantLikes() int

	//
	GuardarLike(string, TDAUsuario.Usuario)
}
