package main
import (
	"errors"
	"log"
	"net/http"

	"example.com/example-project/internal/cookies"
)
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/set", setCookieHandler)
	mux.HandleFunc("/get", getCookieHandler)

	log.Print("Listening...")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal((err))
	}

}

func setCookieHandler(w http.ResponseWriter, r *http.Request)  {
	cookie := http.Cookie{
		Name: "exampleCookie",
		Value: "Hello Zoё!",
		Path: "/",
		MaxAge: 3600,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteLaxMode,
	}
	err := cookies.Write(w, cookie)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// http.SetCookie(w, &cookie)
	w.Write([]byte("cookie set!"))

}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	value, err := cookies.Read(r, "exampleCookie")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		case errors.Is(err, cookies.ErrInvalidValue):
			http.Error(w, "invalid cookie", http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte(value))
}