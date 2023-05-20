package network_tests

import (
	"encoding/json"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestBasicNetworkModule(t *testing.T) {

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

	expectedPeSubnetOutupt := SubnetStruct{
		Address_prefixes: []string{"10.19.1.0/24"},
		Delegation:       []DelegationStruct{},
		Private_endpoint_network_policies_enabled:     false,
		Private_link_service_network_policies_enabled: false,
	}

	expectedAciSubnetOutupt := SubnetStruct{
		Address_prefixes: []string{"10.19.2.0/24"},
		Delegation: []DelegationStruct{
			Service_delegation: []ServiceDelegationStruct{
				Name:    "Microsoft.ContainerInstance/containerGroups",
				Actions: []string{"Microsoft.Network/virtualNetworks/subnets/join/action", "Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action"},
			},
		},
		Private_endpoint_network_policies_enabled:     true,
		Private_link_service_network_policies_enabled: true,
	}

	expectedFeSubnetOutupt := SubnetStruct{
		Address_prefixes: []string{"10.19.3.0/24"},
		Delegation:       []DelegationStruct{},
		Private_endpoint_network_policies_enabled:     true,
		Private_link_service_network_policies_enabled: true,
	}

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