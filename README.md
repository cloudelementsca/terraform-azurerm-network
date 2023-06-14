# terraform-azurerm-network
## Description
This module creates an Azure virtual network in a resource group and location that you specify. It's highly adjustable and takes the same input variables as `azurerm_virtual_network` and `azurerm_subnet` resource definitions have attributes.

## Using the Module
See the `examples/` folders on how to use the module.

### Terraform Registry
```
module "network" {
  source = "cloudelementsca/terraform-azurerm-network"

  resource_group_name = azurerm_resource_group.network_rg.name
  location            = azurerm_resource_group.network_rg.location
}
```

### GitHub
```
module "network" {
  source = "github.com/cloudelementsca/terraform-azurerm-network"

  resource_group_name = azurerm_resource_group.network_rg.name
  location            = azurerm_resource_group.network_rg.location
}
```
## Requirements

The following requirements are needed by this module:

- <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) (=3.54.0)

- <a name="requirement_random"></a> [random](#requirement\_random) (=3.5.1)

## Providers

The following providers are used by this module:

- <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) (=3.54.0)

- <a name="provider_random"></a> [random](#provider\_random) (=3.5.1)

## Modules

No modules.

## Resources

The following resources are used by this module:

- [azurerm_subnet.subnets](https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/subnet) (resource)
- [azurerm_virtual_network.vnet](https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/virtual_network) (resource)
- [random_string.random_string_subnets](https://registry.terraform.io/providers/hashicorp/random/3.5.1/docs/resources/string) (resource)
- [random_string.random_string_vnet](https://registry.terraform.io/providers/hashicorp/random/3.5.1/docs/resources/string) (resource)

## Required Inputs

The following input variables are required:

### <a name="input_location"></a> [location](#input\_location)

Description: Location for all resources.

Type: `string`

### <a name="input_resource_group_name"></a> [resource\_group\_name](#input\_resource\_group\_name)

Description: Name of existing resource group that will contain the vnet.

Type: `string`

## Optional Inputs

The following input variables are optional (have default values):

### <a name="input_network"></a> [network](#input\_network)

Description: Vnet definition.

Type:

```hcl
object({
    address_space = list(string)
    subnets = optional(map(object({
      address_prefixes = optional(list(string))
      name             = optional(string)
      service_delegations = optional(map(object({
        name    = string
        actions = list(string)
      })), {})
      private_endpoint_network_policies_enabled     = optional(bool)
      private_link_service_network_policies_enabled = optional(bool)
      service_endpoints                             = optional(list(string))
      service_endpoint_policy_ids                   = optional(list(string))
    })), { subnet1 = {} })
    name          = optional(string)
    dns_servers   = optional(list(string), [])
    bgp_community = optional(string)
    ddos_protection_plan = optional(object({
      id     = string
      enable = bool
    }))
    edge_zone               = optional(string)
    flow_timeout_in_minutes = optional(number)
  })
```

Default:

```json
{
  "address_space": [
    "10.0.0.0/8"
  ]
}
```

### <a name="input_tags"></a> [tags](#input\_tags)

Description: Tags for all resources.

Type: `map(string)`

Default:

```json
{
  "environment": "dev"
}
```

## Outputs

The following outputs are exported:

### <a name="output_subnets"></a> [subnets](#output\_subnets)

Description: All subnets info.

### <a name="output_vnet"></a> [vnet](#output\_vnet)

Description: All vnet info.
