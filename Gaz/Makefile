sqlup:
	migrate -path migrations -database "postgresql://root:root@localhost:5432/root?sslmode=disable"  up
sqldown:
	migrate -path migrations -database "postgresql://root:root@localhost:5432/root?sslmode=disable"  down
run:
	go run cmd/main.go
.PHONY: sqlup sqldown run
