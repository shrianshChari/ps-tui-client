package utils

import "strings"

// Returns the Websocket URL for the following servers:
// - main (play.pokemonshowdown.com)
// - local (server hosted locally)
// - smogtours
// If the server is not one of these, then it returns an empty string.
func SelectServer(server string) string {
	server = strings.ToLower(server)
	if len(server) == 0 || server == "main" || server == "play" {
		return "sim3.psim.us"
	} else if server == "local" || server == "localhost" {
		return "localhost:8000"
	} else if server == "smogtours" {
		return "sim3.psim.us:8001"
	}
	// If not part of these already-known servers, then
	// use the WebSocket URL of the intended server
	return ""
}
