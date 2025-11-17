package algogramSistema_test

import (
	"fmt"
	"testing"
	TDAAlgogram "tp2/algogramSistema"

	"github.com/stretchr/testify/require"
)

const (
	_ERROR_USUARIO_NO_LOGUEADO = "Error: no habia usuario loggeado"
	_ERROR_USUARIO_YA_LOGUEADO = "Error: Ya habia un usuario loggeado"
	_ERROR_USUARIO_NO_EXISTE   = "Error: usuario no existente"
	_ERROR_FEED                = "Usuario no loggeado o no hay mas posts para ver"
	_ERROR_LIKE                = "Error: Usuario no loggeado o Post inexistente"
	_ERROR_MOSTRAR_LIKES       = "Error: Post inexistente o sin likes"

	_MENSAJE_LOGIN          = "Hola %s" // nombre
	_MENSAJE_LOGOUT         = "Adios"
	_MENSAJE_PUBLICADO      = "Post publicado"
	_MENSAJE_SIGUIENTE_POST = "Post ID %d\n%s dijo: %s\nLikes: %d" // id_post, nombre, texto, cant de likes
	_MENSAJE_LIKEAR         = "Post likeado"
	_MENSAJE_MOSTRAR_LIKES  = "El post tiene %d likes:%s"
)

func TestAlgoGramVacio(t *testing.T) {
	usuarios := make([]string, 0)

	algoG := TDAAlgogram.CrearAlgoGram(usuarios)

	require.Equal(t, _ERROR_USUARIO_NO_EXISTE, algoG.Login("Messi"))
	require.Equal(t, _ERROR_USUARIO_NO_LOGUEADO, algoG.Logout())
	require.Equal(t, _ERROR_USUARIO_NO_LOGUEADO, algoG.Publicar("Hola Mundo!"))
	require.Equal(t, _ERROR_FEED, algoG.VerSiguienteFeed())
	require.Equal(t, _ERROR_LIKE, algoG.LikearPost(0))
	require.Equal(t, _ERROR_MOSTRAR_LIKES, algoG.MostrarLikes(0))
}

func TestLoginLogout(t *testing.T) {
	usuarios := []string{"Messi", "Cuti", "Barbara", "Alan"}

	algoG := TDAAlgogram.CrearAlgoGram(usuarios)

	require.Equal(t, _ERROR_USUARIO_NO_EXISTE, algoG.Login("Dybala"))
	require.Equal(t, _ERROR_USUARIO_NO_LOGUEADO, algoG.Logout())

	require.Equal(t, fmt.Sprintf(_MENSAJE_LOGIN, "Messi"), algoG.Login("Messi"))
	require.Equal(t, _ERROR_USUARIO_YA_LOGUEADO, algoG.Login("Messi"))
	require.Equal(t, _ERROR_USUARIO_YA_LOGUEADO, algoG.Login("Barbara"))
	require.Equal(t, _MENSAJE_LOGOUT, algoG.Logout())

	require.Equal(t, fmt.Sprintf(_MENSAJE_LOGIN, "Barbara"), algoG.Login("Barbara"))
	require.Equal(t, _MENSAJE_LOGOUT, algoG.Logout())
	require.Equal(t, _ERROR_USUARIO_NO_LOGUEADO, algoG.Logout())
}

func TestPublicarYVerFeed(t *testing.T) {
	usuarios := []string{"Messi", "Cuti", "Barbara", "Alan"}

	postMessi := "Hola Mundo!"
	postBarbara := "Hola Mundo Informático!"

	algoG := TDAAlgogram.CrearAlgoGram(usuarios)

	require.Equal(t, fmt.Sprintf(_MENSAJE_LOGIN, "Messi"), algoG.Login("Messi"))
	require.Equal(t, _MENSAJE_PUBLICADO, algoG.Publicar(postMessi))
	require.Equal(t, _ERROR_FEED, algoG.VerSiguienteFeed())
	require.Equal(t, _MENSAJE_LOGOUT, algoG.Logout())

	require.Equal(t, fmt.Sprintf(_MENSAJE_LOGIN, "Barbara"), algoG.Login("Barbara"))
	require.Equal(t, _MENSAJE_PUBLICADO, algoG.Publicar(postBarbara))
	require.Equal(t, fmt.Sprintf(_MENSAJE_SIGUIENTE_POST, 0, "Messi", postMessi, 0), algoG.VerSiguienteFeed())
	require.Equal(t, _MENSAJE_LOGOUT, algoG.Logout())

	require.Equal(t, fmt.Sprintf(_MENSAJE_LOGIN, "Alan"), algoG.Login("Alan"))
	require.Equal(t, fmt.Sprintf(_MENSAJE_SIGUIENTE_POST, 1, "Barbara", postBarbara, 0), algoG.VerSiguienteFeed())
	require.Equal(t, fmt.Sprintf(_MENSAJE_SIGUIENTE_POST, 0, "Messi", postMessi, 0), algoG.VerSiguienteFeed())
	require.Equal(t, _MENSAJE_LOGOUT, algoG.Logout())

	require.Equal(t, fmt.Sprintf(_MENSAJE_LOGIN, "Cuti"), algoG.Login("Cuti"))
	require.Equal(t, fmt.Sprintf(_MENSAJE_SIGUIENTE_POST, 0, "Messi", postMessi, 0), algoG.VerSiguienteFeed())
	require.Equal(t, fmt.Sprintf(_MENSAJE_SIGUIENTE_POST, 1, "Barbara", postBarbara, 0), algoG.VerSiguienteFeed())
	require.Equal(t, _ERROR_FEED, algoG.VerSiguienteFeed())
	require.Equal(t, _MENSAJE_LOGOUT, algoG.Logout())
	require.Equal(t, _ERROR_USUARIO_NO_LOGUEADO, algoG.Logout())
}

func TestLikearYMostrarLikes(t *testing.T) {
	usuarios := []string{"Messi", "Cuti", "Barbara", "Alan"}

	postMessi := "Hola Mundo!"
	postBarbara := "Hola Mundo Informático!"

	algoG := TDAAlgogram.CrearAlgoGram(usuarios)

	require.Equal(t, fmt.Sprintf(_MENSAJE_LOGIN, "Messi"), algoG.Login("Messi"))
	require.Equal(t, _MENSAJE_PUBLICADO, algoG.Publicar(postMessi))
	require.Equal(t, _ERROR_FEED, algoG.VerSiguienteFeed())
	require.Equal(t, _MENSAJE_LOGOUT, algoG.Logout())

	require.Equal(t, fmt.Sprintf(_MENSAJE_LOGIN, "Barbara"), algoG.Login("Barbara"))
	require.Equal(t, _MENSAJE_PUBLICADO, algoG.Publicar(postBarbara))
	require.Equal(t, fmt.Sprintf(_MENSAJE_SIGUIENTE_POST, 0, "Messi", postMessi, 0), algoG.VerSiguienteFeed())
	require.Equal(t, _ERROR_MOSTRAR_LIKES, algoG.MostrarLikes(0))
	require.Equal(t, _MENSAJE_LIKEAR, algoG.LikearPost(0))
	require.Equal(t, fmt.Sprintf(_MENSAJE_MOSTRAR_LIKES, 1, "\n\tBarbara"), algoG.MostrarLikes(0))

	require.Equal(t, _ERROR_LIKE, algoG.LikearPost(3))
	require.Equal(t, _MENSAJE_LOGOUT, algoG.Logout())

	require.Equal(t, fmt.Sprintf(_MENSAJE_LOGIN, "Alan"), algoG.Login("Alan"))
	require.Equal(t, fmt.Sprintf(_MENSAJE_SIGUIENTE_POST, 1, "Barbara", postBarbara, 0), algoG.VerSiguienteFeed())
	require.Equal(t, fmt.Sprintf(_MENSAJE_SIGUIENTE_POST, 0, "Messi", postMessi, 1), algoG.VerSiguienteFeed())
	require.Equal(t, _MENSAJE_LIKEAR, algoG.LikearPost(1))
	require.Equal(t, _MENSAJE_LIKEAR, algoG.LikearPost(0))
	require.Equal(t, fmt.Sprintf(_MENSAJE_MOSTRAR_LIKES, 2, "\n\tAlan\n\tBarbara"), algoG.MostrarLikes(0))
	require.Equal(t, fmt.Sprintf(_MENSAJE_MOSTRAR_LIKES, 1, "\n\tAlan"), algoG.MostrarLikes(1))
	require.Equal(t, _MENSAJE_LOGOUT, algoG.Logout())

	require.Equal(t, fmt.Sprintf(_MENSAJE_LOGIN, "Cuti"), algoG.Login("Cuti"))
	require.Equal(t, fmt.Sprintf(_MENSAJE_SIGUIENTE_POST, 0, "Messi", postMessi, 2), algoG.VerSiguienteFeed())
	require.Equal(t, fmt.Sprintf(_MENSAJE_SIGUIENTE_POST, 1, "Barbara", postBarbara, 1), algoG.VerSiguienteFeed())
	require.Equal(t, _ERROR_FEED, algoG.VerSiguienteFeed())
	require.Equal(t, _MENSAJE_LIKEAR, algoG.LikearPost(1))
	require.Equal(t, _MENSAJE_LIKEAR, algoG.LikearPost(0))
	require.Equal(t, fmt.Sprintf(_MENSAJE_MOSTRAR_LIKES, 3, "\n\tAlan\n\tBarbara\n\tCuti"), algoG.MostrarLikes(0))
	require.Equal(t, fmt.Sprintf(_MENSAJE_MOSTRAR_LIKES, 2, "\n\tAlan\n\tCuti"), algoG.MostrarLikes(1))
	require.Equal(t, _MENSAJE_LOGOUT, algoG.Logout())
	require.Equal(t, _ERROR_USUARIO_NO_LOGUEADO, algoG.Logout())
}
