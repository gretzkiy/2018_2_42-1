package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
	flag "github.com/spf13/pflag" // ради gnu style: --flag='value'

	"github.com/go-park-mail-ru/2018_2_42/authorization_server/accessor"
	"github.com/go-park-mail-ru/2018_2_42/authorization_server/environment"
	"github.com/go-park-mail-ru/2018_2_42/authorization_server/handlers"
)

func registerUsersHandlers(handlersEnv handlers.Environment) {
	http.Handle("/api/v1/users", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// setupResponse(w, r)
			switch r.Method {
			case http.MethodGet:
				handlersEnv.LeaderBoard(w, r)
			default:
				handlersEnv.ErrorMethodNotAllowed(w, r)
			}
		}))
}

func registerSessionHandlers(handlersEnv handlers.Environment) {
	http.Handle("/api/v1/session", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				handlersEnv.Login(w, r)
			case http.MethodDelete:
				handlersEnv.Logout(w, r)
			default:
				handlersEnv.ErrorMethodNotAllowed(w, r)
			}
		}))
}

func registerAvatarHandlers(handlersEnv handlers.Environment) {
	http.Handle("/api/v1/avatar", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				handlersEnv.SetAvatar(w, r)
			default:
				handlersEnv.ErrorMethodNotAllowed(w, r)
			}
		}))
}

func registerUserHandlers(handlersEnv handlers.Environment) {
	http.Handle("/api/v1/user", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				handlersEnv.UserProfile(w, r)
			case http.MethodPost:
				params := r.URL.Query()
				temporary := false
				if isTemporary, ok := params["temporary"]; ok && len(isTemporary) == 1 {
					temporary = isTemporary[0] == "true"
				}
				if temporary {
					handlersEnv.RegistrationTemporary(w, r)
				} else {
					handlersEnv.RegistrationRegular(w, r)
				}
				return
			default:
				handlersEnv.ErrorMethodNotAllowed(w, r)
			}
		}))
}

func main() {
	// получаем конфигурацию из аргументов командной строки
	listeningPort := flag.String("listening-port", "8080", "port on which the server will listen")
	postgresPath := flag.String("postgres-path",
		"full postgres address like 'postgres://postgres:@127.0.0.1:5432/postgres?sslmode=disable' in default",
		"postgres://postgres:@127.0.0.1:5432/postgres?sslmode=disable")
	flag.Parse()

	// подключаемся к базе.
	var err error
	handlersEnv := handlers.Environment(environment.Environment{})
	handlersEnv.DB, err = accessor.ConnectToDatabase(*postgresPath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "accessor.ConnectToDatabase: "))
	}
	err = handlersEnv.DB.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := handlersEnv.DB.Close()
		if err != nil {
			log.Fatal("failed to close Database connection: " + err.Error())
		}
	}()

	// регистрируем обработчики запросов с логикой сервера.
	registerUserHandlers(handlersEnv)
	registerUsersHandlers(handlersEnv)
	registerSessionHandlers(handlersEnv)
	registerAvatarHandlers(handlersEnv)

	// начинаем слушать порт.
	fmt.Println("starting server at :" + *listeningPort)
	log.Println(http.ListenAndServe(":"+*listeningPort, nil))

	return
}
