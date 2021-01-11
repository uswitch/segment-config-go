package segment

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// List Tracking plans
func (c *Client) ListTrackingPlans() (TrackingPlans, error) {
	var tps TrackingPlans
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint),
		nil)
	if err != nil {
		return tps, err
	}
	err = json.Unmarshal(data, &tps)
	fmt.Println("This is the string of data")
	fmt.Println(string(data))
	// fmt.Printf("%+v", tps)
	if err != nil {
		return tps, errors.Wrap(err, "failed to unmarshal tracking plans response")
	}

	return tps, nil
}

// GetTrackingPlan gets a specific tracking plan from segment
func (c *Client) GetTrackingPlan(TrackingPlanID string) (TrackingPlan, error) {
	var tp TrackingPlan
	data, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("%s/%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, TrackingPlanID),
		nil)
	if err != nil {
		return tp, err
	}
	err = json.Unmarshal(data, &tp)
	// fmt.Printf("%+v", tp)
	if err != nil {
		return tp, errors.Wrap(err, "failed to unmarshal tracking plans response")
	}

	return tp, nil
}

// CreateTrackingPlan creates tracking plan
func (c *Client) CreateTrackingPlan(data TrackingPlan) (TrackingPlan, error) {
	var tp TrackingPlan
	tpCreateReq := trackingPlanCreateRequest{
		TrackingPlan: data,
	}
	responseBody, err := c.doRequest(http.MethodPost,
		fmt.Sprintf("%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint),
			tpCreateReq)

	if err != nil {
		return tp, err
	}
	err = json.Unmarshal(responseBody, &tp)

	return tp, nil
}

// UpdateTrackingPlan updates a tracking plan
func (c *Client) UpdateTrackingPlan(TrackingPlanID string, data TrackingPlan) (TrackingPlan, error) {
	var tp TrackingPlan
	
	um := UpdateMask{
		Paths: []string{"tracking_plan.display_name", "tracking_plan.rules"},
	}
	tpUpdateReq := trackingPlanUpdateRequest{
		UpdateMask: um,
		TrackingPlan: data,
	}
	responseBody, err := c.doRequest(http.MethodPut,
		fmt.Sprintf("%s/%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, TrackingPlanID),
		tpUpdateReq)

	if err != nil {
		return tp, err
	}
	err = json.Unmarshal(responseBody, &tp)

	return tp, nil
}

// DeleteTrackingPlan Deletes a tracking plan
func (c *Client) DeleteTrackingPlan(TrackingPlanID string) error {

	_, err := c.doRequest(http.MethodDelete,
		fmt.Sprintf("%s/%s/%s/%s/",
			WorkspacesEndpoint, c.workspace, TrackingPlanEndpoint, TrackingPlanID),
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
