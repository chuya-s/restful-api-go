# make hello
hello:
	echo "hello"

build:
	docker compose build --no-cache=true

# make up
up: build
	docker compose up --detach --force-recreate

# make ps
ps:
	docker compose ps

# make db
mysql:
	docker exec -it db bash

# make down
down:
	docker compose down

# make volume-ls
volume-ls:
	docker volume ls

# make volume-rm
volume-rm:
	docker volume rm mysql-volume