package controllers
import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/northwind_go/api/auth"
	"github.com/northwind_go/api/models"
	"github.com/northwind_go/api/responses"
)
func (server *Server) CreateRegion(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	region := models.Region{}
	err = json.Unmarshal(body, &region)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	user := models.User{}
	userGotten,err := user.FindUserByID(server.DB,uint32(uid))
	if userGotten==nil && err != nil {
		responses.ERROR(w,http.StatusUnauthorized,errors.New(http.StatusText(http.StatusUnauthorized)))
		return;
	}

	regionCreated,err := region.SaveRegion(server.DB)
	
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, regionCreated.ID))
	responses.JSON(w, http.StatusCreated, regionCreated)
}

func (server *Server) GetRegions(w http.ResponseWriter, r *http.Request) {

	region := models.Region{}

	regions, err := region.FindAllRegions(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, regions)
}

