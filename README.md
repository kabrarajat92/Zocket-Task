# Product Management System

This project is a backend implementation for a **Product Management System** built using **Golang**. It demonstrates key architectural principles, including asynchronous processing, caching, and high scalability.

## Features

1. **API Endpoints**:

   - `POST /products`: Add a new product.
   - `GET /products/:id`: Fetch product details by ID (includes processed images).
   - `GET /products`: Fetch all products with optional filters (user ID, price range, product name).

2. **Asynchronous Image Processing**:

   - Images are processed (compressed) asynchronously and stored in AWS S3.

3. **Caching**:

   - Product details are cached in Redis to reduce database load and improve response times.

4. **Logging**:

   - Structured logging for API requests, image processing events, and errors.

5. **Error Handling**:

   - Comprehensive error handling with retry mechanisms for image processing failures.

6. **Testing**:

   - Unit tests for API endpoints.
   - Integration tests for end-to-end functionality.

## Technologies Used

- **Programming Language**: Golang
- **Database**: PostgreSQL
- **Cache**: Redis
- **Message Queue**: RabbitMQ
- **Object Storage**: AWS S3
- **Logging**: Logrus

## Setup Instructions

### Prerequisites

1. Install the following:

   - [Golang](https://golang.org/doc/install)
   - [PostgreSQL](https://www.postgresql.org/download/)
   - [Redis](https://redis.io/download)
   - [RabbitMQ](https://www.rabbitmq.com/download.html)

2. Configure AWS credentials for S3.

3. Clone the repository:

   ```bash
   git clone https://github.com/kabrarajat92/Zocket-task.git
   cd Zocket-task
   ```

### Environment Variables

Create a `.env` file in the project root and add:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=product_management
REDIS_HOST=localhost
REDIS_PORT=6379
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
AWS_REGION=us-east-1
AWS_BUCKET_NAME=your_s3_bucket_name
AWS_ACCESS_KEY_ID=your_aws_access_key
AWS_SECRET_ACCESS_KEY=your_aws_secret_key
```

### Run the Project

1. **Start Services**:

   - Run PostgreSQL, Redis, and RabbitMQ services.
   - For Redis and RabbitMQ, you can use Docker:
     ```bash
     docker run --name redis -p 6379:6379 -d redis
     docker run -d --hostname my-rabbit --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
     ```

2. **Apply Migrations**:

   ```bash
   psql -U your_db_user -d product_management -f migrations/001_init.sql
   ```

3. **Run the Application**:

   ```bash
   go run main.go
   ```

   The server will start at `http://localhost:8080`.

4. **Test the APIs**:

   - Use Postman or cURL to test the endpoints.

### Example API Requests

1. **Add a Product**:

   ```bash
   curl -X POST http://localhost:8080/products \
   -H "Content-Type: application/json" \
   -d '{
       "user_id": 1,
       "name": "Sample Product",
       "description": "A sample product for testing",
       "price": 99.99,
       "product_images": ["http://example.com/image1.jpg"]
   }'
   ```

2. **Get Product by ID**:

   ```bash
   curl http://localhost:8080/products/1
   ```

3. **Get All Products with Filters**:

   ```bash
   curl "http://localhost:8080/products?user_id=1&min_price=50&max_price=150"
   ```

## Testing

1. **Run Unit Tests**:

   ```bash
   go test ./...
   ```

2. **Run Integration Tests**:

   ```bash
   go test ./tests -v
   ```

## Future Improvements

- Add authentication and authorization using JWT.
- Implement rate limiting for API endpoints.
- Expand image processing capabilities (e.g., resizing, format conversion).
- Introduce CI/CD pipelines for automated testing and deployment.

## Author

Rajat Kabra - Backend Developer

Feel free to reach out if you have any questions or suggestions!

