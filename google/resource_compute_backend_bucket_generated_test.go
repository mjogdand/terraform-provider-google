// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeBackendBucket_backendBucketBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeBackendBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucket_backendBucketBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_backend_bucket.image_backend",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendBucket_backendBucketBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "image-backend-bucket-%s"
  description = "Contains beautiful images"
  bucket_name = "${google_storage_bucket.image_bucket.name}"
  enable_cdn  = true
}

resource "google_storage_bucket" "image_bucket" {
  name     = "image-store-bucket-%s"
  location = "EU"
}
`, val, val,
	)
}

func testAccCheckComputeBackendBucketDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_backend_bucket" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeBackendBucket still exists at %s", url)
		}
	}

	return nil
}
