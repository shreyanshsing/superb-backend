package utils

import (
    "fmt"
    "math/rand"
    "time"
)

func GenerateRandomUserName(base string) string {
    randomNumber := rand.Intn(10000) // Generate a random number
    timestamp := time.Now().Unix()   // Get the current timestamp

    // Combine the base string, random number, and timestamp
    username := fmt.Sprintf("%s_%d_%d", base, randomNumber, timestamp)

    return username
}