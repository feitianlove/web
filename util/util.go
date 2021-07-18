package util

import (
	"encoding/base64"
	"fmt"
	"github.com/feitianlove/golib/common/utils"
	"sync/atomic"
	"time"
)

var RequestCounter uint64 = 0

// AllocateRequestId allocate
func AllocateRequestId() (string, error) {
	localIp := utils.GetLocalIP()
	now := time.Now().UnixNano()
	currentRequestCounter := atomic.AddUint64(&RequestCounter, 1)
	requestIdStr := fmt.Sprintf("%x_%x_%x", now, localIp, currentRequestCounter)
	return base64.StdEncoding.EncodeToString([]byte(requestIdStr)), nil
}
