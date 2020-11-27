package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandKubernetesAzureAuthentication(values interface{}) *octopusdeploy.KubernetesAzureAuthentication {
	flattenedValues := values.([]interface{})
	flattenedAuthentication := flattenedValues[0].(map[string]interface{})

	authentication := &octopusdeploy.KubernetesAzureAuthentication{
		ClusterName:          flattenedAuthentication["cluster_name"].(string),
		ClusterResourceGroup: flattenedAuthentication["cluster_resource_group"].(string),
	}

	authentication.AccountID = flattenedAuthentication["account_id"].(string)
	authentication.AuthenticationType = "KubernetesAzure"

	return authentication
}

func flattenKubernetesAzureAuthentication(kubernetesAzureAuthentication *octopusdeploy.KubernetesAzureAuthentication) []interface{} {
	if kubernetesAzureAuthentication == nil {
		return nil
	}

	return []interface{}{map[string]interface{}{
		"account_id":             kubernetesAzureAuthentication.AccountID,
		"cluster_name":           kubernetesAzureAuthentication.ClusterName,
		"cluster_resource_group": kubernetesAzureAuthentication.ClusterResourceGroup,
	}}
}

func getKubernetesAzureAuthenticationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Optional: true,
			Type:     schema.TypeString,
		},
		"cluster_name": {
			Optional: true,
			Type:     schema.TypeString,
		},
		"cluster_resource_group": {
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}
