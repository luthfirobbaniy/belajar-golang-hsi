# Tugas Pertemuan 5

This project is my answer for **Tugas Pertemuan 5** where we (the students) are asked to create a **Saga Choreography Implementation (Kafka): Proses Pendaftaran Mahasiswa Baru**.

## How To Play

1. Run the Docker compose using `docker compose up"

2. Run `academic-service` and `finance-service` using `go run main.go`

3. Once the services are running, open the Kafka UI, make sure the Consumer Groups (of these 2 services) are ready (status: `Stable`)
   
4. If all ready, run the `student-service` using `go run main.go`

5. See the services logs to see the activities of these services.
