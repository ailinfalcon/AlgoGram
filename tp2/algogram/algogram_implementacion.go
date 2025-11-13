package TDAalgogram

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
	TDAPost "tp2/post"
	TDAUsuario "tp2/usuario"
)

type Algogram struct {
	usuarios        TDADiccionario.Diccionario[string, TDAUsuario.Usuario]
	usuarioLoggeado TDAUsuario.Usuario
	posts           TDALista.Lista[TDAPost.Post]
}

func CrearAlgogram(us TDADiccionario.Diccionario[string, int]) AlgoGram {
	usuarios := cargarUsuarios(us)

	return &Algogram{
		usuarioLoggeado: nil,
		usuarios:        usuarios,
		posts:           TDALista.CrearListaEnlazada[TDAPost.Post](),
	}
}

func (algogram *Algogram) HayLoggeado() bool {
	return algogram.usuarioLoggeado != nil
}

func (algogram *Algogram) Login(nombre string) string {
	if !algogram.usuarios.Pertenece(nombre) {
		fmt.Printf("Error: usuario no existente\n")
		return ""
	}

	if algogram.HayLoggeado() {
		fmt.Printf("Error: Ya habia un usuario loggeado\n")
		return ""
	}

	algogram.loggearUsuario(nombre)
	return nombre
}

func (algogram *Algogram) Logout() bool {
	if !algogram.HayLoggeado() {
		fmt.Printf("Error: no habia usuario loggeado\n")
		return false
	}

	algogram.desloggearUsuario()

	return true
}

func (algogram *Algogram) PublicarPost(contenido string) bool {
	if !algogram.HayLoggeado() {
		fmt.Printf("Error: no habia usuario loggeado\n")
		return false
	}

	post := TDAPost.CrearPost(algogram.usuarioLoggeado.ObtenerNombre(), contenido, algogram.posts.Largo())
	algogram.posts.InsertarUltimo(post)

	iter := algogram.usuarios.Iterador()
	for iter.HaySiguiente() {
		nombre, usuario := iter.VerActual()
		if nombre != algogram.usuarioLoggeado.ObtenerNombre() {
			usuario.AgregarPostFeed(post, algogram.usuarioLoggeado.ObtenerAfinidad())
		}
		iter.Siguiente()
	}

	return true
}

func (algogram *Algogram) VerProximoPost() TDAPost.Post {
	if !algogram.HayLoggeado() || !algogram.usuarioLoggeado.TienePostFeed() {
		fmt.Printf("Usuario no loggeado o no hay mas posts para ver\n")
		return nil
	}

	postFeed := algogram.usuarioLoggeado.ObtenerPostFeed()
	return postFeed
}

func (algogram *Algogram) LikearPost(id int) bool {
	if !algogram.HayLoggeado() || id >= algogram.posts.Largo() || id < 0 {
		fmt.Printf("Error: Usuario no loggeado o Post inexistente\n")
		return false
	}

	iter := algogram.posts.Iterador()

	for i := 0; i < id; i++ {
		iter.Siguiente()
	}

	postActual := iter.VerActual()
	postActual.AgregarLike(algogram.usuarioLoggeado.ObtenerNombre(), algogram.usuarioLoggeado.ObtenerNombre())

	return true
}

func (algogram *Algogram) MostrarLikes(id int) ([]string, int) {
	var likes []string
	if id >= algogram.posts.Largo() || id < 0 {
		fmt.Printf("Error: Post inexistente o sin likes\n")
		return likes, 0
	}

	iter := algogram.posts.Iterador()
	for i := 0; i < id; i++ {
		iter.Siguiente()
	}

	postActual := iter.VerActual()
	if postActual.ObtenerCantLikes() == 0 {
		fmt.Printf("Error: Post inexistente o sin likes\n")
		return likes, 0
	}

	iterLikes := postActual.ObtenerLikes().Iterador()
	for iterLikes.HaySiguiente() {
		nombreUsuario, _ := iterLikes.VerActual()
		likes = append(likes, nombreUsuario)
		iterLikes.Siguiente()
	}

	return likes, postActual.ObtenerCantLikes()
}

func (algogram *Algogram) loggearUsuario(nombre string) {
	usuario := algogram.usuarios.Obtener(nombre)
	algogram.usuarioLoggeado = usuario
}

func (algogram *Algogram) desloggearUsuario() {
	algogram.usuarioLoggeado = nil
}

func cargarUsuarios(us TDADiccionario.Diccionario[string, int]) TDADiccionario.Diccionario[string, TDAUsuario.Usuario] {
	usuarios := TDADiccionario.CrearHash[string, TDAUsuario.Usuario](func(a, b string) bool { return a == b })

	iter := us.Iterador()

	for iter.HaySiguiente() {
		nombre, afinidad := iter.VerActual()
		usuarioCreado := TDAUsuario.CrearUsuario(nombre, afinidad)
		usuarios.Guardar(nombre, usuarioCreado)

		iter.Siguiente()
	}

	return usuarios
}
