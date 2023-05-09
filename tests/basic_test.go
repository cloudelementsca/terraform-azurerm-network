package basic_test

import (
	"testing"
	"encoding/json"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestBasicNetworkModule(t *testing.T) {

	type VnetStruct struct {
		Address_space    		[]string
		Bgp_community   		string
		Ddos_protection_plan    []string
		Dns_servers  			[]string
		Edge_zone				string
		Flow_timeout_in_minutes int
		Location				string
		Resource_group_name		string
		Tags					map[string]interface{}
	}

	expectedVnetOutput := VnetStruct{
		Address_space:    		 []string{"10.0.0.0/8"},
		Bgp_community:   		 "",
		Ddos_protection_plan:    []string{},
		Dns_servers: 			 []string{},
		Edge_zone:				 "",
		Flow_timeout_in_minutes: 0,
		Location:				 "canadacentral",
		Resource_group_name:	 "vnet-rg",
		Tags:					 map[string]interface{}{"environment":"dev"},
	}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualObject := VnetStruct{}

	str := terraform.OutputJson(t, terraformOptions, "vnet")

	json.Unmarshal([]byte(str), &actualObject)

	assert.Equal(t, expectedVnetOutput, actualObject, &actualObject)
}