# Go Song Album CRUD API

This project is a simple CRUD (Create, Read, Update, Delete) API for managing song albums, built using the Go programming language and the Gin web framework.

[Tutorial from Go](https://go.dev/doc/tutorial/web-service-gin)

## Features

- Create a new song album
- Retrieve a list of all song albums
- Retrieve a specific song album by ID
- Update an existing song album
- Delete a song album

## Prerequisites

- Go 1.23.4 (Used for development) or higher
- Gin web framework

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/manjushsh/go-song-album.git
    cd go-song-album
    ```

2. Install the dependencies:

    ```sh
    go mod tidy
    ```

## Running the Application

To start the server, run:

```sh
go run main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

### CRUD API Endpoints

#### Create a new album

- **Endpoint:** `POST /albums`
- **Request Body:**

    ```json
    {
        "id": "1",
        "title": "Album Title",
        "artist": "Artist Name",
        "price": 9.99
    }
    ```

- **Response:**

    ```json
    {
        "id": "1",
        "title": "Album Title",
        "artist": "Artist Name",
        "price": 9.99
    }
    ```

#### Retrieve all albums

- **Endpoint:** `GET /albums`
- **Response:**

    ```json
    [
        {
            "id": "1",
            "title": "Album Title",
            "artist": "Artist Name",
            "price": 9.99
        },
        {
            "id": "2",
            "title": "Another Album",
            "artist": "Another Artist",
            "price": 12.99
        }
    ]
    ```

#### Retrieve a specific album by ID

- **Endpoint:** `GET /albums/:id`
- **Response:**

    ```json
    {
        "id": "1",
        "title": "Album Title",
        "artist": "Artist Name",
        "price": 9.99
    }
    ```

#### Update an existing album

- **Endpoint:** `PUT /albums/:id`
- **Request Body:**

    ```json
    {
        "title": "Updated Album Title",
        "artist": "Updated Artist Name",
        "price": 10.99
    }
    ```

- **Response:**

    ```json
    {
        "id": "1",
        "title": "Updated Album Title",
        "artist": "Updated Artist Name",
        "price": 10.99
    }
    ```

#### Delete an album

- **Endpoint:** `DELETE /albums/:id`
- **Response:**

    ```json
    {
        "message": "Album deleted successfully"
    }
    ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Golang](https://golang.org/)
