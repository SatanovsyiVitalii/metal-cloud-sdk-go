package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestSwitchDeviceUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var obj SwitchDevice
	err := json.Unmarshal([]byte(_switchDeviceFixture1), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())
	Expect(obj.NetworkEquipmentIdentifierString).To(Equal("UK_RDG_EVR01_00_0001_00A9_01"))
	Expect(obj.NetworkEquipmentQuarantineSubnetGateway).To(Equal("11.16.0.1"))

}

const _switchDeviceFixture1 = "{\"network_equipment_id\":1,\"datacenter_name\":\"uk-reading\",\"network_equipment_driver\":\"hp5900\",\"network_equipment_position\":\"tor\",\"network_equipment_provisioner_type\":\"vpls\",\"network_equipment_identifier_string\":\"UK_RDG_EVR01_00_0001_00A9_01\",\"network_equipment_description\":\"HP Comware Software, Version 7.1.045, Release 2311P06\",\"network_equipment_management_address\":\"10.0.0.0\",\"network_equipment_management_port\":22,\"network_equipment_management_username\":\"sad\",\"network_equipment_quarantine_vlan\":5,\"network_equipment_quarantine_subnet_start\":\"11.16.0.1\",\"network_equipment_quarantine_subnet_end\":\"11.16.0.00\",\"network_equipment_quarantine_subnet_prefix_size\":24,\"network_equipment_quarantine_subnet_gateway\":\"11.16.0.1\",\"network_equipment_primary_wan_ipv4_subnet_pool\":\"11.24.0.2\",\"network_equipment_primary_wan_ipv4_subnet_prefix_size\":22,\"network_equipment_primary_san_subnet_pool\":\"100.64.0.0\",\"network_equipment_primary_san_subnet_prefix_size\":21,\"network_equipment_primary_wan_ipv6_subnet_pool_id\":1,\"network_equipment_primary_wan_ipv6_subnet_cidr\":\"2A02:0CB8:0000:0000:0000:0000:0000:0000/53\",\"network_equipment_cached_updated_timestamp\":\"2020-08-04T20:11:49Z\",\"network_equipment_management_protocol\":\"ssh\",\"chassis_rack_id\":null,\"network_equipment_cache_wrapper_json\":null,\"network_equipment_cache_wrapper_phpserialize\":\"\",\"network_equipment_tor_linked_id\":null,\"network_equipment_uplink_ip_addresses_json\":null,\"network_equipment_management_address_mask\":null,\"network_equipment_management_address_gateway\":null,\"network_equipment_requires_os_install\":0,\"network_equipment_management_mac_address\":\"00:00:00:00:00:00\",\"volume_template_id\":null,\"network_equipment_country\":null,\"network_equipment_city\":null,\"network_equipment_datacenter\":null,\"network_equipment_datacenter_room\":null,\"network_equipment_datacenter_rack\":null,\"network_equipment_rack_position_upper_unit\":null,\"network_equipment_rack_position_lower_unit\":null,\"network_equipment_serial_numbers\":null,\"network_equipment_info_json\":null,\"network_equipment_management_subnet\":null,\"network_equipment_management_subnet_prefix_size\":null,\"network_equipment_management_subnet_start\":null,\"network_equipment_management_subnet_end\":null,\"network_equipment_management_subnet_gateway\":null,\"datacenter_id_parent\":null,\"network_equipment_dhcp_packet_sniffing_is_enabled\":1,\"network_equipment_driver_dump_cached_json\":null,\"network_equipment_tags\":[],\"network_equipment_management_password\":\"zk3enQ4VXROZyJ9\"}"
