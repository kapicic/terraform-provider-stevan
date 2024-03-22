package runtimegroup

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stevan-sdk/pkg/client"
	"github.com/stevan-sdk/pkg/runtimegroups"
	"github.com/terraform-provider-stevan/internal/shared/models/config"
	"github.com/terraform-provider-stevan/internal/utils"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &RunTimeGroupResource{}
var _ resource.ResourceWithImportState = &RunTimeGroupResource{}

// constructor
func NewRunTimeGroupResource() resource.Resource {
	return &RunTimeGroupResource{}
}

// client wrapper
type RunTimeGroupResource struct {
	client *client.Client
}

type RunTimeGroupResourceModel struct {
	Id          types.String   `tfsdk:"id"`
	Name        types.String   `tfsdk:"name"`
	Description types.String   `tfsdk:"description"`
	Labels      types.Map      `tfsdk:"labels"`
	Config      *config.Config `tfsdk:"config"`
	CreatedAt   types.String   `tfsdk:"created_at"`
	UpdatedAt   types.String   `tfsdk:"updated_at"`
	ClusterType types.String   `tfsdk:"cluster_type"`
	AuthType    types.String   `tfsdk:"auth_type"`
}

func (r *RunTimeGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_run_time_group"
}

func (r *RunTimeGroupResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The runtime group ID.",
				Computed:    true,
				Optional:    true,
			},

			"name": schema.StringAttribute{
				Description: "The name of the runtime group.",
				Required:    true,
			},

			"description": schema.StringAttribute{
				Description: "The description of the runtime group in stevan.",
				Optional:    true,
			},

			"labels": schema.MapAttribute{
				Description: "Labels to facilitate tagged search on runtime groups. Keys must be of length 1-63 characters, and cannot start with 'kong', 'stevan', 'mesh', 'kic', or '_'.",
				Optional:    true,

				ElementType: types.StringType,
			},

			"config": schema.SingleNestedAttribute{
				Description: "CP configuration object for related access endpoints.",
				Computed:    true,
				Optional:    true,

				Attributes: map[string]schema.Attribute{
					"control_plane_endpoint": schema.StringAttribute{
						Description: "Control Plane Endpoint.",
						Optional:    true,
					},

					"telemetry_endpoint": schema.StringAttribute{
						Description: "Telemetry Endpoint.",
						Optional:    true,
					},
				},
			},

			"created_at": schema.StringAttribute{
				Description: "An ISO-8604 timestamp representation of runtime group creation date.",
				Computed:    true,
				Optional:    true,
			},

			"updated_at": schema.StringAttribute{
				Description: "An ISO-8604 timestamp representation of runtime group update date.",
				Computed:    true,
				Optional:    true,
			},

			"cluster_type": schema.StringAttribute{
				Description: "The ClusterType value of the cluster associated with the Runtime Group.",
				Optional:    true,
			},

			"auth_type": schema.StringAttribute{
				Description: "The auth type value of the cluster associated with the Runtime Group.",
				Optional:    true,
			},
		},
	}
}

func (r *RunTimeGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RunTimeGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var dataModel RunTimeGroupResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	Id := dataModel.Id.ValueString()

	runTimeGroup, err := r.client.RuntimeGroups.GetRuntimeGroup(Id)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling RuntimeGroups.GetRuntimeGroup",
			err.Error(),
		)

		return
	}

	dataModel.Id = utils.NullableString(runTimeGroup.GetId())

	dataModel.Name = utils.NullableString(runTimeGroup.GetName())

	dataModel.Description = utils.NullableString(runTimeGroup.GetDescription())

	dataModel.Labels = utils.ToMap(ctx, runTimeGroup.Labels, utils.TypeAtPath(ctx, path.Root("labels"), resp.State, &resp.Diagnostics), &resp.Diagnostics)

	if runTimeGroup.Config != nil {
		dataModel.Config = utils.NullableObject(runTimeGroup.Config, config.Config{
			ControlPlaneEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetControlPlaneEndpoint()),

			TelemetryEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetTelemetryEndpoint()),
		})
	}

	dataModel.CreatedAt = utils.NullableString(runTimeGroup.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(runTimeGroup.GetUpdatedAt())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *RunTimeGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var dataModel RunTimeGroupResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.Plan.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := runtimegroups.CreateRuntimeGroupRequest{
		Name:        dataModel.Name.ValueStringPointer(),
		Description: dataModel.Description.ValueStringPointer(),
		ClusterType: utils.Pointer(runtimegroups.ClusterType(dataModel.ClusterType.ValueString())),
		AuthType:    utils.Pointer(runtimegroups.AuthType(dataModel.AuthType.ValueString())),
		Labels:      utils.FromTypesMapToMap[string](ctx, dataModel.Labels, &resp.Diagnostics),
	}

	runTimeGroup, err := r.client.RuntimeGroups.CreateRuntimeGroup(requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating RunTimeGroup",
			err.Error(),
		)

		return
	}

	dataModel.Id = utils.NullableString(runTimeGroup.GetId())

	dataModel.Name = utils.NullableString(runTimeGroup.GetName())

	dataModel.Description = utils.NullableString(runTimeGroup.GetDescription())

	dataModel.Labels = utils.ToMap(ctx, runTimeGroup.Labels, utils.TypeAtPath(ctx, path.Root("labels"), resp.State, &resp.Diagnostics), &resp.Diagnostics)

	if runTimeGroup.Config != nil {
		dataModel.Config = utils.NullableObject(runTimeGroup.Config, config.Config{
			ControlPlaneEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetControlPlaneEndpoint()),

			TelemetryEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetTelemetryEndpoint()),
		})
	}

	dataModel.CreatedAt = utils.NullableString(runTimeGroup.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(runTimeGroup.GetUpdatedAt())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *RunTimeGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var dataModel = &RunTimeGroupResourceModel{}
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	Id := dataModel.Id.ValueString()

	err := r.client.RuntimeGroups.DeleteRuntimeGroup(Id)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting RunTimeGroup",
			err.Error(),
		)
	}
}

func (r *RunTimeGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var stateModel = &RunTimeGroupResourceModel{}
	var dataModel = &RunTimeGroupResourceModel{}
	utils.PopulateModelData(ctx, &stateModel, resp.Diagnostics, req.State.Get)
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.Plan.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	Id := stateModel.Id.ValueString()

	requestBody := runtimegroups.UpdateRuntimeGroupRequest{
		Name:        dataModel.Name.ValueStringPointer(),
		Description: dataModel.Description.ValueStringPointer(),
		AuthType:    utils.Pointer(runtimegroups.AuthType_1(dataModel.AuthType.ValueString())),
		Labels:      utils.FromTypesMapToMap[string](ctx, dataModel.Labels, &resp.Diagnostics),
	}

	runTimeGroup, err := r.client.RuntimeGroups.UpdateRuntimeGroup(Id, requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating RunTimeGroup",
			err.Error(),
		)

		return
	}

	dataModel.Id = utils.NullableString(runTimeGroup.GetId())

	dataModel.Name = utils.NullableString(runTimeGroup.GetName())

	dataModel.Description = utils.NullableString(runTimeGroup.GetDescription())

	dataModel.Labels = utils.ToMap(ctx, runTimeGroup.Labels, utils.TypeAtPath(ctx, path.Root("labels"), resp.State, &resp.Diagnostics), &resp.Diagnostics)

	if runTimeGroup.Config != nil {
		dataModel.Config = utils.NullableObject(runTimeGroup.Config, config.Config{
			ControlPlaneEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetControlPlaneEndpoint()),

			TelemetryEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetTelemetryEndpoint()),
		})
	}

	dataModel.CreatedAt = utils.NullableString(runTimeGroup.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(runTimeGroup.GetUpdatedAt())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *RunTimeGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
