package actividad

type Actividad struct {
	Codigo       string `json:"codigo" bson:"codigo"`
	Nombre       string `json:"nombre" bson:"nombre"`
	Convenio     string `json:"convenio" bson:"convenio"`
  Moneda       string `json:"moneda" bson:"moneda"`


}
