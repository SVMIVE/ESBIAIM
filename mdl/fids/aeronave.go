package fids

type Aeronave struct {
	Tipo         				string 	`json:"tipo"`
	IDMovimiento 				int64  	`json:"idaeronave"`
	MovimientoAeronave 	int64 	`json:"modeloaeronave"`
	PesoMaximo    			float64	`json:"pesomaximo"`
	ModeloAeronave    	float64	`json:"modeloaeronave"`

}
