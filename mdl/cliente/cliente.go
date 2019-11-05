package cliente

type Cliente struct {
	Codigo      string `json:"codigo" bson:"codigo"`
	Nombre      string `json:"nombre" bson:"nombre"`
	Lote        string `json:"lote" bson:"lote"`
	RazonSocial string `json:"razonsocial" bson:"razonsocial"`
	Rif         string `json:"rif" bson:"rif"`
	Impresion   bool   `json:"impresion" bson:"impresion"`
	FormaPago   string `json:"formapago" bson:"formapago"`
}
