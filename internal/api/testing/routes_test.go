package testing

import (
	"banking-api/internal/api/middleware"
	"banking-api/internal/api/route"
	"banking-api/internal/entity"
	"banking-api/internal/repository/mongodb"
	"banking-api/internal/service"
	"banking-api/pkg/auth"
	"banking-api/pkg/uhttp"
	"banking-api/pkg/umdw"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var router *gin.Engine

type Request struct {
	Method string
	Path   string
	Body   interface{}
}

func newAdminToken() string {
	userPayload := entity.User{
		FirstName: "sys",
		LastName:  "admin",
		Email:     "admin@example.com",
		Enabled:   true,
	}
	jwtToken, _ := api_middleware.NewJwtToken(userPayload)
	return jwtToken
}

func simulateJSONRequest(r http.Handler, reqParams Request, token bool) (*httptest.ResponseRecorder, uhttp.Response, error) {
	structToMap := func(in interface{}) (map[string]interface{}, error) {
		var result map[string]interface{}

		inStr, err := json.Marshal(&in)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(inStr, &result)
		if err != nil {
			return nil, err
		}

		return result, nil
	}
	structToBody := func(o interface{}) *bytes.Buffer {
		m, err := structToMap(o)
		if err != nil {
			log.Fatal(err)
		}
		delete(m, "_id")
		jsonMdwStrByte, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		return bytes.NewBuffer(jsonMdwStrByte)
	}
	req, _ := http.NewRequest(reqParams.Method, reqParams.Path, structToBody(reqParams.Body))
	req.Header.Add("Content-Type", "application/json")
	if token {
		req.Header.Add(auth.JwtAuthorizationHeader, newAdminToken())
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response uhttp.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)

	return w, response, err
}

func simulateFormDataFileRequest(r http.Handler, reqParams Request, filename string, token bool) (*httptest.ResponseRecorder, uhttp.Response, error) {
	createBodyMultiPartFormDataFile := func() (bytes.Buffer, string, error) {
		var body bytes.Buffer
		multipartWriter := multipart.NewWriter(&body)

		fw, err := multipartWriter.CreateFormFile("file", filename)
		if err != nil {
			return bytes.Buffer{}, "", err
		}

		file, err := os.Open(filename)
		if err != nil {
			return bytes.Buffer{}, "", err
		}

		_, err = io.Copy(fw, file)
		if err != nil {
			return bytes.Buffer{}, "", err
		}

		_ = multipartWriter.Close()

		return body, multipartWriter.FormDataContentType(), nil
	}

	body, formDataContentType, err := createBodyMultiPartFormDataFile()
	if err != nil {
		return nil, uhttp.Response{}, err
	}

	req, _ := http.NewRequest(reqParams.Method, reqParams.Path, &body)
	req.Header.Add("Content-Type", formDataContentType)
	if token {
		req.Header.Add(auth.JwtAuthorizationHeader, newAdminToken())
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response uhttp.Response
	err = json.Unmarshal(w.Body.Bytes(), &response)

	return w, response, err
}

func TestSetApiRoutes(t *testing.T) {
	_ = os.Setenv("GOOGLE_API_KEY", "AIzaSyDBzNMNiFqGWklQeijykzRMkO201MMDh0c")

	gin.DisableConsoleColor()
	gin.SetMode(gin.TestMode)
	router = gin.New()
	router.Use(umdw.BodyContext)

	TestInitMongoDBTest(t)

	feeRepository := mongodb_repository.NewFeeMongoDB(&db)
	userRepository := mongodb_repository.NewUserMongoDB(&db)
	queryRepository := mongodb_repository.NewQueryMongoDB(&db)
	orderRepository := mongodb_repository.NewOrderMongoDB(&db)
	driverRepository := mongodb_repository.NewDriverMongoDB(&db)
	workLoadRepository := mongodb_repository.NewWorkLoadMongoDB(&db)
	businessRepository := mongodb_repository.NewBusinessMongoDB(&db)
	municipalityRepository := mongodb_repository.NewMunicipalityMongoDB(&db)
	driverVehicleRepository := mongodb_repository.NewDriverVehicleMongoDB(&db)
	businessApiKeyRepository := mongodb_repository.NewBusinessApiKeyMongoDB(&db)
	businessOriginRepository := mongodb_repository.NewBusinessOriginMongoDB(&db)
	businessDestinationRepository := mongodb_repository.NewBusinessDestinationMongoDB(&db)
	businessCollaboratorRepository := mongodb_repository.NewBusinessCollaboratorMongoDB(&db)

	feeService := service.NewFeeService(feeRepository)
	userService := service.NewUserService(userRepository)
	queryService := service.NewQueryService(queryRepository)
	orderService := service.NewOrderService(orderRepository)
	driverService := service.NewDriverService(driverRepository)
	workLoadService := service.NewWorkLoadService(workLoadRepository)
	businessService := service.NewBusinessService(businessRepository)
	municipalityService := service.NewMunicipalityService(municipalityRepository)
	driverVehicleService := service.NewDriverVehicleService(driverVehicleRepository)
	businessApiKeyService := service.NewBusinessApiKeyService(businessApiKeyRepository)
	businessOriginService := service.NewBusinessOriginService(businessOriginRepository)
	businessDestinationService := service.NewBusinessDestinationService(businessDestinationRepository)
	businessCollaboratorService := service.NewBusinessCollaboratorService(businessCollaboratorRepository)

	services := &service.Services{
		Fee:                  feeService,
		User:                 userService,
		Query:                queryService,
		Order:                orderService,
		Driver:               driverService,
		WorkLoad:             workLoadService,
		Business:             businessService,
		Municipality:         municipalityService,
		DriverVehicle:        driverVehicleService,
		BusinessApiKey:       businessApiKeyService,
		BusinessOrigin:       businessOriginService,
		BusinessDestination:  businessDestinationService,
		BusinessCollaborator: businessCollaboratorService,
	}

	api_route.SetRoutes(router, services, optimizerEngine)

	mt := entity.Municipality{
		Country: entity.MunicipalityCountry,
		Region:  "Región Metropolitana",
		Name:    "Santiago",
		Kind:    "base",
	}

	_, err := municipalityService.CreateMunicipality(mt)

	assert := assert.New(t)
	assert.Nil(err)
}
