package checkapi

import (
	"github.com/ardanlabs/service/foundation/web"
)

func Routes(app *web.App) {
	// Register both forms so probes work regardless of mux pattern mode.

	app.HandleFunc("GET /readiness", readiness)
	app.HandleFunc("GET /liveness", liveness)
}