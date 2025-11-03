package algogram

type AlgoGram interface {
	///
	Login(nombre string) error

	//
	Logout() error

	//
	PublicarPost(contenido string) error

	//
	MostrarLikes(id int) error
}
