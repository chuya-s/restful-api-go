# make hello
hello:
	echo "hello"

# make up
up:
	docker compose up -d

# make ps
ps:
	docker compose ps

# make db
mysql:
	docker exec -it db bash

# make volume-ls
volume-ls:
	docker volume ls

# make volume-rm
volume-rm:
	docker volume rm mysql-volume