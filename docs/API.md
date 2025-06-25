# API Documentation for Octane Performance Analyzer

## Overview

The Octane Performance Analyzer provides a set of APIs to interact with the performance testing tool. This documentation outlines the available endpoints, their functionalities, and how to use them effectively.

## Base URL

The base URL for the API is:

```
https://api.octane-bench.com/v1
```

## Authentication

All API requests require an API key for authentication. You can include your API key in the request headers:

```
Authorization: Bearer YOUR_API_KEY
```

## Endpoints

### 1. System Information

- **GET** `/system/info`

  Retrieves detailed information about the system's hardware and software configuration.

  **Response:**
  ```json
  {
    "os": "Ubuntu 22.04.3 LTS",
    "cpu": {
      "model": "Intel Core i7-13700K",
      "cores": 16,
      "threads": 24
    },
    "memory": {
      "total": "32GB",
      "available": "28GB"
    },
    "gpu": {
      "model": "NVIDIA GeForce RTX 4090",
      "memory": "24GB"
    },
    "storage": [
      {
        "device": "/dev/nvme0n1",
        "type": "NVMe SSD",
        "capacity": "1TB"
      }
    ]
  }
  ```

### 2. Performance Test

- **POST** `/performance/test`

  Initiates a performance test based on the specified parameters.

  **Request Body:**
  ```json
  {
    "test_type": "full",
    "upload": true,
    "tags": ["datacenter", "gpu-server"]
  }
  ```

  **Response:**
  ```json
  {
    "test_id": "octane-20250625-063144-uuid",
    "status": "running",
    "estimated_duration": "00:18:23"
  }
  ```

### 3. Retrieve Test Results

- **GET** `/performance/results/{test_id}`

  Fetches the results of a completed performance test.

  **Response:**
  ```json
  {
    "test_id": "octane-20250625-063144-uuid",
    "results": {
      "cpu": {
        "score": 1847,
        "grade": "A"
      },
      "memory": {
        "score": 45600,
        "grade": "A+"
      },
      "gpu": {
        "score": 19890,
        "grade": "A+"
      }
    },
    "overall_score": 94,
    "octane_rating": "racing_fuel"
  }
  ```

### 4. Upload Test Report

- **POST** `/performance/upload`

  Uploads the test report to the server.

  **Request Body:**
  ```json
  {
    "test_id": "octane-20250625-063144-uuid",
    "report": { /* report data */ },
    "anonymous": true
  }
  ```

  **Response:**
  ```json
  {
    "status": "success",
    "message": "Report uploaded successfully."
  }
  ```

### 5. Get Octane Ratings

- **GET** `/ratings`

  Retrieves the octane ratings for the system based on the latest tests.

  **Response:**
  ```json
  {
    "overall": {
      "ron": 94,
      "grade": "racing_fuel"
    },
    "breakdown": {
      "cpu": {
        "ron": 91,
        "grade": "premium_plus"
      },
      "memory": {
        "ron": 96,
        "grade": "racing_fuel"
      }
    }
  }
  ```

## Error Handling

All API responses include a status code and a message. Common status codes include:

- `200 OK`: Successful request.
- `400 Bad Request`: Invalid request parameters.
- `401 Unauthorized`: Missing or invalid API key.
- `404 Not Found`: Resource not found.
- `500 Internal Server Error`: An error occurred on the server.

## Conclusion

This API documentation provides a comprehensive guide to interacting with the Octane Performance Analyzer. For further assistance, please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file or contact support.