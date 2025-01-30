package infraestructure

import (
	"demo/src/core"
	//"fmt"
	"log"
)

type MySQLSabrita struct {
	conn *core.Conn_MySQL
}

func NewMySQLSabrita() *MySQLSabrita {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQLSabrita{conn: conn}
}

// Recibir name y price como parámetros
func (mysql *MySQLSabrita) Save(name string, price float32) {
	query := "INSERT INTO sabritas (name, price) VALUES (?, ?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, name, price)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Sabrita guardada: %d", rowsAffected)
	}
}

func (mysql *MySQLSabrita) GetAll() ([]map[string]interface{}, error) {
	query := "SELECT * FROM sabritas"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var sabritas []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		var price float32
		if err := rows.Scan(&id, &name, &price); err != nil {
			return nil, err
		}
		refresco := map[string]interface{}{
			"id":    id,
			"name":  name,
			"price": price,
		}
		sabritas = append(sabritas, refresco)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sabritas, nil
}

func (mysql *MySQLSabrita) GetByID(id int) (map[string]interface{}, error) {
	query := "SELECT * FROM sabritas WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var sabritas map[string]interface{}
	if rows.Next() {
		var id int
		var name string
		var price float32
		if err := rows.Scan(&id, &name, &price); err != nil {
			return nil, err
		}
		sabritas = map[string]interface{}{
			"id":    id,
			"name":  name,
			"price": price,
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sabritas, nil
}

func (mysql *MySQLSabrita) Update(id int, name string, price float32) {
	query := "UPDATE sabritas SET name = ?, price = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, name, price, id)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - sabrita actualizada: %d", rowsAffected)
	}
}

func (mysql *MySQLSabrita) Delete(id int) {
	query := "DELETE FROM sabritas WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - sabrita eliminada: %d", rowsAffected)
	}
}