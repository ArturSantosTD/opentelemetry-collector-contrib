// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
		{
			name:        "filter_set_include",
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "filter_set_exclude",
			resAttrsSet: testDataSetAll,
			expectEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopSettings(receivertest.NopType)
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, tt.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsChangeCountDataPoint(ts, 1, "vcs.repository.url.full-val", AttributeVcsChangeStateOpen, "vcs.repository.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsChangeDurationDataPoint(ts, 1, "vcs.repository.url.full-val", "vcs.repository.name-val", "vcs.ref.head.name-val", AttributeVcsChangeStateOpen)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsChangeTimeToApprovalDataPoint(ts, 1, "vcs.repository.url.full-val", "vcs.repository.name-val", "vcs.ref.head.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsChangeTimeToMergeDataPoint(ts, 1, "vcs.repository.url.full-val", "vcs.repository.name-val", "vcs.ref.head.name-val")

			allMetricsCount++
			mb.RecordVcsContributorCountDataPoint(ts, 1, "vcs.repository.url.full-val", "vcs.repository.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRefCountDataPoint(ts, 1, "vcs.repository.url.full-val", "vcs.repository.name-val", AttributeVcsRefHeadTypeBranch)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRefLinesDeltaDataPoint(ts, 1, "vcs.repository.url.full-val", "vcs.repository.name-val", "vcs.ref.head.name-val", AttributeVcsRefHeadTypeBranch, AttributeVcsLineChangeTypeAdded)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRefRevisionsDeltaDataPoint(ts, 1, "vcs.repository.url.full-val", "vcs.repository.name-val", "vcs.ref.head.name-val", AttributeVcsRefHeadTypeBranch, AttributeVcsRevisionDeltaDirectionAhead)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRefTimeDataPoint(ts, 1, "vcs.repository.url.full-val", "vcs.repository.name-val", "vcs.ref.head.name-val", AttributeVcsRefHeadTypeBranch)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordVcsRepositoryCountDataPoint(ts, 1)

			rb := mb.NewResourceBuilder()
			rb.SetOrganizationName("organization.name-val")
			rb.SetVcsVendorName("vcs.vendor.name-val")
			res := rb.Emit()
			metrics := mb.Emit(WithResource(res))

			if tt.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if tt.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if tt.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "vcs.change.count":
					assert.False(t, validatedMetrics["vcs.change.count"], "Found a duplicate in the metrics slice: vcs.change.count")
					validatedMetrics["vcs.change.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of changes (pull requests) in a repository, categorized by their state (either open or merged).", ms.At(i).Description())
					assert.Equal(t, "{change}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.change.state")
					assert.True(t, ok)
					assert.Equal(t, "open", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.repository.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.name-val", attrVal.Str())
				case "vcs.change.duration":
					assert.False(t, validatedMetrics["vcs.change.duration"], "Found a duplicate in the metrics slice: vcs.change.duration")
					validatedMetrics["vcs.change.duration"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The time duration a change (pull request/merge request/changelist) has been in an open state.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.repository.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.ref.head.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.ref.head.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.change.state")
					assert.True(t, ok)
					assert.Equal(t, "open", attrVal.Str())
				case "vcs.change.time_to_approval":
					assert.False(t, validatedMetrics["vcs.change.time_to_approval"], "Found a duplicate in the metrics slice: vcs.change.time_to_approval")
					validatedMetrics["vcs.change.time_to_approval"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The amount of time it took a change (pull request) to go from open to approved.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.repository.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.ref.head.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.ref.head.name-val", attrVal.Str())
				case "vcs.change.time_to_merge":
					assert.False(t, validatedMetrics["vcs.change.time_to_merge"], "Found a duplicate in the metrics slice: vcs.change.time_to_merge")
					validatedMetrics["vcs.change.time_to_merge"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The amount of time it took a change (pull request) to go from open to merged.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.repository.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.ref.head.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.ref.head.name-val", attrVal.Str())
				case "vcs.contributor.count":
					assert.False(t, validatedMetrics["vcs.contributor.count"], "Found a duplicate in the metrics slice: vcs.contributor.count")
					validatedMetrics["vcs.contributor.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of unique contributors to a repository.", ms.At(i).Description())
					assert.Equal(t, "{contributor}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.repository.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.name-val", attrVal.Str())
				case "vcs.ref.count":
					assert.False(t, validatedMetrics["vcs.ref.count"], "Found a duplicate in the metrics slice: vcs.ref.count")
					validatedMetrics["vcs.ref.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of refs of type branch in a repository.", ms.At(i).Description())
					assert.Equal(t, "{ref}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.repository.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.ref.head.type")
					assert.True(t, ok)
					assert.Equal(t, "branch", attrVal.Str())
				case "vcs.ref.lines_delta":
					assert.False(t, validatedMetrics["vcs.ref.lines_delta"], "Found a duplicate in the metrics slice: vcs.ref.lines_delta")
					validatedMetrics["vcs.ref.lines_delta"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of lines added/removed in a ref (branch) relative to the default branch (trunk).", ms.At(i).Description())
					assert.Equal(t, "{line}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.repository.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.ref.head.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.ref.head.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.ref.head.type")
					assert.True(t, ok)
					assert.Equal(t, "branch", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.line_change.type")
					assert.True(t, ok)
					assert.Equal(t, "added", attrVal.Str())
				case "vcs.ref.revisions_delta":
					assert.False(t, validatedMetrics["vcs.ref.revisions_delta"], "Found a duplicate in the metrics slice: vcs.ref.revisions_delta")
					validatedMetrics["vcs.ref.revisions_delta"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of revisions (commits) a ref (branch) is ahead/behind the branch from trunk (default).", ms.At(i).Description())
					assert.Equal(t, "{revision}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.repository.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.ref.head.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.ref.head.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.ref.head.type")
					assert.True(t, ok)
					assert.Equal(t, "branch", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.revision_delta.direction")
					assert.True(t, ok)
					assert.Equal(t, "ahead", attrVal.Str())
				case "vcs.ref.time":
					assert.False(t, validatedMetrics["vcs.ref.time"], "Found a duplicate in the metrics slice: vcs.ref.time")
					validatedMetrics["vcs.ref.time"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Time a ref (branch) created from the default branch (trunk) has existed. The `vcs.ref.head.type` attribute will always be `branch`.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("vcs.repository.url.full")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.url.full-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.repository.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.repository.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.ref.head.name")
					assert.True(t, ok)
					assert.Equal(t, "vcs.ref.head.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("vcs.ref.head.type")
					assert.True(t, ok)
					assert.Equal(t, "branch", attrVal.Str())
				case "vcs.repository.count":
					assert.False(t, validatedMetrics["vcs.repository.count"], "Found a duplicate in the metrics slice: vcs.repository.count")
					validatedMetrics["vcs.repository.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of repositories in an organization.", ms.At(i).Description())
					assert.Equal(t, "{repository}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}
