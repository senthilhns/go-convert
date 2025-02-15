package json

import (
	"errors"
	"fmt"
	harness "github.com/drone/spec/dist/go"
)

const (
	ArtifactoryRtCommandsPluginImage = "plugins/artifactory"
	InputPlaceHolder                 = "<+input>"
	MvnTool                          = "mvn"
	GradleTool                       = "gradle"
)

func ConvertArtifactoryRtCommand(stepType string, node Node, variables map[string]string) *harness.Step {
	switch stepType {
	case "rtMavenRun":
		return convertRtMavenRun(node, variables)
	case "rtGradleRun":
		return convertRtGradleRun(node, variables)
	case "publishBuildInfo":
		return convertPublishBuildInfo(node, variables)
	case "rtPromote":
		return convertRtPromote(node, variables)
	case "xrayScan":
		return convertXrayScan(node, variables)
	}
	return nil
}

var ConvertRtMavenRunParamMapperList = []JenkinsToDroneParamMapper{
	{"pom", "source", StringType, nil},
}

func convertRtMavenRun(node Node, variables map[string]string) *harness.Step {
	step := GetStepWithProperties(&node, ConvertRtMavenRunParamMapperList, ArtifactoryRtCommandsPluginImage)
	if step == nil {
		fmt.Println("Error: failed to convert rtMavenRun")
		return nil
	}
	tmpStepPlugin, ok := step.Spec.(*harness.StepPlugin)
	if !ok {
		fmt.Println("Error: failed to convert rtMavenRun")
		return nil
	}
	if tmpStepPlugin.With == nil {
		tmpStepPlugin.With = map[string]interface{}{}
	}
	tmpStepPlugin.With["build_tool"] = MvnTool
	attributesList := []string{"url", "username", "password", "access_token",
		"build_name", "build_number", "resolver_id", "deployer_id", "resolve_release_repo", "resolve_snapshot_repo"}
	err := SetRtCommandAttributesToInputPlaceHolder(tmpStepPlugin, attributesList)
	if err != nil {
		fmt.Println("Error: failed to set attributes to input placeholder")
		return nil
	}
	return step
}

func convertRtGradleRun(node Node, variables map[string]string) *harness.Step {
	step := GetStepWithProperties(&node, nil, ArtifactoryRtCommandsPluginImage)
	if step == nil {
		fmt.Println("Error: failed to convert rtGradleRun")
		return nil
	}
	tmpStepPlugin, ok := step.Spec.(*harness.StepPlugin)
	if !ok {
		fmt.Println("Error: failed to convert rtGradleRun")
		return nil
	}
	if tmpStepPlugin.With == nil {
		tmpStepPlugin.With = map[string]interface{}{}
	}
	tmpStepPlugin.With["build_tool"] = GradleTool
	attributesList := []string{"url", "username", "password", "access_token", "build_name",
		"build_number", "resolver_id", "deployer_id", "repo_resolve", "repo_deploy"}
	err := SetRtCommandAttributesToInputPlaceHolder(tmpStepPlugin, attributesList)
	if err != nil {
		fmt.Println("Error: failed to set attributes to input placeholder")
		return nil
	}
	return step
}

func convertPublishBuildInfo(node Node, variables map[string]string) *harness.Step {
	step := GetStepWithProperties(&node, nil, ArtifactoryRtCommandsPluginImage)
	if step == nil {
		fmt.Println("Error: failed to convert publishBuildInfo")
		return nil
	}
	tmpStepPlugin, ok := step.Spec.(*harness.StepPlugin)
	if !ok {
		fmt.Println("Error: failed to convert publishBuildInfo")
		return nil
	}
	if tmpStepPlugin.With == nil {
		tmpStepPlugin.With = map[string]interface{}{}
	}
	tmpStepPlugin.With["command"] = "publish"
	attributesList := []string{"build_tool", "url", "username", "password", "access_token", "build_name",
		"build_number", "deployer_id", "deploy_release_repo", "deploy_snapshot_repo"}

	err := SetRtCommandAttributesToInputPlaceHolder(tmpStepPlugin, attributesList)
	if err != nil {
		fmt.Println("Error: failed to set attributes to input placeholder")
		return nil
	}
	return step
}

func convertRtPromote(node Node, variables map[string]string) *harness.Step {
	step := GetStepWithProperties(&node, nil, ArtifactoryRtCommandsPluginImage)
	if step == nil {
		fmt.Println("Error: failed to convert rtPromote")
		return nil
	}
	tmpStepPlugin, ok := step.Spec.(*harness.StepPlugin)
	if !ok {
		fmt.Println("Error: failed to convert rtPromote")
		return nil
	}
	if tmpStepPlugin.With == nil {
		tmpStepPlugin.With = map[string]interface{}{}
	}
	tmpStepPlugin.With["command"] = "promote"
	attributesList := []string{"url", "username", "password", "access_token", "build_name",
		"build_number", "target", "copy"}
	err := SetRtCommandAttributesToInputPlaceHolder(tmpStepPlugin, attributesList)
	if err != nil {
		fmt.Println("Error: failed to set attributes to input placeholder")
		return nil
	}
	return step
}

func convertXrayScan(node Node, variables map[string]string) *harness.Step {
	step := GetStepWithProperties(&node, nil, ArtifactoryRtCommandsPluginImage)
	if step == nil {
		fmt.Println("Error: failed to convert xrayScan")
		return nil
	}
	tmpStepPlugin, ok := step.Spec.(*harness.StepPlugin)
	if !ok {
		fmt.Println("Error: failed to convert xrayScan")
		return nil
	}
	if tmpStepPlugin.With == nil {
		tmpStepPlugin.With = map[string]interface{}{}
	}
	tmpStepPlugin.With["command"] = "scan"
	attributesList := []string{"url", "username", "password", "access_token", "build_name",
		"build_number", "log_level"}
	err := SetRtCommandAttributesToInputPlaceHolder(tmpStepPlugin, attributesList)
	if err != nil {
		fmt.Println("Error: failed to set attributes to input placeholder")
		return nil
	}
	return step
}

func SetRtCommandAttributesToInputPlaceHolder(tmpStepPlugin *harness.StepPlugin, attributeValues []string) error {
	if tmpStepPlugin == nil {
		errStr := "error: rtCommand StepPlugin is nil"
		fmt.Println(errStr)
		return errors.New(errStr)
	}

	for _, attribute := range attributeValues {
		tmpStepPlugin.With[attribute] = InputPlaceHolder
	}
	return nil
}
