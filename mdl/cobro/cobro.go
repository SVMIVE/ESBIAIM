package cobro

type Cobro struct {

	CodTipoDosa    string	`json:"codtipodosa"`
	IdDosaAnulada  int64	`json:"idanulada"`
	IdCobro				 int64	`json:"idcobro"`
	CodTipoMoneda  string	`json:"codtipomoneda"`
	CodFactura     string	`json:"codfactura"`
	CodCierre      string	`json:"codcierre"`
	CodApertura    string	`json:"codapertura"`
	CodEstatus     string	`json:"codestatus"`
	FechaApertura  string	`json:"fechaapertura"`
	Mixta          bool		`json:"mixta"`
	Correlativo		 string	`json:"correlativo"`
}
