## ---------------------------------------------------------------------------------------------------------------------
## ALL MODULE RESOURCES
## Define all module resources in this file.
## ---------------------------------------------------------------------------------------------------------------------

resource "azurerm_resource_group" "network_rg" {
  name     = var.resource_group_name
  location = var.location
  tags     = var.tags
}

module "network" {
  source = "../../"
  resource_group_name = azurerm_resource_group.network_rg.name
  location            = azurerm_resource_group.network_rg.location
  network             = {
    address_space = ["10.19.0.0/16"]
    subnets = {
      pe-subnet = { 
        address_prefixes = ["10.19.1.0/24"] 
        private_endpoint_network_policies_enabled     = false
        private_link_service_network_policies_enabled = false
      }
      aci-subnet = {
        address_prefixes = ["10.19.2.0/24"]
        delegations      = {
          aci_delegation = {
            name    = "Microsoft.ContainerInstance/containerGroups"
            actions = [ "Microsoft.Network/virtualNetworks/subnets/join/action", "Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action" ]
          }
        }
      }
      fe-subnet = { address_prefixes = ["10.19.3.0/24"] }
    }
  }
  tags = var.tags
}