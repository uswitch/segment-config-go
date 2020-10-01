package segment

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrackingPlans_ListTrackingPlans(t *testing.T) {
	setup()
	defer teardown()

	endpoint := fmt.Sprintf("/%s/%s/%s/%s/", apiVersion, WorkspacesEndpoint, testWorkspace, TrackingPlanEndpoint)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
  "tracking_plans": [
    {
      "name": "workspaces/myworkspace/tracking-plans/rs_123",
      "display_name": "Kicks App",
      "rules": {
        "identify_traits": [],
        "group_traits": [],
        "events": []
      },
    }
  ]
}`)
	})

	actual, err := client.ListTrackingPlans()
	assert.NoError(t, err)

	expected := TrackingPlans{TrackingPlans: []TrackingPlan{
		{
			Name:        "workspaces/myworkspace/tracking-plans/rs_123",
			DisplayName: "Kicks App",
		},
	}}
	assert.Equal(t, expected, actual)
}
