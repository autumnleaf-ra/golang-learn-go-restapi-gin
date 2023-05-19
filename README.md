# Simple API

Simple API with CRUD feature

## Tech Stack
1. MySQL
2. Golang
3. External packages :
* SQL driver for MySQL 
* Gin
* Gorm

## Prerequisites
Enter the following command to install sql driver for mysql, gin and gorm in the project.
```
go get github.com/go-sql-driver/mysql
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
```

## Running the REST API
```
go run main.go
```

## API Endpoints
| Route         | HTTP          | Description           |
| ------------- |:-------------:| ---------------------:|
| /api/user     | GET           | Get all users         |
| /api/user:id  | GET           | Get single user       |
| /api/user     | POST          | Create a user         |
| /api/user:id  | PUT           | Update data of a user |
| /api/user:id  | DELETE        | Delete a user         |

## Fungsi apa saja yang saya telah buat pada REST API :
1. INDEX - menampilkan data secara keseluruhan
2. SHOW - menampilkan data berdasarkan ID
3. CREATE - menambahkan data
4. PUT - mengupdate data
5. DELETE - menghapus data

## Reference
[Create Your First Rest API With GOLANG using GIN, GORM and MySql](https://medium.com/wesionary-team/create-your-first-rest-api-with-golang-using-gin-gorm-and-mysql-d439bcc6f987)
