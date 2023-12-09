postgres: 
	docker create -p5432:5432 --name sql1 --network devorabank-network -e POSTGRES_PASSWORD=83postgres19 -e POSTGRES_USER=root postgres:latest
createdb:
	docker exec -it techschool createdb bank
dropdb:
	docker exec -it techschool dropdb bank
migrateup:
	./migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/bank?sslmode=disable" -verbose up 
migratedown:
	./migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/bank?sslmode=disable" -verbose down
migrateup1:
	./migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/bank?sslmode=disable" -verbose up 1
migratedown1:
	./migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go clean -testcache
	go test -v --cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/DMV-Nicolas/DevoraBank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc server mock