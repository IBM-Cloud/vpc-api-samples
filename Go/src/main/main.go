package main

import "github.com/vpc-api-samples/Go/src/core"

func main() {
	core.RetrieveToken(core.API_key)

	// Regions and Zones
	core.GetRegions()
	core.GetZones("us-south")

	// VPC
	core.GetVPCs()
	vpcInput := &core.CreateVPCInput{Name: "VPC_NAME", DefaultACL: &core.ResourceByID{ID: "ACL_ID"}}
	core.PostVPC(vpcInput)

	// SSH Keys
	core.GetSSHKeys()
	sskKeyInput := &core.CreateSSHKeyInput{
		Name:      "KEY_NAME",
		PublicKey: "RSA_PUBLIC_KEY",
		Type:      "rsa"}
	core.PostSSHKey(sskKeyInput)

	// Subnet
	core.GetSubnets()

	vpcID := &core.ResourceByID{ID: "VPC_ID"}
	zone := &core.ResourceByName{Name: "ZONE_ID"}
	subnetCountOnly := &core.CreateSubnetCountOnlyTemplateInput{
		Name: "SUBNET_NAME", Vpc: vpcID, Zone: zone,
		TotalIpv4AddressCount: 8, //number of addresses
	}
	core.PostSubnet(subnetCountOnly)

	// security groups
	core.GetSecurityGroups()

	vpcID = &core.ResourceByID{ID: "VPC_ID"}
	remote := &core.Remote{CidrBlock: "0.0.0.0/0"} //cidr in this format
	rule1 := &core.Rule{Direction: "inbound", Remote: remote, Protocol: "all"}
	rule2 := &core.Rule{Direction: "outbound", Remote: remote, Protocol: "all"}
	sg := &core.CreateSecurityGroupInput{Name: "SG_NAME", Rules: []*core.Rule{rule1, rule2}, Vpc: vpcID}
	core.PostSecurityGroup(sg)

	// Images
	core.GetImages()

	// Instance Profiles
	core.GetProfiles()

	// floating ips
	core.GetFloatingIPs()

	// VSI
	core.GetVSIs()

	vsi := &core.CreateVSIInput{
		Name: "VSI_NAME",
		Keys: []*core.ResourceByID{{ID: "SSH_KEY_ID"}},
		PrimaryNetworkInterface: &core.NetworkInterface{
			Name:           "PRIMARY_NETWORK_INTERFACE_NAME",
			PortSpeed:      1000,
			SecurityGroups: []*core.ResourceByID{{ID: "SG_ID"}},
			Subnet:         &core.ResourceByID{ID: "SUBNET_ID"},
		},
		Profile: &core.ResourceByName{Name: "PROFILE_NAME"},
		Vpc:     &core.ResourceByID{ID: "VPC_ID"},
		Zone:    &core.ResourceByName{Name: "ZONE_NAME"},
		Image:   &core.ResourceByID{ID: "IMAGE_ID"},
	}
	core.PostVSI(vsi)

	// reserve a floating ip
	fip := &core.CreateFloatingIPInput{Name: "FIP_NAME", Target: &core.ResourceByID{ID: "SUBNET_ID"}}
	core.PostReserveFloatingIP(fip)
}
