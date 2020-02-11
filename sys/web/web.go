package web

//Copyright Carlos Peña
//Modulo de negociación WEB
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/svmive/esbiaim/sys/web/api"
)

//Variables de Control
var (
	Enrutador   = mux.NewRouter()
	WsEnrutador = mux.NewRouter()
)

//Cargar los diferentes modulos del sistema
func Cargar() {
	CargarModulosWeb()

	CargarModulosSeguridad()

	WMAdminLTE()
	CargarModulosWebDevel()
	CargarModulosWebSite()

}

//CargarModulosWeb Cargador de modulos web
func CargarModulosWeb() {
	var APi api.API
	var sybase api.WSyBase
	var tasaC api.TasaCambio
	var wFact api.WFactura
	var wBanc api.WBanco
	var wClient api.WCliente
	var wConcep api.WConcepto
	var wActivi api.WActividad
	var wDocume api.WDocumento
	var wServi api.WServicio
	var wAdminC api.WAdminControl

	//Enrutador.HandleFunc("/", Principal)
	Enrutador.HandleFunc("/iaim/api/", APi.Consultar).Methods("GET")
	//
	Enrutador.HandleFunc("/iaim/api/dosa/listar", APi.ListarDosa).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/dosa/listar", APi.Opciones).Methods("OPTIONS")
	//
	Enrutador.HandleFunc("/iaim/api/dosa/noprocesadas", APi.NoProcesadas).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/dosa/noprocesadas", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/listardocumentos", sybase.ListarDocumentos).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/sybase/listardocumentos", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/listartasa", tasaC.Listar).Methods("GET")
	Enrutador.HandleFunc("/iaim/api/sybase/listartasa", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/factura/listar", wFact.Listar).Methods("GET")
	Enrutador.HandleFunc("/iaim/api/sybase/factura/listar", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/banco/listar", wBanc.Listar).Methods("GET")
	Enrutador.HandleFunc("/iaim/api/sybase/banco/listar", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/concepto/listar", wConcep.Listar).Methods("GET")
	Enrutador.HandleFunc("/iaim/api/sybase/concepto/listar", APi.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/iaim/api/sybase/concepto/consultar", wConcep.Consultar).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/sybase/concepto/consultar", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/banco/insertar", wBanc.Agregar).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/sybase/banco/insertar", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/cliente/lstactividad", wClient.ListarActividad).Methods("GET")
	Enrutador.HandleFunc("/iaim/api/sybase/cliente/lstactividad", APi.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/iaim/api/sybase/cliente/listar", wClient.Listar).Methods("GET")
	Enrutador.HandleFunc("/iaim/api/sybase/cliente/listar", APi.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/iaim/api/sybase/cliente/pagos", wClient.Pagos).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/sybase/cliente/pagos", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/cliente/insertar", wClient.Insertar).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/sybase/cliente/insertar", APi.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/iaim/api/sybase/cliente/consultar", wClient.Consultar).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/sybase/cliente/consultar", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/cliente/razonsocial", wClient.RazonSocial).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/sybase/cliente/razonsocial", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/actividad/listar", wActivi.Listar).Methods("GET")
	Enrutador.HandleFunc("/iaim/api/sybase/actividad/listar", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/servicio/listar", wServi.Listar).Methods("GET")
	Enrutador.HandleFunc("/iaim/api/sybase/servicio/listar", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/documento/lstdocactivos", wDocume.ListarDocActivos).Methods("GET")
	Enrutador.HandleFunc("/iaim/api/sybase/documento/lstdocactivos", APi.Opciones).Methods("OPTIONS")

	//----------------------------- AdminControl ---------------------------------//
	Enrutador.HandleFunc("/iaim/api/sybase/admincontrol/autoincrement", wAdminC.AutoIncrement).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/sybase/admincontrol/autoincrement", APi.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/iaim/api/sybase/admincontrol/insertjoin", wAdminC.InsertJoin).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/sybase/admincontrol/insertjoin", APi.Opciones).Methods("OPTIONS")
	Enrutador.HandleFunc("/iaim/api/sybase/admincontrol/insertinto", wAdminC.InsertInto).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/sybase/admincontrol/insertinto", APi.Opciones).Methods("OPTIONS")


}

//CargarModulosSeguridad Y cifrado
func CargarModulosSeguridad() {
	var wUsuario api.WUsuario
	Enrutador.HandleFunc("/iaim/api/wusuario/login", wUsuario.Login).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/wusuario/login", wUsuario.Opciones).Methods("OPTIONS")
	// Enrutador.HandleFunc("/ipsfa/api/wusuario/validar", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("POST")
	// Enrutador.HandleFunc("/ipsfa/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")

}

//Principal Página inicial del sistema o bienvenida
func Principal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Saludos bienvenidos al Bus Empresarial de Datos")
}

//WMAdminLTE OpenSource tema de panel de control
//Tecnología Bootstrap3
func WMAdminLTE() {
	fmt.Println("Cargando Modulos de Ngx-Recaudaciones...")
	prefix := http.StripPrefix("/", http.FileServer(http.Dir("public_web/iaim/")))
	Enrutador.PathPrefix("/").Handler(prefix)
}

//CargarModulosWebDevel Cargador de modulos web
func CargarModulosWebDevel() {

}

//CargarModulosWebSite Cargador de modulos web
func CargarModulosWebSite() {

}
