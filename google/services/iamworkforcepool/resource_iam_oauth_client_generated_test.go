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

package iamworkforcepool_test

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

func TestAccIAMWorkforcePoolOauthClient_iamOauthClientBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMWorkforcePoolOauthClientDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMWorkforcePoolOauthClient_iamOauthClientBasicExample(context),
			},
			{
				ResourceName:            "google_iam_oauth_client.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "oauth_client_id"},
			},
		},
	})
}

func testAccIAMWorkforcePoolOauthClient_iamOauthClientBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_oauth_client" "example" {
  oauth_client_id = "tf-test-example-client-id%{random_suffix}"
  location                  = "global"
  allowed_grant_types       = ["AUTHORIZATION_CODE_GRANT"]
  allowed_redirect_uris     = ["https://www.example.com"]
  allowed_scopes            = ["https://www.googleapis.com/auth/cloud-platform"]
  client_type               = "CONFIDENTIAL_CLIENT"
}
`, context)
}

func TestAccIAMWorkforcePoolOauthClient_iamOauthClientFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMWorkforcePoolOauthClientDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMWorkforcePoolOauthClient_iamOauthClientFullExample(context),
			},
			{
				ResourceName:            "google_iam_oauth_client.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "oauth_client_id"},
			},
		},
	})
}

func testAccIAMWorkforcePoolOauthClient_iamOauthClientFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_oauth_client" "example" {
  oauth_client_id = "tf-test-example-client-id%{random_suffix}"
  display_name              = "Display Name of OAuth client"
  description               = "A sample OAuth client"
  location                  = "global"
  disabled                  = false
  allowed_grant_types       = ["AUTHORIZATION_CODE_GRANT"]
  allowed_redirect_uris     = ["https://www.example.com"]
  allowed_scopes            = ["https://www.googleapis.com/auth/cloud-platform"]
  client_type               = "CONFIDENTIAL_CLIENT"
}
`, context)
}

func testAccCheckIAMWorkforcePoolOauthClientDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_iam_oauth_client" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{IAMWorkforcePoolBasePath}}projects/{{project}}/locations/global/oauthClients/{{oauth_client_id}}")
			if err != nil {
				return err
			}

			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err != nil {
				return nil
			}

			if v := res["state"]; v == "DELETED" {
				return nil
			}

			return fmt.Errorf("IAMOAuthCLient still exists at %s", url)
		}

		return nil
	}
}
