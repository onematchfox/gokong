package gokong

import (
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ConsumersGetById(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	result, err := client.Consumers().GetById(createdConsumer.Id)

	assert.Equal(t, createdConsumer, result)

}

func Test_ConsumersGetByUsername(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	result, err := client.Consumers().GetByUsername(createdConsumer.Username)

	assert.Equal(t, createdConsumer, result)

}

func Test_ConsumersCreate(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	result, err := NewClient(NewDefaultConfig()).Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, consumerRequest.Username, result.Username)
	assert.Equal(t, consumerRequest.CustomId, result.CustomId)

}

func Test_ConsumersCreateOnlyUsername(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
	}

	result, err := NewClient(NewDefaultConfig()).Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, consumerRequest.Username, result.Username)
	assert.Equal(t, "", result.CustomId)

}

func Test_ConsumersCreateOnlyCustomId(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		CustomId: "test-" + uuid.NewV4().String(),
	}

	result, err := NewClient(NewDefaultConfig()).Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "", result.Username)
	assert.Equal(t, consumerRequest.CustomId, result.CustomId)

}

func Test_ConsumersList(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	results, err := client.Consumers().List()

	assert.True(t, results.Total > 0)
	assert.True(t, len(results.Results) > 0)

}

func Test_ConsumersListFilteredById(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	consumerRequest2 := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}
	createdConsumer2, err := client.Consumers().Create(consumerRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer2)

	results, err := client.Consumers().ListFiltered(&ConsumerFilter{Id: createdConsumer.Id})

	assert.True(t, len(results.Results) == 1)
	result := results.Results[0]

	assert.Equal(t, createdConsumer.Id, result.Id)
	assert.Equal(t, createdConsumer.CustomId, result.CustomId)
	assert.Equal(t, createdConsumer.Username, result.Username)

}

func Test_ConsumersListFilteredByCustomId(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	consumerRequest2 := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}
	createdConsumer2, err := client.Consumers().Create(consumerRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer2)

	results, err := client.Consumers().ListFiltered(&ConsumerFilter{CustomId: createdConsumer.CustomId})

	assert.True(t, len(results.Results) == 1)
	result := results.Results[0]

	assert.Equal(t, createdConsumer.Id, result.Id)
	assert.Equal(t, createdConsumer.CustomId, result.CustomId)
	assert.Equal(t, createdConsumer.Username, result.Username)

}

func Test_ConsumersListFilteredByUsername(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	consumerRequest2 := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}
	createdConsumer2, err := client.Consumers().Create(consumerRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer2)

	results, err := client.Consumers().ListFiltered(&ConsumerFilter{Username: createdConsumer.Username})

	assert.True(t, len(results.Results) == 1)
	result := results.Results[0]

	assert.Equal(t, createdConsumer.Id, result.Id)
	assert.Equal(t, createdConsumer.CustomId, result.CustomId)
	assert.Equal(t, createdConsumer.Username, result.Username)

}

func Test_ConsumersListFilteredBySize(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	consumerRequest2 := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}
	createdConsumer2, err := client.Consumers().Create(consumerRequest2)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer2)

	consumerRequest3 := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}
	createdConsumer3, err := client.Consumers().Create(consumerRequest3)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer3)

	results, err := client.Consumers().ListFiltered(&ConsumerFilter{Size: 2})

	assert.True(t, len(results.Results) == 2)

}

func Test_ConsumersDeleteById(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	err = client.Consumers().DeleteById(createdConsumer.Id)
	assert.Nil(t, err)

	deletedConsumer, err := client.Consumers().GetById(createdConsumer.Id)
	assert.Nil(t, deletedConsumer)
}

func Test_ConsumersDeleteByUsername(t *testing.T) {
	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)

	err = client.Consumers().DeleteByUsername(createdConsumer.Username)
	assert.Nil(t, err)

	deletedConsumer, err := client.Consumers().GetById(createdConsumer.Id)
	assert.Nil(t, deletedConsumer)
}

func Test_ConsumersUpdateById(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)
	assert.Equal(t, consumerRequest.CustomId, createdConsumer.CustomId)

	consumerRequest.CustomId = "test-" + uuid.NewV4().String()

	result, err := client.Consumers().UpdateById(createdConsumer.Id, consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, consumerRequest.CustomId, result.CustomId)
	assert.Equal(t, consumerRequest.Username, result.Username)
}

func Test_ConsumersUpdateByUsername(t *testing.T) {

	consumerRequest := &ConsumerRequest{
		Username: "username-" + uuid.NewV4().String(),
		CustomId: "test-" + uuid.NewV4().String(),
	}

	client := NewClient(NewDefaultConfig())
	createdConsumer, err := client.Consumers().Create(consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, createdConsumer)
	assert.Equal(t, consumerRequest.CustomId, createdConsumer.CustomId)

	consumerRequest.Username = "username-" + uuid.NewV4().String()

	result, err := client.Consumers().UpdateByUsername(createdConsumer.Username, consumerRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, consumerRequest.CustomId, result.CustomId)
	assert.Equal(t, consumerRequest.Username, result.Username)
}
