// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// RiskScoringService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRiskScoringService] method instead.
type RiskScoringService struct {
	Options      []option.RequestOption
	Behaviours   *RiskScoringBehaviourService
	Summary      *RiskScoringSummaryService
	Integrations *RiskScoringIntegrationService
}

// NewRiskScoringService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRiskScoringService(opts ...option.RequestOption) (r *RiskScoringService) {
	r = &RiskScoringService{}
	r.Options = opts
	r.Behaviours = NewRiskScoringBehaviourService(opts...)
	r.Summary = NewRiskScoringSummaryService(opts...)
	r.Integrations = NewRiskScoringIntegrationService(opts...)
	return
}

// Get risk event/score information for a specific user
func (r *RiskScoringService) Get(ctx context.Context, userID string, query RiskScoringGetParams, opts ...option.RequestOption) (res *RiskScoringGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/%s", query.AccountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Clear the risk score for a particular user
func (r *RiskScoringService) Reset(ctx context.Context, userID string, body RiskScoringResetParams, opts ...option.RequestOption) (res *RiskScoringResetResponse, err error) {
	var env RiskScoringResetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/%s/reset", body.AccountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RiskScoringGetResponse struct {
	Email         string                          `json:"email,required"`
	Events        []RiskScoringGetResponseEvent   `json:"events,required"`
	Name          string                          `json:"name,required"`
	LastResetTime time.Time                       `json:"last_reset_time,nullable" format:"date-time"`
	RiskLevel     RiskScoringGetResponseRiskLevel `json:"risk_level"`
	JSON          riskScoringGetResponseJSON      `json:"-"`
}

// riskScoringGetResponseJSON contains the JSON metadata for the struct
// [RiskScoringGetResponse]
type riskScoringGetResponseJSON struct {
	Email         apijson.Field
	Events        apijson.Field
	Name          apijson.Field
	LastResetTime apijson.Field
	RiskLevel     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RiskScoringGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringGetResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringGetResponseEvent struct {
	ID           string                                `json:"id,required"`
	Name         string                                `json:"name,required"`
	RiskLevel    RiskScoringGetResponseEventsRiskLevel `json:"risk_level,required"`
	Timestamp    time.Time                             `json:"timestamp,required" format:"date-time"`
	EventDetails interface{}                           `json:"event_details"`
	JSON         riskScoringGetResponseEventJSON       `json:"-"`
}

// riskScoringGetResponseEventJSON contains the JSON metadata for the struct
// [RiskScoringGetResponseEvent]
type riskScoringGetResponseEventJSON struct {
	ID           apijson.Field
	Name         apijson.Field
	RiskLevel    apijson.Field
	Timestamp    apijson.Field
	EventDetails apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RiskScoringGetResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringGetResponseEventJSON) RawJSON() string {
	return r.raw
}

type RiskScoringGetResponseEventsRiskLevel string

const (
	RiskScoringGetResponseEventsRiskLevelLow    RiskScoringGetResponseEventsRiskLevel = "low"
	RiskScoringGetResponseEventsRiskLevelMedium RiskScoringGetResponseEventsRiskLevel = "medium"
	RiskScoringGetResponseEventsRiskLevelHigh   RiskScoringGetResponseEventsRiskLevel = "high"
)

func (r RiskScoringGetResponseEventsRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringGetResponseEventsRiskLevelLow, RiskScoringGetResponseEventsRiskLevelMedium, RiskScoringGetResponseEventsRiskLevelHigh:
		return true
	}
	return false
}

type RiskScoringGetResponseRiskLevel string

const (
	RiskScoringGetResponseRiskLevelLow    RiskScoringGetResponseRiskLevel = "low"
	RiskScoringGetResponseRiskLevelMedium RiskScoringGetResponseRiskLevel = "medium"
	RiskScoringGetResponseRiskLevelHigh   RiskScoringGetResponseRiskLevel = "high"
)

func (r RiskScoringGetResponseRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringGetResponseRiskLevelLow, RiskScoringGetResponseRiskLevelMedium, RiskScoringGetResponseRiskLevelHigh:
		return true
	}
	return false
}

type RiskScoringResetResponse = interface{}

type RiskScoringGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RiskScoringResetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RiskScoringResetResponseEnvelope struct {
	Errors     []shared.ResponseInfo                      `json:"errors,required"`
	Messages   []shared.ResponseInfo                      `json:"messages,required"`
	Success    bool                                       `json:"success,required"`
	Result     RiskScoringResetResponse                   `json:"result,nullable"`
	ResultInfo RiskScoringResetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       riskScoringResetResponseEnvelopeJSON       `json:"-"`
}

// riskScoringResetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RiskScoringResetResponseEnvelope]
type riskScoringResetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringResetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringResetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RiskScoringResetResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                          `json:"total_count,required"`
	JSON       riskScoringResetResponseEnvelopeResultInfoJSON `json:"-"`
}

// riskScoringResetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [RiskScoringResetResponseEnvelopeResultInfo]
type riskScoringResetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringResetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringResetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
