// (c) Copyright 2019 Hewlett Packard Enterprise Development LP

package chapi2

import (
	"fmt"
	"runtime"

	log "github.com/hpe-storage/common-host-libs/logger"
	"github.com/hpe-storage/common-host-libs/model"
	"github.com/hpe-storage/common-host-libs/util"
)

// GetSocketName returns unix socket name (per process)
//Added this to avoid compile error in vscode on mac
func GetSocketName() string {
	log.Errorln("GetSocketName: Not supported on", runtime.GOOS)
	return ""
}

//RunNimbled does nothing
func RunNimbled(c chan error) {
}

// CheckFsCreationInProgress checks if FS creation/formatting is in progress on the device
// Need this to just compile code on Mac
func CheckFsCreationInProgress(device chapi2.Device) (inProgress bool, err error) {
	return false, fmt.Errorf("FS progress check is not implemented for Mac")
}

// GetSerialNumber returns the volume serial number as per platform format
func GetSerialNumber(serialNumber string) string {
	return ""
}
