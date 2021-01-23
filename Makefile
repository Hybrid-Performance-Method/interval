build:
	docker build . -t hpm/interval --no-cache

docker: build
	docker run --rm hpm/interval:latest
push:

interval: build-bin exec
build-bin:
	go build -o ./interval main.go 
exec:
	./interval