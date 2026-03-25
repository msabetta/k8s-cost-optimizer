package pkg

import (
	"fmt"
	"k8s-cost-optimizer/pkg/config"
	"k8s-cost-optimizer/pkg/logger"
	"k8s-cost-optimizer/pkg/utils"
)

func LoadPackage() {
	logger.Init()
	cfg, err := config.LoadConfig("configs/config.yaml", "configs/pricing.yaml")
	if err != nil {
		logger.Error.Println(err)
		return
	}
	logger.Info.Println("Config loaded:", cfg.Server.Port)

	val := utils.RoundFloat(1.234567, 2)
	fmt.Println("Rounded:", val)
}
