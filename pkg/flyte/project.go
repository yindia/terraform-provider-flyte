package flyte

import (
	"context"
	"fmt"
	"regexp"

	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"sigs.k8s.io/structured-merge-diff/v4/schema"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectTypeRead,
		Update: resourceProjectTypeUpdate,
		Delete: resourceProjectTypeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if len(value) > 64 {
						errors = append(errors, fmt.Errorf(
							"%q cannot be longer than 64 characters", k))
					}
					if !regexp.MustCompile(`^[\w+=,.@-]*$`).MatchString(value) {
						errors = append(errors, fmt.Errorf(
							"%q must match [\\w+=,.@-]", k))
					}
					return
				},
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if len(value) > 64 {
						errors = append(errors, fmt.Errorf(
							"%q cannot be longer than 64 characters", k))
					}
					if !regexp.MustCompile(`^[\w+=,.@-]*$`).MatchString(value) {
						errors = append(errors, fmt.Errorf(
							"%q must match [\\w+=,.@-]", k))
					}
					return
				},
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
			},
			"labels": {
				Type: schema.Map,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: false,
			},
			"status": {
				Type:     schema.TypeInt,
				Required: false,
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, meta interface{}) error {
	apiClient := meta.(service.AdminServiceClient)

	name := d.Get("name").(string)
	id := d.Get("id").(string)
	description := d.Get("description").(string)
	state := d.Get("state").(int32)
	labels := d.Get("labels").(map[string]string)

	_, err := apiClient.RegisterProject(context.Background(), &admin.ProjectRegisterRequest{
		Project: &admin.Project{
			Id:          id,
			Name:        name,
			Description: description,
			Labels: &admin.Labels{
				Values: labels,
			},
			State: state,
		},
	})
	if err != nil {
		return nil
	}

	return nil
}

func resourceProjectTypeRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceProjectTypeUpdate(d *schema.ResourceData, meta interface{}) error {
	apiClient := meta.(service.AdminServiceClient)
	name := d.Get("name").(string)
	id := d.Get("id").(string)
	description := d.Get("description").(string)
	labels := d.Get("labels").(map[string]string)
	state := d.Get("state").(int32)

	apiClient.UpdateProject(context.Background(), &admin.Project{
		Id:          id,
		Name:        name,
		Description: description,
		Labels: &admin.Labels{
			Values: labels,
		},
		State: state,
	})

	return nil
}

func resourceProjectTypeDelete(d *schema.ResourceData, meta interface{}) error {
	apiClient := meta.(service.AdminServiceClient)

	name := d.Get("name").(string)
	id := d.Get("id").(string)
	description := d.Get("description").(string)
	labels := d.Get("labels").(map[string]string)

	apiClient.UpdateProject(context.Background(), &admin.Project{
		Id:          id,
		Name:        name,
		Description: description,
		Labels: &admin.Labels{
			Values: labels,
		},
		State: admin.Project_ARCHIVED,
	})

	fmt.Printf("Archived project %s", name)

	return nil
}
