package subsystem

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/filanov/bm-inventory/client"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kelseyhightower/envconfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB
var bmclient *client.BMInventory

var Options struct {
	DBHost        string `envconfig:"DB_HOST"`
	DBPort        string `envconfig:"DB_PORT"`
	InventoryHost string `envconfig:"INVENTORY"`
}

func init() {
	var err error
	err = envconfig.Process("subsystem", &Options)
	if err != nil {
		log.Fatal(err.Error())
	}

	bmclient = client.New(client.Config{
		URL: &url.URL{
			Scheme: client.DefaultSchemes[0],
			Host:   Options.InventoryHost,
			Path:   client.DefaultBasePath,
		},
	})

	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=postgresadmin dbname=postgresdb password=admin123 sslmode=disable",
		Options.DBHost, Options.DBPort))
	if err != nil {
		logrus.Fatal("Fail to connect to DB, ", err)
	}
}

func TestSubsystem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Subsystem Suite")
}