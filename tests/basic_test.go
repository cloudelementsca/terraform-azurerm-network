package test

import (
	"testing"

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
		edge_zone:				 nil,
		flow_timeout_in_minutes: nil,
		location:				 "canadacentral",
		resource_group_name:	 "vnet-rg",
		timeouts				 nil,
		tags					 map[string]interface{}{"environments":"dev"},
	}
	
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualObject := VnetStruct{}
	OutputStruct(t, terraformOptions, "vnet", &actualObject)
	assert.Equal(t, expectedVnetOutput, actualObject)
}