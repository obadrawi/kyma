package listener

import (
	"testing"

	"github.com/kyma-project/kyma/components/console-backend-service/internal/domain/servicecatalogaddons/listener/automock"
	"github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
)

func TestConfigMapAddonsConfiguration_OnAdd(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// given
		gqlAddonsConfiguration := new(gqlschema.AddonsConfiguration)
		cfgMap := new(v1.ConfigMap)
		converter := automock.NewGQLConfigMapAddonsConfigurationConverter()

		channel := make(chan gqlschema.AddonsConfigurationEvent, 1)
		defer close(channel)
		converter.On("ConfigMapToGQL", cfgMap).Return(gqlAddonsConfiguration, nil).Once()
		defer converter.AssertExpectations(t)
		serviceBrokerListener := NewConfigMapAddonsConfiguration(channel, filterAddonsConfigurationTrue, converter)

		// when
		serviceBrokerListener.OnAdd(cfgMap)
		result := <-channel

		// then
		assert.Equal(t, gqlschema.SubscriptionEventTypeAdd, result.Type)
		assert.Equal(t, *gqlAddonsConfiguration, result.AddonsConfiguration)
	})

	t.Run("Filtered out", func(t *testing.T) {
		// given
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationFalse, nil)

		// when
		serviceBrokerListener.OnAdd(new(v1.ConfigMap))
	})

	t.Run("Nil", func(t *testing.T) {
		// given
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationTrue, nil)

		// when
		serviceBrokerListener.OnAdd(nil)
	})

	t.Run("Nil GQL Type", func(t *testing.T) {
		// given
		serviceBroker := new(v1.ConfigMap)
		converter := automock.NewGQLConfigMapAddonsConfigurationConverter()

		converter.On("ConfigMapToGQL", serviceBroker).Return(nil, nil).Once()
		defer converter.AssertExpectations(t)
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationTrue, converter)

		// when
		serviceBrokerListener.OnAdd(serviceBroker)
	})

	t.Run("Invalid type", func(t *testing.T) {
		// given
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationTrue, nil)

		// when
		serviceBrokerListener.OnAdd(new(struct{}))
	})

}

func TestConfigMapAddonsConfiguration_OnDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// given
		gqlClusterServiceBroker := new(gqlschema.AddonsConfiguration)
		serviceBroker := new(v1.ConfigMap)
		converter := automock.NewGQLConfigMapAddonsConfigurationConverter()

		channel := make(chan gqlschema.AddonsConfigurationEvent, 1)
		defer close(channel)
		converter.On("ConfigMapToGQL", serviceBroker).Return(gqlClusterServiceBroker, nil).Once()
		defer converter.AssertExpectations(t)
		serviceBrokerListener := NewConfigMapAddonsConfiguration(channel, filterAddonsConfigurationTrue, converter)

		// when
		serviceBrokerListener.OnDelete(serviceBroker)
		result := <-channel

		// then
		assert.Equal(t, gqlschema.SubscriptionEventTypeDelete, result.Type)
		assert.Equal(t, *gqlClusterServiceBroker, result.AddonsConfiguration)

	})

	t.Run("Filtered out", func(t *testing.T) {
		// given
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationFalse, nil)

		// when
		serviceBrokerListener.OnDelete(new(v1.ConfigMap))
	})

	t.Run("Nil", func(t *testing.T) {
		// given
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationTrue, nil)

		// when
		serviceBrokerListener.OnDelete(nil)
	})

	t.Run("Nil GQL Type", func(t *testing.T) {
		// given
		serviceBroker := new(v1.ConfigMap)
		converter := automock.NewGQLConfigMapAddonsConfigurationConverter()

		converter.On("ConfigMapToGQL", serviceBroker).Return(nil, nil).Once()
		defer converter.AssertExpectations(t)
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationTrue, converter)

		// when
		serviceBrokerListener.OnDelete(serviceBroker)
	})

	t.Run("Invalid type", func(t *testing.T) {
		// given
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationTrue, nil)

		// when
		serviceBrokerListener.OnDelete(new(struct{}))
	})

}

func TestConfigMapAddonsConfiguration_OnUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// given
		gqlClusterServiceBroker := new(gqlschema.AddonsConfiguration)
		configMap := new(v1.ConfigMap)
		converter := automock.NewGQLConfigMapAddonsConfigurationConverter()

		channel := make(chan gqlschema.AddonsConfigurationEvent, 1)
		defer close(channel)
		converter.On("ConfigMapToGQL", configMap).Return(gqlClusterServiceBroker, nil).Once()
		defer converter.AssertExpectations(t)
		serviceBrokerListener := NewConfigMapAddonsConfiguration(channel, filterAddonsConfigurationTrue, converter)

		// when
		serviceBrokerListener.OnUpdate(configMap, configMap)
		result := <-channel

		// then
		assert.Equal(t, gqlschema.SubscriptionEventTypeUpdate, result.Type)
		assert.Equal(t, *gqlClusterServiceBroker, result.AddonsConfiguration)

	})

	t.Run("Filtered out", func(t *testing.T) {
		// given
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationFalse, nil)

		// when
		serviceBrokerListener.OnUpdate(new(v1.ConfigMap), new(v1.ConfigMap))
	})

	t.Run("Nil", func(t *testing.T) {
		// given
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationTrue, nil)

		// when
		serviceBrokerListener.OnUpdate(nil, nil)
	})

	t.Run("Nil GQL Type", func(t *testing.T) {
		// given
		serviceBroker := new(v1.ConfigMap)
		converter := automock.NewGQLConfigMapAddonsConfigurationConverter()

		converter.On("ConfigMapToGQL", serviceBroker).Return(nil, nil).Once()
		defer converter.AssertExpectations(t)
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationTrue, converter)

		// when
		serviceBrokerListener.OnUpdate(nil, serviceBroker)
	})

	t.Run("Invalid type", func(t *testing.T) {
		// given
		serviceBrokerListener := NewConfigMapAddonsConfiguration(nil, filterAddonsConfigurationTrue, nil)

		// when
		serviceBrokerListener.OnUpdate(new(struct{}), new(struct{}))
	})

}

func filterAddonsConfigurationTrue(o *v1.ConfigMap) bool {
	return true
}

func filterAddonsConfigurationFalse(o *v1.ConfigMap) bool {
	return false
}
