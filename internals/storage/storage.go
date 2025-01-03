package storage

import "github.com/Kishan-Thanki/go-students-rest-api/internals/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	DeleteStudentById(id int64) (bool, error)
	UpdateStudentById(id int64, name string, email string, age int) (types.Student, error)
}
