package main

import (
	"banking-api/internal/api"
	"banking-api/internal/service"
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	flag.Int("port", 5001, "server port.")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.AutomaticEnv()
	_ = viper.BindPFlags(pflag.CommandLine)
}

func main() {

	paymentService := service.NewPaymentService()
	refundService := service.NewRefundService()

	services := &service.Services{
		Payment: paymentService,
		Refund:  refundService,
	}

	server := api.NewServer(services)

	err := server.Run(fmt.Sprintf(":%d", viper.GetInt("port")))
	if err != nil {
		panic(err)
	}
}
