build:
	docker build . -t go-server:1.0

run:
	docker container run -d -p 8080:8080 go-server:1.0

deploy:
	gcloud builds submit --tag gcr.io/$(PROJECT_ID)/go-server
