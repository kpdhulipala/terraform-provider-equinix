---
subcategory: "Fabric"
---

# equinix_fabric_route_filter (Resource)

Fabric V4 API compatible resource allows creation and management of Equinix Fabric Route Filter Policy

Additional Documentation:
* Getting Started: https://docs.equinix.com/en-us/Content/Interconnection/FCR/FCR-route-filters.htm
* API: https://developer.equinix.com/dev-docs/fabric/api-reference/fabric-v4-apis#route-filters

## Example Usage

```terraform
resource "equinix_fabric_route_filter" "rf_policy" {
  name = "RF_Policy_Name",
  project {
    projectId = "<project_id>"
  },
  type = "BGP_IPv4_PREFIX_FILTER",
  description = "Route Filter Policy for X Purpose",
}

output "id" {
  value = equinix_fabric_route_filter.rf_policy.id
}

output "type" {
  value = equinix_fabric_route_filter.rf_policy.type
}

output "state" {
  value = equinix_fabric_route_filter.rf_policy.state
}

output "not_matched_rules_action" {
  value = equinix_fabric_route_filter.rf_policy.not_matched_rule_action
}

output "connections_count" {
  value = equinix_fabric_route_filter.rf_policy.connections_count
}

output "rules_count" {
  value = equinix_fabric_route_filter.rf_policy.rules_count
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the Route Filter
- `project` (Block Set, Min: 1, Max: 1) The Project object that contains project_id and href that is related to the Fabric Project containing connections the Route Filter can be attached to (see [below for nested schema](#nestedblock--project))
- `type` (String) Route Filter Type. One of [ "BGP_IPv4_PREFIX_FILTER", "BGP_IPv6_PREFIX_FILTER" ]

### Optional

- `description` (String) Optional description to add to the Route Filter you will be creating
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `change` (Set of Object) An object with the details of the previous change applied on the Route Filter (see [below for nested schema](#nestedatt--change))
- `change_log` (Set of Object) (see [below for nested schema](#nestedatt--change_log))
- `connections_count` (Number) The number of Fabric Connections that this Route Filter is attached to
- `href` (String) Route filter URI
- `id` (String) The ID of this resource.
- `not_matched_rule_action` (String) The action that will be taken on ip ranges that don't match the rules present within the Route Filter
- `rules_count` (Number) The number of Route Filter Rules attached to this Route Filter
- `state` (String) State of the Route Filter in its lifecycle
- `uuid` (String) Equinix Assigned ID for Route Filter

<a id="nestedblock--project"></a>
### Nested Schema for `project`

Required:

- `project_id` (String) Project id associated with Fabric Project

Read-Only:

- `href` (String) URI of the Fabric Project


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedatt--change"></a>
### Nested Schema for `change`

Read-Only:

- `href` (String)
- `type` (String)
- `uuid` (String)


<a id="nestedatt--change_log"></a>
### Nested Schema for `change_log`

Read-Only:

- `created_by` (String)
- `created_by_email` (String)
- `created_by_full_name` (String)
- `created_date_time` (String)
- `deleted_by` (String)
- `deleted_by_email` (String)
- `deleted_by_full_name` (String)
- `deleted_date_time` (String)
- `updated_by` (String)
- `updated_by_email` (String)
- `updated_by_full_name` (String)
- `updated_date_time` (String)
