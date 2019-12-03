package cliente

type Cliente struct {
	Codigo      					string `json:"codigo" bson:"codigo"`
	Nombre      					string `json:"nombre" bson:"nombre"`
	RazonSocial 					string `json:"razonsocial" bson:"razonsocial"`
	Nit         					string `json:"nit" bson:"nit"`
	Rif        	 					string `json:"rif" bson:"rif"`
	TipoCliente 					string `json:"tipocliente" bson:"tipocliente"`
	Estatus								string `json:"estatus" bson:"estatus"`
	Actvidad	  					string   `json:"actividad" bson:"actividad"`
	IngresosBrutos 				string   `json:"ingresosbrutos" bson:"ingresosbrutos"`
	FechaInicioContrato 	string `json:"fechainiciocontrato" bson:"fechainiciocontrato"`
	FechaModifContrato	 	string `json:"fechamodificacioncontrato" bson:"fechamodificacioncontrato"`
	Email   							string `json:"formapago" bson:"formapago"`
	Telefono   						string `json:"formapago" bson:"formapago"`
	CodigoPostal   				string `json:"formapago" bson:"formapago"`
	DireccionFiscal				string `json:"formapago" bson:"formapago"`
	FormaPago   					string `json:"formapago" bson:"formapago"`
	Lote        					string `json:"lote" bson:"lote"`

}
