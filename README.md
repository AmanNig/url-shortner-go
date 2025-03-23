# URL Shortener

A simple URL shortener service built using Go Fiber and PostgreSQL as part of an assignment for Giva.

## Features
- Shorten long URLs using a POST API
- Redirect to original URLs using a GET API
- Built with Go Fiber for fast and efficient performance
- Uses PostgreSQL as the database for storing URL mappings

## Installation & Setup

### Prerequisites
- Go (1.18+ recommended)
- PostgreSQL
- Git

### Clone the Repository
```sh
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
```

### Install Dependencies
```sh
go mod tidy
```

### Configure Database
Create a PostgreSQL database and update your database credentials in the application.
```sql
CREATE DATABASE url_shortener;
```

### Run the Application
```sh
go run main.go
```

## API Endpoints

### 1. Shorten a URL
**Endpoint:** `POST /shorten`

**Request Body:**
```json
{
    "long_url": "https://meet.google.com/jgb"
}
```

**Response:**
```json
{
    "ID": 0,
    "long_url": "https://www.youtube.com/watch?v=qEBiHQTFWQI&list=RD6UiwXthbtZs&index=7",
    "alias": "6nLmEMer"
}
```

### 2. Redirect to Original URL
**Endpoint:** `GET /get/:shortUrl`

**Example Request:**
```sh
curl -X GET http://localhost:3000/get/6nLmEMer
```

**Behavior:** Redirects to the original long URL.

## Technologies Used
- **Go Fiber** - Web framework
- **PostgreSQL** - Database
- **GORM** - ORM for Go

## License
This project is licensed under the MIT License.

