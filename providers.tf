## ---------------------------------------------------------------------------------------------------------------------
## ALL PROVIDERS REQUIRED FOR MODULE
## Add all required providers below.
## ---------------------------------------------------------------------------------------------------------------------

terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.54.0"
    }
  }
}