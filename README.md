# Public API

## Overview

This is a simple RESTful API built using Go and the `gorilla/mux` package. The API provides an endpoint that returns a JSON response containing an email, the current date and time, and a GitHub URL. The server includes CORS support to allow cross-origin requests.

## Features

- Built with Go and `gorilla/mux`
- CORS enabled for external access
- Returns current date and time in UTC format
- Static email response
- Includes a GitHub repository link

## Endpoints

### **1. Get API Response**

**URL:** `/`

**Method:** `GET`

**Response:**
```json
{
  "email": "oguejioforrichard@gmail.com",
  "current_datetime": "2025-01-29T12:00:00Z",
  "github_url": "https://github.com/ikennarichard/hng12/stage-0"
}
```

## Installation & Running Locally

### **1️⃣ Clone the Repository**

```bash
git clone https://github.com/ikennarichard/hng-12.git
cd hng-12
```

### **2️⃣ Install Dependencies**

Make sure you have Go installed, then run:
```bash
go mod tidy
```

### **3️⃣ Run the Server**

```bash
go run main.go
```

The API will start at `http://localhost:8000/`.

## Testing the API

### **Unit Testing in Go**

Run tests to validate the JSON response:
```bash
go test .
```

### **Check Response Time (CLI)**

Use `curl` to measure response time:
```bash
time curl -X GET http://localhost:8000/
```
