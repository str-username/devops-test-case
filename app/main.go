package main

import (
    "context"
    "encoding/base64"
    "fmt"
    "log"
    "net/http"
    "time"
    
    "github.com/redis/go-redis/v9"
)

var redisAddress = "127.0.0.1:6379"
var httpServerHost = ":8080"

func main() {
    // Check if redis address not empty
    if redisAddress == "" {
        log.Fatal("error: oops, i am an config error empty")
    }
    
    // My super magic
    RedisAddressDecoded, err := base64.StdEncoding.DecodeString(redisAddress)
    
    if err != nil {
        log.Fatal("error: oops, i am an config error decode")
    }
    
    decodedRedisAddress := string(RedisAddressDecoded)
    
    ctx := context.Background()
    
    redisClient := redis.NewClient(&redis.Options{
        Addr:     decodedRedisAddress, // I don’t know why, but I wanted the address to be passed as a Base64 string, <addr>:port.
        Password: "",
        DB:       0,
    })
    
    pong, err := redisClient.Ping(ctx).Result()
    if err != nil {
        log.Printf("error: oops, i am an config error, %s", err)
    }
    
    log.Printf("info: redis connect: %s", pong)
    
    err = redisClient.Set(ctx, "response", "OK", 10*time.Second).Err()
    if err != nil {
        log.Printf("error: set key error: %s", err)
    }
    
    res, err := redisClient.Get(ctx, "response").Result()
    if err != nil {
        log.Printf("error: get key error: %s", err)
    }
    
    log.Printf("info: key value %s:", res)
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, res)
    })
    
    log.Printf("info: run server %s", httpServerHost)
    log.Fatal(http.ListenAndServe(httpServerHost, nil))
}
