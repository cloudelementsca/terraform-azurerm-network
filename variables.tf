## ---------------------------------------------------------------------------------------------------------------------
## ALL MODULE VARIABLES
## Define all module variables below.
## ---------------------------------------------------------------------------------------------------------------------

## ---------------------------------------------------------------------------------------------------------------------
## ENVIRONMENT VARIABLES
## Define these secrets as environment variables
## ---------------------------------------------------------------------------------------------------------------------

## TF_VAR_master_password

## ---------------------------------------------------------------------------------------------------------------------
## MODULE PARAMETERS
## These variables are expected to be passed in by the operator
## ---------------------------------------------------------------------------------------------------------------------

## ---------------------------------------------------------------------------------------------------------------------
## OPTIONAL PARAMETERS
## These variables have defaults and may be overridden
## ---------------------------------------------------------------------------------------------------------------------

variable "resource_group_name" {
  description = "Name of the resource group that will contain the vnet."
  type        = string
  default     = "azurerm-vnet-rg"
}

variable "location" {
  description = "Location for all resources."
  type        = string
  default     = "canadacentral"
}

variable "tags" {
  description = "Tags for all resources."
  type        = map(string)
  default     = { environment = "dev" }
}

variable "network" {
  description = "Vnet definition."
  type        = object({
    name                = string
    address_space       = list(string)
    subnets             = map(object({
      name             = string
      address_prefixes = list(string)
      delegations      = optional(map(object({
        name               = string
        service_delegation = object({
          name    = string
          actions = list(string)
        })        
      })))
      private_endpoint_network_policies_enabled     = optional(bool)
      private_link_service_network_policies_enabled = optional(bool)
      service_endpoints                             = optional(list(string), [])
      service_endpoint_policy_ids                   = optional(list(string), [])
    }))
    dns_servers          = optional(list(string), [])
    bgp_community        = optional(string)
    ddos_protection_plan = optional(object({
      id     = string
      enable = bool
    }))
    edge_zone               = optional(string)
    flow_timeout_in_minutes = optional(number)
  })
  default     = {
    name                = "azurerm-vnet"
    address_space       = ["10.0.0.0/8"]
    subnets             = {
      subnet1 = {
        name             = "subnet1"
        address_prefixes = "10.0.10.0/24"
      }
    }
  }
}