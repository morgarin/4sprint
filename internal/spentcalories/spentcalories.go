package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	strings := strings.Split(data, ",")
	if len(strings) != 3 {
		return 0, "", 0, errors.New("not enough arguments: expected data")
	}
	steps, err := strconv.Atoi(strings[0])
	if err != nil {
		return 0, "", 0, err
	}
	duration, err := time.ParseDuration(strings[2])
	if err != nil {
		return 0, "", 0, err
	}
	return steps, strings[1], duration, nil
}

// Дистанция
func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return stepLength * float64(steps) / mInKm

}

// Средняя скорость
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

// Вывод информации о тренировке
func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, trainingType, duration, err := parseTraining(data)
	switch trainingType {
	case "Бег":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		return fmt.Sprintf("Тип тренировки: %s\n Длительность: %.2f ч.\n Дистанция: %.2f км.\n Скорость: %.2f км/ч\n Сожгли калорий: %.2f\n", trainingType, duration.Hours(), dist, speed, calories), err
	case "Ходьба":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		return fmt.Sprintf("Тип тренировки: %s\n Длительность: %.2f ч.\n Дистанция: %.2f км.\n Скорость: %.2f км/ч\n Сожгли калорий: %.2f\n", trainingType, duration.Hours(), dist, speed, calories), err
	}
	return "", err
}

// Количество потраченных калорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if (steps <= 0) || (weight <= 0) || (height <= 0) || (duration <= 0) {
		return 0, fmt.Errorf("not enough arguments: expected data")
	}
	durationInMinutes := float64(duration.Minutes())
	return (weight * meanSpeed(steps, height, duration) * durationInMinutes) / minInH, nil
}

// Количество потраченных калорий при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if (steps <= 0) || (weight <= 0) || (height <= 0) || (duration <= 0) {
		return 0, fmt.Errorf("not enough arguments: expected data")
	}
	durationInMinutes := float64(duration.Minutes())
	return (weight * meanSpeed(steps, height, duration) * durationInMinutes) / minInH * walkingCaloriesCoefficient, nil
}
