password = secret
path = "postgresql://root:secret@localhost:5433/university?sslmode=disable"
migratepath = go/db/migrations


# Creates a database and initializes it 
initdb: check-name
	docker run --name ${name} -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=${password} -d postgres:14-alpine
	timeout 3
	docker exec universityDB psql -U root -d root -c "CREATE DATABASE university;"
	make migrateup
	
# migrates database into newest iteration 
migrateup:
	migrate -path ${migratepath} -database ${path} -verbose up

# migrates database down (complete reset)
migratedown:
	migrate -path ${migratepath} -database ${path} -verbose down

updb:
	make sqlc
	make mock

# Check if an env variable is specified in the command
check-name:
ifndef name 
	$(error name is undefined)
endif

.ONESHELL:
# runs sqlc
sqlc:
	cd go/db/
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

# runs go testing with coverage
test: 
	cd go/
	go test -v -cover ./...

# runs gin API 
server:
	cd go/
	go run .

# runs vue
vue: 
	cd vue/university
	npm run dev

# runs mockgen. for api testing
mock:
	cd go
	mockgen -package mockdb -destination db/mock/store.go github.com/cyanrad/university/db/sqlc Store

.PHONY: initdatabase check-name migratedown migrateup sqlc vue mock updb