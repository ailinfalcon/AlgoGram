package algogram

type AlgoGram interface {
	///
	Login(nombre string) error

	//
	Logout() error

	//
	MostrarLikes(id int)
}
