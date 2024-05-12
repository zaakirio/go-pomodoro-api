
## Go-Pomodoro API

This API provides a RESTful interface for managing pomodoro sessions using Google Firestore

### Features

-   **Get all Pomodoros:** Retrieve a list of all pomodoro sessions.
-   **Get a Pomodoro by ID:** Fetch a specific pomodoro session by its unique identifier.
-   **Create a Pomodoro:** Start a new pomodoro session.
-   **Update a Pomodoro:** Modify an existing pomodoro session (currently only allows for generic updates to the entire Pomodoro object).
-   **Delete a Pomodoro:** Remove a pomodoro session.

### Prerequisites

-   Go 1.17 or later
-   Google Cloud project with Firestore enabled
-   Service account key with permissions to access Firestore

### Installation

1.  Clone this repository.
2.  Set up your Google Cloud project and enable Firestore.
3.  Create a service account key and download the JSON file.
4.  Place the service account key JSON file in the project directory (adjust the path as needed).
5.  Configure the `FirestoreClient` in your `main.go` file to connect to your Firestore instance.

### Running the API

1.  Open a terminal in the project directory.
2.  Run `go run main.go`.

### API Endpoints

**Base URL:**  `http://localhost:8080` (assuming the server runs on port 8080)

**Get All Pomodoros:**

-   Method: GET
-   Path:  `/pomodoros`
-   Response: JSON array of `Pomodoro` objects (see Models section)
-   Status Code:
    
    -   200: Success
    -   500: Internal Server Error (database error)
    

**Get Pomodoro by ID:**

-   Method: GET
-   Path:  `/pomodoros/:id`
-   Parameter:
    
    -   `id`: Unique identifier of the pomodoro session (string)
    
-   Response: JSON object representing a single `Pomodoro` (see Models section)
-   Status Code:
    
    -   200: Success
    -   404: Not Found (pomodoro with specified ID not found)
    -   500: Internal Server Error (database error)
    

**Create Pomodoro:**

-   Method: POST
-   Path:  `/pomodoros/create`
-   Request Body: JSON object representing a new `Pomodoro` (see Models section)
-   Response: JSON object representing the created `Pomodoro`
-   Status Code:
    
    -   201: Created
    -   400: Bad Request (invalid request body)
    -   500: Internal Server Error (database error)
    

**Update Pomodoro:**

-   Method: PUT
-   Path:  `/pomodoros/:id`
-   Parameter:
    
    -   `id`: Unique identifier of the pomodoro session (string)
    
-   Request Body: JSON object representing the updated `Pomodoro` (see Models section)
-   Response: JSON object representing the updated `Pomodoro`
-   Status Code:
    
    -   200: OK (pomodoro updated successfully)
    -   400: Bad Request (invalid request body)
    -   404: Not Found (pomodoro with specified ID not found)
    -   500: Internal Server Error (database error)
    

**Delete Pomodoro:**

-   Method: DELETE
-   Path:  `/pomodoros/:id`
-   Parameter:
    
    -   `id`: Unique identifier of the pomodoro session (string)
    
-   Response: No content
-   Status Code:
    -   204: No Content (pomodoro deleted successfully)
    -   404: Not Found (pomodoro with specified ID not found)
    -   500: Internal Server Error (database error)