package fids

type Fids struct {
	TipoVuelo          	  	 string   `json:"tipovuelo"`
	MovimientoAeronave 			 int64    `json:"movimientoaeronave"`
	HoraReal           			 string   `json:"horareal"`
	NumeroVueloLlegada 			 int64    `json:"numerovuelollegada"`
	NumeroVueloSalida  			 int64    `json:"numerovuelosalida"`
	HoraRealLlegada     		 string   `json:"horarealllegada"`
	HoraRealSalida     			 string   `json:"horarealsalida"`
	Nacionalidad       			 string   `json:"nacionalidad"`
	Matricula          			 string   `json:"matricula"`
	Aeronave          			 Aeronave `json:"Aeronave"`
	Vuelos             			 Vuelos  	`json:"Vuelos"`
	FechaLlegada 			 			 string 	`json:"fechallegada"`
	IdTipoVuelo 			 			 string 	`json:"idtipovuelo"`
	Descripcion							 string		`json:"nombretipovuelo"`
	IdVuelosDiariosSalida	 	 string		`json:"idvuelosdiariossalida"`
	IdVuelosDiariosLlegada	 string		`json:"idvuelosdiariosllegada"`
	InicioAsignacionCorrea	 string		`json:"inicioasignacioncorrea"`
	FinAsignacionCorrea	 		 string		`json:"finasignacioncorrea"`
}
