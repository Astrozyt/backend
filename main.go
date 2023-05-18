package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type experiment struct {
	ID        int
	Name      string
	Device    string
	SensorId  string
	Data      string
	Timestamp time.Time
}

// Initialize array of experiments
var experiments = []experiment{
	{1, "Experiment 1", "Device 1", "Sensor 1", "1.0, 2.0, 3.0", time.Now()},
}

func main() {
	database, _ := sql.Open("sqlite3", "./database.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS experiments (id INTEGER PRIMARY KEY, name TEXT, device TEXT, sensor_id TEXT, data TEXT, timestamp DATETIME)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO experiments (name, device, sensor_id, data, timestamp) VALUES (?, ?, ?, ?, ?)")
	statement.Exec("Experiment 2", "Device 2", "Sensor 2", "1.0, 2.0, 3.0", time.Now())
	rows, _ := database.Query("SELECT id, name, device, sensor_id, data, timestamp FROM experiments")
	for rows.Next() {
		var exp experiment
		err := rows.Scan(&exp.ID, &exp.Name, &exp.Device, &exp.SensorId, &exp.Data, &exp.Timestamp)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		experiments = append(experiments, exp)
	}

	engine := html.New("./views", ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")
	app.Use(cors.New())

	//Return HTML app (via tmpl file)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"myName":      "My medical experiments",
			"experiments": experiments,
		})
	})

	app.Get("/experiment", func(c *fiber.Ctx) error {
		rows, _ := database.Query("SELECT id, name, device, sensor_id, data, timestamp FROM experiments")
		var experiments []experiment
		for rows.Next() {
			var exp experiment
			err := rows.Scan(&exp.ID, &exp.Name, &exp.Device, &exp.SensorId, &exp.Data, &exp.Timestamp)
			if err != nil {
				fmt.Println("Error scanning row:", err)
				continue
			}
			experiments = append(experiments, exp)
		}
		return c.JSON(experiments)
	})

	app.Get("/experiment/:id", func(c *fiber.Ctx) error {
		experimentId := c.Params("id")
		if experimentId == "" {
			return c.Status(400).SendString("Please provide valid experiment id")
		}
		statement, err := database.Prepare("SELECT id, name, device, sensor_id, data, timestamp FROM experiments WHERE id = ?")
		if err != nil {
			return err
		}
		row := statement.QueryRow(experimentId)
		var exp experiment
		err = row.Scan(&exp.ID, &exp.Name, &exp.Device, &exp.SensorId, &exp.Data, &exp.Timestamp)
		if err != nil {
			return c.Status(404).SendString("Could not find experiment with id " + experimentId)
		}
		return c.JSON(exp)
	})

	app.Post("/experiment", func(c *fiber.Ctx) error {
		var newExperiment experiment
		if err := c.BodyParser(&newExperiment); err != nil {
			return c.Status(400).SendString("Please provide valid experiment data")
		}

		statement, err := database.Prepare("INSERT INTO experiments (name, device, sensor_id, data, timestamp) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		_, err = statement.Exec(newExperiment.Name, newExperiment.Device, newExperiment.SensorId, newExperiment.Data, newExperiment.Timestamp)
		if err != nil {
			return err
		}

		return c.Status(201).SendString("Experiment added successfully")

	})

	app.Delete("/experiment/:id", func(c *fiber.Ctx) error {
		experimentId := c.Params("id")
		if experimentId == "" {
			return c.Status(400).SendString("Please provide valid experiment id")
		}
		statement, err := database.Prepare("DELETE FROM experiments WHERE id = ?")
		if err != nil {
			return err
		}
		_, err = statement.Exec(experimentId)
		if err != nil {
			return c.Status(405).SendString("Could not delete experiment with id " + experimentId)
		}

		return c.Status(200).SendString("Successfully deleted experiment with id " + experimentId)
	})

	//Return greeting on get route /greeting
	app.Get("/greeting", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	//Return information on type of hardware I can currently afford
	app.Get("/whatareyou", func(c *fiber.Ctx) error {
		return c.SendStatus(418)
	})

	log.Fatal(app.Listen(":3000"))
}
