package definition

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestFromRawFetchesProperly(t *testing.T) {
	raw := RawGroups{
		"group1": {
			"entity1": {
				"metric_name_1": "metric_value_1",
				"metric_name_2": "metric_value_2",
			},
			"entity2": {
				"metric_name_3": "metric_value_3",
				"metric_name_4": "metric_value_4",
				"metric_name_5": "metric_value_5",
			},
		},
	}

	v, err := FromRaw("metric_name_3")("group1", "entity2", raw)
	assert.NoError(t, err)
	assert.Equal(t, "metric_value_3", v)
}

func TestFromRawErrorsOnNotFound(t *testing.T) {
	raw := RawGroups{
		"group1": {
			"entity1": {
				"metric_name_1": "metric_value_1",
				"metric_name_2": "metric_value_2",
			},
			"entity2": {
				"metric_name_3": "metric_value_3",
				"metric_name_4": "metric_value_4",
				"metric_name_5": "metric_value_5",
			},
		},
	}

	v, err := FromRaw("metric_name_3")("nonExistingGroup", "entity2", raw)
	assert.EqualError(t, err, "FromRaw: group not found: nonExistingGroup")
	assert.Nil(t, v)

	v, err = FromRaw("metric_name_3")("group1", "nonExistingEntity", raw)
	assert.EqualError(t, err, "FromRaw: entity not found. SpecGroup: group1, EntityID: nonExistingEntity")
	assert.Nil(t, v)

	v, err = FromRaw("non_existing_metric")("group1", "entity2", raw)
	assert.EqualError(t, err, "FromRaw: metric not found. SpecGroup: group1, EntityID: entity2, Metric: non_existing_metric")
	assert.Nil(t, v)
}

func TestTransform(t *testing.T) {
	raw := RawGroups{
		"group1": {
			"entity1": {
				"metric_name_1": "metric_value_1",
				"metric_name_2": "metric_value_2",
			},
			"entity2": {
				"metric_name_3": "metric_value_3",
				"metric_name_4": "metric_value_4",
				"metric_name_5": "metric_value_5",
			},
		},
	}

	v, err := Transform(FromRaw("metric_name_3"),
		func(in FetchedValue) FetchedValue {
			return strings.ToUpper(in.(string))
		})("group1", "entity2", raw)
	assert.NoError(t, err)
	assert.Equal(t, "METRIC_VALUE_3", v)
}

func TestTransformBypassesError(t *testing.T) {
	raw := RawGroups{
		"group1": {
			"entity1": {
				"metric_name_1": "metric_value_1",
				"metric_name_2": "metric_value_2",
			},
			"entity2": {
				"metric_name_3": "metric_value_3",
				"metric_name_4": "metric_value_4",
				"metric_name_5": "metric_value_5",
			},
		},
	}

	transformFunc := func(in FetchedValue) FetchedValue {
		return in
	}

	v, err := Transform(FromRaw("metric_name_3"), transformFunc)("nonExistingGroup", "entity2", raw)
	assert.EqualError(t, err, "FromRaw: group not found: nonExistingGroup")
	assert.Nil(t, v)

	v, err = Transform(FromRaw("metric_name_3"), transformFunc)("group1", "nonExistingEntity", raw)
	assert.EqualError(t, err, "FromRaw: entity not found. SpecGroup: group1, EntityID: nonExistingEntity")
	assert.Nil(t, v)

	v, err = Transform(FromRaw("non_existing_metric"), transformFunc)("group1", "entity2", raw)
	assert.EqualError(t, err, "FromRaw: metric not found. SpecGroup: group1, EntityID: entity2, Metric: non_existing_metric")
	assert.Nil(t, v)
}