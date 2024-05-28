# create schema
# go run -mod=mod entgo.io/ent/cmd/ent new User

# Generate ent model api
generate:
	go generate ./ent

# Make migrations
migrations:
	atlas migrate diff $(filter-out $@,$(MAKECMDGOALS)) \
		--dir "file://ent/migrate/migrations" \
		--to "ent://ent/schema" \
		--dev-url "sqlite://file?mode=memory&_fk=1"

# Apply migration
migrate:
	atlas migrate apply \
		--dir "file://ent/migrate/migrations" \
		--url "sqlite://test.sqlite?_fk=1"

# Run the project
run:
	go run main.go