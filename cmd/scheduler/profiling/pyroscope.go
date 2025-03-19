// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package profiling

import (
	"fmt"
	"runtime"

	"github.com/grafana/pyroscope-go"
)

func EnablePyroscope(serverAddress, schedulerName string, mutexProfilerRate int, blockProfilerRate int) error {
	runtime.SetMutexProfileFraction(mutexProfilerRate)
	runtime.SetBlockProfileRate(blockProfilerRate)
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: schedulerName,

		// replace this with the address of pyroscope server
		ServerAddress: serverAddress,

		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,

		// you can provide static tags via a map:
		// Tags: tags,

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to start pyroscope go agnet: %v", err)
	}
	return nil
}
