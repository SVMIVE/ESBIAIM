package concepto

type Concepto struct {
	Codigo      	string 	`json:"codigo" bson:"codigo"`
	Nombre     		string 	`json:"nombre" bson:"nombre"`
	Porcentaje 		string 	`json:"porcentaje" bson:"porcentaje"`
	Monto     		string 	`json:"monto" bson:"monto"`
	CodigoCuenta  string 	`json:"codigocuenta" bson:"codigocuenta"`
	ConIva 				bool 	 	`json:"coniva" bson:"coniva"`
	ConDescuento 	bool	 	`json:"condescuento" bson:"condescuento"`
	Estatus 			string 	`json:"estatus" bson:"estatus"`
	IdServicio		string 	`json:"idservicio" bson:"idservicio"`
	TipoCambio		string 	`json:"tipocambio" bson:"tipocambio"`
	PorcentajeIva	string	`json:"porcentajeiva" bson:"porcentajeiva"`
	TipoConcepto	string	`json:"tipoconcepto" bson:"tipoconcepto"`


}
