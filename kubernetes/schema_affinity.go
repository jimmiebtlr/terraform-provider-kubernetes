package kubernetes

import "github.com/hashicorp/terraform/helper/schema"

func podAffinityTermFields(_ bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"topology_key": {
			Type:        schema.TypeString,
			Description: "idk",
			Optional:    true,
			Default:     1,
		},
		"label_selector": {
			Type:        schema.TypeList,
			Description: "No clue what this does",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: labelSelectorFields(),
			},
		},
	}
	return s
}

func preferredDuringSchedulingIgnoredDuringExecutionFields(isUpdatable bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"weight": {
			Type:        schema.TypeInt,
			Description: "Wieght of the anti affinity.",
			Optional:    true,
			Default:     1,
		},
		"pod_affinity_term": {
			Type:        schema.TypeList,
			Description: "Anti affinity specifies pods that shouldn't exist on a machine together.",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: podAffinityTermFields(isUpdatable),
			},
		},
	}
	return s
}

func podAntiAffinityFields(isUpdatable bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"preferred_during_scheduling_ignored_during_execution_fields": {
			Type:        schema.TypeList,
			Description: "Anti affinity specifies pods that shouldn't exist on a machine together.",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: preferredDuringSchedulingIgnoredDuringExecutionFields(isUpdatable),
			},
		},
	}
	return s
}

func affinityFields(isUpdatable bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"pod_anti_affinity": {
			Type:        schema.TypeList,
			Description: "Anti affinity specifies pods that shouldn't exist on a machine together.",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: podAntiAffinityFields(isUpdatable),
			},
		},
	}
	return s
}
