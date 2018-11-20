package main

import (
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
	"github.com/samze/helmbroker/pkg/broker"
)

func main() {
	logger := lager.NewLogger("helm-broker")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.INFO))

	creds := brokerapi.BrokerCredentials{
		Username: "admin",
		Password: "password",
	}

	broker := broker.HelmBroker{}

	handler := brokerapi.New(&broker, logger, creds)

	s := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	logger.Info("starting web server")

	logger.Fatal("web server stopped", s.ListenAndServe())
}
