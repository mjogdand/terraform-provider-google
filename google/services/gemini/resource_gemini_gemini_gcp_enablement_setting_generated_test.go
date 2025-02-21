// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

package gemini_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccGeminiGeminiGcpEnablementSetting_geminiGeminiGcpEnablementSettingBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGeminiGeminiGcpEnablementSettingDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGeminiGeminiGcpEnablementSetting_geminiGeminiGcpEnablementSettingBasicExample(context),
			},
			{
				ResourceName:            "google_gemini_gemini_gcp_enablement_setting.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"gemini_gcp_enablement_setting_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccGeminiGeminiGcpEnablementSetting_geminiGeminiGcpEnablementSettingBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gemini_gemini_gcp_enablement_setting" "example" {
    gemini_gcp_enablement_setting_id = "tf-test-ls1-tf%{random_suffix}"
    location = "global"
    labels = {"my_key": "my_value"}
    enable_customer_data_sharing = true
}
`, context)
}

func testAccCheckGeminiGeminiGcpEnablementSettingDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_gemini_gemini_gcp_enablement_setting" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/geminiGcpEnablementSettings/{{gemini_gcp_enablement_setting_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("GeminiGeminiGcpEnablementSetting still exists at %s", url)
			}
		}

		return nil
	}
}
