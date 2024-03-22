resource "stevan_api_product_version" "example" {
  name = "FirstVersion"

  gateway_service = {
    id               = "09b4786a-3e48-4631-8f6b-62d1d8e1a7f3"
    control_plane_id = "e4d9ebb1-26b4-426a-b00e-cb67044f3baf"
  }


  publish_status = "publish_status"

  deprecated = false

  notify = true

  api_product_id = "d32d905a-ed33-46a3-a093-d8f536af9a8a"

}
