package routers

import (
	"encoding/json"

	"github.com/Junior_Jurado/gambit/bd"
	"github.com/Junior_Jurado/gambit/models"
)

func InsertAddress(body string, User string) (int, string) {
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if t.AddAddress == "" {
		return 400, "Debe especificar el Address "
	}

	if t.AddName == "" {
		return 400, "Debe especificar el Address "
	}

	if t.AddTitle == "" {
		return 400, "Debe especificar el Title "
	}

	if t.AddCity == "" {
		return 400, "Debe especificar el City "
	}

	if t.AddPhone == "" {
		return 400, "Debe especificar el Phone "
	}

	if t.AddPostalCode == "" {
		return 400, "Debe especificar el Postal Code "
	}

	err = bd.InsertAddress(t, User)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del Address para el ID de usuario " + User + " > " + err.Error()
	}

	return 200, "InsertAddress OK"
	
}