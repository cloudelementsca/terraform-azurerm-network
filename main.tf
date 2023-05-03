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
  name                = var.network.name
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  address_space       = var.network.address_space
  dns_servers         = var.network.dns_servers

  tags = var.tags
}

resource "azurerm_subnet" "subnets" { 
  for_each = var.network.subnets

  name                 = each.value.name
  resource_group_name  = azurerm_resource_group.rg.name 
  virtual_network_name = azurerm_virtual_network.vnet.name 
  address_prefixes     = each.value.address_prefixes
    
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
