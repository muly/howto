// go mod scenario with multiple package, and to create a vendor folder so that the dependencies are part repo

package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/muly/howto/golang/dependency-management/mod/scenario2/api"
)

func main() {
	h := mux.NewRouter()
	h.HandleFunc("/", api.Helloworld)
	http.ListenAndServe(":8080", h)
}

/* cd to project root folder and then go mod commands
cd $GOPATH/src/github.com/muly/howto/golang/dependency-management/mod/scenario2
go mod init github.com/muly/howto/golang/dependency-management/mod/scenario2
go mod vendor

-- to include new libraries
cd $GOPATH/src/github.com/muly/howto/golang/dependency-management/mod/scenario2
go mod vendor
*/
