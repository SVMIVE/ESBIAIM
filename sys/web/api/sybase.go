package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Respuesta struct {
	Mensaje string `json:"mensaje"`
	Tipo    int64  `json:"tipo"`
}

func (SB *WSyBase) ListarDocumentos(w http.ResponseWriter, r *http.Request) {
	var M Respuesta
	Cabecera(w, r)
	base := "http://192.168.10.66/app_dosa/webservices/ws_recaudacion.php?met=sybaseListarDocumentos"

	err := json.NewDecoder(r.Body).Decode(&SB)
	url := base + "|" + SB.Desde + ";FAC;A;cartelera;C312a1aywWev"
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
