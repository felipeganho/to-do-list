package main

import (
	"context"
	"log"
	"os"

	"github.com/felipeganho/to-do-list/pkg/entities"
	"github.com/felipeganho/to-do-list/pkg/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

var mg MongoInstance

func Connect() error {
	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("todo_db").Collection("todos")

	mg = MongoInstance{
		Client:     client,
		Collection: collection,
	}

	return nil
}

func Disconnect() {
	if err := mg.Client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	defer Disconnect()

	app := fiber.New()

	if os.Getenv("ENV") != "production" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost:5173",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
	}

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	if os.Getenv("ENV") == "production" {
		app.Static("/", ".client/dist")
	}

	log.Fatal(app.Listen("0.0.0.0:" + os.Getenv("PORT")))
}

func getTodos(c *fiber.Ctx) error {
	var todos []entities.Todo

	cursor, err := mg.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo entities.Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.JSON(presenter.TodoSuccessResponse(todos))
}

func createTodo(c *fiber.Ctx) error {
	todo := new(entities.Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
	}

	insertResult, err := mg.Collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = mg.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"message": "Todo deleted successfully"})
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}

	_, err = mg.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"message": "Todo updated successfully"})
}
