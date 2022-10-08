package http

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"lillybox-backend/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Run ...
func (server *Server) Run() error {
	server.LoadConfig()
	server.CreateRouter()
	server.SetMiddlewares()
	server.ConnectDatabase()
	server.SetRouters()
	server.SetAccessLogger()
	server.SetErrorLogger()

	return server.App.Listen(fmt.Sprintf(":%d", server.Config.Port))
}

// CreateRouter ...
func (server *Server) CreateRouter() {
	server.App = fiber.New()
}

func (server *Server) SetAccessLogger() {
	logFile := "log/http/access/lillybox-%Y-%m-%d-%H.log"
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationTime(time.Hour))
	FatalWithError(err)
	encoderConfig := map[string]string{
		"levelEncoder": "capital",
		"timeKey":      "date",
		"timeEncoder":  "iso8601",
	}
	data, _ := json.Marshal(encoderConfig)
	var encCfg zapcore.EncoderConfig
	err = json.Unmarshal(data, &encCfg)
	FatalWithError(err)
	w := zapcore.AddSync(rotator)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), w, zap.InfoLevel)
	accessLoger := zap.New(core)
	accessLoger.Info("Now logging in a rotated file")
	server.App.Use(logger.New(logger.Config{
		Format:   "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeZone: "Asia/Seoul",
		Output:   rotator,
	}))
	log.SetOutput(rotator)
	server.Handlers.AccessLogger = accessLoger
}

func (server *Server) SetErrorLogger() {
	logFile := "log/http/error/lillybox-%Y-%m-%d-%H.log"
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationTime(time.Hour))
	FatalWithError(err)
	encoderConfig := map[string]string{
		"levelEncoder": "capital",
		"timeKey":      "date",
		"timeEncoder":  "iso8601",
	}
	data, _ := json.Marshal(encoderConfig)
	var encCfg zapcore.EncoderConfig
	err = json.Unmarshal(data, &encCfg)
	FatalWithError(err)
	w := zapcore.AddSync(rotator)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), w, zap.ErrorLevel)
	logger := zap.New(core)
	logger.Error("Now logging in a rotated file")
	server.Handlers.ErrorLogger = logger
}

// SetMiddlewares ...
func (server *Server) SetMiddlewares() {
	server.App.Use(recover.New())
	server.App.Use(cors.New(cors.ConfigDefault))
}

// SetRouters ...
func (server *Server) SetRouters() {
	server.App.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello")
	})

	server.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	server.App.Get("/swagger/*", swagger.HandlerDefault)
	server.App.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	// Auth Routers
	server.App.Post("/api/v1/auth/request_id", server.Handlers.GetLoginRequestID)
	server.App.Post("/api/v1/auth/user", server.Handlers.GetUserInfo)
	server.App.Post("/api/v1/auth/sign", server.Handlers.SignUp)
	server.App.Post("/api/v1/auth/check", server.Handlers.CheckNickname)
	// Streaming Routers
	server.App.Get("/streamings", server.Handlers.GetStreamings)
	server.App.Post("/streamings", server.Handlers.CreateStreaming)
	server.App.Get("/streamings/:id", server.Handlers.GetStreamingByID)
	server.App.Patch("/streamings/:id", server.Handlers.PatchStreamingByID)
	server.App.Delete("/streamings/:id", server.Handlers.DeleteStreamingByID)

	// OnDemand Routers
	server.App.Get("api/v1/videos", server.Handlers.GetOnDemands)
	server.App.Get("api/v1/videos/:id", server.Handlers.GetOnDemandByID)
	server.App.Patch("api/v1/videos/:id", server.Handlers.PatchOnDemandByID)
	server.App.Delete("api/v1/videos/:id", server.Handlers.DeleteOnDemandByID)

	// Comment Routers
	server.App.Get("api/v1/videos/:id/comments", server.Handlers.GetComments)
	server.App.Post("api/v1/videos/:id/comments", server.Handlers.CreateComment)
	server.App.Patch("api/v1/videos/:id/comments", server.Handlers.PatchCommentByID)
	server.App.Delete("api/v1/videos/:id/comments", server.Handlers.DeleteCommentByID)

	// Report Routers
	server.App.Patch("api/v1/reports/videos/:id", server.Handlers.ReportOnDemandByID)
	server.App.Patch("api/v1/reports/comments/:id", server.Handlers.ReportCommentByID)
}

// LoadConfig ...
func (server *Server) LoadConfig() {

	err := godotenv.Load("./.env")
	FatalWithError(err)

	appPort := os.Getenv("PORT")
	port, err := strconv.Atoi(appPort)
	FatalWithError(err)

	server.Config.Port = port
	server.Config.Livepeer.Host = os.Getenv("LIVEPEER_API_HOST")
	server.Config.Livepeer.APIKey = os.Getenv("LIVEPEER_API_KEY")
}

// ConnectDatabase ...
func (server *Server) ConnectDatabase() {
	db := &database.Database{}

	db.ConnectDatabase()
	server.Handlers.Database = db
}
