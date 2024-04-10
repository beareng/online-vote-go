package backend

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

var ApiInternalError = errors.New("Internal Server Error")

func apiInternalError() (int, interface{}, error) {
	return 500, nil, ApiInternalError
}

func apiInvalidRequest(err error) (int, interface{}, error) {
	return 400, nil, fmt.Errorf("Invalid request: %v", err)
}

func apiUnauthorizedRequest(err error) (int, interface{}, error) {
	return 401, nil, fmt.Errorf("Unauthorized request: %v", err)
}

func apiNotFound(err error) (int, interface{}, error) {
	return 404, nil, fmt.Errorf("Not found: %v", err)
}

type auth struct {
	Token string `json:",omitempty"`
	Name  string `json:",omitempty"`
}

func (etx ElectionsTx) findAuth(a auth) (*User, error) {
	if 0 != len(a.Token) {
		return etx.FindUserByToken(a.Token)
	} else {
		return nil, nil
	}
}

func (etx ElectionsTx) findOrCreateAuth(a auth) (*User, error) {
	if 0 != len(a.Token) {
		return etx.FindUserByToken(a.Token)
	} else {
		return etx.FindOrCreateUnregisteredUser(a.Name)
	}
}

type voteReq struct {
	Auth auth
}

type resultsReq struct {
	Auth auth
}

func (edb ElectionsDb) ApiVoteHandler() http.HandlerFunc {
	return makeApiHandler(edb.apiHandleVote)
}

func (edb ElectionsDb) ApiResultsHandler() http.HandlerFunc {
	return makeApiHandler(edb.apiHandleResults)
}

func (edb ElectionsDb) ApiListingHandler() http.HandlerFunc {
	return makeApiHandler(edb.apiHandleListing)
}

func (edb ElectionsDb) apiHandleVote(query url.Values, jsonBody []byte) (int, interface{}, error) {
	//TODO: TBD
}

func (edb ElectionsDb) apiHandleResults(query url.Values, jsonBody []byte) (int, interface{}, error) {
	//TODO: TBD
}

func (edb ElectionsDb) apiHandleListing(query url.Values, jsonBody []byte) (int, interface{}, error) {
	//TODO: TBD
}

func makeApiHandler(api func(query url.Values, jsonBody []byte) (int, interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		jsonBody, err := io.ReadAll(io.Reader(req.Body))
		if nil != err {
			log.Printf("Couldn't read json request body: %v", err)
			http.Error(w, "400 Bad Request", 400)
			return
		} else if code, result, err := api(req.URL.Query(), jsonBody); nil != err {
			log.Printf("Request[%+q] failed: %d %+q", req.URL.EscapedPath(), code, err)
			http.Error(w, err.Error(), code)
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(code)
			w.Write(types.JsonMustEncode(result))
		}
	}
}

func (edb ElectionsDb) BindServeMux(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/vote", edb.ApiVoteHandler())
	mux.HandleFunc(prefix+"/result", edb.ApiResultsHandler())
	mux.HandleFunc(prefix+"/listPhoto", edb.ApiListingHandler())
}
