type T struct {
AmiLaunchIndex int `json:"AmiLaunchIndex"`
Architecture string `json:"Architecture"`
BlockDeviceMappings []struct    {
DeviceName string `json:"DeviceName"`
Ebs struct        {
AttachTime time.Time `json:"AttachTime"`
DeleteOnTermination bool `json:"DeleteOnTermination"`
Status string `json:"Status"`
VolumeId string `json:"VolumeId"`
} `json:"Ebs"`
} `json:"BlockDeviceMappings"`
BootMode string `json:"BootMode"`
CapacityReservationId interface{} `json:"CapacityReservationId"`
CapacityReservationSpecification struct    {
CapacityReservationPreference string `json:"CapacityReservationPreference"`
CapacityReservationTarget interface{} `json:"CapacityReservationTarget"`
} `json:"CapacityReservationSpecification"`
ClientToken string `json:"ClientToken"`
CpuOptions struct    {
AmdSevSnp string `json:"AmdSevSnp"`
CoreCount int `json:"CoreCount"`
ThreadsPerCore int `json:"ThreadsPerCore"`
} `json:"CpuOptions"`
CurrentInstanceBootMode string `json:"CurrentInstanceBootMode"`
EbsOptimized bool `json:"EbsOptimized"`
ElasticGpuAssociations interface{} `json:"ElasticGpuAssociations"`
ElasticInferenceAcceleratorAssociations interface{} `json:"ElasticInferenceAcceleratorAssociations"`
EnaSupport bool `json:"EnaSupport"`
EnclaveOptions struct    {
Enabled bool `json:"Enabled"`
} `json:"EnclaveOptions"`
HibernationOptions struct    {
Configured bool `json:"Configured"`
} `json:"HibernationOptions"`
Hypervisor string `json:"Hypervisor"`
IamInstanceProfile interface{} `json:"IamInstanceProfile"`
ImageId string `json:"ImageId"`
InstanceId string `json:"InstanceId"`
InstanceLifecycle string `json:"InstanceLifecycle"`
InstanceType string `json:"InstanceType"`
Ipv6Address interface{} `json:"Ipv6Address"`
KernelId interface{} `json:"KernelId"`
KeyName string `json:"KeyName"`
LaunchTime time.Time `json:"LaunchTime"`
Licenses interface{} `json:"Licenses"`
MaintenanceOptions struct    {
AutoRecovery string `json:"AutoRecovery"`
} `json:"MaintenanceOptions"`
MetadataOptions struct    {
HttpEndpoint string `json:"HttpEndpoint"`
HttpProtocolIpv6 string `json:"HttpProtocolIpv6"`
HttpPutResponseHopLimit int `json:"HttpPutResponseHopLimit"`
HttpTokens string `json:"HttpTokens"`
InstanceMetadataTags string `json:"InstanceMetadataTags"`
State string `json:"State"`
} `json:"MetadataOptions"`
Monitoring struct    {
State string `json:"State"`
} `json:"Monitoring"`
NetworkInterfaces []struct    {
Association struct        {
CarrierIp interface{} `json:"CarrierIp"`
CustomerOwnedIp interface{} `json:"CustomerOwnedIp"`
IpOwnerId string `json:"IpOwnerId"`
PublicDnsName string `json:"PublicDnsName"`
PublicIp string `json:"PublicIp"`
} `json:"Association"`
Attachment struct        {
AttachTime time.Time `json:"AttachTime"`
AttachmentId string `json:"AttachmentId"`
DeleteOnTermination bool `json:"DeleteOnTermination"`
DeviceIndex int `json:"DeviceIndex"`
EnaSrdSpecification interface{} `json:"EnaSrdSpecification"`
NetworkCardIndex int `json:"NetworkCardIndex"`
Status string `json:"Status"`
} `json:"Attachment"`
Description string `json:"Description"`
Groups []struct        {
GroupId string `json:"GroupId"`
GroupName string `json:"GroupName"`
} `json:"Groups"`
InterfaceType string `json:"InterfaceType"`
Ipv4Prefixes interface{} `json:"Ipv4Prefixes"`
Ipv6Addresses []interface{} `json:"Ipv6Addresses"`
Ipv6Prefixes interface{} `json:"Ipv6Prefixes"`
MacAddress string `json:"MacAddress"`
NetworkInterfaceId string `json:"NetworkInterfaceId"`
OwnerId string `json:"OwnerId"`
PrivateDnsName interface{} `json:"PrivateDnsName"`
PrivateIpAddress string `json:"PrivateIpAddress"`
PrivateIpAddresses []struct        {
Association struct            {
CarrierIp interface{} `json:"CarrierIp"`
CustomerOwnedIp interface{} `json:"CustomerOwnedIp"`
IpOwnerId string `json:"IpOwnerId"`
PublicDnsName string `json:"PublicDnsName"`
PublicIp string `json:"PublicIp"`
} `json:"Association"`
Primary bool `json:"Primary"`
PrivateDnsName interface{} `json:"PrivateDnsName"`
PrivateIpAddress string `json:"PrivateIpAddress"`
} `json:"PrivateIpAddresses"`
SourceDestCheck bool `json:"SourceDestCheck"`
Status string `json:"Status"`
SubnetId string `json:"SubnetId"`
VpcId string `json:"VpcId"`
} `json:"NetworkInterfaces"`
OutpostArn interface{} `json:"OutpostArn"`
Placement struct    {
Affinity interface{} `json:"Affinity"`
AvailabilityZone string `json:"AvailabilityZone"`
GroupId interface{} `json:"GroupId"`
GroupName string `json:"GroupName"`
HostId interface{} `json:"HostId"`
HostResourceGroupArn interface{} `json:"HostResourceGroupArn"`
PartitionNumber interface{} `json:"PartitionNumber"`
SpreadDomain interface{} `json:"SpreadDomain"`
Tenancy string `json:"Tenancy"`
} `json:"Placement"`
Platform string `json:"Platform"`
PlatformDetails string `json:"PlatformDetails"`
PrivateDnsName string `json:"PrivateDnsName"`
PrivateDnsNameOptions struct    {
EnableResourceNameDnsAAAARecord bool `json:"EnableResourceNameDnsAAAARecord"`
EnableResourceNameDnsARecord bool `json:"EnableResourceNameDnsARecord"`
HostnameType string `json:"HostnameType"`
} `json:"PrivateDnsNameOptions"`
PrivateIpAddress string `json:"PrivateIpAddress"`
ProductCodes []interface{} `json:"ProductCodes"`
PublicDnsName string `json:"PublicDnsName"`
PublicIpAddress string `json:"PublicIpAddress"`
RamdiskId interface{} `json:"RamdiskId"`
RootDeviceName string `json:"RootDeviceName"`
RootDeviceType string `json:"RootDeviceType"`
SecurityGroups []struct    {
GroupId string `json:"GroupId"`
GroupName string `json:"GroupName"`
} `json:"SecurityGroups"`
SourceDestCheck bool `json:"SourceDestCheck"`
SpotInstanceRequestId interface{} `json:"SpotInstanceRequestId"`
SriovNetSupport interface{} `json:"SriovNetSupport"`
State struct    {
Code int `json:"Code"`
Name string `json:"Name"`
} `json:"State"`
StateReason interface{} `json:"StateReason"`
StateTransitionReason string `json:"StateTransitionReason"`
SubnetId string `json:"SubnetId"`
Tags []struct    {
Key string `json:"Key"`
Value string `json:"Value"`
} `json:"Tags"`
TpmSupport interface{} `json:"TpmSupport"`
UsageOperation string `json:"UsageOperation"`
UsageOperationUpdateTime time.Time `json:"UsageOperationUpdateTime"`
VirtualizationType string `json:"VirtualizationType"`
VpcId string `json:"VpcId"`
}
