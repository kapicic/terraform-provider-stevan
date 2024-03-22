package apiproduct

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stevan-sdk/pkg/apiproducts"
	"github.com/stevan-sdk/pkg/client"
	"github.com/terraform-provider-stevan/internal/utils"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &ApiProductResource{}
var _ resource.ResourceWithImportState = &ApiProductResource{}

// constructor
func NewApiProductResource() resource.Resource {
	return &ApiProductResource{}
}

// client wrapper
type ApiProductResource struct {
	client *client.Client
}

type ApiProductResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	PortalIds   types.List   `tfsdk:"portal_ids"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
	Labels      types.Map    `tfsdk:"labels"`
}

func (r *ApiProductResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_product"
}

func (r *ApiProductResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The API product ID.",
				Computed:    true,
				Optional:    true,
			},

			"name": schema.StringAttribute{
				Description: "The name of the API product",
				Required:    true,
			},

			"description": schema.StringAttribute{
				Description: "The description of the API product",
				Optional:    true,
			},

			"portal_ids": schema.ListAttribute{
				Description: "The list of portal identifiers which this API product is published to",
				Computed:    true,
				Optional:    true,

				ElementType: types.StringType,
			},

			"created_at": schema.StringAttribute{
				Description: "An ISO-8601 timestamp representation of entity creation date.",
				Computed:    true,
				Optional:    true,
			},

			"updated_at": schema.StringAttribute{
				Description: "An ISO-8601 timestamp representation of entity update date.",
				Computed:    true,
				Optional:    true,
			},

			"labels": schema.MapAttribute{
				Description: "description: A maximum of 5 user-defined labels are allowed on this resource.Keys must not start with kong, stevan, insomnia, mesh, kic or _, which are reserved for Kong.Keys are case-sensitive.",
				Optional:    true,

				ElementType: types.StringType,
			},
		},
	}
}

func (r *ApiProductResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ApiProductResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var dataModel ApiProductResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	Id := dataModel.Id.ValueString()

	apiProduct, err := r.client.ApiProducts.GetApiProduct(Id)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling ApiProducts.GetApiProduct",
			err.Error(),
		)

		return
	}

	dataModel.Id = utils.NullableString(apiProduct.GetId())

	dataModel.Name = utils.NullableString(apiProduct.GetName())

	dataModel.Description = utils.NullableString(apiProduct.GetDescription())

	dataModel.PortalIds = utils.ToList(ctx, apiProduct.PortalIds, types.StringType, &resp.Diagnostics)

	dataModel.CreatedAt = utils.NullableString(apiProduct.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(apiProduct.GetUpdatedAt())

	dataModel.Labels = utils.ToMap(ctx, apiProduct.Labels, utils.TypeAtPath(ctx, path.Root("labels"), resp.State, &resp.Diagnostics), &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *ApiProductResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var dataModel ApiProductResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.Plan.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := apiproducts.CreateApiProductDto{
		Name:        dataModel.Name.ValueStringPointer(),
		Description: dataModel.Description.ValueStringPointer(),
		Labels:      utils.FromTypesMapToMap[string](ctx, dataModel.Labels, &resp.Diagnostics),
	}

	apiProduct, err := r.client.ApiProducts.CreateApiProduct(requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating ApiProduct",
			err.Error(),
		)

		return
	}

	dataModel.Id = utils.NullableString(apiProduct.GetId())

	dataModel.Name = utils.NullableString(apiProduct.GetName())

	dataModel.Description = utils.NullableString(apiProduct.GetDescription())

	dataModel.PortalIds = utils.ToList(ctx, apiProduct.PortalIds, types.StringType, &resp.Diagnostics)

	dataModel.CreatedAt = utils.NullableString(apiProduct.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(apiProduct.GetUpdatedAt())

	dataModel.Labels = utils.ToMap(ctx, apiProduct.Labels, utils.TypeAtPath(ctx, path.Root("labels"), resp.State, &resp.Diagnostics), &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *ApiProductResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var dataModel = &ApiProductResourceModel{}
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	Id := dataModel.Id.ValueString()

	err := r.client.ApiProducts.DeleteApiProduct(Id)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting ApiProduct",
			err.Error(),
		)
	}
}

func (r *ApiProductResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var stateModel = &ApiProductResourceModel{}
	var dataModel = &ApiProductResourceModel{}
	utils.PopulateModelData(ctx, &stateModel, resp.Diagnostics, req.State.Get)
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.Plan.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	Id := stateModel.Id.ValueString()

	requestBody := apiproducts.UpdateApiProductDto{
		Name:        dataModel.Name.ValueStringPointer(),
		Description: dataModel.Description.ValueStringPointer(),
		Labels:      utils.FromTypesMapToMap[string](ctx, dataModel.Labels, &resp.Diagnostics),
		PortalIds:   utils.FromListToPrimitiveSlice[string](ctx, dataModel.PortalIds, &resp.Diagnostics),
	}

	apiProduct, err := r.client.ApiProducts.UpdateApiProduct(Id, requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating ApiProduct",
			err.Error(),
		)

		return
	}

	dataModel.Id = utils.NullableString(apiProduct.GetId())

	dataModel.Name = utils.NullableString(apiProduct.GetName())

	dataModel.Description = utils.NullableString(apiProduct.GetDescription())

	dataModel.PortalIds = utils.ToList(ctx, apiProduct.PortalIds, types.StringType, &resp.Diagnostics)

	dataModel.CreatedAt = utils.NullableString(apiProduct.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(apiProduct.GetUpdatedAt())

	dataModel.Labels = utils.ToMap(ctx, apiProduct.Labels, utils.TypeAtPath(ctx, path.Root("labels"), resp.State, &resp.Diagnostics), &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *ApiProductResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
