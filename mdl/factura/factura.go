package factura

type Factura struct {

	NumeroDocumento			string 							`json:"numerodocumento" bson:"numerodocumento"`
	FechaDocumento			string 							`json:"fechadocumento" bson:"fechadocumento"`
	TipoDocumento				string 							`json:"tipodocumento" bson:"tipodocumento"`
	NumeroSeniat    		bool   							`json:"numeroseniat" bson:"numeroseniat"`
	Monto								string 							`json:"Monto" bson:"monto"`
	MontoIva						string							`json:"MontoIva" bson:"montoiva"`
	EstatusDocumento 		string							`json:"estatusdocumento" bson:"estatusdocumento"`
	FechaDesde					string							`json:"fechadesde" bson:"fechadesde"`
	FechaHasta					string							`json:"FechaHasta" bson:"fechahasta"`
	Observacion					string							`json:"Observacion" bson:"observacion"`
	IdPagado						string							`json:"IdPagado" bson:"idpagado"`
	IdContabilizado			string							`json:"idcontabilizado" bson:"idcontabilizado"`
	FechaEntrega				string							`json:"fechaentrega" bson:"fechaentrega"`
	FechaAnulada				string							`json:"fechaanulada" bson:"fechaanulada"`
	UsuarioAnulada			string							`json:"usuarioanulada" bson:"usuarioanulada"`
	PorcentajeIva				string							`json:"porcentajeiva" bson:"porcentajeiva"`
	Serie								string 							`json:"serie" bson:"serie"`
	MotivoAnulacion			string							`json:"motivoanulacion" bson:"motivoanulacion"`
	MontoBolivares			string							`json:"montobolivares" bson:"montobolivares"`
	MontoIvaBs					string							`json:"montoivabolivares" bson:"montoivabolivares"`
	CodTerminal					string							`json:"terminal" bson:"terminal"`
	BaseImponibleBs			string							`json:"baseimponiblebolivares" bson:"baseimponiblebolivares"`
	ExentoBolivares			string							`json:"exento" bson:"exento"`
	Oficna							string							`json:"ofcina" bson:"oficina"`
	Moneda							string							`json:"moneda" bson:"moneda"`
	MontoDolares				string							`json:"montodolares" bson:"montodolares"`
	MontoIvaDol					string							`json:"montoivadolares" bson:"montoivadolares"`
	BaseImponibleDol		string							`json:"baseimponibledolares" bson:"baseimponibledolares"`
	ExentoDolares				string							`json:"exentodolares" bson:"exentodolares"`
	TasaCambio					string							`json:"tasacambio" bson:"tasacambio"`
	AplicaCambio				string							`json:"aplicacambio" bson:"aplicacambio"`
	FechaAplicacion			string							`json:"fechaaplicacion" bson:"fechaaplicacion"`
	Istasas							bool								`json:"istasas" bson:"istasas"`
	IdFacturaDuplicada	string							`json:"idfacturaduplicada" bson:"idfacturaduplicada"`
	FechaPago						string							`json:"fechapago" bson:"fechapago"`
	ConDescuento				string							`json:"condescuento" bson:"condescuento"`
	MontoDescuentoBol		string							`json:"montodescuentobolivar" bson:"montodescuentobolivar"`
	MontoDescuentoDol		string							`json:"montodescuentodolar" bson:"montodescuentodolar"`
	PorcentajeDescuento	string							`json:"porcentajedescuentodolar" bson:"pocentajedescuentodolar"`
	CodigoDescuento			string							`json:"codigodescuento" bson:"codigodescuento"`
	NombreDescuento			string							`json:"nombredescuento" bson:"nombredescuento"`
	TasaCambioEuro			string							`json:"tasacambioeuro" bson:"tasacambioeuro"`
	AplicaCambioEuro		string							`json:"aplicacambioeuro" bson:"aplicacambioeuro"`
	FechaAplicacionEuro	string							`json:"fechaaplicacambioeuro" bson:"fechaaplicacambioeuro"`
	MontoEuro						string							`json:"montoeuro" bson:"montoeuro"`
	MontoIvaEuro				string							`json:"montoivaeuro" bson:"montoivaeuro"`
	BaseImponibleEuro		string							`json:"baseimponibleeuro" bson:"baseimponibleeuro"`
	ExentoEuro					string							`json:"exentoeuro" bson:"exentoeuro"`


}
