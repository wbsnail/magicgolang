package server

import (
	"app/pkg/db"
	"app/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Options struct {
	Version string

	Addr     string
	DBDriver string
	DBArgs   string
	GinDebug bool
}

type App struct {
	Version string

	Addr string

	DB *gorm.DB

	// Gin router
	Router *gin.Engine
}

func SetupApp(options *Options) (*App, error) {
	if options == nil {
		options = &Options{}
	}

	if options.GinDebug {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}

	app := &App{}
	app.Version = options.Version
	app.Addr = options.Addr

	d, err := db.NewDB(options.DBDriver, options.DBArgs)
	if err != nil {
		return nil, errors.Wrap(err, "create db connection error")
	}
	app.DB = d

	r := gin.Default()

	r.GET("/", app.Index)
	r.Any("/healthz", app.Healthz)

	r.HandleMethodNotAllowed = true

	app.Router = r

	return app, nil
}

func (app *App) Serve() error {
	log.Infof("start serving on %s", app.Addr)
	return app.Router.Run(app.Addr)
}
