package algogramSistema

type AlgoGram interface {
	// u = cantidad de usuarios en total
	// p = cantidad de post en total
	// Up = cantidad de usuarios que likearon anteriormente un post en particular

	// Si el usuario existe y no hay sesion iniciada, devuelve "Hola" + el nombre de usuario
	Login(string) string // O(1)

	//Si hay sesion iniciada, devuelve "Adios"
	Logout() string // O(1)

	// si hay sescion iniciada, devuelve "Post publicado"
	Publicar(string) string // O(u log(p))

	// si hay siguiente en el feed, devuelve la siguiente publicacion
	VerSiguienteFeed() string // O(log (p))

	// si existe la publicacio, devuelve "Post likeado"
	LikearPost(int) string // O(log (Up))

	// si existe la publicacio, devuelve la cantidad de likes y la lista de usuarios que likearon en orden alfabetico
	MostrarLikes(int) string // O(Up)
}

type usuarioAlgoGram interface {
	// devuelve su posicion en el arreglo que uso para calcular afinidades
	verPosicion() int

	// devuelve el nombre de usuario
	verNombre() string

	// devuelve el siguiente post si lo hay, sino nil
	verSiguienteFeed() publicacionAlgoGram

	// agrega una publucacion con su afinidad entre usuarios al feed del usuario
	agregarFeed(postConAfinidad)

	// devuelve  si es el mismo usuario, un postitivo si el que "llama" es mayor al que se pasa por parametro, y un negativo si es menor
	compararAfinidad(usuarioAlgoGram) int
}

type publicacionAlgoGram interface {
	// devuelve el usuario que publico el post
	verUsuario() usuarioAlgoGram

	// devuelve el id de post que representa su posicion en el arreglo de publicaciones
	verId() int

	// devuelve el texto del string
	verTexto() string

	// le da like al post
	likear(usuarioAlgoGram)

	// devuelve la cantidad de likes
	cantidadLikes() int

	// devuelve la lista de usuarios que dieron like al post
	likeadoPor() []usuarioAlgoGram
}
