package service

import (
	"CyberAssetMapper/src/global"
	"CyberAssetMapper/src/service/dto"
	"fmt"
)

var hostService *HostService

type HostService struct {
	BaseService
}

func NewHostService() *HostService {
	if hostService == nil {
		hostService = &HostService{}
	}
	return hostService
}

func (m *HostService) Shutdown(iShutdownHostDTO dto.ShutdownHostDTO) error {
	stHostIP := iShutdownHostDTO.HostIP
	global.Logger.Info(fmt.Sprintf("Shutdown host: %s", stHostIP))
	var errResult error
	return errResult
}
