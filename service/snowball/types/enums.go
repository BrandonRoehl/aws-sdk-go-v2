// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ClusterState string

// Enum values for ClusterState
const (
	ClusterStateAwaiting_quorum ClusterState = "AwaitingQuorum"
	ClusterStatePending         ClusterState = "Pending"
	ClusterStateIn_use          ClusterState = "InUse"
	ClusterStateComplete        ClusterState = "Complete"
	ClusterStateCancelled       ClusterState = "Cancelled"
)

type JobState string

// Enum values for JobState
const (
	JobStateNew                       JobState = "New"
	JobStatePreparing_appliance       JobState = "PreparingAppliance"
	JobStatePreparing_shipment        JobState = "PreparingShipment"
	JobStateIn_transit_to_customer    JobState = "InTransitToCustomer"
	JobStateWith_customer             JobState = "WithCustomer"
	JobStateIn_transit_to_aws         JobState = "InTransitToAWS"
	JobStateWith_aws_sorting_facility JobState = "WithAWSSortingFacility"
	JobStateWith_aws                  JobState = "WithAWS"
	JobStateIn_progress               JobState = "InProgress"
	JobStateComplete                  JobState = "Complete"
	JobStateCancelled                 JobState = "Cancelled"
	JobStateListing                   JobState = "Listing"
	JobStatePending                   JobState = "Pending"
)

type JobType string

// Enum values for JobType
const (
	JobTypeImport    JobType = "IMPORT"
	JobTypeExport    JobType = "EXPORT"
	JobTypeLocal_use JobType = "LOCAL_USE"
)

type ShippingOption string

// Enum values for ShippingOption
const (
	ShippingOptionSecond_day ShippingOption = "SECOND_DAY"
	ShippingOptionNext_day   ShippingOption = "NEXT_DAY"
	ShippingOptionExpress    ShippingOption = "EXPRESS"
	ShippingOptionStandard   ShippingOption = "STANDARD"
)

type SnowballCapacity string

// Enum values for SnowballCapacity
const (
	SnowballCapacityT50           SnowballCapacity = "T50"
	SnowballCapacityT80           SnowballCapacity = "T80"
	SnowballCapacityT100          SnowballCapacity = "T100"
	SnowballCapacityT42           SnowballCapacity = "T42"
	SnowballCapacityT98           SnowballCapacity = "T98"
	SnowballCapacityT8            SnowballCapacity = "T8"
	SnowballCapacityNo_preference SnowballCapacity = "NoPreference"
)

type SnowballType string

// Enum values for SnowballType
const (
	SnowballTypeStandard SnowballType = "STANDARD"
	SnowballTypeEdge     SnowballType = "EDGE"
	SnowballTypeEdge_c   SnowballType = "EDGE_C"
	SnowballTypeEdge_cg  SnowballType = "EDGE_CG"
	SnowballTypeEdge_s   SnowballType = "EDGE_S"
	SnowballTypeSnc1_hdd SnowballType = "SNC1_HDD"
)
