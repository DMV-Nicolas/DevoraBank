postgres: 
	docker start techschool
createdb:
	docker exec -it techschool createdb sakurabank
dropdb:
	docker exec -it techschool dropdb sakurabank
migrateup:
	./migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/sakurabank?sslmode=disable" -verbose up 
migratedown:
	./migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/sakurabank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go clean -testcache
	go test -v --cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/DMV-Nicolas/sakurabank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock