# Golang REST API

Daftar library yang digunakan:
- Gorm 
- Gin Gonic 
- Jwt
- Godotenv


Pertama, download semua library:

```sh
go get 
```

Lalu sesuaikan seluruh _environtment_ pada .env 


Kemudian jalankan file docker untuk database

```sh
docker-compose up
```

Setelah database sudah running, kemudian build aplikasi dan jalankan dengan perintah berikut:

```sh
go build
./privyid-golang-test
```

Sekarang, gunakan HTTP client (seperti [Postman](https://www.getpostman.com/apps)) dan hit api utk login:

```
POST http://localhost:9090/api/auth/login
------------------------------------------
Request Body:
{
  "email":"felix123@mail.com",
  "password":"123456"
}
 
Response:
{
  "status": "success",
  "data": {
    "token_type": "Bearer",
    "expires_in": 1669315399,
    "access_token": "<token>"
  }
}
```

Lalu tes kembali hit api utk top up balance:

```
GET http://localhost:9090/api/user/balance
--------------------------------------------------
Request Header: Bearer <token>
 
Response:
{
  "status": "success",
  "data": {
    "balance": 150,
    "balance_achieve": 0
  }
}
```

Format header authorization: ```Authorization: Bearer <your_token>``` 

Route yang tersedia:
- ```(POST) /api/auth/login``` | Required body: email _(string)_, password _(string)_
- ```(GET) /api/auth/logout``` | Required header: authorization 
- ```(POST) /api/auth/balance``` | Required header: authorization 
- ```(POST) /api/auth/top-up``` | Required header: authorization; body: amount _(int)_
- ```(POST) /api/auth/transfer``` | Required header: authorization; body: amount _(int)_, to (username) _(string)_

