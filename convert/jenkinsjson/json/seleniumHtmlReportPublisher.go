package json

import (
	harness "github.com/drone/spec/dist/go"
)

func ConvertSeleniumHtmlReportPublisher(node Node, variables map[string]string) *harness.Step {

	step := &harness.Step{
		Id:   SanitizeForId("UploadPublish", node.SpanId),
		Name: "Upload and Publish",
		Type: "plugin",
		Spec: &harness.StepPlugin{
			Connector: "<+input>",
			Image:     DroneS3UploadPublish,
			With: map[string]interface{}{
				"aws_access_key_id":                    "<+input>",
				"aws_secret_access_key":                "<+input>",
				"aws_bucket":                           "<+input>",
				"aws_default_region":                   "<+input>",
				"test_results_dir":                     "<+input>",
				"fail_if_exception_on_parsing_results": "<+input>",
			},
		},
	}
	return step
}
