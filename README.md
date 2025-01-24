# Oskar Brzeski Portfolio

This is the repository for my portfolio.

The website uses [HTMX](https://htmx.org/) to perform AJAX requests to the
webserver. The webserver uses [Go](https://go.dev/), with the
[Echo](https://echo.labstack.com/) package, to take in requests, render the
HTML that the user wants to see, and send it as a response. This allows the
website to update what it displays with barely any latency at all and no
requirement for a fast internet connection or device.

This project is not currently deployed.

# Clone repository

```bash
git clone https://github.com/OskarBrzeski/portfolio.git
cd portfolio
```

# Run locally

To run the project locally, navigate to the project's root directory and run:
```bash
go run cmd/main.go
```
You can also use [Air](https://github.com/air-verse/air). Just run the `air`
executable inside the project's root directory.
```bash
path/to/air
```
The website will be availbale on `localhost:8080`.
