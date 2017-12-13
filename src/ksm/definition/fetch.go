package definition

import (
	"fmt"
)

// RawValue is just any value from a raw metric.
type RawValue interface{}

// RawMetrics is a map of RawValue indexed by metric name.
type RawMetrics map[string]RawValue

// FetchedValue is just any value from an already fetched metric.
type FetchedValue interface{}

// FetchedValues is a map of FetchedValue indexed by metric name.
type FetchedValues map[string]FetchedValue

// FetchFunc fetches values or values from raw metric groups.
// Return FetchedValues if you want to prototype metrics.
type FetchFunc func(groupLabel, entityID string, groups RawGroups) (FetchedValue, error)

// RawGroups are grouped raw metrics.
type RawGroups map[string]map[string]RawMetrics

// FromRaw fetches metrics from raw metrics. Is the most simple use case.
func FromRaw(metricKey string) FetchFunc {
	return func(groupLabel, entityID string, groups RawGroups) (FetchedValue, error) {
		g, ok := groups[groupLabel]
		if !ok {
			return nil, fmt.Errorf("FromRaw: group not found: %v", groupLabel)
		}

		e, ok := g[entityID]
		if !ok {
			return nil, fmt.Errorf("FromRaw: entity not found. Group: %v, EntityID: %v", groupLabel, entityID)
		}

		value, ok := e[metricKey]
		if !ok {
			return nil, fmt.Errorf("FromRaw: metric not found. Group: %v, EntityID: %v, Metric: %v", groupLabel, entityID, metricKey)
		}

		return value, nil
	}
}
