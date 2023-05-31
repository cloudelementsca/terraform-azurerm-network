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

