package mq

import (
	"github.com/mz-eco/mz/kafka"
	"github.com/mz-eco/mz/log"
)

const (
	TOPIC_SERVICE_AGENCY_SERVICE = "service-agency-service"
)

const (
	IDENT_SERVICE_AGENCY_SERVICE_MODIFY_ORGANIZATION_INFO = "modify_organization_info"
	IDENT_SERVICE_AGENCY_SERVICE_ADD_ORGANIZATION         = "add_organization"
)

var (
	topicServiceAgencyService *TopicServiceAgencyService = nil
)

func GetTopicServiceAgencyService() (topic *TopicServiceAgencyService, err error) {
	if topicServiceAgencyService != nil {
		topic = topicServiceAgencyService
		return
	}

	producer, err := kafka.NewAsyncProducer()
	if err != nil {
		log.Warnf("new async producer failed. %s", err)
		return
	}

	topicServiceAgencyService = &TopicServiceAgencyService{
		Producer: producer,
	}

	topic = topicServiceAgencyService

	return
}

type TopicServiceAgencyService struct {
	Producer *kafka.AsyncProducer
}

func (m *TopicServiceAgencyService) send(ident string, msg interface{}) (err error) {
	err = m.Producer.SendMessage(TOPIC_SERVICE_AGENCY_SERVICE, ident, msg)
	if err != nil {
		log.Warnf("send topic message failed. %s", err)
		return
	}
	return
}

// 组织信息修改
type ModifyOrganizationInfoMessage struct {
	OrganizationId uint32
	Values         map[string]interface{}
}

func (m *TopicServiceAgencyService) ModifyOrganizationInfo(msg *ModifyOrganizationInfoMessage) (err error) {
	return m.send(IDENT_SERVICE_AGENCY_SERVICE_MODIFY_ORGANIZATION_INFO, msg)
}
