package controllers

import (
	"encoding/json"
	"github.com/iloncar89/calculation-api/src/services"
	"github.com/iloncar89/calculation-api/src/utils/logger"
	"github.com/iloncar89/calculation-api/src/utils/parser"
	"github.com/iloncar89/calculation-api/src/utils/restErrors"
	"net/http"
)

var CalculationController calculationControllerInterface = &calculationController{}

type calculationControllerInterface interface {
	MathOperation(w http.ResponseWriter, r *http.Request)
}

type calculationController struct{}

//MathOperation handler function for /add /subtract /multiply and /divide url. Resolves GET, POST, OPTION http function on those urls.
func (cont *calculationController) MathOperation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		logger.InfoLogger.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		logger.InfoLogger.Print("parser.CalculationParser.ParseCalculationPathVariables")
		calculationRequest, restErr := parser.CalculationParser.ParseCalculationPathVariables(*r)
		if calculationRequest != nil {
			calculationRequest.Action = r.URL.Path
		}

		if restErr != nil {
			w.WriteHeader(restErr.Status())
			response, err := json.Marshal(restErr)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, err = w.Write(response)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		logger.InfoLogger.Print("Calling services.CalculationService.PerformBasicMathOperation with ", *calculationRequest)
		calculation, restErr := services.CalculationService.BasicMathOperation(*calculationRequest)
		if restErr != nil {
			w.WriteHeader(restErr.Status())
			response, err := json.Marshal(restErr)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, err = w.Write(response)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		response, err := json.Marshal(calculation)
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(response)
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger.InfoLogger.Print("Response sent ", calculation, "for GET request", r.URL.Path, "\n")
		return

	case http.MethodPost:
		logger.InfoLogger.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		logger.InfoLogger.Print("parser.CalculationParser.ParseCalculationPathVariables")
		calculationRequest, restErr := parser.CalculationParser.ParseCalculationPathVariables(*r)
		if calculationRequest != nil {
			calculationRequest.Action = r.URL.Path
		}

		if restErr != nil && restErr.Message() != restErrors.MissingBothCalculationParamsErrorMessage {
			w.WriteHeader(restErr.Status())
			response, err := json.Marshal(restErr)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, err = w.Write(response)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		} else if restErr == nil {
			logger.InfoLogger.Print("Calling services.CalculationService.PerformBasicMathOperation with ", *calculationRequest)
			calculation, restErr := services.CalculationService.BasicMathOperation(*calculationRequest)
			if restErr != nil {
				w.WriteHeader(restErr.Status())
				response, err := json.Marshal(restErr)
				if err != nil {
					logger.ErrorLogger.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				_, err = w.Write(response)
				if err != nil {
					logger.ErrorLogger.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				return
			}

			response, err := json.Marshal(calculation)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, err = w.Write(response)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			logger.InfoLogger.Print("Response sent ", calculation, "for POST request", r.URL.Path, " with constants params\n")
			return
		}

		logger.InfoLogger.Print("Trying to parse body in POST request")
		err := json.NewDecoder(r.Body).Decode(&calculationRequest)
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.WriteHeader(restErr.Status())
			restErr = restErrors.NewBadRequestError(restErrors.MissingParametersInPostRequest)
			response, err := json.Marshal(restErr)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, err = w.Write(response)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
		calculationRequest.Action = r.URL.Path
		logger.InfoLogger.Print("Calling services.CalculationService.PerformBasicMathOperation with ", *calculationRequest)

		calculation, restErr := services.CalculationService.BasicMathOperation(*calculationRequest)
		if restErr != nil {
			w.WriteHeader(restErr.Status())
			response, err := json.Marshal(restErr)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, err = w.Write(response)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		response, err := json.Marshal(calculation)
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(response)
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger.InfoLogger.Print("Response sent ", calculation, "for POST request with body ", r.URL.Path, ",\n")
		return

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
