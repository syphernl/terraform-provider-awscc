// Code generated by generators/resource/main.go; DO NOT EDIT.

package panorama_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSPanoramaPackageVersion_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Panorama::PackageVersion", "awscc_panorama_package_version", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
