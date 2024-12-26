include .env

# Database URL, you can override this by exporting DB_URL in your environment
DB_URL ?= postgres://user:password@localhost:5432/dbname?sslmode=disable

# Directory where migration files are stored
MIGRATIONS_DIR ?= db/migrations

# Command to create a new migration file with a positional argument
migrate-new:
	@echo "Creating new migration: $(filter-out $@,$(MAKECMDGOALS))"
	migrate create -ext sql -dir $(MIGRATIONS_DIR) $(filter-out $@,$(MAKECMDGOALS))

# Command to run migrations up to the latest version
migrate-up:
	@echo "Running migrations up"
	migrate -path "$(MIGRATIONS_DIR)" -database "$(DB_URL)" up

# Command to rollback the last migration
migrate-down:
	@echo "Rolling back the last migration"
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

# Command to reset the entire database
migrate-reset:
	@echo "Resetting the database"
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" drop
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Command to show the migration status
migrate-status:
	@echo "Showing migration status"
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

# Prevent make from thinking 'create_user_table' is a file and trying to build it
%:
	@:
