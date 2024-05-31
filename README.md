# final-task-pbi-rakamin-fullstack-amosaugusto

# Backend API Project with Golang

### Technology used

- **GoLang**: The main programming language used to develop the API.
- **Gin**: Framework to create API endpoints.
- **GORM**: ORM (Object-Relational Mapping) for Go, used to interact with the database.
- **Golang JWT**: Library to implement JWT (JSON Web Tokens) based authentication.
- **MySql**: Database Management System (DBMS)

### Features

- User registration and login
- JWT authentication
- Update User
- Delete User
- Logout User
- Get photos
- Upload profile photos
- Update profile photos
- Delete profile photos

// User <br>
Register: http://localhost:8080/users/register (POST)<br>
Login: http://localhost:8080/users/login (POST)<br>
Update User: http://localhost:8080/users/:id (PUT)<br>
Delete User: http://localhost:8080/users/:id (DELETE)<br>
Logout: http://localhost:8080/users/logout (POST)<br>

// Photo<br>
Get Photos: http://localhost:8080/photos (GET)<br>
Create Photo: http://localhost:8080/photos (POST)<br>
Update Photo: http://localhost:8080/photos/:id (PUT)<br>
Delete Photo: http://localhost:8080/photos/:id (DELETE)<br>

You can use Postman app to make a request to the API. When using POST method, you can parse some data in JSON format. For example, you can write the data like this for registering a new user:<br>
{<br>
&nbsp;&nbsp;&nbsp;&nbsp;"username": "carljohnson",<br>
&nbsp;&nbsp;&nbsp;&nbsp;"email": "carljohnson@gmail.com",<br>
&nbsp;&nbsp;&nbsp;&nbsp;"password": "qweqwe"<br>
}
