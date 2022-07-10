package application

import (
	"fmt"
	"github.com/iloncar89/calculation-api/src/controllers"
	"github.com/iloncar89/calculation-api/src/utils/constants/path/calculation"
	"github.com/iloncar89/calculation-api/src/utils/middleware"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculationController_AddOperation(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/add?&y=7&x=-2.2", nil)
	w := httptest.NewRecorder()
	addOperationUsingPathParamsHandler := http.HandlerFunc(controllers.CalculationController.MathOperation)
	http.Handle(calculation.AddCalculationUrlPath, middleware.Middleware(addOperationUsingPathParamsHandler))

	addOperationUsingPathParamsHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))

	w1 := httptest.NewRecorder()

	addOperationUsingPathParamsHandler(w1, req)

	resp1 := w.Result()
	body1, _ := ioutil.ReadAll(resp1.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(string(body1))

}
