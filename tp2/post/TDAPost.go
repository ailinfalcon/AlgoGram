package TDAPost

type Post interface {

    //
    ObtenerId() int

    //
    ObtenerAutor() string

    //
    ObtenerContenido() string

    //
    ObtenerLikes() []string

    //
    ObtenerCantLikes() int
}
