package plan

import "testing"

func TestDefaulJSONPlan(t *testing.T) {
	// if it errors out, it'll panic; therefore a simple call is enough to test it
	DefaultJSONPlan()
}
