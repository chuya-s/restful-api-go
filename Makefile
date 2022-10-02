# make hello
hello:
	echo "hello"

# make build
build:
	docker build -t backen .

# make run
run:
	docker run -d -it -p 8080:8080 backend

# make ps
ps:
	docker ps