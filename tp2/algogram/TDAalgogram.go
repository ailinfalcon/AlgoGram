package TDAalgogram

type AlgoGram interface {
	///
	Login(nombre string)

	//
	Logout()

	//
	PublicarPost(contenido string)

	//
	VerProximoPost() post

	//
	LikearPost(id int)

	//
	MostrarLikes(id int) ([]string, int)
}
