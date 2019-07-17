// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_limits "github.com/oracle/oci-go-sdk/limits"
)

func LimitsQuotaDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["quota_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(LimitsQuotaResource(), fieldMap, readSingularLimitsQuota)
}

func readSingularLimitsQuota(d *schema.ResourceData, m interface{}) error {
	sync := &LimitsQuotaDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).quotasClient

	return ReadResource(sync)
}

type LimitsQuotaDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_limits.QuotasClient
	Res    *oci_limits.GetQuotaResponse
}

func (s *LimitsQuotaDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LimitsQuotaDataSourceCrud) Get() error {
	request := oci_limits.GetQuotaRequest{}

	if quotaId, ok := s.D.GetOkExists("quota_id"); ok {
		tmp := quotaId.(string)
		request.QuotaId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "limits")

	response, err := s.Client.GetQuota(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LimitsQuotaDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("statements", s.Res.Statements)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
