package TDAalgogram

type AlgoGram interface {

	///
	HayLoggeado() bool

	///
	Login(string) string

	//
	Logout()

	//
	PublicarPost(string)

	//
	VerProximoPost() (int, string, string, int)

	//
	LikearPost(int)

	//
	MostrarLikes(int) ([]string, int)
}
