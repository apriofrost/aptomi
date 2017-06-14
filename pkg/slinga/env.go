package slinga

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

// AptomiOject represents an aptomi entity, which gets stored in aptomi DB
type AptomiOject string

const (
	/*
		The following objects can be added to Aptomi
	*/

	// Clusters is k8s cluster or any other cluster
	Clusters AptomiOject = "policy/clusters"

	// Services is service definitions
	Services AptomiOject = "policy/services"

	// Contexts is how service gets allocated
	Contexts AptomiOject = "policy/contexts"

	// Rules is global rules of the land
	Rules AptomiOject = "policy/rules"

	// Dependencies is who requested what
	Dependencies AptomiOject = "dependencies"

	/*
		The following objects must be configured to point to external resources
	*/

	// Users is where users are stored (later in AD and LDAP)
	Users AptomiOject = "external/users"

	// Secrets is where secret tokens are stored (later in Hashicorp Vault)
	Secrets AptomiOject = "external/secrets"

	// Charts is where binary charts/images are stored (later in external repo)
	Charts AptomiOject = "external/charts"

	/*
		The following objects are generated by aptomi during or after dependency resolution via policy
	*/

	// PolicyResolution holds usage data for components/dependencies
	PolicyResolution AptomiOject = "resolution/usage"

	// Logs contains debug logs
	Logs AptomiOject = "resolution/logs"

	// Graphics contains images generated by graphviz
	Graphics AptomiOject = "resolution/graphics"
)

// AptomiObjectsCanBeModified contains a map of all objects which can be added to aptomi policy
var AptomiObjectsCanBeModified = map[string]AptomiOject{
	"cluster":      Clusters,
	"service":      Services,
	"context":      Contexts,
	"rules":        Rules,
	"dependencies": Dependencies,
	"users":        Users,
	"secrets":      Secrets,
	"chart":        Charts,
}

// Return aptomi DB directory
func getAptomiEnvVarAsDir(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		debug.WithFields(log.Fields{
			"var": key,
		}).Fatal("Environment variable is not present. Must point to a directory")
	}
	if stat, err := os.Stat(value); err != nil || !stat.IsDir() {
		debug.WithFields(log.Fields{
			"var":       key,
			"directory": value,
			"error":     err,
		}).Fatal("Directory doesn't exist or error encountered")
	}
	return value
}

// GetAptomiBaseDir returns base directory, i.e. the value of APTOMI_DB environment variable
func GetAptomiBaseDir() string {
	return getAptomiEnvVarAsDir("APTOMI_DB")
}

// GetAptomiObjectDir returns a directory where Aptomi stores objects of a particular object time
func GetAptomiObjectDir(baseDir string, apt AptomiOject) string {
	dir := baseDir + "/" + string(apt)
	if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
		_ = os.MkdirAll(dir, 0755)
	}
	if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
		debug.WithFields(log.Fields{
			"directory": dir,
			"error":     err,
		}).Fatal("Directory can't be created or error encountered")
	}
	return dir
}
