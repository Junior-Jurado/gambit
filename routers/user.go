package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	// "fmt"
	// "strconv"

	// "github.com/aws/aws-lambda-go/events"
	"github.com/Junior_Jurado/gambit/bd"
	"github.com/Junior_Jurado/gambit/models"
	"github.com/aws/aws-lambda-go/events"
)

func UpdateUser(body string, User string) (int, string) {
	var t models.User
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.UserFirstName) == 0 && len(t.UserLastName) == 0 {
		return 400, "Debe especificar el Nombre (FirtsName) o (LastName) del Usuario"
	}

	_, encontrado := bd.UserExists(User)
	if !encontrado {
		return 400, "No existe un usuario con el UUID '" + User + "'"
	}

	err = bd.UpdateUser(t, User)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar la actualización del usuario " + User + " > " + err.Error()
	}

	return 200, "Update User OK"
}

func SelectUser(body string, User string) (int, string) {
	_, encontrado := bd.UserExists(User)
	if !encontrado {
		return 400, "No existe un usuario con el UUID '" + User + "'"
	}

	row, err := bd.SelectUser(User)
	fmt.Println(row)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el Select del usuario " + User + " > " + err.Error()
	}

	respJson, err := json.Marshal(row)
	if err != nil {
		return 500, "Error al formatear los datos del usuario en JSON"
	}

	return 200, string(respJson)
}

func SelectUsers(body string, User string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var Page int
	if len(request.QueryStringParameters["page"]) == 0 {
		Page = 1
	} else {
		Page, _ = strconv.Atoi(request.QueryStringParameters["page"])
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	user, err := bd.SelectUsers(Page)
	if err != nil {
		return 400, "Ocurrió un error al intentar obtener la lista de usuarios > " + err.Error()
	}

	respJson, err := json.Marshal(user)
	if err != nil {
		return 500, "Error al formatear los datos de los usuarios como JSON"
	}

	return 200, string(respJson)
}