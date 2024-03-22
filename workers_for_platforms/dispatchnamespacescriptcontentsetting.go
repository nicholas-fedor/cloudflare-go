// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// DispatchNamespaceScriptContentSettingService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDispatchNamespaceScriptContentSettingService] method instead.
type DispatchNamespaceScriptContentSettingService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptContentSettingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDispatchNamespaceScriptContentSettingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptContentSettingService) {
	r = &DispatchNamespaceScriptContentSettingService{}
	r.Options = opts
	return
}

// Patch script metadata, such as bindings
func (r *DispatchNamespaceScriptContentSettingService) Edit(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptContentSettingEditParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptContentSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceScriptContentSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptContentSettingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptContentSettingGetParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptContentSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceScriptContentSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DispatchNamespaceScriptContentSettingEditResponse struct {
	// List of bindings attached to this Worker
	Bindings []DispatchNamespaceScriptContentSettingEditResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations DispatchNamespaceScriptContentSettingEditResponseMigrations `json:"migrations"`
	Placement  DispatchNamespaceScriptContentSettingEditResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []DispatchNamespaceScriptContentSettingEditResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                `json:"usage_model"`
	JSON       dispatchNamespaceScriptContentSettingEditResponseJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptContentSettingEditResponse]
type dispatchNamespaceScriptContentSettingEditResponseJSON struct {
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Logpush            apijson.Field
	Migrations         apijson.Field
	Placement          apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2Binding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1Binding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBinding].
type DispatchNamespaceScriptContentSettingEditResponseBinding interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptContentSettingEditResponseBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBinding{}),
		},
	)
}

type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBinding]
type dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingType string

const (
	DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBinding]
type dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingType string

const (
	DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingTypeService DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingType = "service"
)

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                        `json:"script_name"`
	JSON       dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBinding]
type dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingType string

const (
	DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2Binding]
type dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2Binding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingType string

const (
	DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingType = "r2_bucket"
)

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBinding]
type dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingType string

const (
	DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingTypeQueue DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingType = "queue"
)

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1Binding]
type dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1Binding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingType string

const (
	DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingTypeD1 DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingType = "d1"
)

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
type dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                     `json:"service"`
	JSON    dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                              `json:"certificate_id"`
	JSON          dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMtlscertBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMtlscertBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBinding]
type dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMtlscertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBindingType string

const (
	DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrations]
// or
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrations].
type DispatchNamespaceScriptContentSettingEditResponseMigrations interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptContentSettingEditResponseMigrations)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrations struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrations]
type dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseMigrations() {
}

type DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                                 `json:"from"`
	To   string                                                                                                 `json:"to"`
	JSON dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                     `json:"from"`
	FromScript string                                                                                                     `json:"from_script"`
	To         string                                                                                                     `json:"to"`
	JSON       dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrations]
type dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditResponseMigrations() {
}

type DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStep]
type dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                   `json:"from"`
	To   string                                                                                                   `json:"to"`
	JSON dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                       `json:"from"`
	FromScript string                                                                                                       `json:"from_script"`
	To         string                                                                                                       `json:"to"`
	JSON       dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode DispatchNamespaceScriptContentSettingEditResponsePlacementMode `json:"mode"`
	JSON dispatchNamespaceScriptContentSettingEditResponsePlacementJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponsePlacementJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponsePlacement]
type dispatchNamespaceScriptContentSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type DispatchNamespaceScriptContentSettingEditResponsePlacementMode string

const (
	DispatchNamespaceScriptContentSettingEditResponsePlacementModeSmart DispatchNamespaceScriptContentSettingEditResponsePlacementMode = "smart"
)

func (r DispatchNamespaceScriptContentSettingEditResponsePlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditResponsePlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type DispatchNamespaceScriptContentSettingEditResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                            `json:"namespace"`
	JSON      dispatchNamespaceScriptContentSettingEditResponseTailConsumerJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseTailConsumerJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseTailConsumer]
type dispatchNamespaceScriptContentSettingEditResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []DispatchNamespaceScriptContentSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations DispatchNamespaceScriptContentSettingGetResponseMigrations `json:"migrations"`
	Placement  DispatchNamespaceScriptContentSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []DispatchNamespaceScriptContentSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                               `json:"usage_model"`
	JSON       dispatchNamespaceScriptContentSettingGetResponseJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptContentSettingGetResponse]
type dispatchNamespaceScriptContentSettingGetResponseJSON struct {
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Logpush            apijson.Field
	Migrations         apijson.Field
	Placement          apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2Binding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1Binding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [workers_for_platforms.DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBinding].
type DispatchNamespaceScriptContentSettingGetResponseBinding interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptContentSettingGetResponseBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBinding{}),
		},
	)
}

type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBinding]
type dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingType string

const (
	DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBinding]
type dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingType string

const (
	DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingTypeService DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingType = "service"
)

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                       `json:"script_name"`
	JSON       dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBinding]
type dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingType string

const (
	DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2Binding]
type dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2Binding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingType string

const (
	DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingType = "r2_bucket"
)

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBinding]
type dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingType string

const (
	DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingTypeQueue DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingType = "queue"
)

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1Binding]
type dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1Binding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingType string

const (
	DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingTypeD1 DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingType = "d1"
)

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
type dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                    `json:"service"`
	JSON    dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                             `json:"certificate_id"`
	JSON          dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMtlscertBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMtlscertBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBinding]
type dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMtlscertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBindingType string

const (
	DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrations]
// or
// [workers_for_platforms.DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrations].
type DispatchNamespaceScriptContentSettingGetResponseMigrations interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptContentSettingGetResponseMigrations)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrations struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrations]
type dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseMigrations() {
}

type DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                                `json:"from"`
	To   string                                                                                                `json:"to"`
	JSON dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                    `json:"from"`
	FromScript string                                                                                                    `json:"from_script"`
	To         string                                                                                                    `json:"to"`
	JSON       dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrations]
type dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingGetResponseMigrations() {
}

type DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStep]
type dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                  `json:"from"`
	To   string                                                                                                  `json:"to"`
	JSON dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                      `json:"from"`
	FromScript string                                                                                                      `json:"from_script"`
	To         string                                                                                                      `json:"to"`
	JSON       dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode DispatchNamespaceScriptContentSettingGetResponsePlacementMode `json:"mode"`
	JSON dispatchNamespaceScriptContentSettingGetResponsePlacementJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponsePlacementJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponsePlacement]
type dispatchNamespaceScriptContentSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type DispatchNamespaceScriptContentSettingGetResponsePlacementMode string

const (
	DispatchNamespaceScriptContentSettingGetResponsePlacementModeSmart DispatchNamespaceScriptContentSettingGetResponsePlacementMode = "smart"
)

func (r DispatchNamespaceScriptContentSettingGetResponsePlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingGetResponsePlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type DispatchNamespaceScriptContentSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                           `json:"namespace"`
	JSON      dispatchNamespaceScriptContentSettingGetResponseTailConsumerJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseTailConsumerJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseTailConsumer]
type dispatchNamespaceScriptContentSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                                                   `path:"account_id,required"`
	Errors    param.Field[[]DispatchNamespaceScriptContentSettingEditParamsError]   `json:"errors,required"`
	Messages  param.Field[[]DispatchNamespaceScriptContentSettingEditParamsMessage] `json:"messages,required"`
	Result    param.Field[DispatchNamespaceScriptContentSettingEditParamsResult]    `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[DispatchNamespaceScriptContentSettingEditParamsSuccess] `json:"success,required"`
}

func (r DispatchNamespaceScriptContentSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptContentSettingEditParamsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptContentSettingEditParamsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptContentSettingEditParamsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]DispatchNamespaceScriptContentSettingEditParamsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[DispatchNamespaceScriptContentSettingEditParamsResultMigrations] `json:"migrations"`
	Placement  param.Field[DispatchNamespaceScriptContentSettingEditParamsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]DispatchNamespaceScriptContentSettingEditParamsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2Binding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1Binding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCERTBinding].
type DispatchNamespaceScriptContentSettingEditParamsResultBinding interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultBinding()
}

type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType string

const (
	DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBindingType string

const (
	DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBindingTypeService DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBindingType = "service"
)

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBindingType string

const (
	DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2Binding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2BindingType string

const (
	DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2BindingTypeR2Bucket DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2BindingType = "r2_bucket"
)

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBindingType string

const (
	DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBindingTypeQueue DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBindingType = "queue"
)

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1Binding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1BindingType string

const (
	DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1BindingTypeD1 DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1BindingType = "d1"
)

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCERTBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCERTBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCERTBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCERTBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCERTBindingType string

const (
	DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCERTBindingTypeMTLSCertificate DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrations],
// [workers_for_platforms.DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrations].
type DispatchNamespaceScriptContentSettingEditParamsResultMigrations interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultMigrations()
}

// A single set of migrations to apply.
type DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrations struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultMigrations() {
}

type DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptContentSettingEditParamsResultMigrations() {
}

type DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptContentSettingEditParamsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[DispatchNamespaceScriptContentSettingEditParamsResultPlacementMode] `json:"mode"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type DispatchNamespaceScriptContentSettingEditParamsResultPlacementMode string

const (
	DispatchNamespaceScriptContentSettingEditParamsResultPlacementModeSmart DispatchNamespaceScriptContentSettingEditParamsResultPlacementMode = "smart"
)

func (r DispatchNamespaceScriptContentSettingEditParamsResultPlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditParamsResultPlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type DispatchNamespaceScriptContentSettingEditParamsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r DispatchNamespaceScriptContentSettingEditParamsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type DispatchNamespaceScriptContentSettingEditParamsSuccess bool

const (
	DispatchNamespaceScriptContentSettingEditParamsSuccessTrue DispatchNamespaceScriptContentSettingEditParamsSuccess = true
)

func (r DispatchNamespaceScriptContentSettingEditParamsSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditParamsSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingEditResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptContentSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptContentSettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   DispatchNamespaceScriptContentSettingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptContentSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptContentSettingEditResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseEnvelope]
type dispatchNamespaceScriptContentSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingEditResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    dispatchNamespaceScriptContentSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseEnvelopeErrors]
type dispatchNamespaceScriptContentSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingEditResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    dispatchNamespaceScriptContentSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingEditResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingEditResponseEnvelopeMessages]
type dispatchNamespaceScriptContentSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptContentSettingEditResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptContentSettingEditResponseEnvelopeSuccessTrue DispatchNamespaceScriptContentSettingEditResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptContentSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptContentSettingGetResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptContentSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptContentSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DispatchNamespaceScriptContentSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptContentSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptContentSettingGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseEnvelope]
type dispatchNamespaceScriptContentSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingGetResponseEnvelopeErrors struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    dispatchNamespaceScriptContentSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseEnvelopeErrors]
type dispatchNamespaceScriptContentSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentSettingGetResponseEnvelopeMessages struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    dispatchNamespaceScriptContentSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dispatchNamespaceScriptContentSettingGetResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptContentSettingGetResponseEnvelopeMessages]
type dispatchNamespaceScriptContentSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptContentSettingGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptContentSettingGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptContentSettingGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptContentSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
