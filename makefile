# End point for DB [remove this to push changes as oer your aws configuration]
EP = --endpoint=http://localhost:8000

# Table names
TABLE1 = assets/db_schema/customers.json
TABLE2 = assets/db_schema/orders.json
TABLE3 = assets/db_schema/restaurants.json

# RULES FOR LOCAL TESTING OF DB
# rule to create all tables above
tables:
	aws dynamodb create-table --cli-input-json file://$(TABLE1) $(EP)
	aws dynamodb create-table --cli-input-json file://$(TABLE2) $(EP)
	aws dynamodb create-table --cli-input-json file://$(TABLE3) $(EP)

# rule to populate DB with sample data
populate-db:
	go run cmd/DB/main.go assets/sample_data

# one single rule to INIT-DB for testing APIs
db:
	make tables populate-db

# rule to list all tables in DB
list-tables:
	aws dynamodb list-tables $(EP)

# rule to describe a table in DB
show-table:
	aws dynamodb describe-table --table-name $(T) $(EP)

# rule to delete a given table
delete-table:
	aws dynamodb delete-table --table-name $(T) $(EP)

clean:
	rm -fr bin

# RULES FOR BUILDING APPLICATION
# rule to generate stub code from proto files
protos:
	protoc pkg/protos/* --go_out=plugins=grpc:.

# rule to build gin-API-Client
api:
	go build -o bin/api cmd/API/main.go

# rule to build gRPC-Server
server:
	go build -o bin/server cmd/SERVER/main.go

# one single rule to build application ready to run
app:
	make api server

# one single rule to buid application ready to run on linux
app-linux:
	env GOOS=linux make app

docker-server:
	docker build -f Dockerfile.server mygoapp-server .

docker-api:
	docker build -f Dockerfile.api mygoapp-api .

api-tests:
	go test cmd/API -coverprofile=coverage.out