// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package emr provides a client for Amazon Elastic MapReduce.
package emr

import (
	"net/http"
	"time"

	"github.com/hashicorp/aws-sdk-go/aws"
	"github.com/hashicorp/aws-sdk-go/gen/endpoints"
)

// EMR is a client for Amazon Elastic MapReduce.
type EMR struct {
	client *aws.JSONClient
}

// New returns a new EMR client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *EMR {
	if client == nil {
		client = http.DefaultClient
	}

	endpoint, service, region := endpoints.Lookup("elasticmapreduce", region)

	return &EMR{
		client: &aws.JSONClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			}, Client: client,
			Endpoint:     endpoint,
			JSONVersion:  "1.1",
			TargetPrefix: "ElasticMapReduce",
		},
	}
}

// AddInstanceGroups addInstanceGroups adds an instance group to a running
// cluster.
func (c *EMR) AddInstanceGroups(req *AddInstanceGroupsInput) (resp *AddInstanceGroupsOutput, err error) {
	resp = &AddInstanceGroupsOutput{}
	err = c.client.Do("AddInstanceGroups", "POST", "/", req, resp)
	return
}

// AddJobFlowSteps addJobFlowSteps adds new steps to a running job flow. A
// maximum of 256 steps are allowed in each job flow. If your job flow is
// long-running (such as a Hive data warehouse) or complex, you may require
// more than 256 steps to process your data. You can bypass the 256-step
// limitation in various ways, including using the SSH shell to connect to
// the master node and submitting queries directly to the software running
// on the master node, such as Hive and Hadoop. For more information on how
// to do this, go to Add More than 256 Steps to a Job Flow in the Amazon
// Elastic MapReduce Developer's Guide A step specifies the location of a
// JAR file stored either on the master node of the job flow or in Amazon
// S3. Each step is performed by the main function of the main class of the
// JAR file. The main class can be specified either in the manifest of the
// JAR or by using the MainFunction parameter of the step. Elastic
// MapReduce executes each step in the order listed. For a step to be
// considered complete, the main function must exit with a zero exit code
// and all Hadoop jobs started while the step was running must have
// completed and run successfully. You can only add steps to a job flow
// that is in one of the following states: or
func (c *EMR) AddJobFlowSteps(req *AddJobFlowStepsInput) (resp *AddJobFlowStepsOutput, err error) {
	resp = &AddJobFlowStepsOutput{}
	err = c.client.Do("AddJobFlowSteps", "POST", "/", req, resp)
	return
}

// AddTags adds tags to an Amazon EMR resource. Tags make it easier to
// associate clusters in various ways, such as grouping clusters to track
// your Amazon EMR resource allocation costs. For more information, see
// Tagging Amazon EMR Resources .
func (c *EMR) AddTags(req *AddTagsInput) (resp *AddTagsOutput, err error) {
	resp = &AddTagsOutput{}
	err = c.client.Do("AddTags", "POST", "/", req, resp)
	return
}

// DescribeCluster provides cluster-level details including status,
// hardware and software configuration, VPC settings, and so on. For
// information about the cluster steps, see ListSteps
func (c *EMR) DescribeCluster(req *DescribeClusterInput) (resp *DescribeClusterOutput, err error) {
	resp = &DescribeClusterOutput{}
	err = c.client.Do("DescribeCluster", "POST", "/", req, resp)
	return
}

// DescribeJobFlows this API is deprecated and will eventually be removed.
// We recommend you use ListClusters , DescribeCluster , ListSteps ,
// ListInstanceGroups and ListBootstrapActions instead. DescribeJobFlows
// returns a list of job flows that match all of the supplied parameters.
// The parameters can include a list of job flow IDs, job flow states, and
// restrictions on job flow creation date and time. Regardless of supplied
// parameters, only job flows created within the last two months are
// returned. If no parameters are supplied, then job flows matching either
// of the following criteria are returned: Job flows created and completed
// in the last two weeks Job flows created within the last two months that
// are in one of the following states: , , , Amazon Elastic MapReduce can
// return a maximum of 512 job flow descriptions.
func (c *EMR) DescribeJobFlows(req *DescribeJobFlowsInput) (resp *DescribeJobFlowsOutput, err error) {
	resp = &DescribeJobFlowsOutput{}
	err = c.client.Do("DescribeJobFlows", "POST", "/", req, resp)
	return
}

// DescribeStep is undocumented.
func (c *EMR) DescribeStep(req *DescribeStepInput) (resp *DescribeStepOutput, err error) {
	resp = &DescribeStepOutput{}
	err = c.client.Do("DescribeStep", "POST", "/", req, resp)
	return
}

// ListBootstrapActions provides information about the bootstrap actions
// associated with a cluster.
func (c *EMR) ListBootstrapActions(req *ListBootstrapActionsInput) (resp *ListBootstrapActionsOutput, err error) {
	resp = &ListBootstrapActionsOutput{}
	err = c.client.Do("ListBootstrapActions", "POST", "/", req, resp)
	return
}

// ListClusters provides the status of all clusters visible to this AWS
// account. Allows you to filter the list of clusters based on certain
// criteria; for example, filtering by cluster creation date and time or by
// status. This call returns a maximum of 50 clusters per call, but returns
// a marker to track the paging of the cluster list across multiple
// ListClusters calls.
func (c *EMR) ListClusters(req *ListClustersInput) (resp *ListClustersOutput, err error) {
	resp = &ListClustersOutput{}
	err = c.client.Do("ListClusters", "POST", "/", req, resp)
	return
}

// ListInstanceGroups provides all available details about the instance
// groups in a cluster.
func (c *EMR) ListInstanceGroups(req *ListInstanceGroupsInput) (resp *ListInstanceGroupsOutput, err error) {
	resp = &ListInstanceGroupsOutput{}
	err = c.client.Do("ListInstanceGroups", "POST", "/", req, resp)
	return
}

// ListInstances provides information about the cluster instances that
// Amazon EMR provisions on behalf of a user when it creates the cluster.
// For example, this operation indicates when the EC2 instances reach the
// Ready state, when instances become available to Amazon EMR to use for
// jobs, and the IP addresses for cluster instances, etc.
func (c *EMR) ListInstances(req *ListInstancesInput) (resp *ListInstancesOutput, err error) {
	resp = &ListInstancesOutput{}
	err = c.client.Do("ListInstances", "POST", "/", req, resp)
	return
}

// ListSteps is undocumented.
func (c *EMR) ListSteps(req *ListStepsInput) (resp *ListStepsOutput, err error) {
	resp = &ListStepsOutput{}
	err = c.client.Do("ListSteps", "POST", "/", req, resp)
	return
}

// ModifyInstanceGroups modifyInstanceGroups modifies the number of nodes
// and configuration settings of an instance group. The input parameters
// include the new target instance count for the group and the instance
// group ID. The call will either succeed or fail atomically.
func (c *EMR) ModifyInstanceGroups(req *ModifyInstanceGroupsInput) (err error) {
	// NRE
	err = c.client.Do("ModifyInstanceGroups", "POST", "/", req, nil)
	return
}

// RemoveTags removes tags from an Amazon EMR resource. Tags make it easier
// to associate clusters in various ways, such as grouping clusters to
// track your Amazon EMR resource allocation costs. For more information,
// see Tagging Amazon EMR Resources . The following example removes the
// stack tag with value Prod from a cluster:
func (c *EMR) RemoveTags(req *RemoveTagsInput) (resp *RemoveTagsOutput, err error) {
	resp = &RemoveTagsOutput{}
	err = c.client.Do("RemoveTags", "POST", "/", req, resp)
	return
}

// RunJobFlow runJobFlow creates and starts running a new job flow. The job
// flow will run the steps specified. Once the job flow completes, the
// cluster is stopped and the partition is lost. To prevent loss of data,
// configure the last step of the job flow to store results in Amazon S3.
// If the JobFlowInstancesConfig KeepJobFlowAliveWhenNoSteps parameter is
// set to , the job flow will transition to the state rather than shutting
// down once the steps have completed. For additional protection, you can
// set the JobFlowInstancesConfig TerminationProtected parameter to to lock
// the job flow and prevent it from being terminated by API call, user
// intervention, or in the event of a job flow error. A maximum of 256
// steps are allowed in each job flow. If your job flow is long-running
// (such as a Hive data warehouse) or complex, you may require more than
// 256 steps to process your data. You can bypass the 256-step limitation
// in various ways, including using the SSH shell to connect to the master
// node and submitting queries directly to the software running on the
// master node, such as Hive and Hadoop. For more information on how to do
// this, go to Add More than 256 Steps to a Job Flow in the Amazon Elastic
// MapReduce Developer's Guide For long running job flows, we recommend
// that you periodically store your results.
func (c *EMR) RunJobFlow(req *RunJobFlowInput) (resp *RunJobFlowOutput, err error) {
	resp = &RunJobFlowOutput{}
	err = c.client.Do("RunJobFlow", "POST", "/", req, resp)
	return
}

// SetTerminationProtection setTerminationProtection locks a job flow so
// the Amazon EC2 instances in the cluster cannot be terminated by user
// intervention, an API call, or in the event of a job-flow error. The
// cluster still terminates upon successful completion of the job flow.
// Calling SetTerminationProtection on a job flow is analogous to calling
// the Amazon EC2 DisableAPITermination API on all of the EC2 instances in
// a cluster. SetTerminationProtection is used to prevent accidental
// termination of a job flow and to ensure that in the event of an error,
// the instances will persist so you can recover any data stored in their
// ephemeral instance storage. To terminate a job flow that has been locked
// by setting SetTerminationProtection to true , you must first unlock the
// job flow by a subsequent call to SetTerminationProtection in which you
// set the value to false . For more information, go to Protecting a Job
// Flow from Termination in the Amazon Elastic MapReduce Developer's Guide.
func (c *EMR) SetTerminationProtection(req *SetTerminationProtectionInput) (err error) {
	// NRE
	err = c.client.Do("SetTerminationProtection", "POST", "/", req, nil)
	return
}

// SetVisibleToAllUsers sets whether all AWS Identity and Access Management
// users under your account can access the specified job flows. This action
// works on running job flows. You can also set the visibility of a job
// flow when you launch it using the VisibleToAllUsers parameter of
// RunJobFlow . The SetVisibleToAllUsers action can be called only by an
// IAM user who created the job flow or the AWS account that owns the job
// flow.
func (c *EMR) SetVisibleToAllUsers(req *SetVisibleToAllUsersInput) (err error) {
	// NRE
	err = c.client.Do("SetVisibleToAllUsers", "POST", "/", req, nil)
	return
}

// TerminateJobFlows terminateJobFlows shuts a list of job flows down. When
// a job flow is shut down, any step not yet completed is canceled and the
// EC2 instances on which the job flow is running are stopped. Any log
// files not already saved are uploaded to Amazon S3 if a LogUri was
// specified when the job flow was created. The maximum number of JobFlows
// allowed is 10. The call to TerminateJobFlows is asynchronous. Depending
// on the configuration of the job flow, it may take up to 5-20 minutes for
// the job flow to completely terminate and release allocated resources,
// such as Amazon EC2 instances.
func (c *EMR) TerminateJobFlows(req *TerminateJobFlowsInput) (err error) {
	// NRE
	err = c.client.Do("TerminateJobFlows", "POST", "/", req, nil)
	return
}

// Possible values for EMR.
const (
	ActionOnFailureCancelAndWait    = "CANCEL_AND_WAIT"
	ActionOnFailureContinue         = "CONTINUE"
	ActionOnFailureTerminateCluster = "TERMINATE_CLUSTER"
	ActionOnFailureTerminateJobFlow = "TERMINATE_JOB_FLOW"
)

// AddInstanceGroupsInput is undocumented.
type AddInstanceGroupsInput struct {
	InstanceGroups []InstanceGroupConfig `json:"InstanceGroups"`
	JobFlowID      aws.StringValue       `json:"JobFlowId"`
}

// AddInstanceGroupsOutput is undocumented.
type AddInstanceGroupsOutput struct {
	InstanceGroupIDs []string        `json:"InstanceGroupIds,omitempty"`
	JobFlowID        aws.StringValue `json:"JobFlowId,omitempty"`
}

// AddJobFlowStepsInput is undocumented.
type AddJobFlowStepsInput struct {
	JobFlowID aws.StringValue `json:"JobFlowId"`
	Steps     []StepConfig    `json:"Steps"`
}

// AddJobFlowStepsOutput is undocumented.
type AddJobFlowStepsOutput struct {
	StepIDs []string `json:"StepIds,omitempty"`
}

// AddTagsInput is undocumented.
type AddTagsInput struct {
	ResourceID aws.StringValue `json:"ResourceId"`
	Tags       []Tag           `json:"Tags"`
}

// AddTagsOutput is undocumented.
type AddTagsOutput struct {
}

// Application is undocumented.
type Application struct {
	AdditionalInfo map[string]string `json:"AdditionalInfo,omitempty"`
	Args           []string          `json:"Args,omitempty"`
	Name           aws.StringValue   `json:"Name,omitempty"`
	Version        aws.StringValue   `json:"Version,omitempty"`
}

// BootstrapActionConfig is undocumented.
type BootstrapActionConfig struct {
	Name                  aws.StringValue              `json:"Name"`
	ScriptBootstrapAction *ScriptBootstrapActionConfig `json:"ScriptBootstrapAction"`
}

// BootstrapActionDetail is undocumented.
type BootstrapActionDetail struct {
	BootstrapActionConfig *BootstrapActionConfig `json:"BootstrapActionConfig,omitempty"`
}

// Cluster is undocumented.
type Cluster struct {
	Applications            []Application          `json:"Applications,omitempty"`
	AutoTerminate           aws.BooleanValue       `json:"AutoTerminate,omitempty"`
	EC2InstanceAttributes   *EC2InstanceAttributes `json:"Ec2InstanceAttributes,omitempty"`
	ID                      aws.StringValue        `json:"Id,omitempty"`
	LogURI                  aws.StringValue        `json:"LogUri,omitempty"`
	MasterPublicDNSName     aws.StringValue        `json:"MasterPublicDnsName,omitempty"`
	Name                    aws.StringValue        `json:"Name,omitempty"`
	NormalizedInstanceHours aws.IntegerValue       `json:"NormalizedInstanceHours,omitempty"`
	RequestedAMIVersion     aws.StringValue        `json:"RequestedAmiVersion,omitempty"`
	RunningAMIVersion       aws.StringValue        `json:"RunningAmiVersion,omitempty"`
	ServiceRole             aws.StringValue        `json:"ServiceRole,omitempty"`
	Status                  *ClusterStatus         `json:"Status,omitempty"`
	Tags                    []Tag                  `json:"Tags,omitempty"`
	TerminationProtected    aws.BooleanValue       `json:"TerminationProtected,omitempty"`
	VisibleToAllUsers       aws.BooleanValue       `json:"VisibleToAllUsers,omitempty"`
}

// Possible values for EMR.
const (
	ClusterStateBootstrapping        = "BOOTSTRAPPING"
	ClusterStateRunning              = "RUNNING"
	ClusterStateStarting             = "STARTING"
	ClusterStateTerminated           = "TERMINATED"
	ClusterStateTerminatedWithErrors = "TERMINATED_WITH_ERRORS"
	ClusterStateTerminating          = "TERMINATING"
	ClusterStateWaiting              = "WAITING"
)

// ClusterStateChangeReason is undocumented.
type ClusterStateChangeReason struct {
	Code    aws.StringValue `json:"Code,omitempty"`
	Message aws.StringValue `json:"Message,omitempty"`
}

// Possible values for EMR.
const (
	ClusterStateChangeReasonCodeAllStepsCompleted = "ALL_STEPS_COMPLETED"
	ClusterStateChangeReasonCodeBootstrapFailure  = "BOOTSTRAP_FAILURE"
	ClusterStateChangeReasonCodeInstanceFailure   = "INSTANCE_FAILURE"
	ClusterStateChangeReasonCodeInternalError     = "INTERNAL_ERROR"
	ClusterStateChangeReasonCodeStepFailure       = "STEP_FAILURE"
	ClusterStateChangeReasonCodeUserRequest       = "USER_REQUEST"
	ClusterStateChangeReasonCodeValidationError   = "VALIDATION_ERROR"
)

// ClusterStatus is undocumented.
type ClusterStatus struct {
	State             aws.StringValue           `json:"State,omitempty"`
	StateChangeReason *ClusterStateChangeReason `json:"StateChangeReason,omitempty"`
	Timeline          *ClusterTimeline          `json:"Timeline,omitempty"`
}

// ClusterSummary is undocumented.
type ClusterSummary struct {
	ID                      aws.StringValue  `json:"Id,omitempty"`
	Name                    aws.StringValue  `json:"Name,omitempty"`
	NormalizedInstanceHours aws.IntegerValue `json:"NormalizedInstanceHours,omitempty"`
	Status                  *ClusterStatus   `json:"Status,omitempty"`
}

// ClusterTimeline is undocumented.
type ClusterTimeline struct {
	CreationDateTime *aws.UnixTimestamp `json:"CreationDateTime,omitempty"`
	EndDateTime      *aws.UnixTimestamp `json:"EndDateTime,omitempty"`
	ReadyDateTime    *aws.UnixTimestamp `json:"ReadyDateTime,omitempty"`
}

// Command is undocumented.
type Command struct {
	Args       []string        `json:"Args,omitempty"`
	Name       aws.StringValue `json:"Name,omitempty"`
	ScriptPath aws.StringValue `json:"ScriptPath,omitempty"`
}

// DescribeClusterInput is undocumented.
type DescribeClusterInput struct {
	ClusterID aws.StringValue `json:"ClusterId"`
}

// DescribeClusterOutput is undocumented.
type DescribeClusterOutput struct {
	Cluster *Cluster `json:"Cluster,omitempty"`
}

// DescribeJobFlowsInput is undocumented.
type DescribeJobFlowsInput struct {
	CreatedAfter  *aws.UnixTimestamp `json:"CreatedAfter,omitempty"`
	CreatedBefore *aws.UnixTimestamp `json:"CreatedBefore,omitempty"`
	JobFlowIDs    []string           `json:"JobFlowIds,omitempty"`
	JobFlowStates []string           `json:"JobFlowStates,omitempty"`
}

// DescribeJobFlowsOutput is undocumented.
type DescribeJobFlowsOutput struct {
	JobFlows []JobFlowDetail `json:"JobFlows,omitempty"`
}

// DescribeStepInput is undocumented.
type DescribeStepInput struct {
	ClusterID aws.StringValue `json:"ClusterId"`
	StepID    aws.StringValue `json:"StepId"`
}

// DescribeStepOutput is undocumented.
type DescribeStepOutput struct {
	Step *Step `json:"Step,omitempty"`
}

// EC2InstanceAttributes is undocumented.
type EC2InstanceAttributes struct {
	AdditionalMasterSecurityGroups []string        `json:"AdditionalMasterSecurityGroups,omitempty"`
	AdditionalSlaveSecurityGroups  []string        `json:"AdditionalSlaveSecurityGroups,omitempty"`
	EC2AvailabilityZone            aws.StringValue `json:"Ec2AvailabilityZone,omitempty"`
	EC2KeyName                     aws.StringValue `json:"Ec2KeyName,omitempty"`
	EC2SubnetID                    aws.StringValue `json:"Ec2SubnetId,omitempty"`
	EmrManagedMasterSecurityGroup  aws.StringValue `json:"EmrManagedMasterSecurityGroup,omitempty"`
	EmrManagedSlaveSecurityGroup   aws.StringValue `json:"EmrManagedSlaveSecurityGroup,omitempty"`
	IAMInstanceProfile             aws.StringValue `json:"IamInstanceProfile,omitempty"`
}

// HadoopJARStepConfig is undocumented.
type HadoopJARStepConfig struct {
	Args       []string        `json:"Args,omitempty"`
	JAR        aws.StringValue `json:"Jar"`
	MainClass  aws.StringValue `json:"MainClass,omitempty"`
	Properties []KeyValue      `json:"Properties,omitempty"`
}

// HadoopStepConfig is undocumented.
type HadoopStepConfig struct {
	Args       []string          `json:"Args,omitempty"`
	JAR        aws.StringValue   `json:"Jar,omitempty"`
	MainClass  aws.StringValue   `json:"MainClass,omitempty"`
	Properties map[string]string `json:"Properties,omitempty"`
}

// Instance is undocumented.
type Instance struct {
	EC2InstanceID    aws.StringValue `json:"Ec2InstanceId,omitempty"`
	ID               aws.StringValue `json:"Id,omitempty"`
	PrivateDNSName   aws.StringValue `json:"PrivateDnsName,omitempty"`
	PrivateIPAddress aws.StringValue `json:"PrivateIpAddress,omitempty"`
	PublicDNSName    aws.StringValue `json:"PublicDnsName,omitempty"`
	PublicIPAddress  aws.StringValue `json:"PublicIpAddress,omitempty"`
	Status           *InstanceStatus `json:"Status,omitempty"`
}

// InstanceGroup is undocumented.
type InstanceGroup struct {
	BidPrice               aws.StringValue      `json:"BidPrice,omitempty"`
	ID                     aws.StringValue      `json:"Id,omitempty"`
	InstanceGroupType      aws.StringValue      `json:"InstanceGroupType,omitempty"`
	InstanceType           aws.StringValue      `json:"InstanceType,omitempty"`
	Market                 aws.StringValue      `json:"Market,omitempty"`
	Name                   aws.StringValue      `json:"Name,omitempty"`
	RequestedInstanceCount aws.IntegerValue     `json:"RequestedInstanceCount,omitempty"`
	RunningInstanceCount   aws.IntegerValue     `json:"RunningInstanceCount,omitempty"`
	Status                 *InstanceGroupStatus `json:"Status,omitempty"`
}

// InstanceGroupConfig is undocumented.
type InstanceGroupConfig struct {
	BidPrice      aws.StringValue  `json:"BidPrice,omitempty"`
	InstanceCount aws.IntegerValue `json:"InstanceCount"`
	InstanceRole  aws.StringValue  `json:"InstanceRole"`
	InstanceType  aws.StringValue  `json:"InstanceType"`
	Market        aws.StringValue  `json:"Market,omitempty"`
	Name          aws.StringValue  `json:"Name,omitempty"`
}

// InstanceGroupDetail is undocumented.
type InstanceGroupDetail struct {
	BidPrice              aws.StringValue    `json:"BidPrice,omitempty"`
	CreationDateTime      *aws.UnixTimestamp `json:"CreationDateTime"`
	EndDateTime           *aws.UnixTimestamp `json:"EndDateTime,omitempty"`
	InstanceGroupID       aws.StringValue    `json:"InstanceGroupId,omitempty"`
	InstanceRequestCount  aws.IntegerValue   `json:"InstanceRequestCount"`
	InstanceRole          aws.StringValue    `json:"InstanceRole"`
	InstanceRunningCount  aws.IntegerValue   `json:"InstanceRunningCount"`
	InstanceType          aws.StringValue    `json:"InstanceType"`
	LastStateChangeReason aws.StringValue    `json:"LastStateChangeReason,omitempty"`
	Market                aws.StringValue    `json:"Market"`
	Name                  aws.StringValue    `json:"Name,omitempty"`
	ReadyDateTime         *aws.UnixTimestamp `json:"ReadyDateTime,omitempty"`
	StartDateTime         *aws.UnixTimestamp `json:"StartDateTime,omitempty"`
	State                 aws.StringValue    `json:"State"`
}

// InstanceGroupModifyConfig is undocumented.
type InstanceGroupModifyConfig struct {
	EC2InstanceIDsToTerminate []string         `json:"EC2InstanceIdsToTerminate,omitempty"`
	InstanceCount             aws.IntegerValue `json:"InstanceCount,omitempty"`
	InstanceGroupID           aws.StringValue  `json:"InstanceGroupId"`
}

// Possible values for EMR.
const (
	InstanceGroupStateArrested      = "ARRESTED"
	InstanceGroupStateBootstrapping = "BOOTSTRAPPING"
	InstanceGroupStateEnded         = "ENDED"
	InstanceGroupStateProvisioning  = "PROVISIONING"
	InstanceGroupStateResizing      = "RESIZING"
	InstanceGroupStateRunning       = "RUNNING"
	InstanceGroupStateShuttingDown  = "SHUTTING_DOWN"
	InstanceGroupStateSuspended     = "SUSPENDED"
	InstanceGroupStateTerminated    = "TERMINATED"
	InstanceGroupStateTerminating   = "TERMINATING"
)

// InstanceGroupStateChangeReason is undocumented.
type InstanceGroupStateChangeReason struct {
	Code    aws.StringValue `json:"Code,omitempty"`
	Message aws.StringValue `json:"Message,omitempty"`
}

// Possible values for EMR.
const (
	InstanceGroupStateChangeReasonCodeClusterTerminated = "CLUSTER_TERMINATED"
	InstanceGroupStateChangeReasonCodeInstanceFailure   = "INSTANCE_FAILURE"
	InstanceGroupStateChangeReasonCodeInternalError     = "INTERNAL_ERROR"
	InstanceGroupStateChangeReasonCodeValidationError   = "VALIDATION_ERROR"
)

// InstanceGroupStatus is undocumented.
type InstanceGroupStatus struct {
	State             aws.StringValue                 `json:"State,omitempty"`
	StateChangeReason *InstanceGroupStateChangeReason `json:"StateChangeReason,omitempty"`
	Timeline          *InstanceGroupTimeline          `json:"Timeline,omitempty"`
}

// InstanceGroupTimeline is undocumented.
type InstanceGroupTimeline struct {
	CreationDateTime *aws.UnixTimestamp `json:"CreationDateTime,omitempty"`
	EndDateTime      *aws.UnixTimestamp `json:"EndDateTime,omitempty"`
	ReadyDateTime    *aws.UnixTimestamp `json:"ReadyDateTime,omitempty"`
}

// Possible values for EMR.
const (
	InstanceGroupTypeCore   = "CORE"
	InstanceGroupTypeMaster = "MASTER"
	InstanceGroupTypeTask   = "TASK"
)

// Possible values for EMR.
const (
	InstanceRoleTypeCore   = "CORE"
	InstanceRoleTypeMaster = "MASTER"
	InstanceRoleTypeTask   = "TASK"
)

// Possible values for EMR.
const (
	InstanceStateAwaitingFulfillment = "AWAITING_FULFILLMENT"
	InstanceStateBootstrapping       = "BOOTSTRAPPING"
	InstanceStateProvisioning        = "PROVISIONING"
	InstanceStateRunning             = "RUNNING"
	InstanceStateTerminated          = "TERMINATED"
)

// InstanceStateChangeReason is undocumented.
type InstanceStateChangeReason struct {
	Code    aws.StringValue `json:"Code,omitempty"`
	Message aws.StringValue `json:"Message,omitempty"`
}

// Possible values for EMR.
const (
	InstanceStateChangeReasonCodeBootstrapFailure  = "BOOTSTRAP_FAILURE"
	InstanceStateChangeReasonCodeClusterTerminated = "CLUSTER_TERMINATED"
	InstanceStateChangeReasonCodeInstanceFailure   = "INSTANCE_FAILURE"
	InstanceStateChangeReasonCodeInternalError     = "INTERNAL_ERROR"
	InstanceStateChangeReasonCodeValidationError   = "VALIDATION_ERROR"
)

// InstanceStatus is undocumented.
type InstanceStatus struct {
	State             aws.StringValue            `json:"State,omitempty"`
	StateChangeReason *InstanceStateChangeReason `json:"StateChangeReason,omitempty"`
	Timeline          *InstanceTimeline          `json:"Timeline,omitempty"`
}

// InstanceTimeline is undocumented.
type InstanceTimeline struct {
	CreationDateTime *aws.UnixTimestamp `json:"CreationDateTime,omitempty"`
	EndDateTime      *aws.UnixTimestamp `json:"EndDateTime,omitempty"`
	ReadyDateTime    *aws.UnixTimestamp `json:"ReadyDateTime,omitempty"`
}

// JobFlowDetail is undocumented.
type JobFlowDetail struct {
	AMIVersion            aws.StringValue               `json:"AmiVersion,omitempty"`
	BootstrapActions      []BootstrapActionDetail       `json:"BootstrapActions,omitempty"`
	ExecutionStatusDetail *JobFlowExecutionStatusDetail `json:"ExecutionStatusDetail"`
	Instances             *JobFlowInstancesDetail       `json:"Instances"`
	JobFlowID             aws.StringValue               `json:"JobFlowId"`
	JobFlowRole           aws.StringValue               `json:"JobFlowRole,omitempty"`
	LogURI                aws.StringValue               `json:"LogUri,omitempty"`
	Name                  aws.StringValue               `json:"Name"`
	ServiceRole           aws.StringValue               `json:"ServiceRole,omitempty"`
	Steps                 []StepDetail                  `json:"Steps,omitempty"`
	SupportedProducts     []string                      `json:"SupportedProducts,omitempty"`
	VisibleToAllUsers     aws.BooleanValue              `json:"VisibleToAllUsers,omitempty"`
}

// Possible values for EMR.
const (
	JobFlowExecutionStateBootstrapping = "BOOTSTRAPPING"
	JobFlowExecutionStateCompleted     = "COMPLETED"
	JobFlowExecutionStateFailed        = "FAILED"
	JobFlowExecutionStateRunning       = "RUNNING"
	JobFlowExecutionStateShuttingDown  = "SHUTTING_DOWN"
	JobFlowExecutionStateStarting      = "STARTING"
	JobFlowExecutionStateTerminated    = "TERMINATED"
	JobFlowExecutionStateWaiting       = "WAITING"
)

// JobFlowExecutionStatusDetail is undocumented.
type JobFlowExecutionStatusDetail struct {
	CreationDateTime      *aws.UnixTimestamp `json:"CreationDateTime"`
	EndDateTime           *aws.UnixTimestamp `json:"EndDateTime,omitempty"`
	LastStateChangeReason aws.StringValue    `json:"LastStateChangeReason,omitempty"`
	ReadyDateTime         *aws.UnixTimestamp `json:"ReadyDateTime,omitempty"`
	StartDateTime         *aws.UnixTimestamp `json:"StartDateTime,omitempty"`
	State                 aws.StringValue    `json:"State"`
}

// JobFlowInstancesConfig is undocumented.
type JobFlowInstancesConfig struct {
	AdditionalMasterSecurityGroups []string              `json:"AdditionalMasterSecurityGroups,omitempty"`
	AdditionalSlaveSecurityGroups  []string              `json:"AdditionalSlaveSecurityGroups,omitempty"`
	EC2KeyName                     aws.StringValue       `json:"Ec2KeyName,omitempty"`
	EC2SubnetID                    aws.StringValue       `json:"Ec2SubnetId,omitempty"`
	EmrManagedMasterSecurityGroup  aws.StringValue       `json:"EmrManagedMasterSecurityGroup,omitempty"`
	EmrManagedSlaveSecurityGroup   aws.StringValue       `json:"EmrManagedSlaveSecurityGroup,omitempty"`
	HadoopVersion                  aws.StringValue       `json:"HadoopVersion,omitempty"`
	InstanceCount                  aws.IntegerValue      `json:"InstanceCount,omitempty"`
	InstanceGroups                 []InstanceGroupConfig `json:"InstanceGroups,omitempty"`
	KeepJobFlowAliveWhenNoSteps    aws.BooleanValue      `json:"KeepJobFlowAliveWhenNoSteps,omitempty"`
	MasterInstanceType             aws.StringValue       `json:"MasterInstanceType,omitempty"`
	Placement                      *PlacementType        `json:"Placement,omitempty"`
	SlaveInstanceType              aws.StringValue       `json:"SlaveInstanceType,omitempty"`
	TerminationProtected           aws.BooleanValue      `json:"TerminationProtected,omitempty"`
}

// JobFlowInstancesDetail is undocumented.
type JobFlowInstancesDetail struct {
	EC2KeyName                  aws.StringValue       `json:"Ec2KeyName,omitempty"`
	EC2SubnetID                 aws.StringValue       `json:"Ec2SubnetId,omitempty"`
	HadoopVersion               aws.StringValue       `json:"HadoopVersion,omitempty"`
	InstanceCount               aws.IntegerValue      `json:"InstanceCount"`
	InstanceGroups              []InstanceGroupDetail `json:"InstanceGroups,omitempty"`
	KeepJobFlowAliveWhenNoSteps aws.BooleanValue      `json:"KeepJobFlowAliveWhenNoSteps,omitempty"`
	MasterInstanceID            aws.StringValue       `json:"MasterInstanceId,omitempty"`
	MasterInstanceType          aws.StringValue       `json:"MasterInstanceType"`
	MasterPublicDNSName         aws.StringValue       `json:"MasterPublicDnsName,omitempty"`
	NormalizedInstanceHours     aws.IntegerValue      `json:"NormalizedInstanceHours,omitempty"`
	Placement                   *PlacementType        `json:"Placement,omitempty"`
	SlaveInstanceType           aws.StringValue       `json:"SlaveInstanceType"`
	TerminationProtected        aws.BooleanValue      `json:"TerminationProtected,omitempty"`
}

// KeyValue is undocumented.
type KeyValue struct {
	Key   aws.StringValue `json:"Key,omitempty"`
	Value aws.StringValue `json:"Value,omitempty"`
}

// ListBootstrapActionsInput is undocumented.
type ListBootstrapActionsInput struct {
	ClusterID aws.StringValue `json:"ClusterId"`
	Marker    aws.StringValue `json:"Marker,omitempty"`
}

// ListBootstrapActionsOutput is undocumented.
type ListBootstrapActionsOutput struct {
	BootstrapActions []Command       `json:"BootstrapActions,omitempty"`
	Marker           aws.StringValue `json:"Marker,omitempty"`
}

// ListClustersInput is undocumented.
type ListClustersInput struct {
	ClusterStates []string           `json:"ClusterStates,omitempty"`
	CreatedAfter  *aws.UnixTimestamp `json:"CreatedAfter,omitempty"`
	CreatedBefore *aws.UnixTimestamp `json:"CreatedBefore,omitempty"`
	Marker        aws.StringValue    `json:"Marker,omitempty"`
}

// ListClustersOutput is undocumented.
type ListClustersOutput struct {
	Clusters []ClusterSummary `json:"Clusters,omitempty"`
	Marker   aws.StringValue  `json:"Marker,omitempty"`
}

// ListInstanceGroupsInput is undocumented.
type ListInstanceGroupsInput struct {
	ClusterID aws.StringValue `json:"ClusterId"`
	Marker    aws.StringValue `json:"Marker,omitempty"`
}

// ListInstanceGroupsOutput is undocumented.
type ListInstanceGroupsOutput struct {
	InstanceGroups []InstanceGroup `json:"InstanceGroups,omitempty"`
	Marker         aws.StringValue `json:"Marker,omitempty"`
}

// ListInstancesInput is undocumented.
type ListInstancesInput struct {
	ClusterID          aws.StringValue `json:"ClusterId"`
	InstanceGroupID    aws.StringValue `json:"InstanceGroupId,omitempty"`
	InstanceGroupTypes []string        `json:"InstanceGroupTypes,omitempty"`
	Marker             aws.StringValue `json:"Marker,omitempty"`
}

// ListInstancesOutput is undocumented.
type ListInstancesOutput struct {
	Instances []Instance      `json:"Instances,omitempty"`
	Marker    aws.StringValue `json:"Marker,omitempty"`
}

// ListStepsInput is undocumented.
type ListStepsInput struct {
	ClusterID  aws.StringValue `json:"ClusterId"`
	Marker     aws.StringValue `json:"Marker,omitempty"`
	StepIDs    []string        `json:"StepIds,omitempty"`
	StepStates []string        `json:"StepStates,omitempty"`
}

// ListStepsOutput is undocumented.
type ListStepsOutput struct {
	Marker aws.StringValue `json:"Marker,omitempty"`
	Steps  []StepSummary   `json:"Steps,omitempty"`
}

// Possible values for EMR.
const (
	MarketTypeOnDemand = "ON_DEMAND"
	MarketTypeSpot     = "SPOT"
)

// ModifyInstanceGroupsInput is undocumented.
type ModifyInstanceGroupsInput struct {
	InstanceGroups []InstanceGroupModifyConfig `json:"InstanceGroups,omitempty"`
}

// PlacementType is undocumented.
type PlacementType struct {
	AvailabilityZone aws.StringValue `json:"AvailabilityZone"`
}

// RemoveTagsInput is undocumented.
type RemoveTagsInput struct {
	ResourceID aws.StringValue `json:"ResourceId"`
	TagKeys    []string        `json:"TagKeys"`
}

// RemoveTagsOutput is undocumented.
type RemoveTagsOutput struct {
}

// RunJobFlowInput is undocumented.
type RunJobFlowInput struct {
	AdditionalInfo       aws.StringValue          `json:"AdditionalInfo,omitempty"`
	AMIVersion           aws.StringValue          `json:"AmiVersion,omitempty"`
	BootstrapActions     []BootstrapActionConfig  `json:"BootstrapActions,omitempty"`
	Instances            *JobFlowInstancesConfig  `json:"Instances"`
	JobFlowRole          aws.StringValue          `json:"JobFlowRole,omitempty"`
	LogURI               aws.StringValue          `json:"LogUri,omitempty"`
	Name                 aws.StringValue          `json:"Name"`
	NewSupportedProducts []SupportedProductConfig `json:"NewSupportedProducts,omitempty"`
	ServiceRole          aws.StringValue          `json:"ServiceRole,omitempty"`
	Steps                []StepConfig             `json:"Steps,omitempty"`
	SupportedProducts    []string                 `json:"SupportedProducts,omitempty"`
	Tags                 []Tag                    `json:"Tags,omitempty"`
	VisibleToAllUsers    aws.BooleanValue         `json:"VisibleToAllUsers,omitempty"`
}

// RunJobFlowOutput is undocumented.
type RunJobFlowOutput struct {
	JobFlowID aws.StringValue `json:"JobFlowId,omitempty"`
}

// ScriptBootstrapActionConfig is undocumented.
type ScriptBootstrapActionConfig struct {
	Args []string        `json:"Args,omitempty"`
	Path aws.StringValue `json:"Path"`
}

// SetTerminationProtectionInput is undocumented.
type SetTerminationProtectionInput struct {
	JobFlowIDs           []string         `json:"JobFlowIds"`
	TerminationProtected aws.BooleanValue `json:"TerminationProtected"`
}

// SetVisibleToAllUsersInput is undocumented.
type SetVisibleToAllUsersInput struct {
	JobFlowIDs        []string         `json:"JobFlowIds"`
	VisibleToAllUsers aws.BooleanValue `json:"VisibleToAllUsers"`
}

// Step is undocumented.
type Step struct {
	ActionOnFailure aws.StringValue   `json:"ActionOnFailure,omitempty"`
	Config          *HadoopStepConfig `json:"Config,omitempty"`
	ID              aws.StringValue   `json:"Id,omitempty"`
	Name            aws.StringValue   `json:"Name,omitempty"`
	Status          *StepStatus       `json:"Status,omitempty"`
}

// StepConfig is undocumented.
type StepConfig struct {
	ActionOnFailure aws.StringValue      `json:"ActionOnFailure,omitempty"`
	HadoopJARStep   *HadoopJARStepConfig `json:"HadoopJarStep"`
	Name            aws.StringValue      `json:"Name"`
}

// StepDetail is undocumented.
type StepDetail struct {
	ExecutionStatusDetail *StepExecutionStatusDetail `json:"ExecutionStatusDetail"`
	StepConfig            *StepConfig                `json:"StepConfig"`
}

// Possible values for EMR.
const (
	StepExecutionStateCancelled   = "CANCELLED"
	StepExecutionStateCompleted   = "COMPLETED"
	StepExecutionStateContinue    = "CONTINUE"
	StepExecutionStateFailed      = "FAILED"
	StepExecutionStateInterrupted = "INTERRUPTED"
	StepExecutionStatePending     = "PENDING"
	StepExecutionStateRunning     = "RUNNING"
)

// StepExecutionStatusDetail is undocumented.
type StepExecutionStatusDetail struct {
	CreationDateTime      *aws.UnixTimestamp `json:"CreationDateTime"`
	EndDateTime           *aws.UnixTimestamp `json:"EndDateTime,omitempty"`
	LastStateChangeReason aws.StringValue    `json:"LastStateChangeReason,omitempty"`
	StartDateTime         *aws.UnixTimestamp `json:"StartDateTime,omitempty"`
	State                 aws.StringValue    `json:"State"`
}

// Possible values for EMR.
const (
	StepStateCancelled   = "CANCELLED"
	StepStateCompleted   = "COMPLETED"
	StepStateFailed      = "FAILED"
	StepStateInterrupted = "INTERRUPTED"
	StepStatePending     = "PENDING"
	StepStateRunning     = "RUNNING"
)

// StepStateChangeReason is undocumented.
type StepStateChangeReason struct {
	Code    aws.StringValue `json:"Code,omitempty"`
	Message aws.StringValue `json:"Message,omitempty"`
}

// Possible values for EMR.
const (
	StepStateChangeReasonCodeNone = "NONE"
)

// StepStatus is undocumented.
type StepStatus struct {
	State             aws.StringValue        `json:"State,omitempty"`
	StateChangeReason *StepStateChangeReason `json:"StateChangeReason,omitempty"`
	Timeline          *StepTimeline          `json:"Timeline,omitempty"`
}

// StepSummary is undocumented.
type StepSummary struct {
	ActionOnFailure aws.StringValue   `json:"ActionOnFailure,omitempty"`
	Config          *HadoopStepConfig `json:"Config,omitempty"`
	ID              aws.StringValue   `json:"Id,omitempty"`
	Name            aws.StringValue   `json:"Name,omitempty"`
	Status          *StepStatus       `json:"Status,omitempty"`
}

// StepTimeline is undocumented.
type StepTimeline struct {
	CreationDateTime *aws.UnixTimestamp `json:"CreationDateTime,omitempty"`
	EndDateTime      *aws.UnixTimestamp `json:"EndDateTime,omitempty"`
	StartDateTime    *aws.UnixTimestamp `json:"StartDateTime,omitempty"`
}

// SupportedProductConfig is undocumented.
type SupportedProductConfig struct {
	Args []string        `json:"Args,omitempty"`
	Name aws.StringValue `json:"Name,omitempty"`
}

// Tag is undocumented.
type Tag struct {
	Key   aws.StringValue `json:"Key,omitempty"`
	Value aws.StringValue `json:"Value,omitempty"`
}

// TerminateJobFlowsInput is undocumented.
type TerminateJobFlowsInput struct {
	JobFlowIDs []string `json:"JobFlowIds"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
