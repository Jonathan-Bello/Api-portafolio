package storage

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	// once nos sirve para crear el singleton
	once sync.Once
)

// Lógica del factory
// Driver of storage
type Driver string

const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

// newPostgresDB create the conextions with postgreSQL
func newPostgresDB() {
	// once.Do asegura que la función solo se llame una sola vez
	once.Do(func() {
		var err error
		// dsn := "postgres://edteam:edteam@localhost:5432/godb?sslmode=disable"

		// dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
		dsn := os.Getenv("DATABASE_URL")

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		// Controlamos el posible error, al realizar la conexión
		if err != nil {
			log.Fatalf("no se puedo conectar con la base de datos: %v", err)
		}

		// Para verificar que solo se conecta una vez(se imprime una sola vez esta linea)
		fmt.Println("conectado a postgres")
	})
}

// newMySQLDB crea la conexión la base de datos
func newMySQLDB() {
	// once.Do asegura que la función solo se llame una sola vez
	once.Do(func() {
		var err error

		// user:password@/dbname
		// dsn := "root:root@tcp(localhost:3306)/godb?parseTime=true"
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		// Controlamos el posible error, al realizar la conexión
		if err != nil {
			log.Fatalf("no se puedo conectar con la base de datos: %v", err)
		}

		// Para verificar que solo se conecta una vez(se imprime una sola vez esta linea)
		fmt.Println("conectado a MySql")
	})
}

// DB return a unique instance of gorm.db
func DB() *gorm.DB {
	return db
}
