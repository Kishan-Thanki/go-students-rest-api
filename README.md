# Students REST API in Go

This is a simple REST API built in **Go** for managing student records. It supports basic **CRUD** (Create, Read, Update, Delete) operations using a **SQLite** database. The API is designed with simplicity in mind, leveraging Go's standard libraries for routing and database management.

## **Features**
- **Create a Student**: Add new student records with name, email, and age.
- **Get a Student by ID**: Retrieve student details using their unique ID.
- **Get All Students**: Fetch a list of all students.
- **Update a Student by ID**: Modify an existing student's information.
- **Delete a Student by ID**: Remove a student from the database.

## **Technologies Used**
- **Go**: The programming language used to build the API.
- **SQLite**: Lightweight database for storing student records.
- **http.ServeMux**: Used for routing HTTP requests.
- **slog**: Structured logging for tracking requests and errors.

## **Endpoints**
- **POST** `/api/students`: Create a new student.
- **GET** `/api/students/{id}`: Get a student by ID.
- **GET** `/api/students`: Get all students.
- **PUT** `/api/students/{id}`: Update a student by ID.
- **DELETE** `/api/students/{id}`: Delete a student by ID.
