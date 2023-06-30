package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wihdi/mnc/usecase"
	"github.com/wihdi/mnc/pkg"
	"github.com/wihdi/mnc/repository"
	"github.com/wihdi/mnc/api/handler"
)
func SetUpRouter() *gin.Engine {
	userRepo := repository.NewUserRepository("data/user.json")
	historyRepo := repository.NewHistoryRepository("data/history.json")
	userService := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userService)
	customerRepo := repository.NewCustomerRepository("data/user.json")
	paymentUsecase := usecase.NewPaymentUsecase(customerRepo)
	paymentHandler := handler.NewPaymentHandler(paymentUsecase,historyRepo)
	transferRepo := repository.NewTransferRepository("data/user.json")
	transferUsecase := usecase.NewTransferUsecase(transferRepo)
	transferHandler := handler.NewTransferHandler(transferUsecase,historyRepo,transferRepo)
	historyUsecase := usecase.NewHistoryUsecase(historyRepo)
	historyHandler := handler.NewHistoryHandler(*historyUsecase)
	
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	userRouters := apiV1.Group("/users")
	{
		userRouters.POST("/login", userHandler.Login)
		userRouters.Use(pkg.AuthMiddleware())
	}
	paymentRouters := apiV1.Group("/payment")
	{	paymentRouters.Use(pkg.AuthMiddleware())
		paymentRouters.POST("/", paymentHandler.ProcessPayment)
		
	}
	transferRouters := apiV1.Group("/transfer")
	{	transferRouters.Use(pkg.AuthMiddleware())
		transferRouters.POST("/transfers", transferHandler.CreateTransfer)
		
	}

	activityRouters := apiV1.Group("/activities")
	{
		transferRouters.Use(pkg.AuthMiddleware())
		activityRouters.GET("/all", historyHandler.GetAllActivities)
	}

	


	return r
}


