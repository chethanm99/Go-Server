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

3.  **Create HTML Files:** Inside the `static` directory, create the following three HTML files.

    *   `static/index.html`:
        ```html
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Go Web Server</title>
        </head>
        <body>
            <h1>Welcome to the Go Web Server!</h1>
            <p>Choose a form to submit:</p>
            <ul>
                <li><a href="/form.html">General Form</a></li>
                <li><a href="/wellbeing.html">Wellbeing Survey</a></li>
            </ul>
        </body>
        </html>
        ```

    *   `static/form.html`:
        ```html
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Contact Form</title>
        </head>
        <body>
            <h2>Contact Form</h2>
            <form action="/form" method="post">
                <label for="name">Name:</label><br>
                <input type="text" id="name" name="name"><br><br>
                <label for="address">Address:</label><br>
                <input type="text" id="address" name="address"><br><br>
                <input type="submit" value="Submit">
            </form>
        </body>
        </html>
        ```

    *   `static/wellbeing.html`:
        ```html
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Wellbeing Survey</title>
        </head>
        <body>
            <h2>Wellbeing Survey</h2>
            <form action="/wellbeing" method="post">
                <label for="name">Name:</label><br>
                <input type="text" id="name" name="name"><br><br>
                <label for="age">Age:</label><br>
                <input type="number" id="age" name="age"><br><br>
                <label for="health-condition">Health Condition:</label><br>
                <input type="text" id="health-condition" name="health-condition"><br><br>
                <input type="submit" value="Submit">
            </form>
        </body>
        </html>
        ```

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
