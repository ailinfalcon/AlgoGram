package TDAalgogram

type AlgoGram interface {
	///
	Login(nombre string) error

	//
	Logout() error

	//
	PublicarPost(contenido string) error

	//
	VerProximoPost() error

	//
	LikearPost(id int) error

	//
	MostrarLikes(id int) error
}
