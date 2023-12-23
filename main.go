package main

import (
	"fmt"
	"time"
)

// Sensor simulates a fire sensor.
type Sensor struct {
	ID       int
	Location string
}

// Alarm simulates a fire alarm.
type Alarm struct {
	ID       int
	Location string
	Active   bool
}

// FireAlarmController represents the main control panel for the fire alarm system.
type FireAlarmController struct {
	Sensors []Sensor
	Alarms  []Alarm
}

// NewFireAlarmController initializes a new FireAlarmController.
func NewFireAlarmController() *FireAlarmController {
	return &FireAlarmController{
		Sensors: []Sensor{},
		Alarms:  []Alarm{},
	}
}

// AddSensor adds a new fire sensor to the system.
func (fac *FireAlarmController) AddSensor(location string) {
	sensor := Sensor{
		ID:       len(fac.Sensors) + 1,
		Location: location,
	}
	fac.Sensors = append(fac.Sensors, sensor)
	fmt.Printf("Added sensor %d at %s\n", sensor.ID, sensor.Location)
}

// AddAlarm adds a new fire alarm to the system.
func (fac *FireAlarmController) AddAlarm(location string) {
	alarm := Alarm{
		ID:       len(fac.Alarms) + 1,
		Location: location,
		Active:   false,
	}
	fac.Alarms = append(fac.Alarms, alarm)
	fmt.Printf("Added alarm %d at %s\n", alarm.ID, alarm.Location)
}

// ActivateAlarm activates a specific alarm.
func (fac *FireAlarmController) ActivateAlarm(alarmID int) {
	for i := range fac.Alarms {
		if fac.Alarms[i].ID == alarmID {
			fac.Alarms[i].Active = true
			fmt.Printf("Alarm %d at %s activated!\n", fac.Alarms[i].ID, fac.Alarms[i].Location)
			return
		}
	}
	fmt.Printf("Alarm with ID %d not found.\n", alarmID)
}

// SimulateFire simulates the occurrence of a fire by activating a random alarm.
func (fac *FireAlarmController) SimulateFire() {
	if len(fac.Alarms) == 0 {
		fmt.Println("No alarms available to simulate fire.")
		return
	}

	randAlarmIndex := rand.Intn(len(fac.Alarms))
	alarmID := fac.Alarms[randAlarmIndex].ID
	fac.ActivateAlarm(alarmID)
}

func main() {
	// Create a new fire alarm controller
	controller := NewFireAlarmController()

	// Add sensors and alarms
	controller.AddSensor("Kitchen")
	controller.AddSensor("Living Room")
	controller.AddAlarm("Hallway")
	controller.AddAlarm("Bedroom")

	// Simulate a fire
	controller.SimulateFire()

	// Simulate another fire after 5 seconds
	time.Sleep(5 * time.Second)
	controller.SimulateFire()
}
