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

func UpdateAddress(body string, User string, id int) (int, string) {
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	t.AddId = id
	var encontrado bool
	err, encontrado = bd.AddressExist(User, t.AddId)
	if !encontrado {
		if err != nil {
			return 400, "Error al intentar buscar el Address para el usuario " + User + " > " + err.Error()
		}
		return 400, "No se encuentra un registro de ID de Usuario asociado a esa ID de Address"
	}
	

	err = bd.UpdateAddress(t, User)
	if err != nil {
		return 400, "Error al intentar realizar Update Address al Usuario " + User + " > " + err.Error()
	}

	return 200, "Update Address OK"
}