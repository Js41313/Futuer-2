package cmd

import (
	"context"
	"fmt"

	"github.com/Js41313/Futuer-2/pkg/constant"

	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Js41313/Futuer-2/initialize"
	"github.com/Js41313/Futuer-2/internal"
	"github.com/Js41313/Futuer-2/internal/config"
	"github.com/Js41313/Futuer-2/internal/svc"
	"github.com/Js41313/Futuer-2/pkg/conf"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/orm"
	"github.com/Js41313/Futuer-2/pkg/service"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/queue"
	"github.com/Js41313/Futuer-2/scheduler"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	startCmd.Flags().StringVar(&startConfigPath, "config", "etc/ppanel.yaml", "ppanel.yaml directory to read from")
}

var (
	startConfigPath string
)

var startCmd = &cobra.Command{
	Use:   "run",
	Short: "start PPanel",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[PPanel version] v" + fmt.Sprintf("%s(%s)", constant.Version, constant.BuildTime))
		run()
	},
}

func run() {
	services := getServers()
	defer services.Stop()
	go services.Start()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
}
func getServers() *service.Group {
	var c config.Config

	// check config file is exist
	if _, err := os.Stat(startConfigPath); os.IsNotExist(err) {
		// check directory is existed
		if _, err := os.Stat("etc"); os.IsNotExist(err) {
			logger.Errorf("Directory %s does not exist. Creating it...\n", "etc")
			if err = os.MkdirAll("etc", os.ModePerm); err != nil {
				log.Fatalf("Please create the directory %s and place the configuration file %s in it.\n", "etc", startConfigPath)
			}
		}
		// create new config file
		if _, err := os.Create(startConfigPath); err != nil {
			logger.Errorf("Please create the configuration file %s in the directory %s.\n", startConfigPath, "etc")
			panic(fmt.Sprintf("Please create the configuration file %s in the directory %s.\n", startConfigPath, "etc"))
		}
	}
	// check config file is empty, if empty, start init web server
	if initConfig(&c) {
		status, server := initialize.Config(startConfigPath)
		<-status
		if err := server.Shutdown(context.TODO()); err != nil {
			log.Printf("Init Server Shutdown: %s\n", err.Error())
		}
	}
	conf.MustLoad(startConfigPath, &c)
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	// init logger
	if err := logger.SetUp(c.Logger); err != nil {
		logger.Errorf("Logger setup failed: %v", err.Error())
	}

	// init service context
	ctx := svc.NewServiceContext(c)
	services := service.NewServiceGroup()
	services.Add(internal.NewService(ctx))
	services.Add(queue.NewService(ctx))
	services.Add(scheduler.NewService(ctx))
	return services
}

func initConfig(c *config.Config) bool {
	// load config
	conf.MustLoad(startConfigPath, c)
	//  check custom config
	if startConfigPath != "etc/ppanel.yaml" && c.MySQL.Addr == "" {
		return true
	}
	// check access secret
	if c.JwtAuth.AccessSecret == "" && startConfigPath == "etc/ppanel.yaml" {
		c.JwtAuth.AccessSecret = uuid.New().String()
		// Get environment variables
		dsn := os.Getenv("PPANEL_DB")
		if dsn == "" {
			return true
		}
		cfg := orm.ParseDSN(dsn)
		if cfg == nil {
			return true
		} else {
			c.MySQL = *cfg
		}

		// Get environment variables
		uri := os.Getenv("PPANEL_REDIS")
		if uri == "" {
			return true
		}
		addr, pass, db, err := tool.ParseRedisURI(uri)
		if err != nil {
			return true
		} else {
			c.Redis.Host = addr
			c.Redis.Pass = pass
			c.Redis.DB = db
		}
		// save yaml file
		newConfig := config.File{
			Host:    c.Host,
			Port:    c.Port,
			Debug:   c.Debug,
			JwtAuth: c.JwtAuth,
			Logger:  c.Logger,
			MySQL:   c.MySQL,
			Redis:   c.Redis,
		}
		fileData, err := yaml.Marshal(newConfig)
		if err != nil {
			panic(err.Error())
		}
		// write to file
		if err := os.WriteFile(startConfigPath, fileData, 0644); err != nil {
			panic(err.Error())
		}
	}
	return false
}
