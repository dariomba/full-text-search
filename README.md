# Full text search using Elastic Search

## Description

Simple web application built with React that allows users to search through a large dataset of over 9000 movies. The frontend interacts with a backend service developed in Go, which communicates with an Elasticsearch instance to fetch movie data by its title, based on user queries.

## Features

- **React Frontend**: Simple UI for searching movies.
- **Go Backend**: A backend service that handles search queries and communicates with Elasticsearch.
- **Elastic Search**: Search engine that indexes and retrieves movie data.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

To run the project, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/dariomba/full-text-search.git
   ```

2. Build and start the application using Docker Compose:

   ```bash
   docker compose up --build
   ```

   This command will build the Docker images for the frontend, backend, and Elasticsearch, and then start the containers.

## Project Structure

- frontend/: Contains the React application code.
- backend/: Contains the Go backend service code.
- docker-compose.yml: Docker Compose configuration file to set up and run the multi-container application.

## Usage

Once the application is running, open your web browser and go to http://localhost to access the UI. You can start searching for movies using the search bar provided.
