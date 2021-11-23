package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetVarToInt(writer http.ResponseWriter, vars map[string]string, name string) (value int, err error) {
	value, err = strconv.Atoi(vars[name])
	if err != nil {
		message := fmt.Sprintf("'%s' is invalid", name)
		log.Println("ERROR:", message)
		http.Error(writer, message, http.StatusBadRequest)
		return 0, err
	}

	return value, nil
}

func ReturnJson(writer http.ResponseWriter, status int, model interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(model)
}
