//
// Copyright (c) 2021 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//

package cmd

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/devfile/registry-support/tests/integration/pkg/config"

	_ "github.com/devfile/registry-support/tests/integration/pkg/tests"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// Integration/e2e test logic based on https://github.com/devfile/devworkspace-operator/tree/main/test/e2e

//Create Constant file
const (
	testResultsDirectory = "/tmp/artifacts"
	jUnitOutputFilename  = "junit-devfileregistry-operator.xml"
)

//SynchronizedBeforeSuite blocks is executed before run all test suites
var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	fmt.Println("Starting to setup objects before run ginkgo suite")
	registry := os.Getenv("REGISTRY")
	if registry == "" {
		registry = "https://registry.devfile.io"
	}
	config.Registry = registry
	config.RegistryList = registry + "," + "https://registry.stage.devfile.io"
	os.Setenv("REGISTRY_LIST", config.RegistryList)

	config.IsTestRegistry, _ = strconv.ParseBool(os.Getenv("IS_TEST_REGISTRY"))

	return nil
}, func(data []byte) {})

func TestDevfileRegistryController(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)

	fmt.Println("Running Devfile Registry integration tests...")
	ginkgo.RunSpecs(t, "Devfile Registry Tests")
}
