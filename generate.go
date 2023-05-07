package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"time"
)

var ctx = context.Background()
var	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB
	})
const geocodes_to_add = 1000000
	
func main() {


	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	} else {
		fmt.Println("Connected to Redis:", pong)
	}

	// Add some locations
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < geocodes_to_add; i++ {
		lat := randomLatitude()
		lon := randomLongitude()
		geocodeKey := fmt.Sprintf("geo:%d", i)

		fmt.Println("Adding geocode:", geocodeKey, lat, lon)

		// Add geocode to Redis
		err := rdb.GeoAdd(ctx, "geocodes", &redis.GeoLocation{
			Name:      geocodeKey,
			Longitude: lon,
			Latitude:  lat,
		}).Err()

		if err != nil {
			fmt.Println("Error adding geocode to Redis:", err)
			break
		}
	}

	fmt.Println("Process complete. Geocodes added:", geocodes_to_add)
}

func randomLatitude() float64 {
	min := 24.396308
	max := 49.384358
	return min + rand.Float64()*(max-min)
}

func randomLongitude() float64 {
	min := -125.000000
	max := -66.934570
	return min + rand.Float64()*(max-min)
}