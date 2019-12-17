package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/svmive/esbiaim/sys"
)

type WAdminControl struct {
}

func (wac WAdminControl) AutoIncrement(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var M Respuesta
	var interfaz interface{}
	url := sys.HTTPAPISYBASE + "admincontrol/autoincrement"

	errx := json.NewDecoder(r.Body).Decode(&interfaz)
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

	jsonW, ex := json.Marshal(interfaz)
	if ex != nil {
		fmt.Println(ex.Error())
	}
	fmt.Println(interfaz)
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
