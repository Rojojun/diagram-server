package bootstrap

import (
	"context"
	"diagram-server/internal/database"
	"diagram-server/internal/handler"
	"diagram-server/internal/persistance"
	"diagram-server/internal/service"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Application struct {
	port           string
	server         *http.Server
	db             database.Connector
	diagramHandler *handler.DiagramHandler
}

func NewApplication() *Application {
	return &Application{
		port: getEnv("PORT", ":8080"),
	}
}

func (app *Application) Run(ctx context.Context) error {
	if err := app.initDatabase(ctx); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	defer app.shutdown(ctx)

	app.initDependencies()
	app.initWebServer()

	// 배너 및 시스템 정보 출력
	StartUp(app.port)

	go func() {
		<-ctx.Done()
		log.Printf("Shutting down server on %s ...", app.port)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := app.server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("[Error] Server shutdown error: %v", err)
		}
	}()

	if err := app.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func (app *Application) initDatabase(ctx context.Context) error {
	cfg := database.Config{
		Type:     database.MongoDB,
		Uri:      "mongodb://localhost:27017",
		Database: "diagram",
		Pool: database.PoolConfig{
			MinSize:     getEnvUint64("DB_POOL_MIN", 5),
			MaxSize:     getEnvUint64("DB_POOL_MAX", 100),
			MaxIdleTime: getEnvDuration("DB_POOL_IDLE_TIME", 30*time.Second),
		},
	}

	conn, err := database.NewConnector(cfg)
	if err != nil {
		return err
	}

	if err := conn.Connect(ctx); err != nil {
		return err
	}

	if err := conn.Ping(ctx); err != nil {
		return err
	}

	app.db = conn
	log.Println("Database connected successfully")
	return nil
}

func (app *Application) initWebServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	mux.HandleFunc("POST /api/diagrams", app.diagramHandler.Create)
	mux.HandleFunc("GET /api/diagrams/{type}", app.diagramHandler.GetAllByType)
	mux.HandleFunc("GET /api/diagrams/{id}", app.diagramHandler.GetByID)
	mux.HandleFunc("DELETE /api/diagrams/{id}", app.diagramHandler.Delete)

	app.server = &http.Server{
		Addr:    app.port,
		Handler: mux,
	}
}

func (app *Application) initDependencies() {
	mongoConn := app.db.(*database.MongoConnector)
	db := mongoConn.DB()

	diagramRepo := persistance.NewDiagramRepository(db)
	diagramSvc := service.NewDiagramService(diagramRepo)
	app.diagramHandler = handler.NewDiagramHandler(diagramSvc)

	log.Println("[INFO] Dependencies initialized")
}

func (app *Application) shutdown(ctx context.Context) {
	if app.db != nil {
		if err := app.db.Disconnect(ctx); err != nil {
			log.Fatalf("[Error] Database disconnect error: %v", err)
		}

		log.Println("Database disconnected")
	}
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

func getEnvUint64(key string, defaultVal uint64) uint64 {
	if v := os.Getenv(key); v != "" {
		if parsed, err := strconv.ParseUint(v, 10, 64); err == nil {
			return parsed
		}
	}
	return defaultVal
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if parsed, err := time.ParseDuration(v); err == nil {
			return parsed
		}
	}
	return defaultVal
}
