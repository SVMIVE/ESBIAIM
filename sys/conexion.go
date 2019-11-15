//configuraciones del sistema
package sys

import (
	"database/sql"
	"fmt"

	mgo "gopkg.in/mgo.v2"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/svmive/esbiaim/util"
)

//MongoDBConexion Conexion a Mongo DB
func MongoDBConexion(mapa map[string]CadenaDeConexion) {
	c := mapa["mongodb"]
	MGOSession, Error = mgo.Dial(c.Host + ":27017")
	fmt.Println("Cargando Conexión Con MongoDB...")
	util.Error(Error)
}

//ConexionSAMAN Funcion de Conexion a Postgres
func ConexionSINIV(mapa map[string]CadenaDeConexion) {
	c := mapa["siniv"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PostgreSQLSINIV, _ = sql.Open("postgres", cadena)
	if PostgreSQLSINIV.Ping() != nil {
		fmt.Println("[Siniv:   Error...] ", PostgreSQLSINIV.Ping())
	} else {
		fmt.Println("[Siniv: ", c.Host, "  OK...]")
	}
}

//ConexionMYSQL
func ConexionMYSQL(mapa map[string]CadenaDeConexion) {
	c := mapa["mysql"]
	cadena := c.Usuario + ":" + c.Clave + "@tcp(" + c.Host + ":3306)/sssifanb"
	MysqlFullText, _ = sql.Open("mysql", cadena)
	if MysqlFullText.Ping() != nil {
		fmt.Println("[mysql FULLTEXT: Error...] ", MysqlFullText.Ping())
	} else {
		fmt.Println("[mysql FULLTEXT: OK...]")
	}
	return
}
