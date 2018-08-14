package kubernetes

import (
	"k8s.io/api/core/v1"
)

func expandPodAffinityTerm(ctrs []interface{}) v1.PodAffinityTerm {
	m := ctrs[0].(map[string]interface{})
	pat := v1.PodAffinityTerm{}
	if val, ok := m["topology_key"]; ok {
		pat.TopologyKey = val.(string)
	}

	if val, ok := m["label_selector"].([]interface{}); ok {
		// panic(spew.Sdump(val))
		pat.LabelSelector = expandLabelSelector(val)
	}

	return pat
}

func expandIgnDuringExec(ctrs []interface{}) []v1.WeightedPodAffinityTerm {
	wpas := make([]v1.WeightedPodAffinityTerm, len(ctrs))

	for i, ctr := range ctrs {
		m := ctr.(map[string]interface{})
		wpas[i] = v1.WeightedPodAffinityTerm{}

		if val, ok := m["weight"].(int); ok {
			wpas[i].Weight = int32(val)
		}

		if val, ok := m["pod_affinity_term"].([]interface{}); ok {
			wpas[i].PodAffinityTerm = expandPodAffinityTerm(val)
		}
	}

	return wpas
}

func expandPodAntiAffinity(ctrs []interface{}) *v1.PodAntiAffinity {
	// This doesn't seem right, expects a single affinity but passes in an array
	if len(ctrs) == 0 {
		return nil
	}

	aa := &v1.PodAntiAffinity{}

	ctr := ctrs[0].(map[string]interface{})
	if val, ok := ctr["preferred_during_scheduling_ignored_during_execution"].([]interface{}); ok && len(val) > 0 {
		aa.PreferredDuringSchedulingIgnoredDuringExecution = expandIgnDuringExec(val)
	}

	return aa
}

func expandAffinity(c []interface{}) *v1.Affinity {
	// This doesn't seem right, expects a single affinity but passes in an array
	a := &v1.Affinity{}
	ctr := c[0].(map[string]interface{})

	if val, ok := ctr["pod_anti_affinity"].([]interface{}); ok && len(val) > 0 {
		a.PodAntiAffinity = expandPodAntiAffinity(val)
	}
	return a
}
