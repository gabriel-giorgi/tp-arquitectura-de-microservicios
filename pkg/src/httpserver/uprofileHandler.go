package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"loyalty_go/pkg/src/domains/userProfile"
	"net/http"
	"strings"
)


// Handler contains the dependencies needed for using this controller
type ProfileHandler struct {
	service userProfile.Service
}


func NewProfileHandler(p userProfile.Service) *ProfileHandler {
	return &ProfileHandler{
		service: p,
	}
}
/**
 * @apiDefine ParamValidationErrors
 *
 * @apiErrorExample 400 Bad Request
 *     HTTP/1.1 400 Bad Request
 *     {
 *        "messages" : [
 *          {
 *            "path" : "{Nombre de la propiedad}",
 *            "message" : "{Motivo del error}"
 *          },
 *          ...
 *       ]
 *     }
 */
/**
 * @apiDefine OtherErrors
 *
 * @apiErrorExample 500 Server Error
 *     HTTP/1.1 500 Internal Server Error
 *     {
 *        "error" : "Not Found"
 *     }
 *
 */
func (p ProfileHandler) newProfile(c *gin.Context) {

	uProfile := &userProfile.Profile{}
	err := c.BindJSON(uProfile)
	if err!= nil {
		log.Panic(err)
	}

	profileResponse := p.service.CreateNewProfile(uProfile.UserID)
	jsonResponse, err :=json.Marshal(&profileResponse)
	if err != nil {
		log.Panic("Error marshalling the profile " ,err)
	}
	fmt.Println(jsonResponse)
	c.JSON(200, jsonResponse)

}

func (p ProfileHandler) getAPIDoc(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func (p ProfileHandler) getProfile(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, " ")
	if len(splitToken) != 2 {
		log.Println("Bearer token incorrect format")
	}
	reqToken = strings.TrimSpace(splitToken[1])

	url := "http://localhost:3000/v1/users/current"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + reqToken

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	if err!= nil {
		log.Println(err)
	}
	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	user := &userProfile.User{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body , &user)
	log.Println(user.ID)

	profileResponse := p.service.GetProfile(user.ID)
	c.JSON(200, profileResponse)
}