## ---------------------------------------------------------------------------------------------------------------------
## ALL MODULE VARIABLES
## Define all module variables below.
## ---------------------------------------------------------------------------------------------------------------------

## ---------------------------------------------------------------------------------------------------------------------
## MODULE PARAMETERS
## These variables are expected to be passed in by the operator
## ---------------------------------------------------------------------------------------------------------------------

variable "resource_group_name" {
  description = "Name of the resource group that will contain the vnet."
  type        = string
}

variable "location" {
  description = "Location for all resources."
  type        = string
}

## ---------------------------------------------------------------------------------------------------------------------
## OPTIONAL PARAMETERS
## These variables have defaults and may be overridden
## ---------------------------------------------------------------------------------------------------------------------

variable "network" {
  description = "Vnet definition."
  type = object({
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
  default = {
    address_space = ["10.0.0.0/8"]
  }
}

variable "tags" {
  description = "Tags for all resources."
  type        = map(string)
  default     = { environment = "dev" }
}

