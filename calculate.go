package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Calculate(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	connections := [][]string{}
	err = json.Unmarshal(reqBody, &connections)
	if err != nil {
		panic(err)
	}
	// flights map maintains source as key and source to connected destination as value
	flights := make(map[string][]string)

	// read all the input connections and
	// check if the destination of the current connection is already visited
	// if visited then update the current flight destination to visited connection destination
	// and delete the visited connection. This will reduce the map and array size.
	for _, connection := range connections {
		flights[connection[0]] = connection
		if _, ok := flights[connection[1]]; ok {
			flights[connection[0]] = []string{
				flights[connection[0]][0],
				flights[connection[1]][len(flights[connection[1]])-1],
			}
			delete(flights, connection[1])
		}
	}
	result := make([]string, 2)
	//search for the source for which end destination is in flights map
	//add the source's starting element and destination's last element to the result.
	for key, flight := range flights {
		if _, ok := flights[flight[len(flight)-1]]; ok {
			result = []string{flights[key][0], flights[flights[key][1]][len(flights[key])-1]}
		}
		if len(flights) == 1 {
			result = flight
		}
	}
	json.NewEncoder(w).Encode(result)
}
