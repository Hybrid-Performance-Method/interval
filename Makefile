# to run the docker action 
build:
	docker build . -t hpm/interval --no-cache

docker: build
	docker run --env INPUT_NOTEBOOK=test_notebook.ipynb \
	--env INPUT_SECRETS=intervalsecret \
	--env INPUT_PARAMETERFILE=parameters.yml \
	hpm/interval:latest

# for development
interval: build-bin exec
build-bin:
	go build -o ./interval main.go 
exec:
	./interval