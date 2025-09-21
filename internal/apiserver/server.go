package apiserver

import (
	"context"
	"os"

	"github.com/moweilong/mo/pkg/contextx"
	genericoptions "github.com/moweilong/mo/pkg/options"
	"github.com/moweilong/mo/pkg/server"
	"github.com/moweilong/mo/pkg/store/where"
	"github.com/moweilong/mo/pkg/version"
)

var (
	// Name is the name of the compiled software.
	Name = "mo"

	ID, _ = os.Hostname()

	Version = version.Get().String()
)

// Config contains application-related configurations.
type Config struct {
	HTTPOptions *genericoptions.HTTPOptions
	TLSOptions  *genericoptions.TLSOptions
}

// NewServer initializes and returns a new Server instance.
func (cfg *Config) NewServer(ctx context.Context) (*Server, error) {
	where.RegisterTenant("userID", func(ctx context.Context) string {
		return contextx.UserID(ctx)
	})

	// Create the core server instance.
	srv, err := InitializeWebServer(cfg)
	if err != nil {
		return nil, err
	}

	return &Server{srv: srv}, nil
}

// Server represents the web server.
type Server struct {
	srv server.Server
}

// ServerConfig contains the core dependencies and configurations of the server.
type ServerConfig struct {
	cfg *Config
}

// Run starts the server and listens for termination signals.
// It gracefully shuts down the server upon receiving a termination signal.
func (s *Server) Run(ctx context.Context) error {
	return server.Serve(ctx, s.srv)
}

type AggregatorServer struct {
	serverConfig *ServerConfig
	// grpc         server.Server
	gin server.Server
}

var _ server.Server = &AggregatorServer{}

func NewAggregatorServer(serverConfig *ServerConfig) (server.Server, error) {
	// grpcSrv := serverConfig.NewGrpcServer()
	ginSrv := serverConfig.NewGinServer()
	return &AggregatorServer{gin: ginSrv}, nil
}

func (s *AggregatorServer) RunOrDie() {
	// go s.grpc.RunOrDie()
	go s.gin.RunOrDie()
}

func (s *AggregatorServer) GracefulStop(ctx context.Context) {
	// s.grpc.GracefulStop(ctx)
	s.gin.GracefulStop(ctx)
}
