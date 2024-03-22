terraform {
  required_providers {
    stevan = {
      source  = "hashicorp.com/edu/stevan"
      version = "1.0.0"
    }
  }
}

provider "stevan" {

  host = "http://localhost:8000"

  auth_token = "MY_TOKEN"

}
