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

	Enrutador.HandleFunc("/", Principal)
	Enrutador.HandleFunc("/iaim/api/", APi.Consultar).Methods("GET")
	Enrutador.HandleFunc("/iaim/api/dosa/listar", APi.ListarDosa).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/dosa/listar", APi.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/iaim/api/sybase/listardocumentos", sybase.ListarDocumentos).Methods("POST")

}

//CargarModulosSeguridad Y cifrado
func CargarModulosSeguridad() {
	var wUsuario api.WUsuario
	// Enrutador.HandleFunc("/ipsfa/app/api/wusuario/crud/{id}", wUsuario.Consultar).Methods("GET")
	Enrutador.HandleFunc("/iaim/api/wusuario/login", wUsuario.Login).Methods("POST")
	Enrutador.HandleFunc("/iaim/api/wusuario/login", wUsuario.Opciones).Methods("OPTIONS")
	// Enrutador.HandleFunc("/ipsfa/api/wusuario/validar", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("POST")
	// Enrutador.HandleFunc("/ipsfa/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")
	//
	// Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.Crear).Methods("POST")
	// Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.CambiarClave)).Methods("PUT")
	// Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.Opciones)).Methods("OPTIONS")
	//
	// Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.Crear)).Methods("POST")
	// Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.CambiarClave)).Methods("PUT")
	// Enrutador.HandleFunc("/ipsfa/api/wusuario", wUsuario.ValidarToken(wUsuario.Opciones)).Methods("OPTIONS")
	// Enrutador.HandleFunc("/ipsfa/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")
	//
	// Enrutador.HandleFunc("/ipsfa/api/wusuario/validarphp", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("GET")
}

//Principal Página inicial del sistema o bienvenida
func Principal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Saludos bienvenidos al Bus Empresarial de Datos")
}

//WMAdminLTE OpenSource tema de panel de control
//Tecnología Bootstrap3
func WMAdminLTE() {
	fmt.Println("Cargando Modulos de AdminLTE...")
	// var GP = GPanel{}
	// Enrutador.HandleFunc("/sssifanb/{id}", GP.IrA)
	prefix := http.StripPrefix("/iaim", http.FileServer(http.Dir("public_web/SSSIFANB")))
	Enrutador.PathPrefix("/iaim/").Handler(prefix)
	// prefixx := http.StripPrefix("/bdse-admin/public/temp", http.FileServer(http.Dir("public/temp")))
	// Enrutador.PathPrefix("/bdse-admin/public/temp/").Handler(prefixx)
}

//CargarModulosWebDevel Cargador de modulos web
func CargarModulosWebDevel() {

}

//CargarModulosWebSite Cargador de modulos web
func CargarModulosWebSite() {

}
