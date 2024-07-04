# Full text search using Elasticsearch

## Description

Simple web application built with React that allows users to search through a large dataset of over 9000 movies. The frontend interacts with a backend service developed in Go, which communicates with an Elasticsearch instance to fetch movie data by its title, based on user queries.

## Features

- **React Frontend**: Simple UI for searching movies.
- **Go Backend**: A backend service that handles search queries and communicates with Elasticsearch.
- **Elasticsearch**: Search engine that indexes and retrieves movie data.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

First of all, clone this repository:

```bash
git clone https://github.com/dariomba/full-text-search.git
```

### Run with Docker

Build and start the application using Docker Compose:

```bash
docker compose up --build
```

This command will build the Docker images for the frontend, backend, and Elasticsearch, and then start the containers.

### Development mode

1. Run Elasticsearch using Docker with the following command:

   ```bash
   docker run  -p 9200:9200 -e discovery.type=single-node -e xpack.security.enabled=false -d docker.elastic.coelasticsearch/elasticsearch:8.14.2
   ```

2. Move to `backend/` directory and create the `.env` file:
   ```bash
   ELASTIC_ADDRESS=http://127.0.0.1:9200
   ALLOWED_HOSTS_URLS=http://localhost:5173
   ```
3. Start backend service:
   ```bash
   go run ./src/cmd/main.go
   ```
4. In other terminal, move to `frontend/` directory, install the dependencies and start the React application:
   ```bash
   cd frontend/
   npm install
   npm run dev
   ```

Done! Now you can go to http://localhost:5173 and start searching movies.

## Project Structure

- frontend/: Contains the React application code.
- backend/: Contains the Go backend service code.
- docker-compose.yml: Docker Compose configuration file to set up and run the multi-container application.

## Usage

Once the application is running, open your web browser and go to http://localhost to access the UI. You can start searching for movies using the search bar provided.

## Dataset

Dataset used in this project: https://www.kaggle.com/datasets/disham993/9000-movies-dataset/data
