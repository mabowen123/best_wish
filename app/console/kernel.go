package console

import (
	"best_wish/app/console/commands/tipoff"
	tipoffdao "best_wish/app/dao/tipoff"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/schedule"
	"github.com/goravel/framework/facades"
)

type Kernel struct {
}

func (kernel *Kernel) Schedule() []schedule.Event {
	return []schedule.Event{
		facades.Schedule().Command("tip:off:reptile").Cron("@every 5s").SkipIfStillRunning(),
		facades.Schedule().Command("tip:off:notice").Cron("@every 1s").SkipIfStillRunning(),
		facades.Schedule().Call(func() { tipoffdao.DelOldData() }).Daily(),
	}
}

func (kernel *Kernel) Commands() []console.Command {
	return []console.Command{
		&tipoff.Reptile{},
		&tipoff.Notify{},
	}
}
