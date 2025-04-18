---
subcategory: "Fabric"
---

# equinix_fabric_cloud_routers (Data Source)

Fabric V4 API compatible data resource that allow user to fetch Fabric Cloud Routers matching custom search criteria

Additional documentation:
* Getting Started: https://docs.equinix.com/en-us/Content/Interconnection/FCR/FCR-intro.htm#HowItWorks
* API: https://developer.equinix.com/dev-docs/fabric/api-reference/fabric-v4-apis#fabric-cloud-routers

## Example Usage

```terraform
data "equinix_fabric_cloud_routers" "test" {
  filter {
    property = "/name"
    operator = "="
    values 	 = ["Test_PFCR"]
  }
  filter {
    property = "/location/metroCode"
    operator = "="
    values   = ["SV"]
  }
  filter {
    property = "/package/code"
    operator = "="
    values = ["STANDARD"]
    or = true
  }
  filter {
    property = "/state"
    operator = "="
    values = ["ACTIVE"]
    or = true
  }
  pagination {
    offset = 5
    limit = 3
  }
  sort {
    direction = "ASC"
    property = "/name"
  }
}

output "number_of_returned_fcrs" {
  value = length(data.equinix_fabric_cloud_routers.test.data)
}

output "first_fcr_name" {
  value = data.equinix_fabric_cloud_routers.test.data.0.name
}

output "first_fcr_state" {
  value = data.equinix_fabric_cloud_routers.test.data.0.state
}

output "first_fcr_uuid" {
  value = data.equinix_fabric_cloud_routers.test.data.0.uuid
}

output "first_fcr_type" {
  value = data.equinix_fabric_cloud_routers.test.data.0.type
}

output "first_fcr_package_code" {
  value = one(data.equinix_fabric_cloud_routers.test.data.0.package).code
}

output "first_fcr_equinix_asn" {
  value = data.equinix_fabric_cloud_routers.test.data.0.equinix_asn
}

output "first_fcr_location_region" {
  value = one(data.equinix_fabric_cloud_routers.test.data.0.location).region
}

output "first_fcr_location_metro_name" {
  value = one(data.equinix_fabric_cloud_routers.test.data.0.location).metro_name
}

output "first_fcr_location_metro_code" {
  value = one(data.equinix_fabric_cloud_routers.test.data.0.location).metro_code
}

output "first_fcr_project_id" {
  value = one(data.equinix_fabric_cloud_routers.test.data.0.project).project_id
}

output "first_fcr_account_number" {
  value = one(data.equinix_fabric_cloud_routers.test.data.0.account).account_number
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `filter` (Block List, Min: 1, Max: 10) Filters for the Data Source Search Request. Maximum of 8 total filters. (see [below for nested schema](#nestedblock--filter))

### Optional

- `pagination` (Block Set, Max: 1) Pagination details for the Data Source Search Request (see [below for nested schema](#nestedblock--pagination))
- `sort` (Block List) Filters for the Data Source Search Request (see [below for nested schema](#nestedblock--sort))

### Read-Only

- `data` (List of Object) List of Cloud Routers (see [below for nested schema](#nestedatt--data))
- `id` (String) The ID of this resource.

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Required:

- `operator` (String) Possible operators to use on the filter property. Can be one of the following: [= - equal, != - not equal, > - greater than, >= - greater than or equal to, < - less than, <= - less than or equal to, [NOT] BETWEEN - (not) between, [NOT] LIKE - (not) like, [NOT] IN - (not) in
- `property` (String) The API response property which you want to filter your request on. Can be one of the following: "/project/projectId", "/name", "/uuid", "/state", "/location/metroCode", "/location/metroName", "/package/code", "/*"
- `values` (List of String) The values that you want to apply the property+operator combination to in order to filter your data search

Optional:

- `or` (Boolean) Boolean flag indicating whether this filter is included in the OR group. There can only be one OR group and it can have a maximum of 3 filters. The OR group only counts as 1 of the 8 possible filters


<a id="nestedblock--pagination"></a>
### Nested Schema for `pagination`

Optional:

- `limit` (Number) Number of elements to be requested per page. Number must be between 1 and 100. Default is 20
- `offset` (Number) The page offset for the pagination request. Index of the first element. Default is 0.


<a id="nestedblock--sort"></a>
### Nested Schema for `sort`

Optional:

- `direction` (String) The sorting direction. Can be one of: [DESC, ASC], Defaults to DESC
- `property` (String) The property name to use in sorting. Can be one of the following: [/name, /uuid, /state, /location/metroCode, /location/metroName, /package/code, /changeLog/createdDateTime, /changeLog/updatedDateTime], Defaults to /changeLog/updatedDateTime


<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `account` (Set of Object) (see [below for nested schema](#nestedobjatt--data--account))
- `change_log` (Set of Object) (see [below for nested schema](#nestedobjatt--data--change_log))
- `connections_count` (Number)
- `description` (String)
- `equinix_asn` (Number)
- `href` (String)
- `location` (Set of Object) (see [below for nested schema](#nestedobjatt--data--location))
- `marketplace_subscription` (Set of Object) (see [below for nested schema](#nestedobjatt--data--marketplace_subscription))
- `name` (String)
- `notifications` (List of Object) (see [below for nested schema](#nestedobjatt--data--notifications))
- `order` (Set of Object) (see [below for nested schema](#nestedobjatt--data--order))
- `package` (Set of Object) (see [below for nested schema](#nestedobjatt--data--package))
- `project` (Set of Object) (see [below for nested schema](#nestedobjatt--data--project))
- `state` (String)
- `type` (String)
- `uuid` (String)

<a id="nestedobjatt--data--account"></a>
### Nested Schema for `data.account`

Read-Only:

- `account_number` (Number)


<a id="nestedobjatt--data--change_log"></a>
### Nested Schema for `data.change_log`

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


<a id="nestedobjatt--data--location"></a>
### Nested Schema for `data.location`

Read-Only:

- `ibx` (String)
- `metro_code` (String)
- `metro_name` (String)
- `region` (String)


<a id="nestedobjatt--data--marketplace_subscription"></a>
### Nested Schema for `data.marketplace_subscription`

Read-Only:

- `type` (String)
- `uuid` (String)


<a id="nestedobjatt--data--notifications"></a>
### Nested Schema for `data.notifications`

Read-Only:

- `emails` (List of String)
- `send_interval` (String)
- `type` (String)


<a id="nestedobjatt--data--order"></a>
### Nested Schema for `data.order`

Read-Only:

- `billing_tier` (String)
- `order_id` (String)
- `order_number` (String)
- `purchase_order_number` (String)
- `term_length` (Number)


<a id="nestedobjatt--data--package"></a>
### Nested Schema for `data.package`

Read-Only:

- `code` (String)


<a id="nestedobjatt--data--project"></a>
### Nested Schema for `data.project`

Read-Only:

- `href` (String)
- `project_id` (String)
