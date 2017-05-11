package slinga

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
)

func TestEngine(t *testing.T) {
	policy := loadPolicyFromDir("testdata/")
	users := loadUsersFromDir("testdata/")

	alice := users.Users["1"]
	bob := users.Users["2"]

	usageState := NewServiceUsageState(&policy)
	usageState.addDependency(alice, "kafka")
	usageState.addDependency(bob, "kafka")
	err := usageState.resolveUsage(&users)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 14, len(usageState.ResolvedLinks), "Policy resolution should result in correct amount of usage entries")
	// usageState.saveServiceUsageState()
}

func TestServiceComponentsTopologicalOrder(t *testing.T) {
	state := loadPolicyFromDir("testdata/")
	service := state.Services["kafka"]

	err := service.sortComponentsTopologically()
	assert.Equal(t, nil, err, "Service components should be topologically sorted without errors")

	assert.Equal(t, "component3", service.ComponentsOrdered[0].Name, "Component tologogical sort should produce correct order")
	assert.Equal(t, "component2", service.ComponentsOrdered[1].Name, "Component tologogical sort should produce correct order")
	assert.Equal(t, "component1", service.ComponentsOrdered[2].Name, "Component tologogical sort should produce correct order")
}
