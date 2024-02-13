package flyte

import (
	"context"

	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"sigs.k8s.io/structured-merge-diff/v4/schema"
)

func resourceExecution() *schema.Resource {
	return &schema.Resource{
		Create: resourceExecutionCreate,
		Read:   resourceExecutionTypeRead,
		Update: resourceExecutionTypeUpdate,
		Delete: resourceExecutionTypeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"iam_Role_arn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kube_service_account": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project": {
				Type:     schema.TypeString,
				Required: true,
			},
			"task": {
				Type:     schema.TypeString,
				Required: true,
			},
			"workflow": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"spec": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_pool": {
				Type:     schema.TypeString,
				Required: true,
			},
			"env": {
				Type:     schema.Map,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceExecutionCreate(d *schema.ResourceData, meta interface{}) error {
	apiClient := meta.(service.AdminServiceClient)

	_, err := apiClient.CreateExecution(context.Background(), &admin.ExecutionCreateRequest{})
	if err != nil {
		return nil
	}

	return nil
}

func resourceExecutionTypeRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceExecutionTypeUpdate(d *schema.ResourceData, meta interface{}) error {
	apiClient := meta.(service.AdminServiceClient)
	_, err := apiClient.UpdateExecution(context.Background(), &admin.ExecutionUpdateRequest{})
	if err != nil {
		return nil
	}
	return nil
}

func resourceExecutionTypeDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
