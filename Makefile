serve:
	go run cmd/restfull/main.go

migrate-up:
	migrate -database "mysql://root:root@tcp(localhost:3306)/posnote" -path db/migrations up

migrate-down:
	migrate -database "mysql://root:root@tcp(localhost:3306)/posnote" -path db/migrations down
