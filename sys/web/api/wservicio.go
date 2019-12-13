package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/svmive/esbiaim/sys"
)

type WServicio struct {
	Servicio string `json:"servicio" bson:"servicio"`
}

//Listar
func (ws *WServicio) Listar(w http.ResponseWriter, r *http.Request) {
	var M Respuesta
	Cabecera(w, r)
	url := sys.HTTPAPISYBASE + "servicio/listar"
	response, err := http.Get(url)
	if err != nil {
		M.Mensaje = err.Error()
		M.Tipo = 0
		w.WriteHeader(http.StatusForbidden)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			M.Mensaje = err.Error()
			M.Tipo = 0
			j, _ := json.Marshal(M)
			w.Write(j)
			return
		}
		defer response.Body.Close()

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}
