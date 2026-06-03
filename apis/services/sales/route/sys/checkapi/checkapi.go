package checkapi

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

 
 
 func liveness(ctx context.Context ,w http.ResponseWriter, r *http.Request)error{
       status:=struct{
		status string
	   }{
		status:"ok",
	   }
	   return web.Respond(ctx,w,status,http.StatusOK)
}

 func readiness(ctx context.Context ,w http.ResponseWriter, r *http.Request)error{
       status:=struct{
		status string
	   }{
		status:"ok",
	   }
	   return web.Respond(ctx,w,status,http.StatusOK)
}
