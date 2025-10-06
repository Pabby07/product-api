package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func init() {
	var err error
	connStr := "postgres://pabby@localhost:5432/your_db_name"
	db, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
}

type Product struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Embedding   []float32 `json:"embedding"`
}

func CreateProduct(c *gin.Context) {
	var p Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Convert []float32 to []byte (json format)
	embeddingJSON, err := json.Marshal(p.Embedding)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode embedding"})
		return
	}

	// Insert into DB
	query := `INSERT INTO products (name, description, embedding) VALUES ($1, $2, $3)`
	_, err = db.Exec(context.Background(), query, p.Name, p.Description, embeddingJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert into DB", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added"})
}
