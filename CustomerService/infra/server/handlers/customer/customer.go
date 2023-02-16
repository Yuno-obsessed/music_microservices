package customer

import (
	"github.com/Yuno-obsessed/music_microservices/CustomerService/service/customer"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/gin-gonic/gin"
)

type Customer struct {
	customer.CustomerService
	logger.CustomLogger
}

func NewCustomer(service customer.CustomerService) Customer {
	return Customer{
		service,
		logger.NewLogger(),
	}
}

func GetUserById(c *gin.Context) {

}
