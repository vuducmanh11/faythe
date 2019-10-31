package main

import (
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/vCloud-DFTBA/faythe/config"
	"github.com/vCloud-DFTBA/faythe/handlers/basic"
	"github.com/vCloud-DFTBA/faythe/handlers/openstack"
	"github.com/vCloud-DFTBA/faythe/handlers/stackstorm"
)

func newRouter(logger *log.Logger) *mux.Router {
	router := mux.NewRouter()

	conf := config.Get()
	// Create nextRequestID
	nextRequestID := func() string {
		return strconv.FormatInt(time.Now().UnixNano(), 36)
	}

	r, _ := regexp.Compile(conf.ServerConfig.RemoteHostPattern)

	// Init middleware
	mw := &middleware{
		logger:        logger,
		nextRequestID: nextRequestID,
		regexp:        r,
	}

	// Routing
	router.Handle("/", basic.Index()).Methods("GET")
	router.Handle("/healthz", basic.Healthz(&healthy)).Methods("GET")
	router.Handle("/stackstorm/{st-name}/{st-rule}", stackstorm.TriggerSt2Rule()).
		Methods("POST")
	router.Handle("/stackstorm/alertmanager/{st-name}/{st-rule}", stackstorm.TriggerSt2RuleAM()).
		Methods("POST")
	router.Handle("/openstack/autoscaling/{ops-name}", openstack.AutoScaling()).
		Methods("POST")

	// Appends a Middlewarefunc to the chain.
	router.Use(mw.tracing, mw.logging, mw.restrictingDomain, mw.authenticating)

	return router
}
