package banco

type Banco struct {
	Codigo      string `json:"codigo" bson:"codigo"`
	Nombre      string `json:"nombre" bson:"nombre"`
	Abreviatura string `json:"abreviatura" bson:"abreviatura"`
	Estatus     bool   `json:"estatus" bson:"estatus"`
	Acciones    string `json:"acciones" bson:"acciones"`

}
