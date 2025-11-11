package TDAalgogram

import (
	"fmt"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
)

type Algogram struct {
	usuarios        TDADiccionario.Diccionario[string, *usuario]
	usuarioLoggeado *usuario
	posts           TDALista.Lista[*post]
}

type usuario struct {
	nombre   string
	feed     TDAHeap.ColaPrioridad[*postFeed]
	afinidad int
}

type post struct {
	id         int
	publicador *usuario
	contenido  string
	likes      TDADiccionario.DiccionarioOrdenado[string, *usuario]
	cantLikes  int
}

type postFeed struct { // cambiar nombre
	post     *post
	afinidad int // afinidad del usuario al cual le pertenece el feed
}

func CrearAlgogram(us TDADiccionario.Diccionario[string, int]) *Algogram {
	usuarios := cargarUsuarios(us)

	return &Algogram{
		usuarioLoggeado: nil,
		usuarios:        usuarios,
		posts:           TDALista.CrearListaEnlazada[*post](),
	}
}

func cargarUsuarios(us TDADiccionario.Diccionario[string, int]) TDADiccionario.Diccionario[string, *usuario] {
	usuarios := TDADiccionario.CrearHash[string, *usuario](func(a, b string) bool { return a == b })

	iter := us.Iterador()

	for iter.HaySiguiente() {
		nombre, afinidad := iter.VerActual()
		usuarios.Guardar(nombre, crearUsuario(nombre, afinidad))

		iter.Siguiente()
	}

	return usuarios
}

func crearUsuario(nombre string, afinidad int) *usuario {
	usuario := new(usuario)
	usuario.nombre = nombre
	usuario.afinidad = afinidad
	usuario.feed = TDAHeap.CrearHeap[*postFeed](igualdadPostFeed) // Hay que hacer una funcion para que muestre por afinidad

	return usuario
}

func igualdadPostFeed(dato1, dato2 *postFeed) int {
	afinidad1 := abs(dato1.afinidad - dato1.post.publicador.afinidad) // es igual a dato2.afinidad
	afinidad2 := abs(dato2.afinidad - dato2.post.publicador.afinidad)
	res := afinidad2 - afinidad1

	if res == 0 {
		res = dato2.post.id - dato1.post.id
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

func (algogram *Algogram) loggearUsuario(nombre string) {
	usuario := algogram.usuarios.Obtener(nombre)
	algogram.usuarioLoggeado = usuario
}

func (algogram *Algogram) Logout() bool {
	if !algogram.HayLoggeado() {
		fmt.Printf("Error: no habia usuario loggeado\n")
		return false
	}

	algogram.desloggearUsuario()

	return true
}

func (algogram *Algogram) desloggearUsuario() {
	algogram.usuarioLoggeado = nil
}

func crearNuevoPost(u *usuario, contenido string, cant int) *post {
	nuevoPost := new(post)
	nuevoPost.id = cant
	nuevoPost.publicador = u
	nuevoPost.contenido = contenido
	nuevoPost.cantLikes = 0
	nuevoPost.likes = TDADiccionario.CrearABB[string, *usuario](strings.Compare)

	return nuevoPost
}

func crearNuevoPostFeed(post *post, afinidad int) *postFeed {
	nuevoPostFeed := new(postFeed)
	nuevoPostFeed.post = post
	nuevoPostFeed.afinidad = afinidad
	return nuevoPostFeed
}

func (algogram *Algogram) PublicarPost(contenido string) bool {
	if !algogram.HayLoggeado() {
		fmt.Printf("Error: no habia usuario loggeado\n")
		return false
	}

	post := crearNuevoPost(algogram.usuarioLoggeado, contenido, algogram.posts.Largo())
	algogram.posts.InsertarUltimo(post)

	iter := algogram.usuarios.Iterador()
	for iter.HaySiguiente() {
		nombre, usuario := iter.VerActual()
		if nombre != algogram.usuarioLoggeado.nombre {
			postFeed := crearNuevoPostFeed(post, usuario.afinidad)
			usuario.feed.Encolar(postFeed)
		}
		iter.Siguiente()
	}

	return true
}

func (algogram *Algogram) VerProximoPost() (int, string, string, int) {
	if !algogram.HayLoggeado() || algogram.usuarioLoggeado.feed.Cantidad() == 0 {
		fmt.Printf("Usuario no loggeado o no hay mas posts para ver\n")
		return 0, "", "", 0
	}

	postFeed := algogram.usuarioLoggeado.feed.Desencolar()

	return postFeed.post.id, postFeed.post.publicador.nombre, postFeed.post.contenido, postFeed.post.cantLikes // re feo jajaj
	// podria hacerse primitiva ObtenerPost(id) (?
}

func (algogram *Algogram) LikearPost(id int) bool {
	if !algogram.HayLoggeado() || id >= algogram.posts.Largo() || id < 0 {
		fmt.Printf("Error: Usuario no loggeado o Post inexistente\n")
		return false
	}

	iter := algogram.posts.Iterador()

	for i := 0; i < id; i++ { // en el peor de los casos O(p)
		iter.Siguiente()
	}

	postActual := iter.VerActual()

	if !postActual.likes.Pertenece(algogram.usuarioLoggeado.nombre) {
		postActual.likes.Guardar(algogram.usuarioLoggeado.nombre, algogram.usuarioLoggeado) // O(log Up)
		postActual.cantLikes++
	}

	return true
}

func (algogram *Algogram) MostrarLikes(id int) ([]string, int) {
	var likes []string
	if id >= algogram.posts.Largo() || id < 0 { // asumiendo que los posts estan en una lista
		fmt.Printf("Error: Post inexistente o sin likes\n")
		return likes, 0
	}

	iter := algogram.posts.Iterador()

	for i := 0; i < id; i++ { // en el peor de los casos O(p)
		iter.Siguiente()
	}
	postActual := iter.VerActual()
	if postActual.cantLikes == 0 {
		fmt.Printf("Error: Post inexistente o sin likes\n")
		return likes, 0
	}

	iterLikes := postActual.likes.Iterador()
	for iterLikes.HaySiguiente() { // O(Up)
		nombreUsuario, _ := iterLikes.VerActual()
		likes = append(likes, nombreUsuario)
		iterLikes.Siguiente()
	}

	return likes, postActual.cantLikes
}
