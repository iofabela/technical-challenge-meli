package main

import "github.com/iofabela/technical-challenge-meli/cmd/api/app"

// @title MELI CHALLENGE API
// @description This API Handle MELI Products.

// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

// @schemes http https
// @version 1.0.0
func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
