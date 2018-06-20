buildimage:
	az acr login --name openhacktable11
	docker build -t openhacktable11.azurecr.io/openhacktable11:$(VER) .
	docker push openhacktable11.azurecr.io/openhacktable11:$(VER)

deploy:
	kubectl apply -f mcapi-deployment.yaml
