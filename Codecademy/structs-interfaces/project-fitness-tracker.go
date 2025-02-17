package main

import (
	"fmt"
	"time"
)

// TODO: Define the `Workout` interface here
type Workout interface {
	// TODO: Add method signatures here
	Duration() time.Duration
	CaloriesBurned() float64
	RecordStats()
	GetType() string
}

// CardioWorkout struct
type CardioWorkout struct {
	duration     time.Duration
	distance     float64
	avgHeartRate float64
}

func (c CardioWorkout) Duration() time.Duration {
	return c.duration
}

// TODO: Implement other CardioWorkout methods
func (c CardioWorkout) CaloriesBurned() float64 {
	return c.duration.Minutes() * 10 * (c.avgHeartRate / 100)
}

func (c CardioWorkout) RecordStats() {
	fmt.Println("Workout duration:", c.duration)
	fmt.Println("Distance:", c.distance)
	fmt.Println("Avg heart rate:", c.avgHeartRate)
	fmt.Println("Calories burned:", c.CaloriesBurned())
}

func (c CardioWorkout) GetType() string {
	return "Cardio"
}

// StrengthWorkout struct
type StrengthWorkout struct {
	duration time.Duration
	weight   int
	reps     int
	sets     int
}

func (s StrengthWorkout) Duration() time.Duration {
	return s.duration
}

// TODO: Implement other StrengthWorkout methods
func (s StrengthWorkout) CaloriesBurned() float64 {
	// not included in the instruction is to use typecasting to avoid errors:
	return s.duration.Minutes() * 5 * (float64(s.weight) / 10)
}

func (s StrengthWorkout) RecordStats() {
	fmt.Println("Workout duration:", s.duration)
	fmt.Println("Weight lifted:", s.weight)
	fmt.Println("Reps:", s.reps)
	fmt.Println("Sets", s.sets)
	fmt.Println("Calories burned:", s.CaloriesBurned())
}

func (s StrengthWorkout) GetType() string {
	return "Weightlifting"
}

func summarizeWorkouts(workouts []Workout) {
	// TODO: Implement workout summary
	for _, workout := range workouts {
		fmt.Println("New workout:", workout.GetType())
		workout.RecordStats()
	}
}

func main() {
	// TODO: Create workout slice and add workouts
	workouts := []Workout{
		CardioWorkout{
			duration:     44 * time.Minute,
			distance:     3.8,
			avgHeartRate: 140,
		},
		StrengthWorkout{
			duration: 50 * time.Minute,
			weight:   120,
			sets:     9,
			reps:     8,
		},
	}

	fmt.Println("Workout Summary:")
	// TODO: Call summarizeWorkouts function
	summarizeWorkouts(workouts)
}
