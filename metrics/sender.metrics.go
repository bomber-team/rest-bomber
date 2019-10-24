package metrics

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"restbomber/configuration"

	"gitlab.com/truecord_team/common/contents"
)

type SenderMetrics struct {
	Metrics        *MetricsScenario
	GlobalConfig   *configuration.GlobalConfiguration
	SenderActivate chan bool
}

/*SetupGlobalConfig - install global configuration*/
func (sender *SenderMetrics) SetupGlobalConfig(glbc *configuration.GlobalConfiguration) {
	sender.GlobalConfig = glbc
}

/*WaitSend - waiting when channel will fill by flag*/
func (sender *SenderMetrics) WaitSend() {
	go func(sender *SenderMetrics) {
		for {
			<-sender.SenderActivate
			sender.sendMetricsToBackend()
		}
	}(sender)
}

func (sender *SenderMetrics) sendMetricsToBackend() {
	metrics := sender.Metrics.GetWithEraseBufferMetrics()
	body := map[string]interface{}{
		"status": 0,
		"data":   metrics,
	}

	jsonString, err1 := json.Marshal(body)
	if err1 != nil {
		log.Fatalf(err1.Error())
	}

	b := bytes.NewBuffer(jsonString)

	resp, err := http.Post(sender.GlobalConfig.BomberConfigurationService.MetricsAddress, contents.JSON, b)
	if err != nil {
		log.Println("can not send metrics. Save local dump")
	}

	if resp.StatusCode != 200 {
		log.Println("error send metrics. Save local dump")
	}
}
