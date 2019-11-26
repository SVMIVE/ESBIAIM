package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/svmive/esbiaim/sys"
)

type WBanco struct {
}

func (wb *WBanco) Listar(w http.ResponseWriter, r *http.Request) {
	var M Respuesta
	Cabecera(w, r)
	base := sys.HTTPAPISYBASE + "recaudacion.php?met=AdminBancos"

	url := base + "|cartelera;C312a1aywWev"
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
