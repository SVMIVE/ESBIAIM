package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/svmive/esbiaim/mdl/banco"
	"github.com/svmive/esbiaim/sys"
)

type WBanco struct {
	Codigo      string `json:"codigo"`
	Descripcion string `json:"descripcion"`
	Estatus     int    `json:"estatus"`
	Abreviatura string `json:"abreviatura"`
	Usuario     string `json:"usuario"`
}

func (wb *WBanco) Listar(w http.ResponseWriter, r *http.Request) {
	var M Respuesta
	Cabecera(w, r)

	url := sys.HTTPAPISYBASE + "banco/listar"
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

func (wb *WBanco) Agregar(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M Respuesta
	var banco banco.Banco
	url := sys.HTTPAPISYBASE + "banco/insertar"

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

	jsonW, ex := json.Marshal(wb)
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
