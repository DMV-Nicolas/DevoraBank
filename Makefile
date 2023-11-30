postgres: 
	docker start techschool
createdb:
	docker exec -it techschool createdb bank
dropdb:
	docker exec -it techschool dropdb bank
migrateup:
	migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/bank?sslmode=disable" -verbose up 
migratedown:
	migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go clean -testcache
	go test -v --cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc