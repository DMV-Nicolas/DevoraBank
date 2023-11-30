postgres: 
	docker start techschool
createdb:
	docker exec -it techschool createdb sakurabank
dropdb:
	docker exec -it techschool dropdb sakurabank
installmigrate:
	curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
	echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
	apt-get update
	apt-get install -y migrate
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