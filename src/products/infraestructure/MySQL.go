package infraestructure

import (
	"demo/src/core"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save() {
	query := "INSERT INTO productos (name, price) VALUES (?, ?)"
	name := "Big-cola"
	price := 12.50

	result, err := mysql.conn.ExecutePreparedQuery(query, name, price)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
}

func (mysql *MySQL) GetAll() {
	query := "SELECT * FROM productos"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()
	for rows.Next() {
		var idproduct int
		var name string
		var price float32
		if err := rows.Scan(&idproduct, &name, &price); err != nil {
			fmt.Println("error al escanear la fila: %w", err)
		}
		fmt.Printf("ID: %d, Name: %s, Price: %f\n", idproduct, name, price)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("error iterando sobre las filas: %w", err)
	}
}

func (mysql *MySQL) GetByID(id int) {
	query := "SELECT * FROM productos WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()
	for rows.Next() {
		var idproduct int
		var name string
		var price float32
		if err := rows.Scan(&idproduct, &name, &price); err != nil {
			fmt.Println("error al escanear la fila: %w", err)
		}
		fmt.Printf("ID: %d, Name: %s, Price: %f\n", idproduct, name, price)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("error iterando sobre las filas: %w", err)
	}
}

func (mysql *MySQL) Update(id int, name string, price float32) {
	query := "UPDATE productos SET name = ?, price = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, name, price, id)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
}

func (mysql *MySQL) Delete(id int) {
	query := "DELETE FROM productos WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
}