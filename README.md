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


## **How to Run**
To run the API, follow these steps:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Kishan-Thanki/go-students-rest-api.git
   cd go-students-rest-api

2. **Install dependencies (if any)**:
   ```bash
   go mod tidy

3. **Run the application: The following command runs the application with the specified configuration file (config/local.yaml)**:

   ```bash
   go run cmd/students-rest-api/main.go -config config/local.yaml

This will start the API locally, and it will be accessible at http://localhost:8082.


## **How to Build**
To build the application and generate an executable binary, follow these steps:

1. **Build the Go application: Run the following command to compile the code into an executable binary**:

   ```bash
   go build -o bin/student-api cmd/students-rest-api/main.go

This will generate an executable file named student-api in the bin directory.

2. **Run the executable: After building, you can run the API by executing the binary directly**:

   ```bash
   ./bin/student-api -config config/local.yaml

This will start the API in the same way as running with go run, but using the compiled binary.