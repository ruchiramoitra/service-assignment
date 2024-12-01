# Service Assignment

This is a backend application built with Go that interacts with a PostgreSQL database(selected as the persistence system). It demonstrates a simple architecture involving database interactions, environment configuration management, and API server setup. The application is containerized using Docker and Docker Compose for easy setup and deployment.

## Features

- **PostgreSQL Database**: The application connects to a PostgreSQL database running in a Docker container.
- **API Server**: Exposes a REST API for interactions.
- **Configuration Management**: Utilizes `viper` for reading environment variables and `.env` files.
- **Docker & Docker Compose**: Containerized using Docker for easier deployment and dependency management.
- **Mock Data**: Initializes some mock data (services and versions) in the database when the application starts.

## Getting Started

Follow the steps below to get the project up and running.

### 1. Clone the repository

```bash
git clone git@github.com:ruchiramoitra/service-assignment.git
cd service-assignment
```

### 2. Setup and Start the Application

The project includes a `docker-compose.yml` file that defines the services for PostgreSQL and the Go application. You can use Docker Compose to set everything up and run it in a containerized environment.

```bash
docker-compose up --build
```

This will:
- Build the Go application.
- Start the PostgreSQL database and the application server.
- Expose the API server on port `8081`.
- The PostgreSQL database will be available at `localhost:6432`.

### 3. Environment Variables

You can configure the environment variables required for the application by setting them in a `.env` file in the root of the project (optional when running through Docker Compose).

Sample `.env` file:
```env
DB_HOST=localhost
DB_USER=test_user
DB_PASSWORD=test_password
DB_NAME=test_db
DB_PORT=6432
```

Alternatively, you can set the environment variables directly in the `docker-compose.yml`.

### 4. Testing

You can run the tests for this project using the following command:

```bash
go test ./...
```

This will run all the tests defined in the project.

### 5. Stopping the Application

To stop the services, you can use:

```bash
docker-compose down
```

### 6. Accessing the API

Once the services are running, you can access the API at:

```
http://localhost:8081
```

### 7. Database Schema Initialization

When the PostgreSQL database starts, it automatically initializes the schema using the `init.sql` file and inserts some mock data into the `services` and `versions` tables.

### 8. Sample curls for running
```curl
curl --location 'localhost:8081/v1/services?sort=asc&pagination_token=MTox&limit=3'
```
```curl
curl --location 'localhost:8081/v1/service'
```

## 9. Remaining tasks

a. Add authorisation layer for the APIs -- we can add a login method and create a jwt based authentication. When a user logs in, issue a JWT token containing the user’s credentials and other metadata. Add a middleware for the apis where we check if the token is present or not. When logging in it is checked against the db for username and password.


b. Full tests (unit and integrations) -- have added few unit tests. For integration tests. We can spinup postgres db in docker and check end to end results if working correctly or not.


c. Add CRUD operations to the APIs -- like search we can add delete, update create end points too where we would need the corresponding queries to update in the db.


For persistance we can also use a caching mechanism to search services which are used more recently.
