# to run the docker action 
build:
	docker build . -t hpm/interval --no-cache

docker: build
	docker run --env INPUT_NOTEBOOK=test_notebook.ipynb \
	--env INPUT_SECRET=intervalsecret \
	--env INPUT_PARAMETERFILE=parameters.yml \
	hpm/interval:latest
