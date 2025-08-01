package scheduler

import (
	"time"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"

	"zero-fusion/share/cron"
)

const (
	defaultSyncInterval = 10 * time.Second
)

type ShedulerServer struct {
	srv *asynq.Scheduler
}

func (s *ShedulerServer) Register(task cron.PeriodicTask) {
	entryID, err := s.srv.Register(
		task.CronSpec(),
		task.Task(),
		task.Options()...,
	)
	if err != nil {
		logx.Errorf("register task %s error: %v", task.Name(), err)
		return
	}

	logx.Infof("register task %s success, entryID: %s", task.Name(), entryID)
}

func (s *ShedulerServer) Start() {
	if err := s.srv.Run(); err != nil {
		panic(err)
	}
	logx.Info("scheduler started")
}

func (s *ShedulerServer) Stop() {
	s.srv.Shutdown()
	logx.Info("scheduler stopped")
}

func NewShedulerServer(clientOpt asynq.RedisClientOpt, schedulerOpts *asynq.SchedulerOpts) cron.PeriodicScheduler {
	return &ShedulerServer{
		srv: asynq.NewScheduler(
			clientOpt,
			schedulerOpts,
		),
	}
}

type Option func(*ShedulerDynamicServer)

type ShedulerDynamicServer struct {
	srv          *asynq.PeriodicTaskManager
	Provider     asynq.PeriodicTaskConfigProvider
	SyncInterval time.Duration
}

func NewShedulerDynamicServer(clientOpt asynq.RedisClientOpt, schedulerOpts *asynq.SchedulerOpts,
	provider asynq.PeriodicTaskConfigProvider, opts ...Option) cron.PeriodicDynamicScheduler {

	ds := &ShedulerDynamicServer{}
	for _, opt := range opts {
		opt(ds)
	}
	if ds.SyncInterval == 0 {
		ds.SyncInterval = defaultSyncInterval
	}

	srv, err := asynq.NewPeriodicTaskManager(
		asynq.PeriodicTaskManagerOpts{
			RedisConnOpt:               clientOpt,
			PeriodicTaskConfigProvider: provider,
			SchedulerOpts:              schedulerOpts,
			SyncInterval:               ds.SyncInterval,
		})
	if err != nil {
		panic(err)
	}

	return &ShedulerDynamicServer{
		srv: srv,
	}
}

func (s *ShedulerDynamicServer) Register(provider asynq.PeriodicTaskConfigProvider) {
	if s.Provider != nil {
		return
	}
	s.Provider = provider
}

func (s *ShedulerDynamicServer) Start() {
	if err := s.srv.Run(); err != nil {
		panic(err)
	}
}

func (s *ShedulerDynamicServer) Stop() {
	s.srv.Shutdown()
	logx.Info("dynamic scheduler stopped")
}

func WithSyncInterval(interval time.Duration) Option {
	return func(s *ShedulerDynamicServer) {
		s.SyncInterval = interval
	}
}

func PreEnqueueFunc(task *asynq.Task, opts []asynq.Option) {
	logx.Infof("pre handle enqueue taskType: %s, opts: %v", task.Type(), opts)
}

func PostEnqueueFunc(info *asynq.TaskInfo, err error) {
	if err == nil {
		logx.Infof("post handle enqueue taskType: %s success, taskID: %s", info.Type, info.ID)
		return
	}
	logx.Errorf("post handle enqueue error: %v", err)
}
