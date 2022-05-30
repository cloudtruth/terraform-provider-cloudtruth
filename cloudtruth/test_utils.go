package cloudtruth

// For the sake of isolation/bootstrapping
// The provider acceptance tests assume that the following resources already exist in the target
// CloudTruth account, if they are destroyed/altered these tests will fail
// The resource acceptance tests are "full service" tests which instantiate, modify and destroy
// the resources which they reference.

// The following constants are shared across multiple tests
// Constants and variables which are unique per resource/data source type are located in the
// corresponding *_test.go files
const (
	defaultEnv      = "default"
	accTestProject  = "AcceptanceTest"
	regularParam    = "DefaultRegularParam"
	regularParamVal = "notreallyasecret"
	secretParam     = "DefaultSecretParam"
	secretParamVal  = "ultratopsecret"
)
