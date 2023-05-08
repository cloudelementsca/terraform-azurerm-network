
output "vnet" {
  description = "All vnet info."
  value       = azurerm_virtual_network.vnet
}

output "subnets" {
  description = "All subnets info."
  value       = azurerm_subnet.subnets
}