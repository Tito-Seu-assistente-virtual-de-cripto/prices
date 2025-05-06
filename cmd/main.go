package main

import (
	"github.com/Tito-Seu-assistente-virtual-de-cripto/prices/internal/jobs/token_price_job"
)

func main() {
	job := token_price_job.NewTokensPriceJob()
	job.Start()
}
