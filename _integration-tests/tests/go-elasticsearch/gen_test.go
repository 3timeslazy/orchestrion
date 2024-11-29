// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.
//
// Code generated by 'go generate'; DO NOT EDIT.

//go:build integration && (linux || !githubci) && !windows

package go_elasticsearch

import (
	"testing"

	"datadoghq.dev/orchestrion/_integration-tests/utils"
)

func TestV6(t *testing.T) {
	utils.RunTest(t, new(TestCaseV6))
}

func TestV7(t *testing.T) {
	utils.RunTest(t, new(TestCaseV7))
}

func TestV8(t *testing.T) {
	utils.RunTest(t, new(TestCaseV8))
}
