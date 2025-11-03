package errores

type ErrorUsuarioLoggeado struct{}

func (e ErrorUsuarioLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioInexistente struct{}

func (e ErrorUsuarioInexistente) Error() string {
	return "Error: usuario no existente"
}

type ErrorUsuarioNoLoggeado struct{}

func (e ErrorUsuarioNoLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorVerProximoPost struct{}

type ErrorMostrarLikes struct{}

func (e ErrorMostrarLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}

type ErrorLikearPost struct{}

func (e ErrorLikearPost) Error() string {
	return "Error: usuario no loggeado o post inexistente"
}
