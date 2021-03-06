package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/naoina/toml"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
	"io/ioutil"
	"net/http"
	"text/template"
	"github.com/nirasan/gae-admin/bindata"
)

func init() {
	r := mux.NewRouter()

	// Management page
	admin := r.PathPrefix("/admin").Subrouter()
	admin.Path("/").HandlerFunc(adminIndexHandler)
	admin.Path("/toml").HandlerFunc(adminTomlImportHander)

	// API page
	api := r.PathPrefix("/api").Subrouter()
	api.Path("/").HandlerFunc(apiIndexHandler)

	http.Handle("/", r)
}

// Top of management pages
func adminIndexHandler(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)
	u := user.Current(ctx)

	t, err := template.New("").Parse(string(bindata.MustAsset("index.html")))
	if err != nil {
		panic(err)
	}

	loginUrl, _ := user.LoginURL(ctx, "/admin/")
	logoutUrl, _ := user.LogoutURL(ctx, "/admin/")

	data := struct {
		User      *user.User
		LoginUrl  string
		LogoutUrl string
	}{
		User:      u,
		LoginUrl:  loginUrl,
		LogoutUrl: logoutUrl,
	}

	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

// Toml file importer
func adminTomlImportHander(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	if r.Method == "POST" {

		// read uploaded file
		f, _, err := r.FormFile("uploadfile")
		if err != nil {
			panic(err)
		}
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		// struct of datadtore kind
		type Table1 struct {
			Id   int64
			Col1 string
			Col2 int64
		}

		// struct of toml file
		var input struct {
			Tables []Table1
		}

		// parse toml
		if err = toml.Unmarshal(buf, &input); err != nil {
			panic(err)
		}

		// import toml data
		for _, table := range input.Tables {
			key := datastore.NewKey(ctx, "Table1", "", table.Id, nil)
			if _, err := datastore.Put(ctx, key, &table); err != nil {
				panic(err)
			}
		}
	}

	// render form
	t, err := template.New("").Parse(string(bindata.MustAsset("toml.html")))
	if err != nil {
		panic(err)
	}

	t.Execute(w, nil)
}

// Api page (dummy)
func apiIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := struct{ Message string }{Message: "hello world"}
	json.NewEncoder(w).Encode(data)
}
