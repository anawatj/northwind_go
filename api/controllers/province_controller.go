package controllers
import (
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/northwind_go/api/models"
	"github.com/northwind_go/api/responses"
)

func (server *Server) GetProvinces(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rid, err := strconv.ParseUint(vars["rid"], 10, 32)
	province := models.Province{}

	provinces, err := province.FindAllProvince(server.DB,uint32(rid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, provinces)
}
