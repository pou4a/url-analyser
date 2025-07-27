# URL Analyzer

This repository contains both the **frontend** and **backend** applications for the url-analyser Challenge. The project is designed to manage URLs, analyze their content, and display crawl results.

## Project Structure

The repository is organized as follows:

-   **Frontend**: A React-based application located in the `frontend-app` directory. It provides a user interface for submitting URLs, viewing the list of URLs, and analyzing crawl results.
-   **Backend**: A Go-based application located in the `go-backend-app` directory. It handles API requests, interacts with the database, and performs URL crawling.
-   **Database**: A MySQL database used to store URL data and crawl results.

## Prerequisites

Before running the project, ensure you have the following installed:

-   [Docker](https://www.docker.com/) (for running the app with Docker Compose)
-   [Docker Compose](https://docs.docker.com/compose/) (v3.8 or higher)

## Running the Project with Docker Compose

To run the entire project (frontend, backend, and database) using Docker Compose:

1. Clone the repository:

    ```bash
    git clone https://github.com/pou4a/url-analyser.git
    cd url-analyser
    ```

2. Build and start the services:

    ```bash
    docker-compose up --build
    ```

3. Access the services:
    - **Frontend**: Open [http://localhost:3000](http://localhost:3000) in your browser.
    - **Backend**: The API is available at [http://localhost:8080](http://localhost:8080).
    - **Database**: MySQL is running on port `3307` (default credentials are defined in `docker-compose.yml`).

## Services Overview

### Frontend

-   **Location**: `frontend-app`
-   **Technology**: React
-   **Features**:
    -   Submit URLs for analysis.
    -   View the list of submitted URLs.
    -   Analyze crawl results for specific URLs.
-   **Development Commands**:
    ```bash
    cd frontend-app
    npm install
    npm start
    ```

### Backend

-   **Location**: `go-backend-app`
-   **Technology**: Go
-   **Features**:
    -   REST API for managing URLs and crawl results.
    -   Interacts with the MySQL database.
    -   Performs URL crawling and analysis.
-   **Development Commands**:
    ```bash
    cd go-backend-app
    go run main.go
    ```

### Database

-   **Technology**: MySQL
-   **Default Credentials**:
    -   **User**: `appuser`
    -   **Password**: `appsecret`
    -   **Database**: `url-analyser`

## Environment Variables

The following environment variables are used in the project:

### Backend

-   `DB_USER`: Database username.
-   `DB_PASSWORD`: Database password.
-   `DB_HOST`: Database host (default: `db`).
-   `DB_PORT`: Database port (default: `3306`).
-   `DB_NAME`: Database name.

### Frontend

-   `CHOKIDAR_USEPOLLING`: Fixes file watching issues in Docker.

## Docker Compose Configuration

The `docker-compose.yml` file defines the following services:

-   **Backend**:
    -   Built from the `go-backend-app` directory.
    -   Exposes port `8080`.
-   **Frontend**:
    -   Built from the `frontend-app` directory.
    -   Exposes port `3000`.
-   **Database**:
    -   Uses the `mysql:8.0` image.
    -   Exposes port `3307`.

## Learn More

-   [React Documentation](https://reactjs.org/)
-   [Go Documentation](https://golang.org/doc/)
-   [Docker Compose Documentation](https://docs.docker.com/compose/)

