# backend
The backend app created in the course "Medical Software Development", 2nd semester of Medical Informatics @ FHNW.

This is a simple application that demonstrates a basic web server built using the Go programming language and the Fiber web framework. The application interacts with a SQLite database to perform CRUD (Create, Read, Update, Delete) operations on experimental data.

## Prerequisites
To run this application, you need to have the following installed:

- Go (1.16 or higher)
- SQLite
- Installation
- Clone the repository or download the source code.
- Open a terminal and navigate to the project directory.
- Setup
- Install the necessary dependencies by running the following command:

go get -u github.com/gofiber/fiber/v2
go get -u github.com/gofiber/template/html
go get -u github.com/mattn/go-sqlite3

### Build the application by running the following command:

go build
Run the application:

./experiment
Open your web browser and visit http://localhost:3000 to access the application.

## Functionality
The application provides the following functionality:

Display a list of experiments on the homepage.
Add a new experiment using the provided form.
Retrieve experiment data from the SQLite database.
Serve static files from the public directory.

## Database
The application uses an SQLite database to store experiment data. The database file database.db will be created in the same directory as the application. If the file already exists, the application will use it; otherwise, it will create a new database file.

The database schema includes a table named experiments with the following columns:

id: Unique identifier for each experiment (integer).
name: Name of the experiment (text).
device: Device used for the experiment (text).
sensor_id: Sensor identifier (text).
data: Experimental data (text).
timestamp: Timestamp of the experiment (datetime).


## API Endpoints
The application exposes the following API endpoints:

GET /: Displays the homepage with a list of experiments and a form to add new experiments.
POST /experiment: Adds a new experiment to the database.
GET /greeting: Returns a greeting message.
Views
The application uses HTML templates for rendering views. The template files are located in the views directory. The template engine used is based on the standard Go html/template package.

## Dependencies
The application uses the following third-party dependencies:

github.com/gofiber/fiber/v2: Web framework for Go.
github.com/gofiber/template/html: HTML template engine for Fiber.
github.com/mattn/go-sqlite3: SQLite driver for Go.
License
This project is licensed under the MIT License.
