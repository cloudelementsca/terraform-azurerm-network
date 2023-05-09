package basic_test

import (
	"testing"
	"encoding/json"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestBasicNetworkModule(t *testing.T) {

	type VnetStruct struct {
		Address_space    		[]string
		Bgp_community   		string
		Ddos_protection_plan    []string
		Dns_servers  			[]string
		Edge_zone				string
		Flow_timeout_in_minutes int
		//Guid					string
		//Id						string
		Location				string
		Resource_group_name		string
		//Timeouts				[]string
		Tags					map[string]interface{}
	}

	expectedVnetOutput := VnetStruct{
		Address_space:    		 []string{"10.0.0.0/8"},
		Bgp_community:   		 "",
		Ddos_protection_plan:    []string{},
		Dns_servers: 			 []string{},
		Edge_zone:				 "",
		Flow_timeout_in_minutes: 0,
		//Guid:					 "",
		//Id:						 "",
		Location:				 "canadacentral",
		Resource_group_name:	 "vnet-rg",
		//Timeouts:				 nil,
		Tags:					 map[string]interface{}{"environment":"dev"},
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
	//terraform.OutputStruct(t, terraformOptions, "vnet", &actualObject)
	str := terraform.OutputJson(t, terraformOptions, "vnet")
	json.Unmarshal([]byte(str), &actualObject)
	assert.Equal(t, expectedVnetOutput, actualObject, &actualObject)
	//require.Equal(t, str, expectedVnetOutput, "JSON %q should match %q", expectedVnetOutput, str)
}