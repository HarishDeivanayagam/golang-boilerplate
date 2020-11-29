do_migrations:
	migrate -path migrations -database mysql://root:root@tcp(localhost:3306)/boilerplate up
reverse_migrations:
	migrate -path migrations -database mysql://root:root@tcp(localhost:3306)/boilerplate down
build:
	go build
