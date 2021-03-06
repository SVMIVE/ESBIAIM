package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/svmive/esbiaim/mdl/dosa"
	"github.com/svmive/esbiaim/mdl/sssifanb"
	"github.com/svmive/esbiaim/mdl/sssifanb/fanb"
	"github.com/svmive/esbiaim/sys"
	"github.com/svmive/esbiaim/sys/seguridad"
	"github.com/svmive/esbiaim/util"
)

var UsuarioConectado seguridad.Usuario

//API Estructura
type API struct {
	Ruta     string
	Conexion string
	Tiempo   time.Time
}

//DOSA Estructura
type WDosa struct {
	Desde string `json:"desde"`
	Hasta string `json:"hasta"`
	//Tiempo time.Time
}

//WSyBase Estructura
type WSyBase struct {
	Desde string `json:"desde"`
	Hasta string `json:"hasta"`
}

func (A *API) Consultar(w http.ResponseWriter, r *http.Request) {
	var a API
	Cabecera(w, r)
	ip := strings.Split(r.RemoteAddr, ":")
	a.Conexion = "Conexion Remota desde el host " + ip[0]
	a.Tiempo = time.Now()
	a.Ruta = "/iaim/"

	j, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

func (A *API) ListarDosa(w http.ResponseWriter, r *http.Request) {

	var dosa dosa.DOSA
	Cabecera(w, r)
	//fmt.Println("Conexion")
	//var wdosa WDosa
	// e := json.NewDecoder(r.Body).Decode(&wdosa)
	// util.Error(e)
	fmt.Printf("Bien pues")
	j, err := dosa.Listar()
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error de conexión"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	return
}

// NoProcesadas
func (A *API) NoProcesadas(w http.ResponseWriter, r *http.Request) {
	var dosa dosa.DOSA
	Cabecera(w, r)
	j, err := dosa.NoProcesadas()
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error de conexión"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	return
}

//Opciones Militar
func (A *API) Opciones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
}

func (wyb *WSyBase) SubirArchivos(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var traza fanb.Traza
	var M sssifanb.Mensaje
	var militarF sssifanb.Militar

	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login

	er := r.ParseMultipartForm(32 << 20)
	if er != nil {
		fmt.Println(er)
		return
	}
	m := r.MultipartForm
	files := m.File["archivo"]
	cedula := r.FormValue("txtFileID")
	directorio := "./public_web/SSSIFANB/afiliacion/temp/" + cedula + "/"
	if cedula == "" {
		M.Mensaje = "Carga fallida"
		M.Tipo = -1
		j, _ := json.Marshal(M)
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		return
	} else if cedula == "DESERTOR" {
		directorio = "./tmp/"

	}

	errr := os.Mkdir(directorio, 0775)
	if errr != nil {
		fmt.Println("la carpeta ya existe.")
	}
	cadena := ""
	for i, _ := range files {
		file, errf := files[i].Open()
		defer file.Close()
		if errf != nil {
			fmt.Println(errf)
			return
		}
		out, er := os.Create(directorio + files[i].Filename)
		defer out.Close()
		if er != nil {
			fmt.Println("No se pudo escribir el archivo verifique los privilegios.")
			return
		}
		// fmt.Println("Subiendo Archivo")
		_, err := io.Copy(out, file) // file not files[i] !
		if err != nil {
			fmt.Println("Entrando en un Erro...", err)
			return
		}
		// fmt.Println("Instertando....")
		cadena += files[i].Filename + ";"
		if cedula == "DESERTOR" {
			var desertor sssifanb.Desertor
			go desertor.LeerDesertores(files[i].Filename)
		}

	}

	if UsuarioConectado.Login[:3] != "act" {
		traza.Documento = "Agregando Historial Digital ( " + cedula + " )"
		traza.Log = cadena
		traza.CrearHistoricoConsulta("hmilitar")
		M.Mensaje = "Carga exitosa"
		M.Tipo = 2
		militarF.ActualizarFoto(cedula)
	} else {
		M.Mensaje = "Carga fallida"
		M.Tipo = -1
	}

	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//SubirArchivosTXTPensiones Pensionados
func (wyb *WSyBase) SubirArchivosTXTPensiones(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var traza fanb.Traza
	var M sssifanb.Mensaje

	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login

	er := r.ParseMultipartForm(32 << 20)
	if er != nil {
		fmt.Println(er)
		return
	}
	m := r.MultipartForm
	files := m.File["input-folder-2"]
	codigo := r.FormValue("txtFileID")
	directorio := "./public_web/SSSIFANB/pensiones/temp/nomina/"
	errr := os.Mkdir(directorio, 0777)
	if errr != nil {
		fmt.Println("El directorio ya existe!")
	}
	cadena := ""
	for i, _ := range files {
		file, errf := files[i].Open()
		defer file.Close()
		if errf != nil {
			fmt.Println(errf)
			return
		}
		out, er := os.Create(directorio + files[i].Filename)
		defer out.Close()
		if er != nil {
			fmt.Println(er.Error())
			return
		}
		_, err := io.Copy(out, file) // file not files[i] !
		if err != nil {
			fmt.Println(err)
			return
		}
		cadena += files[i].Filename + ";"
		ProcesarTxt(files[i].Filename, codigo)

	} // Fin de archivos
	M.Mensaje = cadena
	M.Tipo = 2

	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//ProcesarTxt Proceso de archivo
func ProcesarTxt(doc string, codigo string) {
	var a util.Archivo
	a.Ruta = "./public_web/SSSIFANB/pensiones/temp/nomina/" + doc
	a.LeerCA(sys.PostgreSQLPENSION, codigo)

}
