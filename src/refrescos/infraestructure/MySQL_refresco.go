package infraestructure

import (
	"demo/src/core"
	"log"
)

type MySQLRefresco struct {
	conn *core.Conn_MySQL
}

func NewMySQLRefresco() *MySQLRefresco {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQLRefresco{conn: conn}
}

func (mysql *MySQLRefresco) Save(name string, price float32) {
	query := "INSERT INTO refrescos (name, price) VALUES (?, ?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, name, price)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Refresco guardado: %d", rowsAffected)
	}
}

func (mysql *MySQLRefresco) GetAll() ([]map[string]interface{}, error) {
	query := "SELECT * FROM refrescos"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var refrescos []map[string]interface{}
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
		refrescos = append(refrescos, refresco)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return refrescos, nil
}

func (mysql *MySQLRefresco) GetByID(id int) (map[string]interface{}, error) {
	query := "SELECT * FROM refrescos WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var refresco map[string]interface{}
	if rows.Next() {
		var id int
		var name string
		var price float32
		if err := rows.Scan(&id, &name, &price); err != nil {
			return nil, err
		}
		refresco = map[string]interface{}{
			"id":    id,
			"name":  name,
			"price": price,
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return refresco, nil
}
func (mysql *MySQLRefresco) Update(id int, name string, price float32) {
	query := "UPDATE refrescos SET name = ?, price = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, name, price, id)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Refresco actualizado exitosamente: %d", rowsAffected)
	}
}

func (mysql *MySQLRefresco) Delete(id int) {
	query := "DELETE FROM refrescos WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Refresco eliminado exitosamente: %d", rowsAffected)
	}
}