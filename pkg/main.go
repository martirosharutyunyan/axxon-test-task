package start

import (
	"github.com/martirosharutyunyan/axxon-test-task/pkg/config"
	"github.com/martirosharutyunyan/axxon-test-task/pkg/database"
	gooseDB "github.com/martirosharutyunyan/axxon-test-task/pkg/database/migration-up"
	"github.com/martirosharutyunyan/axxon-test-task/pkg/modules/controllers"
	"log"
)

func Application() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	config.Load("../../.env")
	database.ConnectDB()
	if config.GetEnv() == "development" {
		gooseDB.MigrationsUp()
	}
	controllers.RunServer()
}
