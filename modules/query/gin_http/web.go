package ginHttp

import (
	"net/http"
	"time"

	cmodel "github.com/Cepave/open-falcon-backend/common/model"

	"github.com/Cepave/open-falcon-backend/modules/query/gin_http/computeFunc"
	grahttp "github.com/Cepave/open-falcon-backend/modules/query/gin_http/grafana"
	"github.com/Cepave/open-falcon-backend/modules/query/gin_http/openFalcon"
	"github.com/gin-gonic/gin"
)

type QueryInput struct {
	StartTs       time.Time
	EndTs         time.Time
	ComputeMethod string
	Endpoint      string
	Counter       string
}

//this function will generate query string obj for QueryRRDtool
func getq(q QueryInput) cmodel.GraphQueryParam {
	request := cmodel.GraphQueryParam{
		Start:     q.StartTs.Unix(),
		End:       q.EndTs.Unix(),
		ConsolFun: q.ComputeMethod,
		Endpoint:  q.Endpoint,
		Counter:   q.Counter,
	}
	return request
}

//accept cross domain request
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ConfigWeb() {

	computeHandler := gin.Default()
	computeHandler.Use(CORSMiddleware())
	compute := computeHandler.Group("/lambda/func")
	compute.GET("/compute", computeFunc.Compute)
	compute.GET("/funcations", computeFunc.GetAvaibleFun)
	compute.GET("/smapledata", computeFunc.GetTestData)
	http.Handle("/lambda/func/", computeHandler)

	openfalconHandler := gin.Default()
	openfalconHandler.Use(CORSMiddleware())
	openfalcon := openfalconHandler.Group("/lambda/owl")
	openfalcon.GET("/endpoints", openFalcon.GetEndpoints)
	openfalcon.GET("/queryrrd", openFalcon.QueryData)
	http.Handle("/lambda/owl/", openfalconHandler)

	grafanaHandler := gin.Default()
	grafanaHandler.Use(CORSMiddleware())
	grafana := grafanaHandler.Group("/lambda/api/grafana")
	grafana.GET("/", grahttp.GrafanaMain)
	grafana.GET("/metrics/find", grahttp.GrafanaMain)
	grafana.POST("/render", grahttp.GetQueryTargets)
	http.Handle("/lambda/api/grafana/", grafanaHandler)
}
