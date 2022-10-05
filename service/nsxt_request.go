package service

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chmenegatti/get-nsxt-profile-id/configs"
)

func RequestNSXTApi(url, edge string) (r *http.Response, err error) {
	var (
		user = configs.GetEnvKeys(fmt.Sprintf("%s_USER", edge))
		pass = configs.GetEnvKeys(fmt.Sprintf("%s_PASS", edge))
	)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	d := http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal("Erro: ", err.Error())
	}

	req.SetBasicAuth(user, pass)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := d.Do(req)

	if err != nil {
		log.Fatal("Erro: ", err.Error())
	}

	return res, err
}
