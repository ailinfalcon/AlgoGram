package TDAalgogram

type AlgoGram interface {

	///
	HayLoggeado() bool

	///
	Login(string) string

	//
	Logout() bool

	//
	PublicarPost(string) bool

	//
	VerProximoPost() (int, string, string, int)

	//
	LikearPost(int) bool

	//
	MostrarLikes(int) ([]string, int)
}
