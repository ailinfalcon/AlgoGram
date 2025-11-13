package TDAalgogram

import TDAPost "tp2/post"

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
	VerProximoPost() TDAPost.Post

	//
	LikearPost(int) bool

	//
	MostrarLikes(int) ([]string, int)
}
