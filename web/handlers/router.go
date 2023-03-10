package handlers

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/vspaz/tracker-rest-api/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Router struct {
	Logger        *logrus.Logger
	Conf          *config.Conf
	mux           *chi.Mux
	kafkaProducer *kafka.Producer
}

func NewRouter(conf *config.Conf, logger *logrus.Logger) *Router {
	return &Router{
		Logger: logger,
		Conf:   conf,
		mux:    chi.NewRouter(),
	}
}

func (r *Router) SetKafkaProducer(kafkaProducer *kafka.Producer) {
	r.kafkaProducer = kafkaProducer
}

func (r *Router) RegisterHandlers() {
	apiV1Prefix := "/api/v1/"
	r.mux.Get("/", r.Index)
	r.mux.Get(apiV1Prefix, r.Index)

	r.mux.Post(apiV1Prefix+"track", r.Track)
	r.mux.Post(apiV1Prefix+"t", r.Track)

	r.mux.Post(apiV1Prefix+"identify", r.Identify)
	r.mux.Post(apiV1Prefix+"i", r.Identify)

	r.mux.Post(apiV1Prefix+"alias", r.Alias)
	r.mux.Post(apiV1Prefix+"a", r.Alias)

	r.mux.Post(apiV1Prefix+"page", r.Page)
	r.mux.Post(apiV1Prefix+"p", r.Page)

	r.mux.Post(apiV1Prefix+"screen", r.Screen)
	r.mux.Post(apiV1Prefix+"s", r.Screen)

	r.mux.Post(apiV1Prefix+"batch", r.Batch)
	r.mux.Post(apiV1Prefix+"import", r.Import)

	r.mux.Get("/ping", r.Ping)
	r.mux.Post("/ping", r.Ping)
}

func (r *Router) handleShutDownGracefully(server *http.Server) {
	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, signals...)
	r.Logger.Infof("'%s' signal received, stopping server...", <-signalChannel)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		r.Logger.Errorf("error shutting down server: %s", err)
	}
}

func (r *Router) StartServer() {
	server := &http.Server{
		Addr:         r.Conf.Http.Server.HostAndPort,
		Handler:      http.TimeoutHandler(r.mux, r.Conf.Http.Server.RequestExecutionTimeout, "timeout occurred"),
		ReadTimeout:  r.Conf.Http.Server.ReadTimeout,
		WriteTimeout: r.Conf.Http.Server.WriteTimeout,
		IdleTimeout:  r.Conf.Http.Server.IdleTimeout,
	}
	go r.handleShutDownGracefully(server)
	r.Logger.Infof("starting server pid='%d' at port '%s'.", os.Getpid(), r.Conf.Http.Server.HostAndPort)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		r.Logger.Fatalf("error occurred: %s", err)
	}
	r.Logger.Info("server stopped.")
}
