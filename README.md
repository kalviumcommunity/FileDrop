# Email Validation API

The Email Validation API is a Go-based API that allows you to validate email addresses and check their domain properties such as MX records, SPF, and DMARC.

## Getting Started

To use the Email Validation API, follow the instructions below.

### Prerequisites

- Go programming language (version 1.13 or higher)
- Required dependencies: `github.com/badoux/checkmail`, `github.com/gin-gonic/gin`
- Internet connectivity to perform domain lookups

### Installation

1.  Clone the repository:

    ```bash
    git clone https://github.com/Ahamed1846/FileDrop.git
    ```

2.  Navigate to the project directory:

    ```bash
    cd FileDrop
    ```

3.  Switch to the API branch:
    ```bash
    git checkout go-api
    ```
4.  Build and run the Go API:
    ```bash
    go build .
    ./FileDrop
    ```
By default, the API server will start listening on port 8080.

# API Endpoints
The Email Validation API exposes the following endpoint:

## Check Email
    GET /check/:email

**Parameters:**
- `email`: The email address to validate.

**Example Request:**
```bash
GET /check/johndoe@example.com
```
**Example Response:**
```bash
{
  "email": "johndoe@example.com",
  "format": true,
  "domain": "example.com",
  "mx": true,
  "mx_records": ["smtp.example.com", "mx1.example.com"],
  "spf": true,
  "spf_record": "v=spf1 mx ~all",
  "dmarc": true,
  "dmarc_record": "v=DMARC1; p=none",
  "validity": "Valid"
}
```

# Error Handling
In case of errors, the API will return the appropriate HTTP status code along with an error message in the response body.

- `400 Bad Request`: When the email address parameter is missing or empty.
- `404 Not Found`: When the requested endpoint does not exist.
