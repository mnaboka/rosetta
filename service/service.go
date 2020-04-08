package service

import (
	"context"
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/log"
)

var ErrAlreadyRunning = errors.New("Service already running")

type Service interface {
	// Name retrieves the name of the service, that will be used
	// to identify the service in log messages
	Name() string

	// Running indicates if the service is currently running
	Running() bool

	// Start runs the service and blocks until the service finishes,
	// returns an error when service failed
	Start(ctx context.Context) error
}

// Rosetta HTTP Service
// Geth Node Service
// Monitor Service (package monitor)

func RunServices(ctx context.Context, services ...Service) error {
	ctx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()

	var wg sync.WaitGroup
	wg.Add(len(services))

	errors := NewErrorCollector()
	for _, _srv := range services {
		go func(srv Service) {
			defer wg.Done()
			log.Info("Starting service", "srv", srv.Name())
			err := srv.Start(ctx)
			if err != nil {
				// Stop Services on the first error occured
				cancelFn()
				errors.Add(err)
			}
		}(_srv)
	}

	wg.Wait()
	return errors.Error()
}

type ServiceManager struct {
	wg             sync.WaitGroup
	errorCollector ErrorCollector
	ctx            context.Context
	stopAll        context.CancelFunc
}

func NewServiceManager(ctx context.Context) *ServiceManager {
	ctx, stopAll := context.WithCancel(ctx)
	return &ServiceManager{
		ctx:     ctx,
		stopAll: stopAll,
	}

}

func (sm *ServiceManager) Add(srv Service) {
	sm.wg.Add(1)
	go func() {
		defer sm.wg.Done()
		err := srv.Start(sm.ctx)
		if err != nil {
			sm.errorCollector.Add(err)
			sm.stopAll()
		}
	}()
}

func (sm *ServiceManager) Wait() error {
	sm.wg.Wait()
	return sm.errorCollector.Error()
}