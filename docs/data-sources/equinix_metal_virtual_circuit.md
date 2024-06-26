---
subcategory: "Metal"
---

# equinix_metal_virtual_circuit (Data Source)

Use this data source to retrieve a virtual circuit resource from
[Equinix Fabric - software-defined interconnections](https://deploy.equinix.com/developers/docs/metal/interconnections/introduction/)

See the [Virtual Routing and Forwarding documentation](https://deploy.equinix.com/developers/docs/metal/layer2-networking/vrf/) for product details and API reference material.

## Example Usage

```hcl
data "equinix_metal_connection" "example_connection" {
  connection_id = "4347e805-eb46-4699-9eb9-5c116e6a017d"
}

data "equinix_metal_virtual_circuit" "example_vc" {
  virtual_circuit_id = data.equinix_metal_connection.example_connection.ports[1].virtual_circuit_ids[0]
}

```

## Argument Reference

The following arguments are supported:

* `virtual_circuit_id` - (Required) ID of the virtual circuit resource

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Name of the virtual circuit resource.
* `connection_id` - UUID of Connection where the VC is scoped to.
* `status` - Status of the virtal circuit.
* `port_id` - UUID of the Connection Port where the VC is scoped to.
* `project_id` - ID of project to which the VC belongs.
* `vnid`, `nni_vlan`, `nni_nvid` - VLAN parameters, see the
[documentation for Equinix Fabric](https://deploy.equinix.com/developers/docs/metal/interconnections/introduction/).
* `description` - Description for the Virtual Circuit resource.
* `tags` - Tags for the Virtual Circuit resource.
* `speed` - Speed of the Virtual Circuit resource.
* `vrf_id` - UUID of the VLAN to associate.
* `peer_asn` - The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
* `subnet` - A subnet from one of the IP
  blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
  * For a /31 block, it will only have two IP addresses, which will be used for
  the metal_ip and customer_ip.
  * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
* `metal_ip` - The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
* `customer_ip` - The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
* `md5` - The password that can be set for the VRF BGP peer
