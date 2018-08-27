package mq

import (
	db2 "business/agency/common/db"
	"business/user/common/mq"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/mz-eco/mz/kafka"
	"github.com/mz-eco/mz/log"
	"github.com/mz-eco/mz/settings"
)

var (
	topicUserServiceGroupSetting kafka.TopicGroupSetting
)

func init() {
	settings.RegisterWith(func(viper *settings.Viper) error {
		err := viper.Unmarshal(&topicUserServiceGroupSetting)
		if err != nil {
			panic(err)
			return err
		}
		return nil
	}, "kafka.subscribe.service_user_service")
}

type TopicUserServiceHandler struct {
	kafka.TopicHandler
}

func NewTopicUserServiceHandler() (handler *TopicUserServiceHandler, err error) {
	handler = &TopicUserServiceHandler{
		TopicHandler: kafka.TopicHandler{
			Topics:  []string{mq.TOPIC_SERVICE_USER_SERVICE},
			Brokers: topicUserServiceGroupSetting.Address,
			Group:   topicUserServiceGroupSetting.Group,
		},
	}

	handler.MessageHandle = handler.handleMessage

	return
}

func (m *TopicUserServiceHandler) handleMessage(identMessage *kafka.IdentMessage) (err error) {
	switch identMessage.Ident {
	case mq.IDENT_SERVICE_USER_SERVICE_MODIFY_USER_INFO:
		modifyInfo := &mq.ModifyUserInfoMessage{}
		err = json.Unmarshal(identMessage.Msg, modifyInfo)
		if err != nil {
			log.Warnf("unmarshal modify info message failed. %s", err)
			return
		}

		err = m.ModifyStaffUserInfo(modifyInfo)
		if err != nil {
			log.Warnf("modify staff user info failed. %s", err)
			return
		}

		err = m.ModifyOrganizationUserInfo(modifyInfo)
		if err != nil {
			log.Warnf("modify organization user info failed. %s", err)
			return
		}
	}
	return
}

func (m *TopicUserServiceHandler) ModifyStaffUserInfo(msg *mq.ModifyUserInfoMessage) (err error) {
	if len(msg.Values) == 0 {
		err = errors.New("empty update field")
		return
	}

	var updateFields []string
	var updateValues []interface{}
	var filtFields = map[string]string{
		"mobile": "mobile",
		"name":   "name",
	}
	for key, value := range msg.Values {
		field, ok := filtFields[key]
		if !ok {
			continue
		}

		updateField := fmt.Sprintf("%s=?", field)
		updateFields = append(updateFields, updateField)
		updateValues = append(updateValues, value)
	}

	if len(updateFields) == 0 {
		return
	}

	updateValues = append(updateValues, msg.UserId)

	strSql := "UPDATE agc_staff SET %s WHERE uid=?"
	strSql = fmt.Sprintf(strSql, strings.Join(updateFields, ","))
	dbAgency := db2.NewMallAgency()
	_, err = dbAgency.DB.Exec(strSql, updateValues...)
	if err != nil {
		log.Warnf("update modified staff user info failed. %s", err)
		return
	}

	return
}

func (m *TopicUserServiceHandler) ModifyOrganizationUserInfo(msg *mq.ModifyUserInfoMessage) (err error) {
	if len(msg.Values) == 0 {
		err = errors.New("empty update field")
		return
	}

	var updateFields []string
	var updateValues []interface{}
	var filtFields = map[string]string{
		"mobile": "manager_mobile",
		"name":   "manager_name",
	}
	for key, value := range msg.Values {
		field, ok := filtFields[key]
		if !ok {
			continue
		}

		updateField := fmt.Sprintf("%s=?", field)
		updateFields = append(updateFields, updateField)
		updateValues = append(updateValues, value)
	}

	if len(updateFields) == 0 {
		return
	}

	updateValues = append(updateValues, msg.UserId)

	strSql := "UPDATE agc_organization SET %s WHERE manager_uid=?"
	strSql = fmt.Sprintf(strSql, strings.Join(updateFields, ","))
	dbAgency := db2.NewMallAgency()
	_, err = dbAgency.DB.Exec(strSql, updateValues...)
	if err != nil {
		log.Warnf("update modified staff user info failed. %s", err)
		return
	}

	return
}
