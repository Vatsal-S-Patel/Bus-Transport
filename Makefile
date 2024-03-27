hello:
	echo "working"

go-server:
	cd Backend &&	go run main.go;

react:
	cd client-frontend && npm start;

all:
	poetry run $(make) go-server;
	poetry run $(make) react;
	poetry run $(make) hello;