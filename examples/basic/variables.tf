## ---------------------------------------------------------------------------------------------------------------------
## ALL MODULE VARIABLES
## Define all module variables below.
## ---------------------------------------------------------------------------------------------------------------------


variable "resource_group_name" {
  description = "Name of the resource group that will contain the vnet."
  type        = string
}

variable "location" {
  description = "Location for all resources."
  type        = string
}

variable "tags" {
  description = "Tags to be applied to all resources."
  type        = map(string)
  default = {
    "environment" = "dev"
  }
} 

