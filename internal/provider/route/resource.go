package route

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stevan-sdk/pkg/client"
	"github.com/stevan-sdk/pkg/routes"
	"github.com/terraform-provider-stevan/internal/shared/models/service"
	"github.com/terraform-provider-stevan/internal/utils"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &RouteResource{}
var _ resource.ResourceWithImportState = &RouteResource{}

// constructor
func NewRouteResource() resource.Resource {
	return &RouteResource{}
}

// client wrapper
type RouteResource struct {
	client *client.Client
}

type RouteResourceModel struct {
	CreatedAt               types.Int64      `tfsdk:"created_at"`
	Headers                 types.Map        `tfsdk:"headers"`
	Hosts                   types.List       `tfsdk:"hosts"`
	HttpsRedirectStatusCode types.Int64      `tfsdk:"https_redirect_status_code"`
	Id                      types.String     `tfsdk:"id"`
	Methods                 types.List       `tfsdk:"methods"`
	Name                    types.String     `tfsdk:"name"`
	PathHandling            types.String     `tfsdk:"path_handling"`
	Paths                   types.List       `tfsdk:"paths"`
	PreserveHost            types.Bool       `tfsdk:"preserve_host"`
	Protocols               types.List       `tfsdk:"protocols"`
	RegexPriority           types.Int64      `tfsdk:"regex_priority"`
	RequestBuffering        types.Bool       `tfsdk:"request_buffering"`
	ResponseBuffering       types.Bool       `tfsdk:"response_buffering"`
	Service                 *service.Service `tfsdk:"service"`
	Snis                    types.List       `tfsdk:"snis"`
	StripPath               types.Bool       `tfsdk:"strip_path"`
	Tags                    types.List       `tfsdk:"tags"`
	UpdatedAt               types.Int64      `tfsdk:"updated_at"`
	RuntimeGroupId          types.String     `tfsdk:"runtime_group_id"`
	RouteId                 types.String     `tfsdk:"route_id"`
}

func (r *RouteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route"
}

func (r *RouteResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Route entities define rules to match client requests. Every request matching a given route will be proxied to its associated service.",
		Attributes: map[string]schema.Attribute{
			"created_at": schema.Int64Attribute{
				Description: "Unix epoch when the resource was created.",
				Optional:    true,
			},

			"headers": schema.MapAttribute{
				Description: "One or more lists of values indexed by header name that will cause this route to match if present in the request. The `Host` header cannot be used with this attribute: hosts should be specified using the `hosts` attribute. When `headers` contains only one value and that value starts with the special prefix `~*`, the value is interpreted as a regular expression.",
				Optional:    true,

				ElementType: types.StringType,
			},

			"hosts": schema.ListAttribute{
				Description: "A list of domain names that match this route. Note that the hosts value is case sensitive.",
				Optional:    true,

				ElementType: types.StringType,
			},

			"https_redirect_status_code": schema.Int64Attribute{
				Description: "The status code Kong responds with when all properties of a route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the route is configured to only accept the `https` protocol.",
				Optional:    true,
			},

			"id": schema.StringAttribute{
				Description: "id",
				Optional:    true,
			},

			"methods": schema.ListAttribute{
				Description: "A list of HTTP methods that match this route.",
				Optional:    true,

				ElementType: types.StringType,
			},

			"name": schema.StringAttribute{
				Description: "The name of the route. Route names must be unique, and they are case sensitive. For example, there can be two different routes named test and Test.",
				Optional:    true,
			},

			"path_handling": schema.StringAttribute{
				Description: "Controls how the service path, route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior.",
				Optional:    true,
			},

			"paths": schema.ListAttribute{
				Description: "A list of paths that match this route.",
				Optional:    true,

				ElementType: types.StringType,
			},

			"preserve_host": schema.BoolAttribute{
				Description: "When matching a route via one of the `hosts` domain names, use the request `Host` header in the upstream request headers. If set to `false`, the upstream `Host` header will be that of the service's `host`.",
				Optional:    true,
			},

			"protocols": schema.ListAttribute{
				Description: "An array of the protocols this route should allow. See the [route Object](#route-object) section for a list of accepted protocols. When set to only `https`, HTTP requests are answered with an upgrade error. When set to only `http`, HTTPS requests are answered with an error.",
				Optional:    true,

				ElementType: types.StringType,
			},

			"regex_priority": schema.Int64Attribute{
				Description: "A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same `regex_priority`, the older one (lowest `created_at`) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).",
				Optional:    true,
			},

			"request_buffering": schema.BoolAttribute{
				Description: "Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.",
				Optional:    true,
			},

			"response_buffering": schema.BoolAttribute{
				Description: "Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.",
				Optional:    true,
			},

			"service": schema.SingleNestedAttribute{
				Description: "The service this route is associated to. This is where the route proxies traffic to.",
				Optional:    true,

				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "id",
						Optional:    true,
					},
				},
			},

			"snis": schema.ListAttribute{
				Description: "A list of SNIs that match this route when using stream routing.",
				Optional:    true,

				ElementType: types.StringType,
			},

			"strip_path": schema.BoolAttribute{
				Description: "When matching a route via one of the `paths`, strip the matching prefix from the upstream request URL.",
				Optional:    true,
			},

			"tags": schema.ListAttribute{
				Description: "An optional set of strings associated with the route for grouping and filtering.",
				Optional:    true,

				ElementType: types.StringType,
			},

			"updated_at": schema.Int64Attribute{
				Description: "Unix epoch when the resource was last updated.",
				Optional:    true,
			},

			"runtime_group_id": schema.StringAttribute{
				Description: "The ID of your runtime group. This variable is available in the stevan manager",
				Optional:    true,
			},

			"route_id": schema.StringAttribute{
				Description: "The unique identifier or the name of the route to retrieve.",
				Optional:    true,
			},
		},
	}
}

func (r *RouteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	apiClient, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = apiClient
}

func (r *RouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var dataModel RouteResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	RuntimeGroupId := dataModel.RuntimeGroupId.ValueString()
	RouteId := dataModel.RouteId.ValueString()

	route, err := r.client.Routes.GetRoute(RuntimeGroupId, RouteId)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling Routes.GetRoute",
			err.Error(),
		)

		return
	}

	dataModel.CreatedAt = utils.NullableInt64(route.GetCreatedAt())

	dataModel.Headers = utils.ToMap(ctx, route.Headers, utils.TypeAtPath(ctx, path.Root("headers"), resp.State, &resp.Diagnostics), &resp.Diagnostics)

	dataModel.Hosts = utils.ToList(ctx, route.Hosts, types.StringType, &resp.Diagnostics)

	dataModel.HttpsRedirectStatusCode = utils.NullableInt64(route.GetHttpsRedirectStatusCode())

	dataModel.Id = utils.NullableString(route.GetId())

	dataModel.Methods = utils.ToList(ctx, route.Methods, types.StringType, &resp.Diagnostics)

	dataModel.Name = utils.NullableString(route.GetName())

	dataModel.PathHandling = utils.NullableString(route.GetPathHandling())

	dataModel.Paths = utils.ToList(ctx, route.Paths, types.StringType, &resp.Diagnostics)

	dataModel.PreserveHost = utils.NullableBool(route.GetPreserveHost())

	dataModel.Protocols = utils.ToList(ctx, route.Protocols, types.StringType, &resp.Diagnostics)

	dataModel.RegexPriority = utils.NullableInt64(route.GetRegexPriority())

	dataModel.RequestBuffering = utils.NullableBool(route.GetRequestBuffering())

	dataModel.ResponseBuffering = utils.NullableBool(route.GetResponseBuffering())

	if route.Service != nil {
		dataModel.Service = utils.NullableObject(route.Service, service.Service{
			Id: utils.NullableString(route.GetService().GetId()),
		})
	}

	dataModel.Snis = utils.ToList(ctx, route.Snis, types.StringType, &resp.Diagnostics)

	dataModel.StripPath = utils.NullableBool(route.GetStripPath())

	dataModel.Tags = utils.ToList(ctx, route.Tags, types.StringType, &resp.Diagnostics)

	dataModel.UpdatedAt = utils.NullableInt64(route.GetUpdatedAt())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *RouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var dataModel RouteResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.Plan.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	RuntimeGroupId := dataModel.RuntimeGroupId.ValueString()

	requestBody := routes.Route{
		CreatedAt:               dataModel.CreatedAt.ValueInt64Pointer(),
		Headers:                 utils.FromTypesMapToMap[string](ctx, dataModel.Headers, &resp.Diagnostics),
		Hosts:                   utils.FromListToPrimitiveSlice[string](ctx, dataModel.Hosts, &resp.Diagnostics),
		HttpsRedirectStatusCode: dataModel.HttpsRedirectStatusCode.ValueInt64Pointer(),
		Id:                      dataModel.Id.ValueStringPointer(),
		Methods:                 utils.FromListToPrimitiveSlice[string](ctx, dataModel.Methods, &resp.Diagnostics),
		Name:                    dataModel.Name.ValueStringPointer(),
		PathHandling:            dataModel.PathHandling.ValueStringPointer(),
		Paths:                   utils.FromListToPrimitiveSlice[string](ctx, dataModel.Paths, &resp.Diagnostics),
		PreserveHost:            dataModel.PreserveHost.ValueBoolPointer(),
		Protocols:               utils.FromListToPrimitiveSlice[string](ctx, dataModel.Protocols, &resp.Diagnostics),
		RegexPriority:           dataModel.RegexPriority.ValueInt64Pointer(),
		RequestBuffering:        dataModel.RequestBuffering.ValueBoolPointer(),
		ResponseBuffering:       dataModel.ResponseBuffering.ValueBoolPointer(),

		Service: utils.NullableTfStateObject(dataModel.Service, func(from *service.Service) routes.Service {
			return routes.Service{
				Id: from.Id.ValueStringPointer(),
			}
		}),
		Snis:      utils.FromListToPrimitiveSlice[string](ctx, dataModel.Snis, &resp.Diagnostics),
		StripPath: dataModel.StripPath.ValueBoolPointer(),
		Tags:      utils.FromListToPrimitiveSlice[string](ctx, dataModel.Tags, &resp.Diagnostics),
		UpdatedAt: dataModel.UpdatedAt.ValueInt64Pointer(),
	}

	route, err := r.client.Routes.CreateRoute(RuntimeGroupId, requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Route",
			err.Error(),
		)

		return
	}

	dataModel.CreatedAt = utils.NullableInt64(route.GetCreatedAt())

	dataModel.Headers = utils.ToMap(ctx, route.Headers, utils.TypeAtPath(ctx, path.Root("headers"), resp.State, &resp.Diagnostics), &resp.Diagnostics)

	dataModel.Hosts = utils.ToList(ctx, route.Hosts, types.StringType, &resp.Diagnostics)

	dataModel.HttpsRedirectStatusCode = utils.NullableInt64(route.GetHttpsRedirectStatusCode())

	dataModel.Id = utils.NullableString(route.GetId())

	dataModel.Methods = utils.ToList(ctx, route.Methods, types.StringType, &resp.Diagnostics)

	dataModel.Name = utils.NullableString(route.GetName())

	dataModel.PathHandling = utils.NullableString(route.GetPathHandling())

	dataModel.Paths = utils.ToList(ctx, route.Paths, types.StringType, &resp.Diagnostics)

	dataModel.PreserveHost = utils.NullableBool(route.GetPreserveHost())

	dataModel.Protocols = utils.ToList(ctx, route.Protocols, types.StringType, &resp.Diagnostics)

	dataModel.RegexPriority = utils.NullableInt64(route.GetRegexPriority())

	dataModel.RequestBuffering = utils.NullableBool(route.GetRequestBuffering())

	dataModel.ResponseBuffering = utils.NullableBool(route.GetResponseBuffering())

	if route.Service != nil {
		dataModel.Service = utils.NullableObject(route.Service, service.Service{
			Id: utils.NullableString(route.GetService().GetId()),
		})
	}

	dataModel.Snis = utils.ToList(ctx, route.Snis, types.StringType, &resp.Diagnostics)

	dataModel.StripPath = utils.NullableBool(route.GetStripPath())

	dataModel.Tags = utils.ToList(ctx, route.Tags, types.StringType, &resp.Diagnostics)

	dataModel.UpdatedAt = utils.NullableInt64(route.GetUpdatedAt())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *RouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var dataModel = &RouteResourceModel{}
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	RuntimeGroupId := dataModel.RuntimeGroupId.ValueString()
	RouteId := dataModel.RouteId.ValueString()

	err := r.client.Routes.DeleteRoute(RuntimeGroupId, RouteId)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Route",
			err.Error(),
		)
	}
}

func (r *RouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var stateModel = &RouteResourceModel{}
	var dataModel = &RouteResourceModel{}
	utils.PopulateModelData(ctx, &stateModel, resp.Diagnostics, req.State.Get)
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.Plan.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	RuntimeGroupId := stateModel.RuntimeGroupId.ValueString()
	RouteId := stateModel.RouteId.ValueString()

	requestBody := routes.Route{
		CreatedAt:               dataModel.CreatedAt.ValueInt64Pointer(),
		Headers:                 utils.FromTypesMapToMap[string](ctx, dataModel.Headers, &resp.Diagnostics),
		Hosts:                   utils.FromListToPrimitiveSlice[string](ctx, dataModel.Hosts, &resp.Diagnostics),
		HttpsRedirectStatusCode: dataModel.HttpsRedirectStatusCode.ValueInt64Pointer(),
		Id:                      dataModel.Id.ValueStringPointer(),
		Methods:                 utils.FromListToPrimitiveSlice[string](ctx, dataModel.Methods, &resp.Diagnostics),
		Name:                    dataModel.Name.ValueStringPointer(),
		PathHandling:            dataModel.PathHandling.ValueStringPointer(),
		Paths:                   utils.FromListToPrimitiveSlice[string](ctx, dataModel.Paths, &resp.Diagnostics),
		PreserveHost:            dataModel.PreserveHost.ValueBoolPointer(),
		Protocols:               utils.FromListToPrimitiveSlice[string](ctx, dataModel.Protocols, &resp.Diagnostics),
		RegexPriority:           dataModel.RegexPriority.ValueInt64Pointer(),
		RequestBuffering:        dataModel.RequestBuffering.ValueBoolPointer(),
		ResponseBuffering:       dataModel.ResponseBuffering.ValueBoolPointer(),

		Service: utils.NullableTfStateObject(dataModel.Service, func(from *service.Service) routes.Service {
			return routes.Service{
				Id: from.Id.ValueStringPointer(),
			}
		}),
		Snis:      utils.FromListToPrimitiveSlice[string](ctx, dataModel.Snis, &resp.Diagnostics),
		StripPath: dataModel.StripPath.ValueBoolPointer(),
		Tags:      utils.FromListToPrimitiveSlice[string](ctx, dataModel.Tags, &resp.Diagnostics),
		UpdatedAt: dataModel.UpdatedAt.ValueInt64Pointer(),
	}

	route, err := r.client.Routes.UpsertRoute(RuntimeGroupId, RouteId, requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Route",
			err.Error(),
		)

		return
	}

	dataModel.CreatedAt = utils.NullableInt64(route.GetCreatedAt())

	dataModel.Headers = utils.ToMap(ctx, route.Headers, utils.TypeAtPath(ctx, path.Root("headers"), resp.State, &resp.Diagnostics), &resp.Diagnostics)

	dataModel.Hosts = utils.ToList(ctx, route.Hosts, types.StringType, &resp.Diagnostics)

	dataModel.HttpsRedirectStatusCode = utils.NullableInt64(route.GetHttpsRedirectStatusCode())

	dataModel.Id = utils.NullableString(route.GetId())

	dataModel.Methods = utils.ToList(ctx, route.Methods, types.StringType, &resp.Diagnostics)

	dataModel.Name = utils.NullableString(route.GetName())

	dataModel.PathHandling = utils.NullableString(route.GetPathHandling())

	dataModel.Paths = utils.ToList(ctx, route.Paths, types.StringType, &resp.Diagnostics)

	dataModel.PreserveHost = utils.NullableBool(route.GetPreserveHost())

	dataModel.Protocols = utils.ToList(ctx, route.Protocols, types.StringType, &resp.Diagnostics)

	dataModel.RegexPriority = utils.NullableInt64(route.GetRegexPriority())

	dataModel.RequestBuffering = utils.NullableBool(route.GetRequestBuffering())

	dataModel.ResponseBuffering = utils.NullableBool(route.GetResponseBuffering())

	if route.Service != nil {
		dataModel.Service = utils.NullableObject(route.Service, service.Service{
			Id: utils.NullableString(route.GetService().GetId()),
		})
	}

	dataModel.Snis = utils.ToList(ctx, route.Snis, types.StringType, &resp.Diagnostics)

	dataModel.StripPath = utils.NullableBool(route.GetStripPath())

	dataModel.Tags = utils.ToList(ctx, route.Tags, types.StringType, &resp.Diagnostics)

	dataModel.UpdatedAt = utils.NullableInt64(route.GetUpdatedAt())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *RouteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
