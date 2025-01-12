package main

import (
	"github.com/thanhminhdev2000/ecommerce-backend-go/internal/initialize"
)



func main() {
  // r := routers.NewRouter()
  // InitMySQL
  // InitRedis
  // InitKafka
  // ...
  // r.Run()

  initialize.Run()
}