package utils

import (
	"errors"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	godotenv.Load(".env")

	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("FATAL: %s environment variable is not set", key)
	}

	return value
}

func GetFibonacciNumberWithPosition(pos int) (int, error) {
	if pos < 0 {
		return 0, errors.New("The position must be a non-negative integer")
	}

	if pos == 0 {
		return 0, nil
	}

	if pos == 1 {
		return 1, nil
	}

	if pos >= 44 {
		return 0, errors.New("The number must be less than 44, my developer did not buy more computing power :(")
	}

	a, _ := GetFibonacciNumberWithPosition(pos - 1)
	b, _ := GetFibonacciNumberWithPosition(pos - 2)

	return a + b, nil
}

func StringToInt(input string) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Printf("Cannot convert input \"%s\" to an integer: %s", input, err)
		return 0, errors.New("The position must be a integer")
	}
	return num, nil
}

func GetFibonacciNumberWithPositionBinet(pos int) (int, error) {
	if pos < 0 {
		return 0, errors.New("The position must be a non-negative integer")
	}

	// Binet formula constants
	phi := (1 + math.Sqrt(5)) / 2
	psi := (1 - math.Sqrt(5)) / 2

	// Calculate the nth Fibonacci number using Binet formula
	fibN := (math.Pow(phi, float64(pos)) - math.Pow(psi, float64(pos))) / math.Sqrt(5)

	return int(fibN + 0.5), nil
}
