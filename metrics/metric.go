package metrics

import "time"

type (
	MetricsWithOneAttack struct {
		Time int64
		Date time.Time
	}

	MetricsStage struct {
		MetricsByAttack []MetricsWithOneAttack
	}

	MetricsScenario struct {
		MetricsStages []MetricsStage
	}
)

/*CreateMetricsBufferForStageSize - create metrics by amount stages*/
func (scen *MetricsScenario) CreateMetricsBufferForStageSize(amountStages int) {
	scen.MetricsStages = make([]MetricsStage, amountStages)
}

/*CreateBufferForOneStageByIndex - create buffer for one attack by amount metrics size*/
func (scen *MetricsScenario) CreateBufferForOneStageByIndex(indexStage, amountMetrics int) {
	scen.MetricsStages[indexStage].MetricsByAttack = make([]MetricsWithOneAttack, amountMetrics)
}

/*GetWithEraseBufferMetrics - get all metrics and remove current metrics buffer*/
func (scen *MetricsScenario) GetWithEraseBufferMetrics() *MetricsScenario {
	metrics := scen.MetricsStages
	scen = nil
	return &MetricsScenario{
		MetricsStages: metrics,
	}
}
