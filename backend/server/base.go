package server

import (
	"context"
	"sync"
	"time"

	library "github.com/Z2hMedia/backend-go-library/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Base is the basis of an application
type Base struct {
	// Startup is when this service was created
	Startup time.Time
	// Config contains the service viper configuration
	Config *viper.Viper
	// Logger is a structured logger
	Logger *logrus.Logger

	grpc *grpcServer
	// erRep *errors.Reporter
	// stTr  *stats.Tracer

	unaryIntercept  []grpc.UnaryServerInterceptor
	_uiLock         *sync.Mutex
	streamIntercept []grpc.StreamServerInterceptor
	_siLock         *sync.Mutex

	handlers []GRPCHandler
	_hLock   *sync.Mutex
}

func New(ctx context.Context, conf Config) (*Base, error) {
	vc, err := getConfig(conf)

	if err != nil {
		return nil, err
	}

	// tr, err := stats.Setup(vc)
	// if err != nil {
	//   log.Printf("unable to setup tracing: %v", err)
	//   return nil, err
	// }

	// ec, err := errors.Setup(ctx, vc)
	// if err != nil {
	//   log.Printf("unable to setup error notification: %v", err)
	//   return nil, err
	// }

	ba := &Base{
		Config:  vc,
		Startup: time.Now(),

		// erRep:           ec,
		// stTr:            tr,

		handlers: []GRPCHandler{},
		_hLock:   &sync.Mutex{},

		unaryIntercept: conf.UnaryInterceptors,
		_uiLock:        &sync.Mutex{},

		streamIntercept: conf.StreamInterceptors,
		_siLock:         &sync.Mutex{},
	}
	ba.setupLogger()

	return ba, nil
}

func (ba *Base) setupLogger() {
	level := logrus.TraceLevel
	if ba.IsProd() {
		level = logrus.ErrorLevel
	}
	// WARNING: remove before deploying
	level = logrus.FatalLevel

	logrusLogger := logrus.New()
	logrusLogger.SetLevel(level)

	ba.Logger = logrusLogger
}

// IsProd ...
func (ba *Base) IsProd() bool {
	env := ba.Config.GetString("env")
	return env == library.EnvProduction
}

// AddUnaryInterceptor ...
func (ba *Base) AddUnaryInterceptor(i ...grpc.UnaryServerInterceptor) {
	ba._uiLock.Lock()
	ba.unaryIntercept = append(ba.unaryIntercept, i...)
	ba._uiLock.Unlock()
}

// AddStreamInterceptor ...
func (ba *Base) AddStreamInterceptor(s ...grpc.StreamServerInterceptor) {
	ba._siLock.Lock()
	ba.streamIntercept = append(ba.streamIntercept, s...)
	ba._siLock.Unlock()
}

// RegisterHandler ...
func (ba *Base) RegisterHandler(hn GRPCHandler) {
	ba._hLock.Lock()
	ba.handlers = append(ba.handlers, hn)
	ba._hLock.Unlock()
}

// Uptime ...
func (ba *Base) Uptime() time.Duration {
	return time.Since(ba.Startup)
}
