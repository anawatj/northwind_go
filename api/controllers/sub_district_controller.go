package controllers
import (
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/northwind_go/api/models"
	"github.com/northwind_go/api/responses"
)

func (server *Server) GetSubDistricts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rid, err := strconv.ParseUint(vars["rid"], 10, 32)
	pid, err := strconv.ParseUint(vars["pid"], 10, 32)
	did, err := strconv.ParseUint(vars["did"], 10, 32)
	subDistrict := models.SubDistrict{}

	subDistricts, err := subDistrict.FindAllSubDistrict(server.DB,uint32(rid),uint32(pid),uint32(did))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, subDistricts)
}