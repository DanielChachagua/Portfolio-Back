package database

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	"os"
	"time"

	"github.com/DanielChachagua/Portfolio-Back/models"
	"github.com/DanielChachagua/Portfolio-Back/utils"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConectDB(uri string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Project{})

	var email string 
	db.Model(&models.User{}).Select("email").Where("email = ?", os.Getenv("ADMIN_EMAIL")).Scan(&email)

	if email != "" {
		log.Println("El usuario ya existe en la base de datos")
		return db, nil
	}
	newId := uuid.NewString()

	pass, err := utils.HashPassword(os.Getenv("ADMIN_PASSWORD"))

	if err != nil {
		return nil, err
	}

	db.Create(&models.User{ID: newId, FirstName: os.Getenv("ADMIN_FIRST_NAME"), LastName: os.Getenv("ADMIN_LAST_NAME"), Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: pass, UrlImage: "", CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()})

	return db, nil
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("No se pudo obtener la conexión de bajo nivel:", err)
	}

	if sqlDB != nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Error al cerrar la conexión:", err)
		}
	}
	return nil
}

func ExecuteTransaction(db *sql.DB, query string, args ...interface{}) error {
	tx, err := db.Begin()
	if err != nil {
			log.Printf("Error starting transaction: %v", err)
			return err
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
			log.Printf("Error executing query: %v", err)
			tx.Rollback()
			return err
	}

	err = tx.Commit()
	if err != nil {
			log.Printf("Error committing transaction: %v", err)
			return err
	}

	return nil
}

func ExecuteGroupTransactions(db *sql.DB, queries []string, args [][]interface{}) error {
	tx, err := db.Begin()
	if err != nil {
					log.Printf("Error starting transaction: %v", err)
					return err
	}

	for i, query := range queries {
					_, err := tx.Exec(query, args[i]...)
					if err != nil {
									log.Printf("Error executing query: %v", err)
									tx.Rollback()
									return err
					}
	}

	err = tx.Commit()
	if err != nil {
					log.Printf("Error committing transaction: %v", err)
					return err
	}

	return nil
}

func GetRow(db *sql.DB, query string, args ...interface{}) *sql.Row  {
	row := db.QueryRow(query, args...)
	return row
}

func GetRows(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}

	return rows, nil
}

func MapRowsToStruct(rows *sql.Rows, dest interface{}) error {
	// Obtener el tipo del destino (debe ser un slice de estructuras)
	sliceType := reflect.TypeOf(dest).Elem()
	if sliceType.Kind() != reflect.Slice {
		return fmt.Errorf("el destino debe ser un slice de estructuras")
	}

	// Obtener la referencia al slice destino
	sliceValue := reflect.ValueOf(dest).Elem()

	// Obtener las columnas de la consulta
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// Crear un mapa de nombres de columna a índice
	columnIndex := make(map[string]int)
	for i, col := range columns {
		columnIndex[col] = i
	}

	// Iterar sobre las filas
	for rows.Next() {
		// Crear una nueva instancia del struct
		structPtr := reflect.New(sliceType.Elem()) // Crear un nuevo struct del tipo destino
		structValue := structPtr.Elem()

		// Crear un slice de interfaces para Scan
		fieldValues := make([]interface{}, len(columns))
		for i := range fieldValues {
			fieldValues[i] = new(interface{}) // Puntero a interface{} para recibir los valores
		}

		// Escanear los valores en el slice
		if err := rows.Scan(fieldValues...); err != nil {
			return err
		}

		// Asignar valores a los campos del struct
		for i, colName := range columns {
			// Buscar el campo por el nombre de la etiqueta JSON
			for j := 0; j < structValue.NumField(); j++ {
				fieldStruct := structValue.Type().Field(j)
				jsonTag := fieldStruct.Tag.Get("json")
				if jsonTag == colName {
					field := structValue.Field(j)

					// Si el campo es asignable, establecer el valor
					if field.CanSet() {
						val := reflect.ValueOf(*(fieldValues[i].(*interface{})))
						if val.IsValid() {
							field.Set(val.Convert(field.Type())) // Convertir y asignar el valor
						}
					}
					break
				}
			}
		}

		// Agregar el struct al slice destino
		sliceValue.Set(reflect.Append(sliceValue, structValue))
	}

	// Manejar caso de consulta vacía
	if sliceValue.Len() == 0 {
		return sql.ErrNoRows
	}

	return nil
}




func MapRowToStruct(rows *sql.Rows, dest interface{}) error {
	// Obtener el tipo del destino (debe ser un puntero a una estructura)
	destType := reflect.TypeOf(dest)
	if destType.Kind() != reflect.Ptr || destType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("el destino debe ser un puntero a una estructura")
	}

	// Obtener la referencia al valor de destino
	destValue := reflect.ValueOf(dest).Elem()

	// Obtener las columnas de la consulta
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// Crear un slice de interfaces para Scan
	fieldValues := make([]interface{}, len(columns))
	for i := range fieldValues {
		fieldValues[i] = new(interface{}) // Puntero a interface{} para recibir los valores
	}

	// Leer una sola fila
	if !rows.Next() {
		return sql.ErrNoRows
	}

	// Escanear los valores en el slice
	if err := rows.Scan(fieldValues...); err != nil {
		return err
	}

	// Asignar valores a los campos de la estructura
	for i, colName := range columns {
		// Buscar el campo por el nombre de la etiqueta JSON
		for j := 0; j < destValue.NumField(); j++ {
			fieldStruct := destValue.Type().Field(j)
			jsonTag := fieldStruct.Tag.Get("json")
			if jsonTag == colName {
				field := destValue.Field(j)

				// Si el campo es asignable, establecer el valor
				if field.CanSet() {
					val := reflect.ValueOf(*(fieldValues[i].(*interface{})))
					if val.IsValid() {
						field.Set(val.Convert(field.Type())) // Convertir y asignar el valor
					}
				}
				break
			}
		}
	}

	return nil
}