build:
	docker build --pull --rm -f "Dockerfile" -t ruiblaese/bot:latest "."

push:
	docker push ruiblaese/bot:latest

build-push:
	docker build --pull --rm -f "Dockerfile" -t ruiblaese/bot:latest "."
	docker tag ruiblaese/bot:latest ruiblaese/bot:0.0.2
	docker push ruiblaese/bot:latest	
	docker push ruiblaese/bot:0.0.2