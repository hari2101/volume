# volume

This microservice will take [][]string{} as input and give the calculated []string.
The output is the connections between source and destination.

There are some assumptions to this api.
1. In the input for every source there will be one and only one destination
2. Except for final source and final destination, there will be ultimate connection.
3. Input is always [][]string{}

Usage:
1. Clone the repo and run 'go run main.go' from root. Service will be started in 8080 port.
2. From postman make api call with body as [][]string{}

example(cURL):
curl --location --request GET 'http://localhost:8080/calculate' \
--header 'Content-Type: application/json' \
--data-raw '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'

output: ["SFO","EWR"]