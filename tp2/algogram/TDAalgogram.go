package TDAalgogram

import TDAPost "tp2/post"

type AlgoGram interface {

	// HayLoggeado devuelve true si hay un usuario loggeado, o false en caso contrario.
	HayLoggeado() bool

	// Login
	Login(string) string

	// Logout desloggea al usuario actual, devuelve true si se desloggeo el usuario o false en caso contrario.
	Logout() bool

	// PublicarPost crea un post con el contenido recibido por parametro, devuelve true si se pudo publicar, o false si no hay un usuario loggeado.
	PublicarPost(string) bool

	// VerProximoPost devuelve el siguiente post a mostrar del feed del usuario loggeado.
	VerProximoPost() TDAPost.Post

	// LikearPost agrega el like al post según el id recibido por parametro, devuelve true si se agrego el like del usuario loggeado, false en caso que no haya un usuario loggeado o el post sea inexistente.
	LikearPost(int) bool

	// MostrarLikes devuelve un arreglo de string con los usuarios que le dieron like al post segun el id recibido y la cantidad de likes. Si el post es inexistente o no tiene likes, devuelve 0 y un arreglo vacio.
	MostrarLikes(int) ([]string, int)
}
