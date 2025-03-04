package main

import (
	"database/sql"
	"log"
	"os"
	"simple-go-jwt-api/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	initDatabase()

	app := fiber.New()

	app.Post("/login", login)
	app.Post("/register", register)
	app.Get("/profile", middleware.Authenticate, profile)

	log.Fatal(app.Listen(":" + port))

}

func register(c *fiber.Ctx) error {

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Requisição inválida"})
	}

	if req.Username == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Nome de usuário e senha são obrigatórios"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Erro ao gerar hash da senha: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao processar a senha"})
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", req.Username, hashedPassword)

	if err != nil {
		log.Printf("Erro ao inserir usuário no banco de dados: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Usuário já existe"})
	}

	return c.JSON(fiber.Map{"message": "Usuário registrado com sucesso"})

}

func login(c *fiber.Ctx) error {

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	secretKey := []byte(os.Getenv("SECRET_KEY"))

	log.Println("secretKey", secretKey)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Requisição inválida"})
	}

	if req.Username == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Nome de usuário e senha são obrigatórios"})
	}

	var storedPassword string

	err := db.QueryRow("SELECT password FROM users WHERE username = ?", req.Username).Scan(&storedPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(401).JSON(fiber.Map{"error": "Credenciais inválidas"})
		}
		log.Printf("Erro no servidor ao buscar usuário: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro no servidor"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Credenciais inválidas"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		log.Printf("Erro ao gerar token: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao gerar token"})
	}

	return c.JSON(fiber.Map{"token": tokenString})

}

func profile(c *fiber.Ctx) error {

	userClaims, ok := c.Locals("user").(jwt.MapClaims)

	if !ok {
		log.Printf("Erro ao obter claims do usuário")
		return c.Status(500).JSON(fiber.Map{"error": "Erro interno do servidor"})
	}

	return c.JSON(fiber.Map{
		"message":  "Bem-vindo ao seu perfil",
		"username": userClaims["username"],
	})

}
