package handeller

import (
	controller "Univercity/controllers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Configuration struct
type Configuration struct {
	Port             string // port no
	ConnectionString string // connection string
	Database         string // database name
	Collection       string // collection
}

/*ReadConfig Reading the configs from  db.properties
 */
func ReadConfig() Configuration {
	var configfile = "config.properties"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Configuration
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}

func HandleRequests() {

	config := ReadConfig()
	var port = ":" + config.Port

	Router := mux.NewRouter()
	//corsObj := handlers.AllowedOrigins([]string{"*"})
	Router.HandleFunc("/", controller.Homepage).Methods(http.MethodGet)
	Router.HandleFunc("/students", controller.Poststudent).Methods(http.MethodPost)
	Router.HandleFunc("/students", controller.GetStudent).Methods(http.MethodGet)
	// Router.HandleFunc("/students", controller.GetStudentbyName).Methods(http.MethodGet)
	Router.HandleFunc("/students/{id}", controller.GetStudentbyid).Methods(http.MethodGet)
	Router.HandleFunc("/students/{id}", controller.Deletestudent).Methods(http.MethodDelete)
	Router.HandleFunc("/students/{id}", controller.Updatestudent).Methods(http.MethodPut)
	Router.HandleFunc("/students/{id}", controller.Updatestudent_patch).Methods(http.MethodPatch)
	//log.Fatal(http.ListenAndServe(":8091", Router))
	fmt.Printf("application listening port%s\n", port)
	http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(Router))

}
