# Go Static Web Server & Form Handler

A simple web server built with Go's standard library. It serves static HTML files and provides several API endpoints to handle form submissions and basic GET requests. This project demonstrates fundamental Go web development concepts, including file serving, request routing, and parsing form data.

## Features

*   Serves static files from a `./static` directory.
*   Handles form submissions on multiple endpoints.
*   Provides simple text-based responses for basic GET requests.
*   Built entirely with the standard Go `net/http` package.

## Project Structure

The server is configured to serve static files from a `./static` directory. For the server to function correctly, you must create this directory and place your HTML files inside it.


## Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

*   [Go](https://golang.org/dl/) (version 1.18 or newer recommended)

### Setup

1.  **Save the Code:** Save the provided Go source code as `main.go`.

2.  **Create the Static Directory:** In the same folder as `main.go`, create a new directory named `static`.
    ```sh
    mkdir static
    ```

3.  **Create HTML Files:** Inside the `static` directory, paste the above three HTML files in the `static` directory which you created.

4.  **Run the Server:** Open your terminal, navigate to the project directory, and run the following command:
    ```sh
    go run main.go
    ```
    You will see the message: `Starting server at port 8080`.

## Endpoints

| Path            | Method | Description                                                                 |
| --------------- | ------ | --------------------------------------------------------------------------- |
| `/`             | `GET`  | Serves static files from the `./static` directory (defaults to `index.html`). |
| `/hello`        | `GET`  | Returns a simple `hello!` text response.                                    |
| `/goodbye`      | `GET`  | Returns a simple farewell text response.                                    |
| `/form`         | `POST` | Processes form data with `name` and `address` fields and echoes them back.  |
| `/wellbeing`    | `POST` | Processes form data with `name`, `age`, and `health-condition` fields.      |

## How to Use

#### Using a Web Browser
1.  Open your web browser and navigate to `http://localhost:8080`.
2.  You will see the `index.html` page.
3.  Click the links to navigate to the forms, fill them out, and submit to see the server's response.

#### Using `curl` (for testing endpoints)

You can also test the form handlers directly from your terminal.

*   **Test the `/form` endpoint:**
    ```sh
    curl -X POST -d "name=John Doe&address=123 Main St" http://localhost:8080/form
    ```

*   **Test the `/wellbeing` endpoint:**
    ```sh
    curl -X POST -d "name=Jane Doe&age=30&health-condition=Excellent" http://localhost:8080/wellbeing
    ```

*   **Test the `/hello` endpoint:**
    ```sh
    curl http://localhost:8080/hello
    ```

    ---

    *Thank you for exploring the project*
