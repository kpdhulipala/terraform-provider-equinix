package service_token_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/equinix/terraform-provider-equinix/internal/acceptance"
	"github.com/equinix/terraform-provider-equinix/internal/fabric/testing_helpers"
	"github.com/equinix/terraform-provider-equinix/internal/resources/fabric/service_token"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccFabricZsideVirtualDeviceServiceToken_PNFV(t *testing.T) {
	connectionTestData := testing_helpers.GetFabricEnvConnectionTestData(t)
	var virtualDevice string
	if len(connectionTestData) > 0 {
		virtualDevice = connectionTestData["pnfv"]["virtualDevice"]
	}
	serviceTokenName, serviceTokenUpdatedName := "Service_token_PNFV", "UP_Service_Token_PNFV"
	serviceTokenDescription, serviceTokenUpdatedDescription := "zside vd token", "Updated zside vd token"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: CheckServiceTokenDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricZsideVirtualDeviceServiceTokenConfig(serviceTokenName, serviceTokenDescription, virtualDevice),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("equinix_fabric_service_token.test", "uuid"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "name", serviceTokenName),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "type", "VC_TOKEN"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "description", serviceTokenDescription),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "expiration_date_time", "2025-01-18T06:43:49.981Z"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.supported_bandwidths.#", "3"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.z_side.0.access_point_selectors.0.virtual_device.0.type", "EDGE"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.z_side.0.access_point_selectors.0.virtual_device.0.uuid", virtualDevice),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.z_side.0.access_point_selectors.0.interface.0.type", "NETWORK"),
					resource.TestCheckResourceAttrSet("equinix_fabric_service_token.test", "service_token_connection.0.z_side.0.access_point_selectors.0.interface.0.id"),
				),
				ExpectNonEmptyPlan: false,
			},
			{
				Config: testAccFabricZsideVirtualDeviceServiceTokenConfig(serviceTokenUpdatedName, serviceTokenUpdatedDescription, virtualDevice),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("equinix_fabric_service_token.test", "uuid"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "name", serviceTokenUpdatedName),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "type", "VC_TOKEN"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "description", serviceTokenUpdatedDescription),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "expiration_date_time", "2025-01-18T06:43:49.981Z"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.supported_bandwidths.#", "3"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.z_side.0.access_point_selectors.0.virtual_device.0.type", "EDGE"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.z_side.0.access_point_selectors.0.virtual_device.0.uuid", virtualDevice),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.z_side.0.access_point_selectors.0.interface.0.type", "NETWORK"),
					resource.TestCheckResourceAttrSet("equinix_fabric_service_token.test", "service_token_connection.0.z_side.0.access_point_selectors.0.interface.0.id"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccFabricAsidePortServiceToken_PNFV(t *testing.T) {
	ports := testing_helpers.GetFabricEnvPorts(t)
	var portUuid string
	if len(ports) > 0 {
		portUuid = ports["pnfv"]["dot1q"][0].GetUuid()
	}
	serviceTokenName, serviceTokenUpdatedName := "token_port_PNFV", "UP_Token_port_PNFV"
	serviceTokenDescription, serviceTokenUpdatedDescription := "aside port token", "Updated aside port token"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: CheckServiceTokenDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricAsidePortServiceTokenConfig(serviceTokenName, serviceTokenDescription, portUuid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("equinix_fabric_service_token.test", "uuid"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "name", serviceTokenName),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "type", "VC_TOKEN"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "description", serviceTokenDescription),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "expiration_date_time", "2025-01-18T06:43:49.981Z"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.bandwidth_limit", "1000"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.a_side.0.access_point_selectors.0.port.0.uuid", portUuid),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.a_side.0.access_point_selectors.0.link_protocol.0.type", "DOT1Q"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.a_side.0.access_point_selectors.0.link_protocol.0.vlan_tag", "2987"),
				),
				ExpectNonEmptyPlan: false,
			},
			{
				Config: testAccFabricAsidePortServiceTokenConfig(serviceTokenUpdatedName, serviceTokenUpdatedDescription, portUuid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("equinix_fabric_service_token.test", "uuid"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "name", serviceTokenUpdatedName),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "type", "VC_TOKEN"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "description", serviceTokenUpdatedDescription),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "expiration_date_time", "2025-01-18T06:43:49.981Z"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.bandwidth_limit", "1000"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.a_side.0.access_point_selectors.0.port.0.uuid", portUuid),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.a_side.0.access_point_selectors.0.link_protocol.0.type", "DOT1Q"),
					resource.TestCheckResourceAttr("equinix_fabric_service_token.test", "service_token_connection.0.a_side.0.access_point_selectors.0.link_protocol.0.vlan_tag", "2987"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func testAccFabricZsideVirtualDeviceServiceTokenConfig(serviceTokenName string, serviceTokenDescription string, virtualDeviceUuid string) string {
	return fmt.Sprintf(
		`resource "equinix_fabric_service_token" "test"{
			type = "VC_TOKEN"
			name = "%s"
			description = "%s"
			expiration_date_time = "2025-01-18T06:43:49.981Z"
			service_token_connection {
				type = "EVPL_VC"
				supported_bandwidths = [50, 200, 10000]
				z_side {
					access_point_selectors{
						type = "VD"
						virtual_device{
							type = "EDGE"
							uuid = "%s"
						}
						interface{
							type = "NETWORK"
						}
					}
				}
			}
			notifications {
    			type   = "ALL"
    			emails = ["example@equinix.com", "test1@equinix.com"]
  			}

		}
    `, serviceTokenName, serviceTokenDescription, virtualDeviceUuid)
}

func testAccFabricAsidePortServiceTokenConfig(serviceTokenName string, serviceTokenDescription string, portUuid string) string {
	return fmt.Sprintf(
		`resource "equinix_fabric_service_token" "test"{
			type = "VC_TOKEN"
			name = "%s"
			description = "%s"
			expiration_date_time = "2025-01-18T06:43:49.981Z"
			service_token_connection {
				type = "EVPL_VC"
				bandwidth_limit = 1000
				a_side {
					access_point_selectors{
						type = "COLO"
						port {
							uuid = "%s"
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
    `, serviceTokenName, serviceTokenDescription, portUuid)
}

func CheckServiceTokenDelete(s *terraform.State) error {
	ctx := context.Background()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "equinix_fabric_service_token" {
			continue
		}

		err := service_token.WaitForDeletion(rs.Primary.ID, acceptance.TestAccProvider.Meta(), &schema.ResourceData{}, ctx, 10*time.Minute)
		if err != nil {
			return fmt.Errorf("API call failed while waiting for resource deletion")
		}
	}
	return nil
}
