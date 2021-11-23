# Starts the database container.
db-start:
	docker-compose up -d database

# Stops and removes the database container.
db-stop:
	docker-compose down database

# Stops, removes and starts the database container.
db-restart:
	make db-stop
	make db-start

# Creates the database.
db-create:
	docker exec -it budget_database_dev createdb --username=postgres --owner=postgres budget

# Drops the database.
db-drop:
	docker exec -it budget_database_dev dropdb --username=postgres budget -f

# Drops and creates the database.
db-reset:
	make db-drop
	make db-create

# Seeds the database.
db-seed:
	go run . --seed

# Create the test database.
db-test-create:
	docker exec -it budget_database_dev createdb --username=postgres --owner=postgres budget_test

# Drops the test database.
db-test-drop:
	docker exec -it budget_database_dev dropdb --username=postgres budget_test -f

# Drops and creates the test database.
db-test-reset:
	make db-drop
	make db-create

# Executes the tests.
test:
	go test ./...

.PHONY: db-start db-stop db-restart db-create db-drop db-reset db-seed db-test-create db-test-drop db-test-reset test
