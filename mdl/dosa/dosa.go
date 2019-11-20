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
	Cliente       cliente.Cliente 		`json:"Cliente" bson:"cliente"`
	Tipo          string          		`json:"tipo" bson:"tipo"`
	IdAnulada     int64          		 	`json:"idanulada" bson:"idanulada"`
	TipoMoneda    string         	 		`json:"tipomoneda" bson:"tipomoneda"`
	FechaApertura string          		`json:"fechaapertura" bson:"fechaapertura"`
	Descripcion   string          		`json:"descripcion" bson:"descripcion"`
	Duplicada     int64          	 		`json:"duplicada" bson:"duplicada"`
	Especial      int64           		`json:"especial" bson:"especial"`
	Fids          fids.Fids       		`json:"Fids" bson:"fids"`
	Cobro         cobro.Cobro     		`json:"Cobro" bson:"cobro"`
	Vuelos        fids.Vuelos   			`json:"Vuelos" bson:"vuelos"`
	Aeronave    	fids.Aeronave 				`json:"Aeronave" bson:"aeronave"`
}

type LDOSA struct {
	Fecha time.Time `json:"fecha" bson:"fecha"`
	ID    string    `json:"id" bson:"id"`
	Lista []DOSA    `json:"Lista" bson:"lista"`
}

func (D *DOSA) Listar() (j []byte, err error) {
	var lstDosa LDOSA
	var sql SQLDosa
	consulta := sql.Todas("2019-10-10", "2019-10-10")
	rs, err := sys.PostgreSQLSINIV.Query(consulta)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		for rs.Next() {
			var ds DOSA
			var ccli, ncli, lote, formap, cotipo, comon, fchllev, tipvuelo, descpcion, horareal string
			var hrlleg, hrsal, nac, codfac, codcie, codaper, codest, fecaper string
			var idvuelo, idvdsali string
			var idcoanu, idarn, idcobro, numlleg, numsal, espec, dupli int64
			var mixta bool
			var pesomax, modaero float64
			rs.Scan(&ccli, &ncli, &lote, &formap, &cotipo, &idcoanu, &comon, &idarn, &fchllev, &tipvuelo,
				 			&idvuelo, &descpcion, &horareal, &numlleg, &numsal, &hrlleg, &hrsal, &nac, &idcobro,
						 	&codfac, &codcie, &codaper, &codest, &fecaper, &modaero, &espec, &dupli, &idvdsali,
							&pesomax, &mixta)

			ds.Cliente.Codigo = ccli
			ds.Cliente.Nombre = ncli
			ds.Cliente.Lote = lote
			ds.Cliente.FormaPago = formap
			ds.Cobro.CodTipoDosa = cotipo
			ds.Cobro.IdDosaAnulada = idcoanu
			ds.Cobro.CodTipoMoneda = comon
			ds.Fids.Aeronave.IDMovimiento = idarn
			ds.Fids.FechaLlegada = fchllev
			ds.Fids.TipoVuelo = tipvuelo
			ds.Fids.IdTipoVuelo = idvuelo
			ds.Fids.Descripcion = descpcion
			ds.Fids.HoraReal = horareal
			ds.Fids.NumeroVueloLlegada = numlleg
			ds.Fids.NumeroVueloSalida = numsal
			ds.Fids.HoraRealLlegada = hrlleg
			ds.Fids.HoraRealSalida = hrsal
			ds.Fids.Nacionalidad = nac
			ds.Cobro.IdCobro = idcobro
			ds.Cobro.CodFactura = codfac
			ds.Cobro.CodCierre = codcie
			ds.Cobro.CodApertura = codaper
			ds.Cobro.CodEstatus = codest
			ds.Cobro.FechaApertura = fecaper
			ds.Fids.Aeronave.ModeloAeronave = modaero
			ds.Especial = espec
			ds.Duplicada = dupli
			ds.Fids.IdVuelosDiariosSalida= idvdsali
			ds.Fids.Aeronave.PesoMaximo= pesomax
			ds.Cobro.Mixta= mixta





			lstDosa.Lista = append(lstDosa.Lista, ds)
		}
	}
	lstDosa.Fecha = time.Now()
	j, err = json.Marshal(lstDosa)

	return
}
