package plan

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed default_plan.json
var defaultJSONPlanData []byte

func DefaultJSONPlan() JSONPlan {
	var jplan JSONPlan

	if err := json.Unmarshal(defaultJSONPlanData, &jplan); err != nil {
		errorMessage := fmt.Sprintf("failed to unmarshal default JSONPlan: %s", err.Error())
		panic(errorMessage)
	}
	if err := jplan.Validate(); err != nil {
		errorMessage := fmt.Sprintf("default JSONPlan is invalid: %s", err.Error())
		panic(errorMessage)
	}
	return jplan
}
