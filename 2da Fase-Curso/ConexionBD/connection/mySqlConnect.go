package connection

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DbMysql *sql.DB

func Conexion() {

	errLoadEnv := godotenv.Load()

	if errLoadEnv != nil {
		log.Panicln("Error dirante la carga de env ...", errLoadEnv)
	}

	connetionBd, errorBd := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))

	if errorBd != nil {
		log.Panicln("Ocurrió un error al abrir la conexión a BD", errorBd)
	}

	DbMysql = connetionBd
}

func CloseConnection() {
	DbMysql.Close()
}
