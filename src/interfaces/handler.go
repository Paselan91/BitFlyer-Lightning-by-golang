package interfaces

import (
	"app/src/config"
	"app/src/interfaces/bitflyer"
	"app/src/usecase"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Run start server
func Run(e *echo.Echo, port string) {
	log.Printf("Server running at http://localhost:%s/", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func BindValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	if err := c.Validate(i); err != nil {
		return c.String(http.StatusBadRequest, "Validate is failed: "+err.Error())
	}
	return nil
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	//log.Printf("Request Body: %v\n", string(reqBody))
	//log.Printf("Response Body: %v\n", string(resBody))
}

// Routes returns the initialized router
func Routes(e *echo.Echo) {
	e.Use(middleware.BodyDump(bodyDumpHandler))
	e.Validator = &Validator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Good morning !")
	})

	e.GET("/get_mybalance", func(c echo.Context) error {
		apiClient := bitflyer.New(os.Getenv("api_key"), os.Getenv("api_secret"))
		result, err := apiClient.GetBalance()
		if err != nil {
			return c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		return c.JSON(http.StatusOK, result)
	})

	e.GET("/get_ticker", func(c echo.Context) error {
		apiClient := bitflyer.New(os.Getenv("api_key"), os.Getenv("api_secret"))

		tickerChannel := make(chan bitflyer.Ticker)
		go apiClient.GetRealTimeTicker(os.Getenv("product_code"), tickerChannel)
		for ticker := range tickerChannel {
			fmt.Println(ticker)
			fmt.Println(ticker.GetMidPrice())
			fmt.Println(ticker.DateTime())
			fmt.Println(ticker.TruncateDateTime(time.Second))
			fmt.Println(ticker.TruncateDateTime(time.Minute))
			fmt.Println(ticker.TruncateDateTime(time.Hour))
		}
		return c.JSON(http.StatusOK, "ok")
	})

	e.GET("/stream_ingestion", func(c echo.Context) error {
		usecase.StreamIngestionData()
		return c.JSON(http.StatusOK, "ok")
	})

	e.GET("/view_chart_google", func(c echo.Context) error {

		productCode := c.Response().Header().Get("product_code")
		if productCode == "" {
			productCode = "BTC_JPY"
			// return c.JSON(http.StatusBadRequest,"No product_code param")
		}
		strLimit := c.Response().Header().Get("limit")
		limit, err := strconv.Atoi(strLimit)
		if strLimit == "" || err != nil || limit < 0 || limit > 1000 {
			limit = 1000
		}

		duration := c.Response().Header().Get("duration")
		if duration == "" {
			duration = "1s"
		}
		durationTime := config.Config.Durations[duration]

		df, _ := usecase.GetAllCandle(productCode, durationTime, limit)

		return c.Render(http.StatusOK, "google", df.Candles)
	})

	e.GET("/view_chart", func(c echo.Context) error {
		return c.Render(http.StatusOK, "chart", nil)
	})

	e.GET("/view_combo_chart", func(c echo.Context) error {
		return c.Render(http.StatusOK, "combo_chart", nil)
	})

	e.GET("/api/v1/candle", func(c echo.Context) error {

		productCode := c.QueryParam("product_code")
		if productCode == "" {
			productCode = "BTC_JPY"
		}

		strLimit := c.QueryParam("limit")
		limit, err := strconv.Atoi(strLimit)
		if strLimit == "" || err != nil || limit < 0 || limit > 1000 {
			limit = 1000
		}

		duration := c.QueryParam("duration")
		if duration == "" {
			duration = "1s"
		}
		durationTime := config.Config.Durations[duration]

		df, _ := usecase.GetAllCandle(productCode, durationTime, limit)

		return c.JSON(http.StatusOK, df.Candles)
	})

	// HTML render test
	e.GET("/html_test", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello", "HTML !!")
	})

	// HTML render test
	e.GET("/google", func(c echo.Context) error {
		return c.Render(http.StatusOK, "google", "")
	})

	// Migration Route
	e.GET("/api/v1/migrate", migrate)
	e.GET("/api/v1/seed", Seeds)
}

// =============================
//    MIGRATE
// =============================
func migrate(c echo.Context) error {
	_, err := config.DBMigrate()
	if err != nil {
		return c.String(http.StatusNotFound, "failed migrate")
	} else {
		return c.String(http.StatusOK, "success migrate")
	}
}

func Seeds(c echo.Context) error {
	_, err := config.Seeds()
	if err != nil {
		return c.String(http.StatusNotFound, "failed seed")
	} else {
		return c.String(http.StatusOK, "success seed")
	}
}
