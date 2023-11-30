postgres: 
	docker start techschool
createdb:
	docker exec -it techschool createdb sakurabank
dropdb:
	docker exec -it techschool dropdb sakurabank
installmigrate:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate-4.16.2.zip | tar xvz
    sudo mv migrate.linux-amd64 /usr/bin/migrate
    which migrate
migrateup:
	migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/sakurabank?sslmode=disable" -verbose up 
migratedown:
	migrate -path db/migration -database "postgresql://root:83postgres19@localhost:5432/sakurabank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go clean -testcache
	go test -v --cover ./...
.PHONY: postgres createdb dropdb installmigrate migrateup migratedown sqlc