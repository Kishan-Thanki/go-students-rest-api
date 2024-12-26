package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Kishan-Thanki/go-students-rest-api/internals/config"
	"github.com/Kishan-Thanki/go-students-rest-api/internals/types"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*SQLite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &SQLite{
		Db: db,
	}, nil

}

func (S *SQLite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := S.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (S *SQLite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := S.Db.Prepare("SELECT id, name, email, age FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query Error: %w", err)
	}

	return student, nil
}

func (S *SQLite) GetStudents() ([]types.Student, error) {
	stmt, err := S.Db.Prepare("SELECT id, name, email, age FROM students")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student
	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (S *SQLite) UpdateStudentById(id int64, name string, email string, age int) (types.Student, error) {
	stmt, err := S.Db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to prepare update statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age, id)
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to execute update statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to retrieve affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return types.Student{}, fmt.Errorf("no student found with id %d", id)
	}

	// Retrieve the updated student record
	return S.GetStudentById(id)
}

func (S *SQLite) DeleteStudentById(id int64) (bool, error) {
	stmt, err := S.Db.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		return false, fmt.Errorf("failed to prepare delete statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return false, fmt.Errorf("failed to execute delete statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to retrieve affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return false, fmt.Errorf("no student found with id %d", id)
	}

	return true, nil
}
