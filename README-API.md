# Product API Documentation

This API allows you to manage products with images. Images are stored in MinIO while product details are stored in PostgreSQL.

## Prerequisites

1. PostgreSQL database running
2. MinIO server running
3. Environment variables configured in `.env` file:
```env
API_PORT=8080
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=your_access_key
MINIO_SECRET_KEY=your_secret_key
DB_HOST=localhost
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database_name
DB_PORT=5432
```

## API Endpoints

### 1. Create Product
- **URL**: `/products`
- **Method**: `POST`
- **Content-Type**: `multipart/form-data`
- **Form Fields**:
  - `name` (required): Product name
  - `description`: Product description
  - `expiry_date` (optional): Product expiry date in ISO format (e.g., "2024-12-31T00:00:00Z")
  - `image` (required): Product image file
- **Success Response**: 
  ```json
  {
    "status": true,
    "message": "Product created successfully",
    "payload": {
      "id": 1,
      "name": "Product Name",
      "description": "Product Description",
      "image_path": "image_filename.jpg",
      "expiry_date": "2024-12-31T00:00:00Z",
      "created_at": "2024-12-18T11:43:11+07:00",
      "updated_at": "2024-12-18T11:43:11+07:00"
    }
  }
  ```

### 2. Get All Products
- **URL**: `/products`
- **Method**: `GET`
- **Success Response**:
  ```json
  {
    "status": true,
    "message": "Products retrieved successfully",
    "payload": [
      {
        "id": 1,
        "name": "Product Name",
        "description": "Product Description",
        "image_path": "image_filename.jpg",
        "expiry_date": "2024-12-31T00:00:00Z",
        "created_at": "2024-12-18T11:43:11+07:00",
        "updated_at": "2024-12-18T11:43:11+07:00"
      }
    ]
  }
  ```

### 3. Get Product by ID
- **URL**: `/products/{id}`
- **Method**: `GET`
- **Success Response**:
  ```json
  {
    "status": true,
    "message": "Product retrieved successfully",
    "payload": {
      "id": 1,
      "name": "Product Name",
      "description": "Product Description",
      "image_path": "image_filename.jpg",
      "expiry_date": "2024-12-31T00:00:00Z",
      "created_at": "2024-12-18T11:43:11+07:00",
      "updated_at": "2024-12-18T11:43:11+07:00"
    }
  }
  ```

### 4. Update Product
- **URL**: `/products/{id}`
- **Method**: `PUT`
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "name": "Updated Product Name",
    "description": "Updated Description",
    "expiry_date": "2024-12-31T00:00:00Z"
  }
  ```
- **Success Response**:
  ```json
  {
    "status": true,
    "message": "Product updated successfully",
    "payload": {
      "id": 1,
      "name": "Updated Product Name",
      "description": "Updated Description",
      "image_path": "image_filename.jpg",
      "expiry_date": "2024-12-31T00:00:00Z",
      "created_at": "2024-12-18T11:43:11+07:00",
      "updated_at": "2024-12-18T11:43:11+07:00"
    }
  }
  ```

### 5. Delete Product
- **URL**: `/products/{id}`
- **Method**: `DELETE`
- **Success Response**:
  ```json
  {
    "status": true,
    "message": "Product deleted successfully"
  }
  ```

## Error Responses
All endpoints return error responses in this format:
```json
{
  "status": false,
  "message": "Error message",
  "error": "Detailed error information"
}
```

## Testing
1. Import the provided Postman collection (`product-api.postman_collection.json`)
2. Set up your environment variables
3. Test each endpoint using the collection
