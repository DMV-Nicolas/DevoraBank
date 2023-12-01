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
.PHONY: postgres createdb dropdb migrateup migratedown sqlc server