package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/svmive/esbiaim/sys"
	"github.com/svmive/esbiaim/util"
)

//WFactura Familiares
type WFactura struct {
	Rif    string
	Numero string
}

//FacturasPorPagar Facturas Por Pagar
func (wfactura *WFactura) FacturasPorPagar(w http.ResponseWriter, r *http.Request) {
	var M Respuesta
	Cabecera(w, r)
	e := json.NewDecoder(r.Body).Decode(&wfactura)
	util.Error(e)
	base := sys.HTTPAPISYBASE + "recaudacion.php?met=FacturasPorPagar"

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
