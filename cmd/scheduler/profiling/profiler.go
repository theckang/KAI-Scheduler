// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package profiling

import (
	"fmt"
	"os"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func RegisterProfiler(port string) {
	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	pprof.Register(router)
	err = router.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
