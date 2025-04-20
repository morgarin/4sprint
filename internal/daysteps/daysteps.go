package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	strings := strings.Split(data, ",")
	if len(strings) != 2 {
		return 0, 0, errors.New("not enough arguments: expected data")
	}
	steps, err := strconv.Atoi(strings[0])
	if err != nil {
		return 0, 0, errors.New("cant convert steps(str) to int")
	}
	duration, err := time.ParseDuration(strings[1])
	if err != nil {
		return 0, 0, errors.New("cant convert duration")
	}
	if steps <= 0 {
		return 0, 0, errors.New("steps count: <= 0")
	}
	return steps, duration, err
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distance := stepLength * float64(steps)
	distance = distance / float64(mInKm)
	calories := WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d\n Дистанция составила %.2f км.\n Вы сожгли %.2f ккал.", steps, distance, calories)
}
