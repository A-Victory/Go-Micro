# Go Microservices Project

This project is a learning exercise to improve my skills in building microservices architecture using Go. The project consists of several services, each with its own Docker container.

## Services

- **Caddy**: A web server and reverse proxy.
- **Front-end**: The user interface of the application.
- **Broker Service**: Manages communication between services.
- **Listener Service**: Listens for specific events.
- **Authentication Service**: Handles user authentication.
- **Logger Service**: Logs application events.
- **Mail Service**: Sends emails.
- **RabbitMQ**: Message broker.
- **Mailhog**: Email testing tool.
- **MongoDB**: NoSQL database.
- **PostgreSQL**: SQL database.

## Docker Setup

### Docker Compose

To start the services using Docker Compose, run:
```sh
docker-compose up -d
```

### Docker Swarm

To deploy the services using Docker Swarm, use:
```sh
docker stack deploy -c swarm.yml <stack_name>
```

## Makefile Commands

- `make up`: Starts all containers in the background.
- `make up_build`: Builds and starts all containers.
- `make down`: Stops all containers.
- `make build_<service>`: Builds the specified service binary.
- `make start`: Starts the front-end service.
- `make stop`: Stops the front-end service.

## Environment Variables

Ensure to set the necessary environment variables for each service as defined in the `docker-compose.yml` and `swarm.yml` files.

## Volumes

Persistent data is stored in Docker volumes and local directories as specified in the configuration files.

## Conclusion

This project is a practical approach to learning and improving microservices architecture with Go. Feel free to explore and modify the services to fit your learning needs.
