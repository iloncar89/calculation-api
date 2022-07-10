package application

import (
	"github.com/iloncar89/calculation-api/src/controllers"
	"github.com/iloncar89/calculation-api/src/utils/constants/path/calculation"
	"github.com/iloncar89/calculation-api/src/utils/constants/path/ping"
	"github.com/iloncar89/calculation-api/src/utils/middleware"
	"net/http"
)

//mapUrls maps url in application.
func mapUrls() {
	pingHandler := http.HandlerFunc(controllers.PingController.Ping)
	mathOperationHandler := http.HandlerFunc(controllers.CalculationController.MathOperation)

	http.Handle(ping.Ping, pingHandler)
	http.Handle(calculation.AddCalculationUrlPath, middleware.Middleware(mathOperationHandler))
	http.Handle(calculation.SubtractCalculationUrlPath, middleware.Middleware(mathOperationHandler))
	http.Handle(calculation.MultiplyCalculationUrlPath, middleware.Middleware(mathOperationHandler))
	http.Handle(calculation.DivideCalculationUrlPath, middleware.Middleware(mathOperationHandler))

}
