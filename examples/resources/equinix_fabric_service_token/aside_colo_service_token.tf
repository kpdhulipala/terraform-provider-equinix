resource "equinix_fabric_service_token" "test" {
  type                 = "VC_TOKEN"
  description          = "Aside COLO Service Token"
  expiration_date_time = "2025-01-18T06:43:49.981Z"
  service_token_connection {
    type = "EVPL_VC"
    bandwidth_limit = 1000
    a_side {
      access_point_selectors{
        type = "COLO"
        port {
          uuid = "<port_uuid>"
        }
        link_protocol {
          type = "DOT1Q"
          vlan_tag = "2987"
        }
      }
    }
  }
  notifications {
    type   = "ALL"
    emails = ["example@equinix.com", "test1@equinix.com"]
  }
}
