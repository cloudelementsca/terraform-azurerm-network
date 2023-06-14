## ---------------------------------------------------------------------------------------------------------------------
## ALL MODULE RESOURCES
## Define all module resources in this file.
## ---------------------------------------------------------------------------------------------------------------------

resource "azurerm_virtual_network" "vnet" {
  name                    = var.network.name != null ? var.network.name : "vnet-${random_string.random_string_vnet.result}"
  location                = var.location
  resource_group_name     = var.resource_group_name
  address_space           = var.network.address_space
  dns_servers             = var.network.dns_servers
  bgp_community           = var.network.bgp_community
  edge_zone               = var.network.edge_zone
  flow_timeout_in_minutes = var.network.flow_timeout_in_minutes

  dynamic "ddos_protection_plan" {
    for_each = var.network.ddos_protection_plan != null ? var.network.ddos_protection_plan : {}

    content {
      id     = ddos_protection_plan.value.id
      enable = ddos_protection_plan.value.enable
    }
  }

  tags = var.tags
}

resource "azurerm_subnet" "subnets" {
  for_each = var.network.subnets

  name                                          = each.value.name != null ? each.value.name : "subnet-${random_string.random_string_subnets[each.key].result}"
  resource_group_name                           = var.resource_group_name
  virtual_network_name                          = azurerm_virtual_network.vnet.name
  address_prefixes                              = each.value.address_prefixes != null ? each.value.address_prefixes : azurerm_virtual_network.vnet.address_space
  private_endpoint_network_policies_enabled     = each.value.private_endpoint_network_policies_enabled
  private_link_service_network_policies_enabled = each.value.private_link_service_network_policies_enabled
  service_endpoints                             = each.value.service_endpoints
  service_endpoint_policy_ids                   = each.value.service_endpoint_policy_ids

  dynamic "delegation" {
    for_each = each.value.service_delegations != null ? each.value.service_delegations : {}
    content {
      name = delegation.key
      service_delegation {
        name    = delegation.value.name
        actions = delegation.value.actions
      }
    }
  }
}

resource "random_string" "random_string_vnet" {
  length  = 8
  special = false
  upper   = false
}

resource "random_string" "random_string_subnets" {
  for_each = var.network.subnets

  length  = 8
  special = false
  upper   = false
}
