package core

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/bomber-team/bomber-proto-contracts/golang/rest_contracts"
	"github.com/bomber-team/bomber-proto-contracts/golang/system"
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type Core struct {
	publisher              *nats_listener.Publisher
	config                 *nats_listener.NatsConnectionConfiguration
	currentStatusBomber    system.StatusBomber
	dataAttack             []http.Request
	httpClient             *http.Client
	resultsAttack          map[int32]int64 // amount statuses per status
	resultTimeouts         int64           // amount time out requests
	resultTimesForRequests []int64         // amount ms for one request
	attackReady            bool            // ready for attack?
	bomberIp               string
	formId                 string
}

func (core *Core) CheckReady() bool {
	return core.attackReady
}

const (
	topicName    = "bomber.results"
	bomberResult = "bomber.result"
)

const (
	currentWorkers = 100
)

func NewCore(conn *nats.Conn) *Core {
	return &Core{
		publisher:           nats_listener.NewPublisher(conn),
		currentStatusBomber: system.StatusBomber_UP,
	}
}

type RequestPayload struct {
	Request *http.Request
	Id      int
}

func (core *Core) preparingBody(bodyParams []*rest_contracts.BodyParam) ([]byte, error) {
	var resultBody map[string]interface{}
	for _, value := range bodyParams {
		if value.IsGenerated {
			continue // TODO: Need change to call generating
		} else {
			resultBody[value.Name] = value
		}
	}
	resultMarshaled, err := json.Marshal(resultBody)
	if err != nil {
		logrus.Error("error whilte marshaled body..")
		return nil, err
	}
	return resultMarshaled, nil
}

func (core *Core) prepareRequestParams(requestParams []*rest_contracts.RequestParam) string {
	var resultUrlQueries string = "?"
	for _, value := range requestParams {
		if value.IsGeneratorNeed {
			continue // TODO: need change to call generating
		} else {
			resultUrlQueries += value.Name + "=" + value.Value
		}
	}
	return resultUrlQueries
}

func (core *Core) enhancedHeadersInRequest(request *http.Request, task rest_contracts.Task) *http.Request {
	for key, value := range task.Schema.Headers {
		request.Header.Set(key, value)
	}
	return request
}

func (core *Core) preparingRequest(restTask *rest_contracts.Task) (*http.Request, error) {
	body, err := core.preparingBody(restTask.Schema.Body)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(body)
	urlParams := core.prepareRequestParams(restTask.Schema.Request)
	newRequest, err := http.NewRequest(restTask.Script.RequestMethod, restTask.Script.Address+urlParams, reader)
	if err != nil {
		logrus.Error("Error while building new request")
		return nil, err
	}
	return core.enhancedHeadersInRequest(newRequest, *restTask), nil
}

func (core *Core) cleanCurrentResults() {
	core.dataAttack = []http.Request{}
	core.resultTimeouts = 0
	core.resultTimesForRequests = []int64{}
	core.resultsAttack = map[int32]int64{}
	core.attackReady = false
}

func (core *Core) PreparingData(task rest_contracts.Task) {
	core.cleanCurrentResults()
	var index int64 = 0
	amountRequests := task.Script.Config.Rps * task.Script.Config.Time
	resultSliceRequests := make([]http.Request, amountRequests)
	core.resultTimesForRequests = make([]int64, amountRequests)
	for ; index < amountRequests; index++ {
		newRequest, errFormRequest := core.preparingRequest(&task)
		if errFormRequest != nil {
			logrus.Error("Can not forming request: ", errFormRequest)
			continue
		}
		resultSliceRequests[index] = *newRequest
	}
	core.dataAttack = resultSliceRequests
	core.attackReady = true
}

func (core *Core) runWorkers(task chan RequestPayload, completed chan bool) {
	for {
		select {
		case newRequest := <-task:
			timeStart := time.Now()
			resp, err := core.httpClient.Do(newRequest.Request)
			if err != nil {
				logrus.Error("Error while executing request: ", err)
				core.resultTimeouts++
				continue
			}
			durationTime := time.Since(timeStart)
			core.resultsAttack[int32(resp.StatusCode)]++
			core.resultTimesForRequests[newRequest.Id] = durationTime.Milliseconds()
		case <-completed:
			logrus.Info("Completed requests")
		}
	}
}

func (core *Core) startAttack(taskRunner chan RequestPayload, completed chan bool) error {
	core.currentStatusBomber = system.StatusBomber_WORKING
	for index, request := range core.dataAttack {
		taskRunner <- RequestPayload{
			Request: &request,
			Id:      index,
		}
	}
	completed <- true
	return nil
}

func (core *Core) FormResultAttack() *rest_contracts.BomberResult {
	return &rest_contracts.BomberResult{
		BomberIp:                core.bomberIp,
		FormId:                  core.formId,
		AmountTimeoutsRequests:  core.resultTimeouts,
		AmountStatusesPerStatus: core.resultsAttack,
		MsPerRequest:            core.resultTimesForRequests,
	}
}

func (core *Core) Start(task rest_contracts.Task) {
	taskRunner := make(chan RequestPayload, currentWorkers)
	completed := make(chan bool)
	var index int64 = 0
	for ; index < task.Script.Config.Rps*task.Script.Config.Time; index++ {
		go core.runWorkers(taskRunner, completed)
	}
	core.startAttack(taskRunner, completed)
}

func (core *Core) InitializeService() {
	core.changeStatusBomber(core.currentStatusBomber)
}

func (core *Core) handlingChangeStatusBomber() {
	currentStatus := core.currentStatusBomber
	for {
		time.Sleep(time.Second * 5)
		if currentStatus != core.currentStatusBomber {
			logrus.Info("Handled changing current status worker: ", core.currentStatusBomber.String())
			core.changeStatusBomber(core.currentStatusBomber)
			currentStatus = core.currentStatusBomber
		}
	}
}

func (core *Core) gracefullDownService() {
	logrus.Info("Graceful down service")
	core.changeStatusBomber(system.StatusBomber_DOWN)
}

func (core *Core) changeStatusBomber(status system.StatusBomber) {
	statusBomberInitialized := system.BomberStatusChange{
		BomberId:     core.config.CurrentServiceID,
		StatusBomber: status,
	}
	data, errMarshaling := statusBomberInitialized.Marshal()
	if errMarshaling != nil {
		logrus.Error("Can not marshaled payload for bomber server: ", errMarshaling)
	}
	if errPublish := core.publisher.PublishNewMessage(topicName, data); errPublish != nil {
		logrus.Error("Can not publish message into broker nats")
	}
}
