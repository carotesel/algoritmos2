package algogramSistema

import (
	"fmt"
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
)

const (
	// mensajes de error
	_ERROR_USUARIO_NO_LOGUEADO = "Error: no habia usuario loggeado"
	_ERROR_USUARIO_YA_LOGUEADO = "Error: Ya habia un usuario loggeado"
	_ERROR_USUARIO_NO_EXISTE   = "Error: usuario no existente"
	_ERROR_FEED                = "Usuario no loggeado o no hay mas posts para ver"
	_ERROR_LIKE                = "Error: Usuario no loggeado o Post inexistente"
	_ERROR_MOSTRAR_LIKES       = "Error: Post inexistente o sin likes"

	// mensajes de respuesta/exito
	_MENSAJE_LOGIN          = "Hola %s" // nombre
	_MENSAJE_LOGOUT         = "Adios"
	_MENSAJE_PUBLICADO      = "Post publicado"
	_MENSAJE_SIGUIENTE_POST = "Post ID %d\n%s dijo: %s\nLikes: %d" // id_post, nombre, texto, cant de likes
	_MENSAJE_LIKEAR         = "Post likeado"
	_MENSAJE_MOSTRAR_LIKES  = "El post tiene %d likes:%s" // cantidad de likes, en la funcion le agregamos \n\t + nombre por cada usuario que dio like

	// redimencion del arreglo de publicaciones, solo para arriba ya que no se borran las publicaciones
	_TAMANIO_INICIAL       = 2
	_FACTOR_DE_REDIMENCION = 2
)

// estructuras
type postConAfinidad struct {
	post     publicacionAlgoGram
	afinidad int
}

type publicacion struct {
	usuario usuarioAlgoGram
	id      int
	texto   string
	likes   TDADiccionario.DiccionarioOrdenado[usuarioAlgoGram, *int] // *int para poder pasarle nil
}

type usuario struct {
	nombre   string
	posicion int
	feed     TDAColaPrioridad.ColaPrioridad[postConAfinidad]
}

type sistema struct {
	posts           []publicacionAlgoGram
	cantidadPosts   int
	usuarios        TDADiccionario.Diccionario[string, usuarioAlgoGram]
	usuarioLogueado usuarioAlgoGram
}

// los crear
func CrearAlgoGram(arrUsuarios []string) AlgoGram {
	hash := TDADiccionario.CrearHash[string, usuarioAlgoGram](cmpStringRetornaBool)

	for i, usuario := range arrUsuarios {
		hash.Guardar(usuario, crearUsuario(i, usuario))
	}

	return &sistema{posts: make([]publicacionAlgoGram, _TAMANIO_INICIAL), cantidadPosts: 0, usuarios: hash, usuarioLogueado: nil}
}

func crearUsuario(i int, nombre string) usuarioAlgoGram {
	return &usuario{
		nombre:   nombre,
		feed:     TDAColaPrioridad.CrearHeap(cmpPostafinidad),
		posicion: i,
	}
}

func crearPublicacion(texto string, usuario usuarioAlgoGram, id int) publicacionAlgoGram {
	return &publicacion{
		texto:   texto,
		id:      id,
		likes:   TDADiccionario.CrearABB[usuarioAlgoGram, *int](cmpUsuarioNombreRetornaInt),
		usuario: usuario,
	}
}

// auxiliares privadas
func cmpPostafinidad(p1, p2 postConAfinidad) int {
	if p1.afinidad != p2.afinidad {
		return p2.afinidad - p1.afinidad
	}
	return p2.post.verId() - p1.post.verId()
}

func cmpUsuarioNombreRetornaInt(u1, u2 usuarioAlgoGram) int {
	return strings.Compare(u1.verNombre(), u2.verNombre())
}

func cmpStringRetornaBool(usuario1, usuario2 string) bool {
	return usuario1 == usuario2
}

func absoluto(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// primitivas de la interfaz privada publicacionAlgoGram
func (post *publicacion) verUsuario() usuarioAlgoGram {
	return post.usuario
}

func (post *publicacion) cantidadLikes() int {
	return post.likes.Cantidad()
}

func (post *publicacion) verTexto() string {
	return post.texto
}

func (post *publicacion) verId() int {
	return post.id
}

func (post *publicacion) likear(usuario usuarioAlgoGram) {
	post.likes.Guardar(usuario, nil)
}

func (post *publicacion) likeadoPor() []usuarioAlgoGram {
	res := make([]usuarioAlgoGram, post.likes.Cantidad())

	i := 0
	iter := post.likes.Iterador()

	for iter.HaySiguiente() {
		usuario, _ := iter.VerActual()

		res[i] = usuario

		i++
		iter.Siguiente()
	}

	return res
}

// primitivas de la interfaz privada usuarioAlgoGram
func (usuario *usuario) verPosicion() int {
	return usuario.posicion
}

func (usuario *usuario) compararAfinidad(otro usuarioAlgoGram) int {
	return usuario.posicion - otro.verPosicion()
}

func (usuario *usuario) verNombre() string {
	return usuario.nombre
}

func (usuario *usuario) verSiguienteFeed() publicacionAlgoGram {
	if usuario.feed.EstaVacia() {
		return nil
	}
	postConDist := usuario.feed.Desencolar()
	return postConDist.post
}

func (usuario *usuario) agregarFeed(post postConAfinidad) {
	usuario.feed.Encolar(post)
}

// primitivas privadas de AlgoGram
func (algoG *sistema) sesionIniciada() bool {
	return algoG.usuarioLogueado != nil
}

func (algoG *sistema) redimencion(nuevo_tam int) {
	nuevo := make([]publicacionAlgoGram, nuevo_tam)
	copy(nuevo, algoG.posts[:algoG.cantidadPosts])
	algoG.posts = nuevo
}

// primitivas de AlgoGram
func (algoG *sistema) Login(nombre string) string {
	if algoG.sesionIniciada() {
		return _ERROR_USUARIO_YA_LOGUEADO
	}

	if !algoG.usuarios.Pertenece(nombre) {
		return _ERROR_USUARIO_NO_EXISTE
	}

	algoG.usuarioLogueado = algoG.usuarios.Obtener(nombre)

	return fmt.Sprintf(_MENSAJE_LOGIN, nombre)
}

func (algoG *sistema) Logout() string {
	if !algoG.sesionIniciada() {
		return _ERROR_USUARIO_NO_LOGUEADO
	}

	algoG.usuarioLogueado = nil
	return _MENSAJE_LOGOUT
}

func (algoG *sistema) Publicar(mensaje string) string {
	if !algoG.sesionIniciada() {
		return _ERROR_USUARIO_NO_LOGUEADO
	}

	if algoG.cantidadPosts == cap(algoG.posts) {
		algoG.redimencion(algoG.cantidadPosts * _FACTOR_DE_REDIMENCION)
	}

	post := crearPublicacion(mensaje, algoG.usuarioLogueado, algoG.cantidadPosts)
	algoG.posts[algoG.cantidadPosts] = post
	algoG.cantidadPosts++

	iter := algoG.usuarios.Iterador()

	for iter.HaySiguiente() {
		_, usuario := iter.VerActual()

		if usuario != algoG.usuarioLogueado {
			afinidad := absoluto(algoG.usuarioLogueado.compararAfinidad(usuario))

			usuario.agregarFeed(postConAfinidad{post: post, afinidad: afinidad})
		}

		iter.Siguiente()
	}

	return _MENSAJE_PUBLICADO
}

func (algoG *sistema) VerSiguienteFeed() string {
	if !algoG.sesionIniciada() {
		return _ERROR_FEED
	}

	postSiguiente := algoG.usuarioLogueado.verSiguienteFeed()

	if postSiguiente == nil {
		return _ERROR_FEED
	}

	return fmt.Sprintf(_MENSAJE_SIGUIENTE_POST,
		postSiguiente.verId(),
		postSiguiente.verUsuario().verNombre(),
		postSiguiente.verTexto(),
		postSiguiente.cantidadLikes())
}

func (algoG *sistema) LikearPost(id_post int) string {
	if !algoG.sesionIniciada() || algoG.cantidadPosts == 0 || !(id_post >= 0 && id_post <= algoG.cantidadPosts-1) {
		return _ERROR_LIKE
	}

	post := algoG.posts[id_post]

	post.likear(algoG.usuarioLogueado)

	return _MENSAJE_LIKEAR
}

func (algoG *sistema) MostrarLikes(id_post int) string {
	if algoG.cantidadPosts == 0 || !(id_post >= 0 && id_post <= algoG.cantidadPosts-1) {
		return _ERROR_MOSTRAR_LIKES
	}

	post := algoG.posts[id_post]

	if post.cantidadLikes() == 0 {
		return _ERROR_MOSTRAR_LIKES
	}

	arrUsuarios := post.likeadoPor()
	var stringRes string

	for _, usuario := range arrUsuarios {
		stringRes += "\n\t" + usuario.verNombre()
	}

	return fmt.Sprintf(_MENSAJE_MOSTRAR_LIKES, post.cantidadLikes(), stringRes)
}
