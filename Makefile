build-docker: 
	docker build . -t tcp-server:v1

container: 
	docker run -it --rm -d -p 8080:8080 tcp-server:v1

clean:
	docker container rm $$(docker ps -aq) -f
