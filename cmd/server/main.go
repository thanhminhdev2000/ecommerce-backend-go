package main

import (
	"github.com/thanhminhdev2000/ecommerce-backend-go/internal/routers"
)



func main() {
  r := routers.NewRouter()
  r.Run()
}