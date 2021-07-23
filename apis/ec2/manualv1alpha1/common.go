package manualv1alpha1

// BlockDeviceMapping describes a block device mapping.
type BlockDeviceMapping struct {
	// The device name (for example, /dev/sdh or xvdh).
	DeviceName *string `json:"deviceName"`

	// Parameters used to automatically set up EBS volumes when the instance is
	// launched.
	EBS *EBSBlockDevice `json:"ebs"`

	// Suppresses the specified device included in the block device mapping of the
	// AMI.
	NoDevice *string `json:"noDevice"`

	// The virtual device name (ephemeralN). Instance store volumes are numbered
	// starting from 0. An instance type with 2 available instance store volumes
	// can specify mappings for ephemeral0 and ephemeral1. The number of available
	// instance store volumes depends on the instance type. After you connect to
	// the instance, you must mount the volume.
	//
	// NVMe instance store volumes are automatically enumerated and assigned a device
	// name. Including them in your block device mapping has no effect.
	//
	// Constraints: For M3 instances, you must specify instance store volumes in
	// the block device mapping for the instance. When you launch an M3 instance,
	// we ignore any instance store volumes specified in the block device mapping
	// for the AMI.
	VirtualName *string `json:"virtualName"`
}

// CapacityReservationSpecification describes an instance's Capacity Reservation targeting option. You can specify
// only one parameter at a time. If you specify CapacityReservationPreference
// and CapacityReservationTarget, the request fails.
//
// Use the CapacityReservationPreference parameter to configure the instance
// to run as an On-Demand Instance or to run in any open Capacity Reservation
// that has matching attributes (instance type, platform, Availability Zone).
// Use the CapacityReservationTarget parameter to explicitly target a specific
// Capacity Reservation.
type CapacityReservationSpecification struct {
	// Indicates the instance's Capacity Reservation preferences. Possible preferences
	// include:
	//
	//    * open - The instance can run in any open Capacity Reservation that has
	//    matching attributes (instance type, platform, Availability Zone).
	//
	//    * none - The instance avoids running in a Capacity Reservation even if
	//    one is available. The instance runs as an On-Demand Instance.
	CapacityReservationPreference *string `json:"capacityReservationsPreference"`

	// Information about the target Capacity Reservation.
	CapacityReservationTarget *CapacityReservationTarget `json:"capacityReservationTarget"`
}

// CapacityReservationTarget describes a target Capacity Reservation.
type CapacityReservationTarget struct {
	// The ID of the Capacity Reservation.
	CapacityReservationID *string `json:"capacityReservationId"`
}

// CPUOptionsRequest defines the options for the instance. Both the core count and threads per core
// must be specified in the request.
type CPUOptionsRequest struct {
	// The number of CPU cores for the instance.
	CoreCount *int64 `json:"coreCount"`

	// The number of threads per CPU core. To disable multithreading for the instance,
	// specify a value of 1. Otherwise, specify the default value of 2.
	ThreadsPerCore *int64 `type:"threadsPerCore"`
}

// CreditSpecificationRequest describes the credit option for CPU usage of a T2 or T3 instance.
type CreditSpecificationRequest struct {
	// The credit option for CPU usage of a T2 or T3 instance. Valid values are
	// standard and unlimited.
	//
	// CPUCredits is a required field
	CPUCredits *string `json:"cpuCredits"`
}

// EBSBlockDevice describes a block device for an EBS volume.
type EBSBlockDevice struct {
	// Indicates whether the EBS volume is deleted on instance termination. For
	// more information, see Preserving Amazon EBS Volumes on Instance Termination
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#preserving-volumes-on-termination)
	// in the Amazon Elastic Compute Cloud User Guide.
	DeleteOnTermination *bool `json:"deleteOnTermination"`

	// Indicates whether the encryption state of an EBS volume is changed while
	// being restored from a backing snapshot. The effect of setting the encryption
	// state to true depends on the volume origin (new or from a snapshot), starting
	// encryption state, ownership, and whether encryption by default is enabled.
	// For more information, see Amazon EBS Encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-parameters)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// In no case can you remove encryption from an encrypted volume.
	//
	// Encrypted volumes can only be attached to instances that support Amazon EBS
	// encryption. For more information, see Supported Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).
	//
	// This parameter is not returned by .
	Encrypted *bool `json:"encrypted"`

	// The number of I/O operations per second (IOPS) that the volume supports.
	// For io1 volumes, this represents the number of IOPS that are provisioned
	// for the volume. For gp2 volumes, this represents the baseline performance
	// of the volume and the rate at which the volume accumulates I/O credits for
	// bursting. For more information, see Amazon EBS Volume Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// Constraints: Range is 100-16,000 IOPS for gp2 volumes and 100 to 64,000IOPS
	// for io1 volumes in most Regions. Maximum io1 IOPS of 64,000 is guaranteed
	// only on Nitro-based instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances).
	// Other instance families guarantee performance up to 32,000 IOPS. For more
	// information, see Amazon EBS Volume Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// Condition: This parameter is required for requests to create io1 volumes;
	// it is not used in requests to create gp2, st1, sc1, or standard volumes.
	IOps *int64 `json:"iops"`

	// Identifier (key ID, key alias, ID ARN, or alias ARN) for a customer managed
	// CMK under which the EBS volume is encrypted.
	//
	// This parameter is only supported on BlockDeviceMapping objects called by
	// RunInstances (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html),
	// RequestSpotFleet (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RequestSpotFleet.html),
	// and RequestSpotInstances (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RequestSpotInstances.html).
	KmsKeyID *string `json:"kmsKeyId"`

	// The ID of the snapshot.
	SnapshotID *string `json:"snapshotId"`

	// The size of the volume, in GiB.
	//
	// Default: If you're creating the volume from a snapshot and don't specify
	// a volume size, the default is the snapshot size.
	//
	// Constraints: 1-16384 for General Purpose SSD (gp2), 4-16384 for Provisioned
	// IOPS SSD (io1), 500-16384 for Throughput Optimized HDD (st1), 500-16384 for
	// Cold HDD (sc1), and 1-1024 for Magnetic (standard) volumes. If you specify
	// a snapshot, the volume size must be equal to or larger than the snapshot
	// size.
	VolumeSize *int64 `json:"volumeSize"`

	// The volume type. If you set the type to io1, you must also specify the Iops
	// parameter. If you set the type to gp2, st1, sc1, or standard, you must omit
	// the Iops parameter.
	//
	// Default: gp2
	VolumeType *string `json:"volumeType"`
}

// ElasticGpuSpecification is a specification for an Elastic Graphics accelerator.
type ElasticGpuSpecification struct {
	// The type of Elastic Graphics accelerator. For more information about the
	// values to specify for Type, see Elastic Graphics Basics (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html#elastic-graphics-basics),
	// specifically the Elastic Graphics accelerator column, in the Amazon Elastic
	// Compute Cloud User Guide for Windows Instances.
	//
	// Type is a required field
	Type *string `json:"type"`
}

// ElasticInferenceAccelerator describes an elastic inference accelerator.
type ElasticInferenceAccelerator struct {
	// The number of elastic inference accelerators to attach to the instance.
	//
	// Default: 1
	Count *int64 `json:"count,omitempty"`

	// The type of elastic inference accelerator. The possible values are eia1.medium,
	// eia1.large, and eia1.xlarge.
	//
	// Type is a required field
	Type *string `json:"type"`
}

// HibernationOptionsRequest indicates whether your instance is configured for hibernation. This parameter
// is valid only if the instance meets the hibernation prerequisites (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html#hibernating-prerequisites).
// For more information, see Hibernate Your Instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html)
// in the Amazon Elastic Compute Cloud User Guide.
type HibernationOptionsRequest struct {
	// If you set this parameter to true, your instance is enabled for hibernation.
	//
	// Default: false
	Configured *bool `json:"configured,omitempty"`
}

// IamInstanceProfileSpecification describes an IAM instance profile.
type IamInstanceProfileSpecification struct {
	// The Amazon Resource Name (ARN) of the instance profile.
	ARN *string `json:"arn"`

	// The name of the instance profile.
	Name *string `json:"name"`
}

// InstanceMarketOptionsRequest describes the market (purchasing) option for the instances.
type InstanceMarketOptionsRequest struct {
	// The market type.
	MarketType *string `json:"marketType"`

	// The options for Spot Instances.
	SpotOptions *SpotMarketOptions `json:"spotOptions"`
}

// InstanceIPV6Address describes an IPv6 address.
type InstanceIPV6Address struct {
	// The IPv6 address.
	IPV6Address *string `json:"ipv6Address"`
}

// InstanceNetworkInterfaceSpecification describes a network interface.
type InstanceNetworkInterfaceSpecification struct {
	// Indicates whether to assign a public IPv4 address to an instance you launch
	// in a VPC. The public IP address can only be assigned to a network interface
	// for eth0, and can only be assigned to a new network interface, not an existing
	// one. You cannot specify more than one network interface in the request. If
	// launching into a default subnet, the default value is true.
	AssociatePublicIPAddress *bool `json:"associatePublicIpAddress"`

	// If set to true, the interface is deleted when the instance is terminated.
	// You can specify true only if creating a new network interface when launching
	// an instance.
	DeleteOnTermination *bool `json:"deleteOnTermination"`

	// The description of the network interface. Applies only if creating a network
	// interface when launching an instance.
	Description *string `json:"description"`

	// The position of the network interface in the attachment order. A primary
	// network interface has a device index of 0.
	//
	// If you specify a network interface when launching an instance, you must specify
	// the device index.
	DeviceIndex *int64 `json:"deviceIndex"`

	// The IDs of the security groups for the network interface. Applies only if
	// creating a network interface when launching an instance.
	Groups []string `json:"groups"`

	// The type of network interface. To create an Elastic Fabric Adapter (EFA),
	// specify efa. For more information, see Elastic Fabric Adapter (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/efa.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// If you are not creating an EFA, specify interface or omit this parameter.
	//
	// Valid values: interface | efa
	InterfaceType *string `json:"interfaceType"`

	// A number of IPv6 addresses to assign to the network interface. Amazon EC2
	// chooses the IPv6 addresses from the range of the subnet. You cannot specify
	// this option and the option to assign specific IPv6 addresses in the same
	// request. You can specify this option if you've specified a minimum number
	// of instances to launch.
	Ipv6AddressCount *int64 `json:"ipv6AddressCount"`

	// One or more IPv6 addresses to assign to the network interface. You cannot
	// specify this option and the option to assign a number of IPv6 addresses in
	// the same request. You cannot specify this option if you've specified a minimum
	// number of instances to launch.
	IPV6Addresses []InstanceIPV6Address `json:"ipv6Addresses"`

	// The ID of the network interface.
	//
	// If you are creating a Spot Fleet, omit this parameter because you can’t
	// specify a network interface ID in a launch specification.
	NetworkInterfaceID *string `json:"networkInterfaceId"`

	// The private IPv4 address of the network interface. Applies only if creating
	// a network interface when launching an instance. You cannot specify this option
	// if you're launching more than one instance in a RunInstances
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html)
	// request.
	PrivateIPAddress *string `json:"privateIpAddress"`

	// One or more private IPv4 addresses to assign to the network interface. Only
	// one private IPv4 address can be designated as primary. You cannot specify
	// this option if you're launching more than one instance in a RunInstances
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html)
	// request.
	PrivateIPAddresses []PrivateIPAddressSpecification `json:"privateIpAddresses"`

	// The number of secondary private IPv4 addresses. You can't specify this option
	// and specify more than one private IP address using the private IP addresses
	// option. You cannot specify this option if you're launching more than one
	// instance in a RunInstances
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html)
	// request.
	SecondaryPrivateIPAddressCount *int64 `json:"secondaryPrivateIpAddressCount"`

	// The ID of the subnet associated with the network interface. Applies only
	// if creating a network interface when launching an instance.
	SubnetID *string `json:"subnetId"`
}

// LicenseConfigurationRequest describes a license configuration
type LicenseConfigurationRequest struct {
	// Amazon Resource Name (ARN) of the license configuration
	LicenseConfigurationARN *string `json:"licenseConfigurationArn"`
}

// LaunchTemplateSpecification defines the launch template to use.
// You must specify either the launch template ID or launch template
// name in the request, but not both.
type LaunchTemplateSpecification struct {
	// The ID of the launch template.
	LaunchTemplateID *string `json:"launchTemplateId"`

	// The name of the launch template.
	LaunchTemplateName *string `json:"launchTemplateName"`

	// The version number of the launch template.
	//
	// Default: The default version for the launch template.
	Version *string `json:"version"`
}

// Placement describes the placement of an instance.
type Placement struct {
	// The affinity setting for the instance on the Dedicated Host. This parameter
	// is not supported for the ImportInstance
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportInstance.html)
	// command.
	//
	// This parameter is not supported by CreateFleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet).
	Affinity *string `json:"affinity"`

	// The Availability Zone of the instance.
	//
	// If not specified, an Availability Zone will be automatically chosen for you
	// based on the load balancing criteria for the Region.
	//
	// This parameter is not supported by CreateFleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet).
	AvailabilityZone *string `json:"availabilityZone"`

	// The name of the placement group the instance is in.
	GroupName *string `json:"groupName"`

	// The ID of the Dedicated Host on which the instance resides. This parameter
	// is not supported for the ImportInstance
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportInstance.html)
	// command.
	//
	// This parameter is not supported by CreateFleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet).
	HostID *string `json:"hostId"`

	// The ARN of the host resource group in which to launch the instances. If you
	// specify a host resource group ARN, omit the Tenancy parameter or set it to
	// host.
	//
	// This parameter is not supported by CreateFleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet).
	HostResourceGroupARN *string `json:"hostResourceGroupArn"`

	// The number of the partition the instance is in. Valid only if the placement
	// group strategy is set to partition.
	//
	// This parameter is not supported by CreateFleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet).
	PartitionNumber *int64 `json:"partitionNumber"`

	// Reserved for future use.
	//
	// This parameter is not supported by CreateFleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet).
	SpreadDomain *string `json:"spreadDomain"`

	// The tenancy of the instance (if the instance is running in a VPC). An instance
	// with a tenancy of dedicated runs on single-tenant hardware. The host tenancy
	// is not supported for the ImportInstance
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportInstance.html)
	// command.
	//
	// This parameter is not supported by CreateFleet
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet).
	Tenancy *string `json:"tenancy"`
}

// PrivateIPAddressSpecification describes a secondary private IPv4 address for a network interface.
type PrivateIPAddressSpecification struct {
	// Indicates whether the private IPv4 address is the primary private IPv4 address.
	// Only one IPv4 address can be designated as primary.
	Primary *bool `json:"primary"`

	// The private IPv4 addresses.
	PrivateIPAddress *string `json:"privateIPAddress"`
}

// RunInstancesMonitoringEnabled describes the monitoring of an instance.
type RunInstancesMonitoringEnabled struct {
	// Indicates whether detailed monitoring is enabled. Otherwise, basic monitoring
	// is enabled.
	//
	// Enabled is a required field
	Enabled *bool `json:"enabled"`
}

// SpotMarketOptions are the options for Spot Instances.
type SpotMarketOptions struct {
	// The required duration for the Spot Instances (also known as Spot blocks),
	// in minutes. This value must be a multiple of 60 (60, 120, 180, 240, 300,
	// or 360).
	BlockDurationMinutes *int64 `json:"blockDurationMinutes"`

	// The behavior when a Spot Instance is interrupted. The default is terminate.
	InstanceInterruptionBehavior *string `json:"instanceInterruptionBehavior"`

	// The maximum hourly price you're willing to pay for the Spot Instances. The
	// default is the On-Demand price.
	MaxPrice *string `json:"maxPrice"`

	// The Spot Instance request type. For RunInstances (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances),
	// persistent Spot Instance requests are only supported when InstanceInterruptionBehavior
	// is set to either hibernate or stop.
	SpotInstanceType *string `json:"spotInstanceType"`

	// The end date of the request. For a one-time request, the request remains
	// active until all instances launch, the request is canceled, or this date
	// is reached. If the request is persistent, it remains active until it is canceled
	// or this date and time is reached. The default end date is 7 days from the
	// current date.
	// Must be in UTC format (YYYY-MM-DDTHH:MM:SSZ)
	ValidUntil *string `json:"validUntil"`
}

// Tag defines a tag
type Tag struct {
	// Key is the name of the tag.
	Key string `json:"key"`

	// Value is the value of the tag.
	Value string `json:"value"`
}

// TagSpecification defines the tags to apply to a resource when the resource is being created.
type TagSpecification struct {
	// The type of resource to tag. Currently, the resource types that support tagging
	// on creation are: capacity-reservation | client-vpn-endpoint | dedicated-host
	// | fleet | fpga-image | instance | ipv4pool-ec2 | ipv6pool-ec2 | key-pair
	// | launch-template | natgateway | spot-fleet-request | placement-group | snapshot
	// | traffic-mirror-filter | traffic-mirror-session | traffic-mirror-target
	// | transit-gateway | transit-gateway-attachment | transit-gateway-route-table
	// | vpc-endpoint (for interface VPC endpoints)| vpc-endpoint-service (for gateway
	// VPC endpoints) | volume | vpc-flow-log.
	//
	// To tag a resource after it has been created, see CreateTags
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html).
	//
	// +kubebuilder:validation:Enum=capacity-reservation;client-vpn-endpoint;dedicated-host;fleet;fpga-image;instance;ipv4pool-ec2;ipv6pool-ec2;key-pair;launch-template;natgateway;spot-fleet-request;placement-group;snapshot;traffic-mirror-filter;traffic-mirror-session;traffic-mirror-target;transit-gateway;transit-gateway-attachment;transit-gateway-route-table;vpc-endpoint;vpc-endpoint-service;volume;vpc-flow-log
	ResourceType *string `json:"resourceType"`

	// The tags to apply to the resource
	Tags []Tag `json:"tags"`
}
