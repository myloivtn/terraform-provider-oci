// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateOnDemandPingProbeDetails The request body used to create an on-demand ping probe.
type CreateOnDemandPingProbeDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Targets []string `mandatory:"true" json:"targets"`

	Protocol PingProbeProtocolEnum `mandatory:"true" json:"protocol"`

	VantagePointNames []string `mandatory:"false" json:"vantagePointNames"`

	// The port on which to probe endpoints. If unspecified, probes will use the
	// default port of their protocol.
	Port *int `mandatory:"false" json:"port"`

	// The probe timeout in seconds. Valid values: 10, 20, 30, and 60.
	// The probe timeout must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`
}

func (m CreateOnDemandPingProbeDetails) String() string {
	return common.PointerString(m)
}

// CreateOnDemandPingProbeDetailsProtocolEnum is an alias to type: PingProbeProtocolEnum
// Consider using PingProbeProtocolEnum instead
// Deprecated
type CreateOnDemandPingProbeDetailsProtocolEnum = PingProbeProtocolEnum

// Set of constants representing the allowable values for PingProbeProtocolEnum
// Deprecated
const (
	CreateOnDemandPingProbeDetailsProtocolIcmp PingProbeProtocolEnum = "ICMP"
	CreateOnDemandPingProbeDetailsProtocolTcp  PingProbeProtocolEnum = "TCP"
)

// GetCreateOnDemandPingProbeDetailsProtocolEnumValues Enumerates the set of values for PingProbeProtocolEnum
// Consider using GetPingProbeProtocolEnumValue
// Deprecated
var GetCreateOnDemandPingProbeDetailsProtocolEnumValues = GetPingProbeProtocolEnumValues
