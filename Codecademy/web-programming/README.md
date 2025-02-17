# Web Programming With Go

In this project, you will practice your knowledge of Go’s net/http package. You will build a Restaurant Menu application to display menu items, collect user reviews, and display those reviews.

Currently, the application is missing proper routing and request handling. To fix this, you will need to:

- Start the server
- Set up HTTP routes
- Handle GET requests
- Handle POST requests
- Parse form data
- Display data dynamically

The project includes the following structure:

- `main.go`: Main entry point for the application.
- `handlers.go`: Contains handler functions for different routes.
- `models.go`: Contains data models for menu items and reviews.
- `requests.go`: Handles the requests for menu and review data. This file is already set up, and you do not need to modify it.
- `templates`: Directory containing the HTML template files for different pages (home.html, menu.html, review_form.html, reviews.html).
- `static/styles.css`: CSS styling for the application.

The styling in the CSS file has already been provided. To start the application, use `go run .` in the workspace terminal.

After completing these tasks, you should be able to display the review submission form, parse the submitted data, and send it to the server.

Run `go run server/requests.go` to start a server to handle `POST` and `GET` requests. Open another terminal by clicking the + button next to bash and run `go run .` to start the web server that we created.

Test and QA on the project to make sure it’s running properly.