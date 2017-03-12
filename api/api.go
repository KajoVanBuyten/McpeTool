package api

import (
	"encoding/hex"
	"log"
	"net/http"

	"github.com/midnightfreddie/McpeTool/world"
)

var apiVersion = "1.0"

// TODO: this moved to world/keys.go; remove its use from this package and delete
// convertKey takes a byte array and returns a string if all characters are printable (else "")  hex-string-encoded versions of key
func convertKey(k []byte) (stringKey, hexKey string) {
	allAscii := true
	for i := range k {
		if k[i] < 0x20 || k[i] > 0x7e {
			allAscii = false
		}
	}
	if allAscii {
		stringKey = string(k[:])
	}
	hexKey = hex.EncodeToString(k)
	return
}

// Server is the http REST API server
func Server(world *world.World) error {

	// http handler functions defined in other files in this package
	dbApi(world, "/api/v1/db/")
	worldApi(world, "/api/v1/world/")
	playerApi(world, "/api/v1/player/")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
	return nil
}
