package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/svmive/esbiaim/mdl/banco"
	"github.com/svmive/esbiaim/sys"
	"github.com/svmive/esbiaim/util"
)

//WFactura
type WFactura struct {
	Rif    string
	Numero string
}

func (wf *WFactura) Listar(w http.ResponseWriter, r *http.Request) {
	var M Respuesta
	Cabecera(w, r)

	url := sys.HTTPAPISYBASE + "factura/listar"
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

func (wf *WFactura) Agregar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M Respuesta
	var banco banco.Banco
	url := sys.HTTPAPISYBASE + "recaudacion.php?met=AdminBancos"

	errx := json.NewDecoder(r.Body).Decode(&banco)
	M.Tipo = 1
	if errx != nil {
		M.Mensaje = errx.Error()
		M.Tipo = 0
		fmt.Println(M.Mensaje)
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusForbidden)
		w.Write(j)
		return
	}

	jsonW, ex := json.Marshal(wf)
	if ex != nil {
		fmt.Println(ex.Error())
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonW))
	if err != nil {
		M.Mensaje = err.Error()
		M.Tipo = 0
		w.WriteHeader(http.StatusOK)
		j, _ := json.Marshal(M)
		w.Write(j)
		return
	} else {
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {

			w.WriteHeader(http.StatusOK)
			M.Mensaje = err.Error()
			M.Tipo = 0
			j, _ := json.Marshal(M)
			w.Write(j)
			return
		}
		defer response.Body.Close()
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		return
	}
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
