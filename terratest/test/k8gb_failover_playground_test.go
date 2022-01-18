// +build failover all

package test

/*
Copyright 2022 The k8gb Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Generated by GoLic, for more details see: https://github.com/AbsaOSS/golic
*/

import (
	"k8gbterratest/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestFailoverPlayground is equal to k8gb failover test running on local playground.
// see: https://github.com/k8gb-io/k8gb/blob/master/docs/local.md#failover
func TestFailoverPlayground(t *testing.T) {
	t.Parallel()
	const host = "playground-failover.cloud.example.com"
	const gslbPath = "../examples/failover-playground.yaml"
	const euGeoTag = "eu"
	const usGeoTag = "us"

	instanceEU, err := utils.NewWorkflow(t, "k3d-test-gslb1", 5053).
		WithGslb(gslbPath, host).
		WithTestApp(euGeoTag).
		Start()
	require.NoError(t, err)
	defer instanceEU.Kill()
	instanceUS, err := utils.NewWorkflow(t, "k3d-test-gslb2", 5054).
		WithGslb(gslbPath, host).
		WithTestApp(usGeoTag).
		Start()
	require.NoError(t, err)
	defer instanceUS.Kill()

	actAndAssert := func(test, geoTag string, localTargets []string) {
		// waiting for DNS sync
		err = instanceEU.WaitForExpected(localTargets)
		require.NoError(t, err)
		err = instanceUS.WaitForExpected(localTargets)
		require.NoError(t, err)
		// hit testApp from both clusters
		httpResult := instanceEU.HitTestApp()
		assert.Equal(t, geoTag, httpResult.Message)
		httpResult = instanceUS.HitTestApp()
		assert.Equal(t, geoTag, httpResult.Message)
	}

	t.Run("failover on two concurrent clusters with TestApp running", func(t *testing.T) {
		err = instanceEU.WaitForAppIsRunning()
		require.NoError(t, err)
		err = instanceUS.WaitForAppIsRunning()
		require.NoError(t, err)
	})

	euLocalTargets := instanceEU.GetLocalTargets()
	usLocalTargets := instanceUS.GetLocalTargets()

	t.Run("stop podinfo on eu cluster", func(t *testing.T) {
		instanceEU.StopTestApp()
		require.NoError(t, instanceEU.WaitForAppIsStopped())
		actAndAssert(t.Name(), usGeoTag, usLocalTargets)
	})

	t.Run("start podinfo again on eu cluster", func(t *testing.T) {
		instanceEU.StartTestApp()
		require.NoError(t, instanceEU.WaitForAppIsRunning())
		actAndAssert(t.Name(), euGeoTag, euLocalTargets)
	})
}
