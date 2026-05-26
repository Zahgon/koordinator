/*
Copyright 2022 The Koordinator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package routes

import (
	"html/template"
	"net/http"
	"sync"

	"k8s.io/apiserver/pkg/server/mux"
)

var (
	lock            = &sync.RWMutex{}
	registeredFlags = map[string]debugFlag{}
)

// DebugFlags adds handlers for flags under /debug/flags.
type DebugFlags struct {
	c *mux.PathRecorderMux
}

func NewDebugFlags(c *mux.PathRecorderMux) DebugFlags {
	_ = "STUB: not implemented"
	return *new(DebugFlags)
}

// Install registers the APIServer's flags handler.
func (f DebugFlags) Install(flag string, handler func(http.ResponseWriter, *http.Request)) {
	_ = "STUB: not implemented"
	return
}

// Index responds with the `/debug/flags` request.
// For example, "/debug/flags/v" serves the "--v" flag.
// Index responds to a request for "/debug/flags/" with an HTML page
// listing the available flags.
func (f DebugFlags) Index(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

var indexTmpl = template.Must(template.New("index").Parse(`<html>
<head>
<title>/debug/flags/</title>
</head>
<body>
/debug/flags/<br>
<br>
flags:<br>
<table>
{{range .}}
<tr>{{.Flag}}<br>
{{end}}
</table>
<br>
full flags configurable<br>
</body>
</html>
`))

type debugFlag struct {
	Flag string
}

func (f DebugFlags) addFlag(flag string) { _ = "STUB: not implemented"; return }

// StringFlagSetterFunc is a func used for setting string type flag.
type StringFlagSetterFunc func(string) (string, error)

// StringFlagPutHandler wraps an http Handler to set string type flag.
func StringFlagPutHandler(setter StringFlagSetterFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// writePlainText renders a simple string response.
func writePlainText(statusCode int, text string, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}
