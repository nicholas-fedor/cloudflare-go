// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessServiceTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessServiceTokenService] method
// instead.
type AccessServiceTokenService struct {
	Options []option.RequestOption
}

// NewAccessServiceTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessServiceTokenService(opts ...option.RequestOption) (r *AccessServiceTokenService) {
	r = &AccessServiceTokenService{}
	r.Options = opts
	return
}

// Generates a new service token. **Note:** This is the only time you can get the
// Client Secret. If you lose the Client Secret, you will have to rotate the Client
// Secret or create a new service token.
func (r *AccessServiceTokenService) New(ctx context.Context, params AccessServiceTokenNewParams, opts ...option.RequestOption) (res *AccessServiceTokenNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessServiceTokenNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured service token.
func (r *AccessServiceTokenService) Update(ctx context.Context, uuid string, params AccessServiceTokenUpdateParams, opts ...option.RequestOption) (res *AccessServiceTokenUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessServiceTokenUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all service tokens.
func (r *AccessServiceTokenService) List(ctx context.Context, query AccessServiceTokenListParams, opts ...option.RequestOption) (res *[]AccessServiceTokenListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessServiceTokenListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a service token.
func (r *AccessServiceTokenService) Delete(ctx context.Context, uuid string, body AccessServiceTokenDeleteParams, opts ...option.RequestOption) (res *AccessServiceTokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessServiceTokenDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Refreshes the expiration of a service token.
func (r *AccessServiceTokenService) Refresh(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccessServiceTokenRefreshResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessServiceTokenRefreshResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/refresh", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Generates a new Client Secret for a service token and revokes the old one.
func (r *AccessServiceTokenService) Rotate(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccessServiceTokenRotateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessServiceTokenRotateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/rotate", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessServiceTokenNewResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                            `json:"name"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenNewResponseJSON `json:"-"`
}

// accessServiceTokenNewResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenNewResponse]
type accessServiceTokenNewResponseJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Duration     apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenUpdateResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                               `json:"name"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenUpdateResponseJSON `json:"-"`
}

// accessServiceTokenUpdateResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenUpdateResponse]
type accessServiceTokenUpdateResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenListResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                             `json:"name"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenListResponseJSON `json:"-"`
}

// accessServiceTokenListResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenListResponse]
type accessServiceTokenListResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenDeleteResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                               `json:"name"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenDeleteResponseJSON `json:"-"`
}

// accessServiceTokenDeleteResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenDeleteResponse]
type accessServiceTokenDeleteResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRefreshResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                `json:"name"`
	UpdatedAt time.Time                             `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenRefreshResponseJSON `json:"-"`
}

// accessServiceTokenRefreshResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenRefreshResponse]
type accessServiceTokenRefreshResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRotateResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                               `json:"name"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenRotateResponseJSON `json:"-"`
}

// accessServiceTokenRotateResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenRotateResponse]
type accessServiceTokenRotateResponseJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Duration     apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenNewParams struct {
	// The name of the service token.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
}

func (r AccessServiceTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessServiceTokenNewResponseEnvelope struct {
	Errors   []AccessServiceTokenNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessServiceTokenNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessServiceTokenNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessServiceTokenNewResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenNewResponseEnvelope]
type accessServiceTokenNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessServiceTokenNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessServiceTokenNewResponseEnvelopeErrors]
type accessServiceTokenNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessServiceTokenNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessServiceTokenNewResponseEnvelopeMessages]
type accessServiceTokenNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenNewResponseEnvelopeSuccess bool

const (
	AccessServiceTokenNewResponseEnvelopeSuccessTrue AccessServiceTokenNewResponseEnvelopeSuccess = true
)

type AccessServiceTokenUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
	// The name of the service token.
	Name param.Field[string] `json:"name"`
}

func (r AccessServiceTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessServiceTokenUpdateResponseEnvelope struct {
	Errors   []AccessServiceTokenUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessServiceTokenUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessServiceTokenUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessServiceTokenUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenUpdateResponseEnvelope]
type accessServiceTokenUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenUpdateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessServiceTokenUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessServiceTokenUpdateResponseEnvelopeErrors]
type accessServiceTokenUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenUpdateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accessServiceTokenUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessServiceTokenUpdateResponseEnvelopeMessages]
type accessServiceTokenUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenUpdateResponseEnvelopeSuccess bool

const (
	AccessServiceTokenUpdateResponseEnvelopeSuccessTrue AccessServiceTokenUpdateResponseEnvelopeSuccess = true
)

type AccessServiceTokenListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessServiceTokenListResponseEnvelope struct {
	Errors   []AccessServiceTokenListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessServiceTokenListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessServiceTokenListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessServiceTokenListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessServiceTokenListResponseEnvelopeJSON       `json:"-"`
}

// accessServiceTokenListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenListResponseEnvelope]
type accessServiceTokenListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessServiceTokenListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessServiceTokenListResponseEnvelopeErrors]
type accessServiceTokenListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessServiceTokenListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessServiceTokenListResponseEnvelopeMessages]
type accessServiceTokenListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenListResponseEnvelopeSuccess bool

const (
	AccessServiceTokenListResponseEnvelopeSuccessTrue AccessServiceTokenListResponseEnvelopeSuccess = true
)

type AccessServiceTokenListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       accessServiceTokenListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessServiceTokenListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AccessServiceTokenListResponseEnvelopeResultInfo]
type accessServiceTokenListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessServiceTokenDeleteResponseEnvelope struct {
	Errors   []AccessServiceTokenDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessServiceTokenDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessServiceTokenDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessServiceTokenDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenDeleteResponseEnvelope]
type accessServiceTokenDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessServiceTokenDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessServiceTokenDeleteResponseEnvelopeErrors]
type accessServiceTokenDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accessServiceTokenDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessServiceTokenDeleteResponseEnvelopeMessages]
type accessServiceTokenDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenDeleteResponseEnvelopeSuccess bool

const (
	AccessServiceTokenDeleteResponseEnvelopeSuccessTrue AccessServiceTokenDeleteResponseEnvelopeSuccess = true
)

type AccessServiceTokenRefreshResponseEnvelope struct {
	Errors   []AccessServiceTokenRefreshResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenRefreshResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessServiceTokenRefreshResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessServiceTokenRefreshResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessServiceTokenRefreshResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenRefreshResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenRefreshResponseEnvelope]
type accessServiceTokenRefreshResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRefreshResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessServiceTokenRefreshResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenRefreshResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessServiceTokenRefreshResponseEnvelopeErrors]
type accessServiceTokenRefreshResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRefreshResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accessServiceTokenRefreshResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenRefreshResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessServiceTokenRefreshResponseEnvelopeMessages]
type accessServiceTokenRefreshResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenRefreshResponseEnvelopeSuccess bool

const (
	AccessServiceTokenRefreshResponseEnvelopeSuccessTrue AccessServiceTokenRefreshResponseEnvelopeSuccess = true
)

type AccessServiceTokenRotateResponseEnvelope struct {
	Errors   []AccessServiceTokenRotateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenRotateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessServiceTokenRotateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessServiceTokenRotateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessServiceTokenRotateResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenRotateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenRotateResponseEnvelope]
type accessServiceTokenRotateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRotateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessServiceTokenRotateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenRotateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessServiceTokenRotateResponseEnvelopeErrors]
type accessServiceTokenRotateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRotateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accessServiceTokenRotateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenRotateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessServiceTokenRotateResponseEnvelopeMessages]
type accessServiceTokenRotateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenRotateResponseEnvelopeSuccess bool

const (
	AccessServiceTokenRotateResponseEnvelopeSuccessTrue AccessServiceTokenRotateResponseEnvelopeSuccess = true
)
