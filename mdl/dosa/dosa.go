package dosa

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/svmive/esbiaim/mdl/cliente"
	"github.com/svmive/esbiaim/mdl/cobro"
	"github.com/svmive/esbiaim/mdl/fids"
	"github.com/svmive/esbiaim/sys"
)

type DOSA struct {
	Cliente       cliente.Cliente `json:"Cliente" bson:"cliente"`
	Tipo          string          `json:"tipo" bson:"tipo"`
	IdAnulada     int64           `json:"idanulada" bson:"idanulada"`
	TipoMoneda    string          `json:"tipomoneda" bson:"tipomoneda"`
	FechaApertura string          `json:"fechaapertura" bson:"fechaapertura"`
	Descripcion   string          `json:"descripcion" bson:"descripcion"`
	Duplicada     int64           `json:"duplicada" bson:"duplicada"`
	Fids          fids.Fids       `json:"Fids" bson:"fids"`
	Cobro         cobro.Cobro     `json:"Cobro" bson:"cobro"`
}

type LDOSA struct {
	Fecha time.Time `json:"fecha" bson:"fecha"`
	ID    string    `json:"id" bson:"id"`
	Lista []DOSA    `json:"Lista" bson:"lista"`
}

func (D *DOSA) Listar() (j []byte, err error) {
	var lstDosa LDOSA
	rs, err := sys.PostgreSQLSINIV.Query(SQLDosaTodas())
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		for rs.Next() {
			var ds DOSA
			var ccli, ncli, lote, formap string
			var cotipo, comon string
			var idcoanu int64
			rs.Scan(&ccli, &ncli, &lote, &formap, &cotipo, &idcoanu, &comon)

			ds.Cliente.Codigo = ccli
			ds.Cliente.Nombre = ncli
			ds.Cliente.Lote = lote
			ds.Cliente.FormaPago = formap
			ds.Cobro.CodTipoDosa = cotipo
			ds.Cobro.IdDosaAnulada = idcoanu
			ds.Cobro.CodTipoMoneda = comon

			lstDosa.Lista = append(lstDosa.Lista, ds)
		}
	}
	lstDosa.Fecha = time.Now()
	j, err = json.Marshal(lstDosa)

	return
}
