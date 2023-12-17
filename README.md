# Golang Message Queuing System

## Overview

This project implements a simple message queuing system using the Go programming language with either RabbitMQ or Kafka as the underlying message broker. The system comprises three main components: an API, a producer, and a consumer. The API receives product data, stores it in a database, and then passes the product ID to the message queue. The consumer, upon receiving the product ID, downloads and compresses product images, updating the product record with the local path to the compressed images.

## Components

### 1. API

The API is responsible for handling incoming requests to store product data. It exposes endpoints to receive information such as user ID, product name, description, images, and price. The API stores this data in a database and then sends the product ID to the message queue.

### 2. Producer

The producer component receives the product ID from the API and sends it to the message queue. This ensures that the consumer is notified of the newly stored product.

### 3. Consumer

The consumer, upon receiving the product ID from the message queue, downloads the associated product images, compresses them, and updates the product record in the database with the local path to the compressed images.

## Database Schema

### Users Table

- **id**: int, primary key
- **name**: Name of the user
- **mobile**: Contact number of the user
- **latitude**: Latitude of the user's location
- **longitude**: Longitude of the user's location
- **created_at**: Timestamp of user creation
- **updated_at**: Timestamp of user update

### Products Table

- **product_id**: int, primary key
- **product_name**: Name of the product
- **product_description**: Text describing the product
- **product_images**: Array of image URLs
- **product_price**: Numeric value representing the product price
- **compressed_product_images**: Array of local paths to compressed images
- **created_at**: Timestamp of product creation
- **updated_at**: Timestamp of product update

## Testing

The project includes comprehensive unit tests for API, database, and message queue components. Additionally, integration tests and benchmark tests are provided to ensure the reliability and performance of the system.

