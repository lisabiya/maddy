package main

import (
	"fmt"
	_ "github.com/foxcpp/maddy"
	maddycli "github.com/foxcpp/maddy/internal/cli"
	_ "github.com/foxcpp/maddy/internal/cli/ctl"
	"github.com/kardianos/service"
	"log"
	"os"
)

func main() {
	RunService()
}

func RunService() {
	//logger.ResetLog(true)
	//服务信息
	options := make(service.KeyValue)
	options["LimitNOFILE"] = 1000000
	svcConfig := &service.Config{
		Name:        "server_proxy",
		DisplayName: "server_proxy",
		Description: "基础代理程序",
		Option:      options,
	}
	prg := &daemonService{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Printf("%s 启动失败: %s", svcConfig.DisplayName, err)
		return
	}
	//监听指令
	if len(os.Args) > 1 {
		var err error
		verb := os.Args[1]
		switch verb {
		case "install":
			err = s.Install()
			if err != nil {
				fmt.Printf("Failed to install: %s\n", err)
				return
			}
			log.Printf("【ServerInstalled】 Service %s install.\n", svcConfig.DisplayName)
		case "uninstall":
			err = s.Uninstall()
			if err != nil {
				fmt.Printf("Failed to remove: %s\n", err)
				return
			}
			log.Printf("【ServerUninstall】 Service %s uninstall.\n", svcConfig.DisplayName)
		case "run":
			_ = s.Run()
		case "start":
			err = s.Start()
			//err = service.Control(s, os.Args[1])
			if err != nil {
				fmt.Printf("Failed to start: %s\n", err)
				return
			}
			log.Printf("【ServerStart】 Service %s started.\n", svcConfig.DisplayName)
		case "restart":
			err = s.Restart()
			if err != nil {
				fmt.Printf("Failed to restart: %s\n", err)
				return
			}
			log.Printf("【ServerRestart】 Service %s restarted.\n", svcConfig.DisplayName)
		case "stop":
			err = s.Stop()
			if err != nil {
				fmt.Printf("Failed to stop: %s\n", err)
				return
			}
			log.Printf("【ServerStopped】 Service %s stop.\n", svcConfig.DisplayName)
		case "status":
			sta, err := s.Status()
			if err != nil {
				fmt.Printf("Failed to status: %s\n", err)
				return
			}
			var status = "StatusUnknown"
			if sta == service.StatusRunning {
				status = "Running"
			} else if sta == service.StatusStopped {
				status = "Stopped"
			}
			log.Printf("【ServerStatus】 Service %s  status=%s \n", svcConfig.DisplayName, status)
		case "v":
			log.Printf("【ServerStatus】 Service %s  version=1.0 \n", svcConfig.DisplayName)
		}
	} else {
		log.Printf("【ServerRun】 服务 %s 启动\n", svcConfig.DisplayName)
		var err = s.Run()
		log.Printf("【ServerRun】 服务 %s 启动成功\n", svcConfig.DisplayName)
		if err != nil {
			fmt.Println("启动失败", err.Error())
		}
	}
}

type daemonService struct{}

func (svr *daemonService) Start(s service.Service) error {
	if service.Interactive() {
		log.Printf("【Start】 Running in terminal")
	} else {
		log.Printf("【Start】 Running under service manager")
	}
	log.Printf("【Start】 启动服务")
	go func() {
		err := svr.run()
		if err != nil {
			log.Printf("【Start】 启动服务失败")
		} else {
			log.Printf("【Start】 启动服务成功")
		}
	}()
	return nil
}

func (svr *daemonService) Stop(s service.Service) error {
	_, _ = s.Status()
	log.Printf("【Clean】监听程序")
	if service.Interactive() {
		os.Exit(0)
	}
	return nil
}

func (svr *daemonService) run() (err error) {
	maddycli.Run()
	return nil
}
