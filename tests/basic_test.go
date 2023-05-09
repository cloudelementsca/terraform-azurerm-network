package basic_test

import (
	"testing"
	"encoding/json"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestBasicNetworkModule(t *testing.T) {

	type VnetStruct struct {
		address_space    		[]string
		bgp_community   		string
		ddos_protection_plan    []string
		dns_servers  			[]string
		edge_zone				string
		flow_timeout_in_minutes int
		location				string
		resource_group_name		string
		timeouts				[]string
		tags					map[string]interface{}
	}

	expectedVnetOutput := VnetStruct{
		address_space:    		 []string{"10.0.0.0/8"},
		bgp_community:   		 "",
		ddos_protection_plan:    []string{},
		dns_servers: 			 []string{},
		edge_zone:				 "",
		flow_timeout_in_minutes: 0,
		location:				 "canadacentral",
		resource_group_name:	 "vnet-rg",
		timeouts:				 nil,
		tags:					 map[string]interface{}{"environments":"dev"},
	}

//	expectedVnetOutput := `{
//		"address_space":["10.0.0.0/8"],
//		"bgp_community":"",
//		"ddos_protection_plan":[],
//		"dns_servers":[],
//		"edge_zone":"",
//		"flow_timeout_in_minutes":0,
//		"guid":"549bfbaf-a8b7-4e5d-8419-8a400c1961d2",
//		"id":"/subscriptions/***/resourceGroups/vnet-rg/providers/Microsoft.Network/virtualNetworks/vnet-qzonuvo9",
//		"location":"canadacentral",
//		"name":"vnet-qzonuvo9",
//		"resource_group_name":"vnet-rg",
//		"subnet":[],
//		"tags":{"environment":"dev"},
//		"timeouts":null
//	}`
	
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualObject := VnetStruct{}
	terraform.OutputStruct(t, terraformOptions, "vnet", &actualObject)
	assert.Equal(t, expectedVnetOutput, json.Unmarshal(actualObject, &actualObject))
}