package TDAalgogram

import (
	"fmt"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
)

type Algogram struct {
	usuarios        TDADiccionario.Diccionario[string, *usuario]
	usuarioLoggeado *usuario
	hayLoggeado     bool
	posts           TDALista.Lista[post]
}

type usuario struct {
	nombre   string
	feed     TDAHeap.ColaPrioridad[post]
	afinidad int
}

type post struct {
	id         int
	publicador *usuario
	contenido  string
	likes      TDADiccionario.DiccionarioOrdenado[string, *usuario]
	cantLikes  int
}

func CrearAlgogram(us TDADiccionario.Diccionario[string, int]) *Algogram {
	usuarios := cargarUsuarios(us)

	return &Algogram{
		usuarioLoggeado: nil,
		hayLoggeado:     false,
		usuarios:        usuarios,
		posts:           TDALista.CrearListaEnlazada[post](),
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
	usuario.feed = TDAHeap.CrearHeap[post](igualdadPost) // Hay que hacer una funcion para que muestre por afinidad

	return usuario
}

func igualdadPost(post1, post2 post) int {
	return post2.id - post1.id
}

func (algogram *Algogram) HayLoggeado() bool {
	if algogram.hayLoggeado {
		return true
	}
	return false
}

func (algogram *Algogram) Login(nombre string) string {
	if !algogram.usuarios.Pertenece(nombre) {
		fmt.Printf("Error: usuario no existente\n")
		return ""
	}

	if algogram.hayLoggeado {
		fmt.Printf("Error: Ya habia un usuario loggeado\n")
		return ""
	}

	algogram.loggearUsuario(nombre)
	return nombre
}

func (algogram *Algogram) loggearUsuario(nombre string) {
	usuario := algogram.usuarios.Obtener(nombre)
	algogram.usuarioLoggeado = usuario
	algogram.hayLoggeado = true
}

func (algogram *Algogram) Logout() {
	if !algogram.hayLoggeado {
		fmt.Printf("Error: no habia usuario loggeado\n")
		return
	}

	algogram.desloggearUsuario()
}

func (algogram *Algogram) desloggearUsuario() {
	algogram.usuarioLoggeado = nil
	algogram.hayLoggeado = false
}

func crearNuevoPost(u *usuario, contenido string, cant int) post {
	nuevoPost := new(post)
	nuevoPost.id = cant
	nuevoPost.publicador = u
	nuevoPost.contenido = contenido
	nuevoPost.cantLikes = 0
	nuevoPost.likes = nil

	return *nuevoPost
}

func (algogram *Algogram) PublicarPost(contenido string) {
	if !algogram.hayLoggeado {
		fmt.Printf("Error: no habia usuario loggeado\n")
		return
	}

	post := crearNuevoPost(algogram.usuarioLoggeado, contenido, algogram.posts.Largo())
	algogram.posts.InsertarUltimo(post)

	iter := algogram.usuarios.Iterador()
	for iter.HaySiguiente() {
		nombre, usuario := iter.VerActual()
		if nombre != algogram.usuarioLoggeado.nombre {
			usuario.feed.Encolar(post)
		}
		iter.Siguiente()
	}
}

func (algogram *Algogram) VerProximoPost() (int, string, string, int) {
	if !algogram.hayLoggeado || algogram.usuarioLoggeado.feed.Cantidad() == 0 {
		fmt.Printf("Usuario no loggeado o no hay mas posts para ver\n")
		return 0, "", "", 0
	}

	p := algogram.usuarioLoggeado.feed.Desencolar()

	return p.id, p.publicador.nombre, p.contenido, p.cantLikes // re feo jajaj
	// podria hacerse primitiva ObtenerPost(id) (?
}

func (algogram *Algogram) LikearPost(id int) {
	if !algogram.hayLoggeado || algogram.posts.Largo() >= id {
		fmt.Printf("Error: Usuario no loggeado o Post inexistente\n")
		return
	}

	iter := algogram.posts.Iterador()

	for i := 0; i <= id; i++ { // en el peor de los casos O(p)
		iter.Siguiente()
	}

	postActual := iter.VerActual()

	if !postActual.likes.Pertenece(algogram.usuarioLoggeado.nombre) {
		postActual.likes.Guardar(algogram.usuarioLoggeado.nombre, algogram.usuarioLoggeado) // O(log Up)
		postActual.cantLikes++
	}
}

func (algogram *Algogram) MostrarLikes(id int) ([]string, int) {
	var likes []string
	if algogram.posts.Largo() >= id { // asumiendo que los posts estan en una lista
		fmt.Printf("Error: Post inexistente o sin likes\n")
		return likes, 0
	}

	iter := algogram.posts.Iterador()

	for i := 0; i <= id; i++ { // en el peor de los casos O(p)
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
