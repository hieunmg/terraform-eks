package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"
	"weshare/components/appcontext"
	"weshare/components/uploadprovider"
	"weshare/middleware"
	"weshare/utils"

	uploadcontroller "weshare/modules/upload/controller"

	"weshare/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "weshare_requests_total",
			Help: "Total number of HTTP requests to weshare.",
		},
		[]string{"method", "path", "status"},
	)
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	s3Provider := uploadprovider.NewS3Provider(
		config.S3BucketName,
		config.S3Region,
		config.S3ApiKey,
		config.S3SecretKey,
		config.S3Domain,
	)

	appCtx := appcontext.NewAppContext(
		db,
		config.TokenSymmetricKey,
		config.AccessTokenDuration,
		config.RefreshTokenDuration,
		s3Provider)

	runServer(appCtx, config)

}

func runServer(appCtx appcontext.AppContext, config utils.Config) error {
	prometheus.MustRegister(httpRequestsTotal)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.Recover(appCtx))
	r.Use(middleware.Authorize(config.TokenSymmetricKey))

	r.Use(func(c *gin.Context) {
		httpRequestsTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			fmt.Sprintf("%d", c.Writer.Status()),
		).Inc()
		c.Next()
	})
	// Expose the Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": "v3",
		})
	})

	v1 := r.Group("/api/v1")
	v1.POST("/upload", uploadcontroller.Upload(appCtx))

	router.AuthRouter(v1, appCtx)

	return r.Run(config.HTTPServerAddress)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}
