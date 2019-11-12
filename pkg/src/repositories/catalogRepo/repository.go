package catalogRepo

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"loyalty_go/pkg/src/domains/userProfile"
	"net/http"
)
const(
	url = "http://localhost:3002/v1/articles/"
)

type catalogRepo struct{
	ctx context.Context
}

func NewRepo() catalogRepo{
	return catalogRepo{}
}

func (repo catalogRepo) GetArticle(artID string) userProfile.CatalogResponse{
	art := userProfile.CatalogResponse{}
	req, err := http.NewRequest("GET", url + artID, nil)
	if err!= nil {
		log.Panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	payload, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(payload, &art)
	if err != nil {
		panic(err)
	}
	return art
}