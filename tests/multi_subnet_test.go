package basic_test

import (
	"encoding/json"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicNetworkModule(t *testing.T) {

	type VnetStruct struct {
		Address_space           []string
		Bgp_community           string
		Ddos_protection_plan    []string
		Dns_servers             []string
		Edge_zone               string
		Flow_timeout_in_minutes int
		Location                string
		Resource_group_name     string
		Tags                    map[string]interface{}
	}

	type ServiceDelegationStruct struct {
		Name    string
		Actions []string
	}

	type DelegationStruct struct {
		Name               string
		Service_delegation []ServiceDelegationStruct
	}

	type SubnetStruct struct {
		Address_prefixes                              []string
		Delegation                                    []DelegationStruct
		Private_endpoint_network_policies_enabled     bool
		Private_link_service_network_policies_enabled bool
	}

	expectedVnetOutput := VnetStruct{
		Address_space:           []string{"10.19.0.0/16"},
		Bgp_community:           "",
		Ddos_protection_plan:    []string{},
		Dns_servers:             []string{},
		Edge_zone:               "",
		Flow_timeout_in_minutes: 0,
		Location:                "canadacentral",
		Resource_group_name:     "vnet-rg",
		Tags:                    map[string]interface{}{"environment": "dev"},
	}

	expectedPeSubnetOutupt := SubnetStruct{}

	expectedFeSubnetOutupt := SubnetStruct{}

	expectedAciSubnetOutupt := SubnetStruct{}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/multi_subnet",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualObject := VnetStruct{}

	str := terraform.OutputJson(t, terraformOptions, "vnet")

	json.Unmarshal([]byte(str), &actualObject)

	assert.Equal(t, expectedVnetOutput, actualObject, &actualObject)
}
