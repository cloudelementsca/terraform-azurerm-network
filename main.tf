## ---------------------------------------------------------------------------------------------------------------------
## ALL MODULE RESOURCES
## Define all module resources in this file.
## ---------------------------------------------------------------------------------------------------------------------


resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = var.location
  tags     = var.tags
}

resource "azurerm_virtual_network" "vnet" {
  name                    = var.network.name
  location                = azurerm_resource_group.rg.location
  resource_group_name     = azurerm_resource_group.rg.name
  address_space           = var.network.address_space
  dns_servers             = var.network.dns_servers
  bgp_community           = var.network.bgp_community
  edge_zone               = var.network.edge_zone
  flow_timeout_in_minutes = var.network.flow_timeout_in_minutes
  
  ddos_protection_plan {
    id     = var.network.ddos_protection_plan.id
    enable = var.network.ddos_protection_plan.enable
  }  

  tags = var.tags
}

resource "azurerm_subnet" "subnets" { 
  for_each = var.network.subnets

  name                                          = each.value.name
  resource_group_name                           = azurerm_resource_group.rg.name 
  virtual_network_name                          = azurerm_virtual_network.vnet.name 
  address_prefixes                              = each.value.address_prefixes
  private_endpoint_network_policies_enabled     = each.value.private_endpoint_network_policies_enabled
  private_link_service_network_policies_enabled = each.value.private_link_service_network_policies_enabled
  service_endpoints                             = each.value.service_endpoints
  service_endpoint_policy_ids                   = each.value.service_endpoint_policy_ids
    
  dynamic "delegation" { 
    for_each = each.value.delegations
    content {
      name = delegation.value.name
      service_delegation { 
        name = delegation.value.service_delegation.name
        actions = delegation.value.service_delegation.actions
      } 
    }      
  } 
}
