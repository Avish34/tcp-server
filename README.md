<p align="center">
  <p align="center">
    </br>
     <img src="https://github.com/Avish34/tcp-server/assets/45288918/f4c7a31b-ac33-42b8-a24b-164f2197a0a0"  height="64">
      &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
      <img src="https://github.com/Avish34/tcp-server/assets/45288918/80755845-f314-4d00-82c3-0ad9c6bb821b"  height="64">
        &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
  </p>
</p>


# tcp-server
An adaptable multi-threaded TCP server equipped with built-in rate limiting, affording you the flexibility to tailor thread configurations and rate limits to your specific needs. It comes with metrics support like total request processed by server at a given time, useful for users to analyse the server performance.

## Installation
### Prerequisite
1. [Docker](https://docs.docker.com/desktop/install/mac-install/)
2. Code editor.

### Steps to run server
1. git clone https://github.com/Avish34/tcp-server
2. cd tcp-server
3. make start-service
   
After running (3) command, you should see this output.
Server will be listening on 8080. To edit the port and other settings, you can find customise section below in the README.


>      ⠿ Container tools-prometheus-1  Started                                                                                                                          
>      ⠿ Container tools-tcp-server-1  Started  

Send request to the server. You should see "Hello world as response"

> curl http://localhost:8080                                         
> Hello world !

### Metrics
To check the total request accepted by the server. Open http://localhost:9090/. Search for *total_request* in Expression search bar. Check the below image for reference.

<img width="1336" alt="Screenshot 2023-09-10 at 12 08 50 PM" src="https://github.com/Avish34/tcp-server/assets/45288918/b5548648-fd7c-4517-b068-415f33d37c17">


## Customisation
This server comes with capability to throttle request using rate limiting, serving multiple request using workers. This is completly customisable, it can be done via changing the .env file present in

>  tools/dockerfiles/tcp-server/.env

You should be able to see the below parameters and can change it according to use system capabilites and the requirements.
> SERVER_PORT=8080
> SERVER_URL=0.0.0.0
> SERVER_QUEUE_SIZE=5
> SERVER_WORKERS=2
> SERVER_TOKEN_LIMIT=1
> SERVER_TOKEN_RATE=5
