# Soal 2 HTTP Validasi (Tugas Pekan 2)

This project is my answer for **Soal 2 HTTP Validasi (Tugas Pekan 2)** where we (the students) are asked to create a HTTP server, `GET /validate` endpoint and validate the endpoint's query.

## How To Play

Run this project using `go run main.go` and the terminal will give you an output like this

```
Server running at http://localhost:8080
```

You can open this url `http://localhost:8080/validate?email=user@example.com&age=20` directly in your browser or use `curl` on your terminal:

```
curl http://localhost:8080/validate?email=user@example.com&age=20
```

If your `email` is not empty and `umur` is greater than or equal to 18, you will see a JSON output like this:

```
{
  "status": "ok"
}
```

Otherwise the output will be like this:

```
{
  "status": "email kosong atau umur kurang dari 18"
}
```

## Endpoint

### 1. Validate Input

#### Method:

`GET`

#### Endpoint:

`/validate`

#### Query:

- `email=string` Your email

- `umur=number` Your age 
