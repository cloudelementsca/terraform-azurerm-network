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
  source              = "../../"
  resource_group_name = azurerm_resource_group.network_rg.name
  location            = azurerm_resource_group.network_rg.location
}