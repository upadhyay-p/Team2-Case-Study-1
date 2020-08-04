package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/auth"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/customer"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/order"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/restaurant"
	"google.golang.org/grpc"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_in_celsius",
		Help: "Current teamperature of CPU in degree celsius",
	})
	apiHitcount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_api_hit_count",
		Help: "Number of times APIs were hitted",
	})
)

func init() {
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(apiHitcount)
}

func apiHitCountTracker(c *gin.Context) {
	apiHitcount.Inc()
	c.Next()
}

// homepage of API
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Team 2": "Hello from Aadithya, Abhishek, Priya, Shashi!",
	})
}

func main() {

	cpuTemp.Set(65.3)
	fmt.Println("Hello from the ginAPI :)")
	conn, err := grpc.Dial("0.0.0.0:5051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v", err)
		os.Exit(1)
	}

	defer conn.Close()

	customer.CSC = customer.NewCustomerServiceClient(conn)
	order.OSC = order.NewOrderServiceClient(conn)
	restaurant.RSC = restaurant.NewRestaurantServiceClient(conn)

	router := SetupRoutes(true) // called if flag = true to enable authentication

	router.Run("localhost:9001")
}

func SetupRoutes(authFlag bool) *gin.Engine {
	router := gin.Default()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	loginRouter := router.Group("/login")
	loginRouter.POST("/", auth.Login)

	apiRouter := router.Group("/api")
	apiRouter.Use(apiHitCountTracker)

	if authFlag {
		apiRouter.Use(auth.VerifyUser)
	}

	apiRouter.GET("/", Index)

	apiRouter.GET("/orders", order.GetAll)
	apiRouter.GET("/order/:id", order.GetOne)
	apiRouter.POST("/order", order.Add)

	apiRouter.GET("/customers", customer.GetAll)
	apiRouter.GET("/customer/:id", customer.GetOne)
	apiRouter.POST("/customer", customer.Add)

	apiRouter.GET("/restaurants", restaurant.GetAll)
	apiRouter.GET("/restaurant/:id", restaurant.GetOne)
	// apiRouter.POST("/restaurant", restaurant.Add) [TODO]
	return router
}
