package job

// import (
// 	"time"
// 	"weKnow/model"
// 	"weKnow/service"

// 	"github.com/go-co-op/gocron"
// )

// type KnownJob struct {
// 	Jobs      model.Job
// 	Scheduler *gocron.Scheduler
// 	Service   *service.KnownService
// }

// func NewJob(service *service.KnownService) *KnownJob {
// 	return &KnownJob{
// 		Scheduler: gocron.NewScheduler(time.UTC),
// 		Service:   service}
// }

// func (js *KnownJob) ScheduleJobs() error {
// 	// js.Service.SendWhatsApp()
// 	_, err := js.Scheduler.Cron("* * * * *").Do(js.Service.SendJobEmail)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (js *KnownJob) StartScheduler() {
// 	js.Scheduler.StartAsync()
// }
