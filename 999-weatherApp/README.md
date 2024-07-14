# Weather Forecast Application

This application fetches weather data for a given city, caches the data in Redis, and displays it in a formatted table with color coding using `go-pretty` and `fatih/color`.

## Features

- Fetches weather data from the National Weather Service API.
- Caches the weather data in Redis for faster subsequent requests.
- Displays weather data in a colored, formatted table.

## Prerequisites

- Go (1.16 or higher)
- Redis (Running locally or accessible remotely)

## Installation

1. **Download the repository:**

    From the provided upload

2. **Install dependencies:**

    ```sh
    go get github.com/fatih/color
    go get github.com/go-redis/redis/v8
    go get github.com/jedib0t/go-pretty/v6/table
    ```

3. **Set up Redis:**

    Ensure Redis is running locally or accessible remotely. If you don't have Redis installed, you can run it using Docker:

    ```sh
    docker run --name redis -p 6379:6379 -d redis
    ```

## Usage

Run the application with the desired city coordinates:

```sh
go run .\coding-interview.go 31.8484,-106.4270
```
