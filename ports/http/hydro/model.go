package hydro

import (
	"github.com/boreq/hydro/domain"
)

type Controller struct {
	UUID    string `json:"uuid"`
	Address string `json:"address"`
}

type Device struct {
	UUID     string   `json:"uuid"`
	ID       string   `json:"id"`
	Schedule []Period `json:"schedule"`
}

type Period struct {
	Start Time `json:"start"`
	End   Time `json:"end"`
}

type Time struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

func toControllers(controllers []*domain.Controller) []Controller {
	rv := make([]Controller, 0)
	for _, controller := range controllers {
		rv = append(rv, toController(controller))
	}
	return rv
}

func toController(controller *domain.Controller) Controller {
	return Controller{
		UUID:    controller.UUID().String(),
		Address: controller.Address().String(),
	}
}

func toDevices(devices []*domain.Device) []Device {
	rv := make([]Device, 0)
	for _, device := range devices {
		rv = append(rv, toDevice(device))
	}
	return rv
}

func toDevice(device *domain.Device) Device {
	return Device{
		UUID:     device.UUID().String(),
		ID:       device.ID().String(),
		Schedule: toSchedule(device.Schedule()),
	}
}

func toSchedule(schedule domain.Schedule) []Period {
	rv := make([]Period, 0)
	for _, period := range schedule.Periods() {
		rv = append(rv, toPeriod(period))
	}
	return rv
}

func toPeriod(period domain.Period) Period {
	return Period{
		Start: toTime(period.Start()),
		End:   toTime(period.End()),
	}
}

func toTime(time domain.Time) Time {
	return Time{
		Hour:   time.Hour(),
		Minute: time.Minute(),
	}
}
