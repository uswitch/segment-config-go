package segment

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// ListTrackingPlans returns all tracking plans for a workspace
func (c *Client) ListTrackingPlans() (TrackingPlans, error) {
	var tps TrackingPlans
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s", WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint),
		nil)
	if err != nil {
		return tps, err
	}
	err = json.Unmarshal(data, &tps)
	fmt.Println(string(data))
	fmt.Printf("%+v", tps)
	if err != nil {
		return tps, errors.Wrap(err, "failed to unmarshal tracking plans response")
	}

	return tps, nil
}

// Get Tracking plan
func (c *Client) GetTrackingPlan(TrackingPlanId string) (TrackingPlan, error) {
	var tp TrackingPlan
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, TrackingPlanId),
		nil)
	if err != nil {
		return tp, err
	}
	err = json.Unmarshal(data, &tp)
	fmt.Println(string(data))
	// fmt.Printf("%+v", tp)
	if err != nil {
		return tp, errors.Wrap(err, "failed to unmarshal tracking plans response")
	}

	return tp, nil
}

// Create tracking plan
func (c *Client) CreateTrackingPlan(data interface{}) error {
	_, err := c.doRequest(http.MethodPost,
		fmt.Sprintf("%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint),
			data)

	if err != nil {
		return err
	}
	
	return nil
}

// Update Tracking plan
func (c *Client) UpdateTrackingPlan(TrackingPlanId string, data interface{}) error {
	_, err := c.doRequest(http.MethodPut,
		fmt.Sprintf("%s/%s/%s/%s/",
		WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, TrackingPlanId),
		data)
	if err != nil {
		return(err)
	}

	return nil
}

// Delete tracking plan (The response is an empty json object on delete)
func (c *Client) DeleteTrackingPlan(TrackingPlanId string) error {
	_, err := c.doRequest(http.MethodDelete,
		fmt.Sprintf("%s/%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, TrackingPlanId),
		nil)
	if err != nil {
		return err
	}
	
	return nil
}
	

// Batch create tracking plan source connections

// create tracking plan source connection

//List tracking plan source connections

// Delete tracking plan source connection