# Privy ID Test Golang REST API - 26 Oct 2022

List of used library:
- Gorm 
- Gin Gonic 
- Jwt
- Godotenv


First, download all library:

```sh
go get 
```

Then, adjust all _environtment_ at .env 


Next run the docker compose for create database

```sh
docker-compose up
```

After the database is running, then build the application and run it with the following command:

```sh
go build
./privyid-golang-test
```

Now, use an HTTP client (like [Postman](https://www.getpostman.com/apps)) and hit login endpoint:

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

Then, hit user balance endpoint again to top up the balance:

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

Route available:
- ```(POST) /api/auth/login``` | Required body: email _(string)_, password _(string)_
- ```(GET) /api/auth/logout``` | Required header: authorization 
- ```(POST) /api/auth/balance``` | Required header: authorization 
- ```(POST) /api/auth/top-up``` | Required header: authorization; body: amount _(int)_
- ```(POST) /api/auth/transfer``` | Required header: authorization; body: amount _(int)_, to (username) _(string)_

