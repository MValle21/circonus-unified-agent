package kube_inventory

import (
	"context"
	"time"

	"github.com/circonus-labs/circonus-unified-agent/cua"
	v1 "github.com/ericchiang/k8s/apis/apps/v1"
)

func collectStatefulSets(ctx context.Context, acc cua.Accumulator, ki *KubernetesInventory) {
	list, err := ki.client.getStatefulSets(ctx)
	if err != nil {
		acc.AddError(err)
		return
	}
	for _, s := range list.Items {
		if err = ki.gatherStatefulSet(*s, acc); err != nil {
			acc.AddError(err)
			return
		}
	}
}

func (ki *KubernetesInventory) gatherStatefulSet(s v1.StatefulSet, acc cua.Accumulator) error {
	status := s.Status
	fields := map[string]interface{}{
		"created":             time.Unix(s.Metadata.CreationTimestamp.GetSeconds(), int64(s.Metadata.CreationTimestamp.GetNanos())).UnixNano(),
		"generation":          *s.Metadata.Generation,
		"replicas":            *status.Replicas,
		"replicas_current":    *status.CurrentReplicas,
		"replicas_ready":      *status.ReadyReplicas,
		"replicas_updated":    *status.UpdatedReplicas,
		"spec_replicas":       *s.Spec.Replicas,
		"observed_generation": *s.Status.ObservedGeneration,
	}
	tags := map[string]string{
		"statefulset_name": *s.Metadata.Name,
		"namespace":        *s.Metadata.Namespace,
	}
	for key, val := range s.GetSpec().GetSelector().GetMatchLabels() {
		if ki.selectorFilter.Match(key) {
			tags["selector_"+key] = val
		}
	}

	acc.AddFields(statefulSetMeasurement, fields, tags)

	return nil
}
