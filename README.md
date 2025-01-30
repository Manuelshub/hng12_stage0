# HNG 12 Stage 0

This Project is about implementing a public API to retrieve basic information

## Description

- The API should return the following information in JSON format:

1. Your registered email address
2. The current datetime as an ISO 8601 formatted timestamp.
3. The GitHub URL of the project's codebase.

## Requirements

1. Technology Stack:
2. Programming Language/Framework: Use any of the following: See Sharp (C#), PHP, Python, Go, Java, JavaScript/TypeScript.
3. Deployment: The API must be deployed to a publicly accessible endpoint.
4. CORS Handling: Ensure the API handles Cross-Origin Resource Sharing (CORS) appropriately.
5. Response Format: All responses must be in JSON format.

## API Specification

- Endpoint: GET** <your-url>
- Required JSON Response Format (200 OK):
```json
"email": "your-email@example.com",
"current_datetime": "2025-01-30T09:30:00Z",
"github_url": "<https://github.com/yourusername/your-repo>"
```

## Acceptance Criteria

### Functionality
- The API must accept GET requests and return the required JSON response.
- The current_datetime field must be dynamically generated in ISO 8601 format (UTC).
- Provides appropriate HTTP status code