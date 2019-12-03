package banco

type Banco struct {
	Codigo      string `json:"codigo" bson:"codigo"`
	Descripcion string `json:"nombre" bson:"nombre"`
	Abreviatura string `json:"abreviatura" bson:"abreviatura"`
	Estatus     int	   `json:"estatus" bson:"estatus"`
	Acciones    string `json:"acciones" bson:"acciones"`

}
