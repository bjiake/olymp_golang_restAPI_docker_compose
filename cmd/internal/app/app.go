package app

import (
	requestsToMongoDB "go_project/cmd/internal/database/requestToMongoDB"
	"go_project/cmd/internal/transport"
)

func Start() {
	transport.Transport()
	requestsToMongoDB.CloseConnection()
}
