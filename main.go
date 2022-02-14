package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CorporateCampaignBaseResponse struct {
	TipoCliente                string `json:"tipoCliente" `
	SegmentoPD                 string `json:"segmentoPD"`
	SegmentoPDML               string `json:"segmentoPDML"`
	Microsegmento              string `json:"microsegmento"`
	MLPuntajeClienteNoConocido string `json:"mlPuntajeClienteNoConocido"`
	TipoClienteOferta          string `json:"tipoClienteOferta"`
	MontoCupo                  string `json:"montoCupo"`
	TipoEstrategia             string `json:"tipoEstrategia"`
	TipoProducto               string `json:"tipoProducto"`
	CanalDeVenta               string `json:"canalVenta"`
	Libre1                     string `json:"libre1"`
	Libre2                     string `json:"libre2"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", createEvent)
	log.Fatal(http.ListenAndServe(":8087", router))
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	ev := CorporateCampaignBaseResponse{
		TipoCliente:                "PJ",
		SegmentoPD:                 "test",
		SegmentoPDML:               "test",
		Microsegmento:              "test",
		MLPuntajeClienteNoConocido: "test",
		TipoClienteOferta:          "test",
		MontoCupo:                  "test",
		TipoEstrategia:             "test",
		TipoProducto:               "test",
		CanalDeVenta:               "test",
		Libre1:                     "test",
		Libre2:                     "test",
	}
	json.NewEncoder(w).Encode(ev)

}
