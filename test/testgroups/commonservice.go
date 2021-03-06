//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package testgroups

import (
	"testing"

	"github.com/operator-framework/operator-sdk/pkg/test"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	"github.com/operator-framework/operator-sdk/pkg/test/e2eutil"

	"github.com/IBM/operand-deployment-lifecycle-manager/test/config"
	"github.com/IBM/operand-deployment-lifecycle-manager/test/helpers"
	"github.com/IBM/operand-deployment-lifecycle-manager/test/testsuits"
)

// ODLM is the test group for testing Operand Deployment Lifecycle Manager
func ODLM(t *testing.T) {
	t.Parallel()
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup()

	err := deployOperator(t, ctx)
	helpers.AssertNoError(t, err)

	t.Run("OperandRequest Create", testsuits.OperandRequestCreate)
	t.Run("OperandRequest Create Update", testsuits.OperandRequestCreateUpdate)
	t.Run("OperandConfig Update", testsuits.OperandConfigUpdate)
	t.Run("OperandRegistry Create Update", testsuits.OperandRegistryUpdate)
	t.Run("OperandRequest Create Delete", testsuits.OperandRequestCreateDelete)

}

func deployOperator(t *testing.T, ctx *test.TestCtx) error {
	err := ctx.InitializeClusterResources(
		&test.CleanupOptions{
			TestContext:   ctx,
			Timeout:       config.CleanupTimeout,
			RetryInterval: config.CleanupRetry,
		},
	)
	if err != nil {
		t.Fatalf("failed to initialize cluster resources")
		return err
	}

	namespace, err := ctx.GetNamespace()
	if err != nil {
		t.Fatalf("failed to get the namespace")
		return err
	}

	return e2eutil.WaitForOperatorDeployment(
		t,
		test.Global.KubeClient,
		namespace,
		config.TestOperatorName,
		1,
		config.APIRetry,
		config.APITimeout,
	)
}
