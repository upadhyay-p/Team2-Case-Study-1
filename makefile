# End point for DB
EP = --endpoint http://localhost:8000

# Table names
TABLE1 = assets/db_schema/customers.json
TABLE2 = assets/db_schema/orders.json
TABLE3 = assets/db_schema/restaurants.json

# rule to create all tables above
tables:
	aws dynamodb create-table --cli-input-json file://$(TABLE1) $(EP)
	aws dynamodb create-table --cli-input-json file://$(TABLE2) $(EP)
	aws dynamodb create-table --cli-input-json file://$(TABLE3) $(EP)

# rule to list all tables in DB
list-tables:
	aws dynamodb list-tables $(EP)

# rule to delete a given table
delete-table:
	aws dynamodb delete-table --table-name $(T) $(EP)