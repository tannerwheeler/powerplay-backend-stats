# Power Play Backend

![image](https://github.com/jak103/powerplay/assets/16627408/4ec3df62-d760-40c6-aa57-fa63eaaaf61b)


[![Go](https://github.com/jak103/powerplay/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/jak103/powerplay/actions/workflows/go.yml)

## Table of Contents
- [Adding and Updating API Endpoints](#adding-and-updating-api-endpoints)
- [Adding and Updating Unit Testing for Database Model](#adding-and-updating-unit-testing-for-database-model)
- [Adding and Updating Unit Testing for API Endpoint](#adding-and-updating-unit-testing-for-api-endpoint)
## Adding and Updating API Endpoints

### Ensure Model is Created and Up to Date
1. Navigate to the models directory: \
   `backend/internal/models`

### Add API Endpoint
1. Navigate to the server directory: \
    `backend/internal/server/apis`

### Add DB methods used in API endpoint
1. Navigate to the db directory: \
    `backend/internal/db`

### Update Open API docs**
1. Navigate to the open api spec directory: \
   `static/oas/v1`

## Adding and Updating Unit Testing for Database Model
Taken from an in class walkthrough of adding a unit test for a database model. These tests will be using a docker spin up of the actual database for testing the database interfacing functions.

Uses the dockertest go package

## Adding and Updating Unit Testing for API Endpoint
Taken from an in class walkthrough of adding a unit test for an API endpoint. These tests will be using a mock up of the database models involved.

