package nexusiq

import (
	"encoding/json"
	"io/ioutil"
)

// IqCliResults encapsulates the JSON object generated by an evaluation with the Nexus IQ CLI
type IqCliResults struct {
	ApplicationID          string                 `json:"applicationId"`
	ScanID                 string                 `json:"scanId"`
	ReportHTMLURL          string                 `json:"reportHtmlUrl"`
	ReportPDFURL           string                 `json:"reportPdfUrl"`
	ReportDataURL          string                 `json:"reportDataUrl"`
	PolicyAction           string                 `json:"policyAction"`
	PolicyEvaluationResult policyEvaluationResult `json:"policyEvaluationResult"`
}

type policyEvaluationResult struct {
	Alerts                            []alert `json:"alerts"`
	AffectedComponentCount            int64   `json:"affectedComponentCount"`
	CriticalComponentCount            int64   `json:"criticalComponentCount"`
	SevereComponentCount              int64   `json:"severeComponentCount"`
	ModerateComponentCount            int64   `json:"moderateComponentCount"`
	CriticalPolicyViolationCount      int64   `json:"criticalPolicyViolationCount"`
	SeverePolicyViolationCount        int64   `json:"severePolicyViolationCount"`
	ModeratePolicyViolationCount      int64   `json:"moderatePolicyViolationCount"`
	GrandfatheredPolicyViolationCount int64   `json:"grandfatheredPolicyViolationCount"`
}

type alert struct {
	Trigger trigger  `json:"trigger"`
	Actions []action `json:"actions"`
}

type action struct {
	ActionTypeID string      `json:"actionTypeId"`
	Target       interface{} `json:"target"`
}

type trigger struct {
	PolicyID          string          `json:"policyId"`
	PolicyName        string          `json:"policyName"`
	ThreatLevel       int64           `json:"threatLevel"`
	PolicyViolationID string          `json:"policyViolationId"`
	ComponentFacts    []componentFact `json:"componentFacts"`
}

type componentFact struct {
	Component
	ConstraintFacts []constraintFact `json:"constraintFacts"`
	DisplayName     displayName      `json:"displayName"`
}

// TODO: merge with the same ones in webhooks?
type constraintFact struct {
	ConstraintID   string          `json:"constraintId"`
	ConstraintName string          `json:"constraintName"`
	OperatorName   string          `json:"operatorName"`
	ConditionFacts []conditionFact `json:"conditionFacts"`
}

type conditionFact struct {
	ConditionTypeID string      `json:"conditionTypeId"`
	ConditionIndex  int64       `json:"conditionIndex"`
	Summary         string      `json:"summary"`
	Reason          string      `json:"reason"`
	Reference       *reference  `json:"reference"`
	TriggerJSON     interface{} `json:"triggerJson"`
}

type reference struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type displayName struct {
	Parts []part `json:"parts"`
}

type part struct {
	Field string `json:"field,omitempty"`
	Value string `json:"value"`
}

func ReadIqCliResultFile(filename string) (IqCliResults, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return IqCliResults{}, err
	}

	var cli IqCliResults
	if err = json.Unmarshal(f, &cli); err != nil {
		return IqCliResults{}, err
	}

	return cli, nil
}
