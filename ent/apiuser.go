// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"roofix/ent/apiuser"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ApiUser is the model entity for the ApiUser schema.
type ApiUser struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// PwdHash holds the value of the "pwd_hash" field.
	PwdHash string `json:"pwd_hash,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// CbAPIURL holds the value of the "cb_api_url" field.
	CbAPIURL *string `json:"cb_api_url,omitempty"`
	// CbAPIAuth holds the value of the "cb_api_auth" field.
	CbAPIAuth apiuser.CbAPIAuth `json:"cb_api_auth,omitempty"`
	// CbAPIUser holds the value of the "cb_api_user" field.
	CbAPIUser string `json:"cb_api_user,omitempty"`
	// CbAPIPwd holds the value of the "cb_api_pwd" field.
	CbAPIPwd string `json:"cb_api_pwd,omitempty"`
	// CbAPIToken holds the value of the "cb_api_token" field.
	CbAPIToken string `json:"cb_api_token,omitempty"`
	// CbAPIEndpoints holds the value of the "cb_api_endpoints" field.
	CbAPIEndpoints map[string]string `json:"cb_api_endpoints,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ApiUserQuery when eager-loading is set.
	Edges        ApiUserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ApiUserEdges holds the relations/edges for other nodes in the graph.
type ApiUserEdges struct {
	// Tokens holds the value of the tokens edge.
	Tokens []*ApiUserToken `json:"tokens,omitempty"`
	// AuditLogs holds the value of the audit_logs edge.
	AuditLogs []*AuditLog `json:"audit_logs,omitempty"`
	// CreatedEstimates holds the value of the created_estimates edge.
	CreatedEstimates []*Estimate `json:"created_estimates,omitempty"`
	// CreatedJobs holds the value of the created_jobs edge.
	CreatedJobs []*Job `json:"created_jobs,omitempty"`
	// CreatedPartners holds the value of the created_partners edge.
	CreatedPartners []*Partner `json:"created_partners,omitempty"`
	// SurveyProgress holds the value of the survey_progress edge.
	SurveyProgress []*SurveyProgress `json:"survey_progress,omitempty"`
	// EstimateActivities holds the value of the estimate_activities edge.
	EstimateActivities []*EstimateActivity `json:"estimate_activities,omitempty"`
	// UserActivities holds the value of the user_activities edge.
	UserActivities []*UserActivity `json:"user_activities,omitempty"`
	// PartnerActivities holds the value of the partner_activities edge.
	PartnerActivities []*PartnerActivity `json:"partner_activities,omitempty"`
	// JobActivities holds the value of the job_activities edge.
	JobActivities []*JobActivity `json:"job_activities,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*ChannelMessage `json:"notifications,omitempty"`
	// JobProgressHistory holds the value of the job_progress_history edge.
	JobProgressHistory []*JobProgressHistory `json:"job_progress_history,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [12]bool
	// totalCount holds the count of the edges above.
	totalCount [12]map[string]int

	namedTokens             map[string][]*ApiUserToken
	namedAuditLogs          map[string][]*AuditLog
	namedCreatedEstimates   map[string][]*Estimate
	namedCreatedJobs        map[string][]*Job
	namedCreatedPartners    map[string][]*Partner
	namedSurveyProgress     map[string][]*SurveyProgress
	namedEstimateActivities map[string][]*EstimateActivity
	namedUserActivities     map[string][]*UserActivity
	namedPartnerActivities  map[string][]*PartnerActivity
	namedJobActivities      map[string][]*JobActivity
	namedNotifications      map[string][]*ChannelMessage
	namedJobProgressHistory map[string][]*JobProgressHistory
}

// TokensOrErr returns the Tokens value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) TokensOrErr() ([]*ApiUserToken, error) {
	if e.loadedTypes[0] {
		return e.Tokens, nil
	}
	return nil, &NotLoadedError{edge: "tokens"}
}

// AuditLogsOrErr returns the AuditLogs value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) AuditLogsOrErr() ([]*AuditLog, error) {
	if e.loadedTypes[1] {
		return e.AuditLogs, nil
	}
	return nil, &NotLoadedError{edge: "audit_logs"}
}

// CreatedEstimatesOrErr returns the CreatedEstimates value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) CreatedEstimatesOrErr() ([]*Estimate, error) {
	if e.loadedTypes[2] {
		return e.CreatedEstimates, nil
	}
	return nil, &NotLoadedError{edge: "created_estimates"}
}

// CreatedJobsOrErr returns the CreatedJobs value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) CreatedJobsOrErr() ([]*Job, error) {
	if e.loadedTypes[3] {
		return e.CreatedJobs, nil
	}
	return nil, &NotLoadedError{edge: "created_jobs"}
}

// CreatedPartnersOrErr returns the CreatedPartners value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) CreatedPartnersOrErr() ([]*Partner, error) {
	if e.loadedTypes[4] {
		return e.CreatedPartners, nil
	}
	return nil, &NotLoadedError{edge: "created_partners"}
}

// SurveyProgressOrErr returns the SurveyProgress value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) SurveyProgressOrErr() ([]*SurveyProgress, error) {
	if e.loadedTypes[5] {
		return e.SurveyProgress, nil
	}
	return nil, &NotLoadedError{edge: "survey_progress"}
}

// EstimateActivitiesOrErr returns the EstimateActivities value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) EstimateActivitiesOrErr() ([]*EstimateActivity, error) {
	if e.loadedTypes[6] {
		return e.EstimateActivities, nil
	}
	return nil, &NotLoadedError{edge: "estimate_activities"}
}

// UserActivitiesOrErr returns the UserActivities value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) UserActivitiesOrErr() ([]*UserActivity, error) {
	if e.loadedTypes[7] {
		return e.UserActivities, nil
	}
	return nil, &NotLoadedError{edge: "user_activities"}
}

// PartnerActivitiesOrErr returns the PartnerActivities value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) PartnerActivitiesOrErr() ([]*PartnerActivity, error) {
	if e.loadedTypes[8] {
		return e.PartnerActivities, nil
	}
	return nil, &NotLoadedError{edge: "partner_activities"}
}

// JobActivitiesOrErr returns the JobActivities value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) JobActivitiesOrErr() ([]*JobActivity, error) {
	if e.loadedTypes[9] {
		return e.JobActivities, nil
	}
	return nil, &NotLoadedError{edge: "job_activities"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) NotificationsOrErr() ([]*ChannelMessage, error) {
	if e.loadedTypes[10] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// JobProgressHistoryOrErr returns the JobProgressHistory value or an error if the edge
// was not loaded in eager-loading.
func (e ApiUserEdges) JobProgressHistoryOrErr() ([]*JobProgressHistory, error) {
	if e.loadedTypes[11] {
		return e.JobProgressHistory, nil
	}
	return nil, &NotLoadedError{edge: "job_progress_history"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApiUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apiuser.FieldCbAPIEndpoints:
			values[i] = new([]byte)
		case apiuser.FieldActive:
			values[i] = new(sql.NullBool)
		case apiuser.FieldID, apiuser.FieldUsername, apiuser.FieldPwdHash, apiuser.FieldCbAPIURL, apiuser.FieldCbAPIAuth, apiuser.FieldCbAPIUser, apiuser.FieldCbAPIPwd, apiuser.FieldCbAPIToken:
			values[i] = new(sql.NullString)
		case apiuser.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApiUser fields.
func (au *ApiUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apiuser.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				au.ID = value.String
			}
		case apiuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				au.CreatedAt = value.Time
			}
		case apiuser.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				au.Username = value.String
			}
		case apiuser.FieldPwdHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pwd_hash", values[i])
			} else if value.Valid {
				au.PwdHash = value.String
			}
		case apiuser.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				au.Active = value.Bool
			}
		case apiuser.FieldCbAPIURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cb_api_url", values[i])
			} else if value.Valid {
				au.CbAPIURL = new(string)
				*au.CbAPIURL = value.String
			}
		case apiuser.FieldCbAPIAuth:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cb_api_auth", values[i])
			} else if value.Valid {
				au.CbAPIAuth = apiuser.CbAPIAuth(value.String)
			}
		case apiuser.FieldCbAPIUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cb_api_user", values[i])
			} else if value.Valid {
				au.CbAPIUser = value.String
			}
		case apiuser.FieldCbAPIPwd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cb_api_pwd", values[i])
			} else if value.Valid {
				au.CbAPIPwd = value.String
			}
		case apiuser.FieldCbAPIToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cb_api_token", values[i])
			} else if value.Valid {
				au.CbAPIToken = value.String
			}
		case apiuser.FieldCbAPIEndpoints:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field cb_api_endpoints", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &au.CbAPIEndpoints); err != nil {
					return fmt.Errorf("unmarshal field cb_api_endpoints: %w", err)
				}
			}
		default:
			au.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ApiUser.
// This includes values selected through modifiers, order, etc.
func (au *ApiUser) Value(name string) (ent.Value, error) {
	return au.selectValues.Get(name)
}

// QueryTokens queries the "tokens" edge of the ApiUser entity.
func (au *ApiUser) QueryTokens() *ApiUserTokenQuery {
	return NewApiUserClient(au.config).QueryTokens(au)
}

// QueryAuditLogs queries the "audit_logs" edge of the ApiUser entity.
func (au *ApiUser) QueryAuditLogs() *AuditLogQuery {
	return NewApiUserClient(au.config).QueryAuditLogs(au)
}

// QueryCreatedEstimates queries the "created_estimates" edge of the ApiUser entity.
func (au *ApiUser) QueryCreatedEstimates() *EstimateQuery {
	return NewApiUserClient(au.config).QueryCreatedEstimates(au)
}

// QueryCreatedJobs queries the "created_jobs" edge of the ApiUser entity.
func (au *ApiUser) QueryCreatedJobs() *JobQuery {
	return NewApiUserClient(au.config).QueryCreatedJobs(au)
}

// QueryCreatedPartners queries the "created_partners" edge of the ApiUser entity.
func (au *ApiUser) QueryCreatedPartners() *PartnerQuery {
	return NewApiUserClient(au.config).QueryCreatedPartners(au)
}

// QuerySurveyProgress queries the "survey_progress" edge of the ApiUser entity.
func (au *ApiUser) QuerySurveyProgress() *SurveyProgressQuery {
	return NewApiUserClient(au.config).QuerySurveyProgress(au)
}

// QueryEstimateActivities queries the "estimate_activities" edge of the ApiUser entity.
func (au *ApiUser) QueryEstimateActivities() *EstimateActivityQuery {
	return NewApiUserClient(au.config).QueryEstimateActivities(au)
}

// QueryUserActivities queries the "user_activities" edge of the ApiUser entity.
func (au *ApiUser) QueryUserActivities() *UserActivityQuery {
	return NewApiUserClient(au.config).QueryUserActivities(au)
}

// QueryPartnerActivities queries the "partner_activities" edge of the ApiUser entity.
func (au *ApiUser) QueryPartnerActivities() *PartnerActivityQuery {
	return NewApiUserClient(au.config).QueryPartnerActivities(au)
}

// QueryJobActivities queries the "job_activities" edge of the ApiUser entity.
func (au *ApiUser) QueryJobActivities() *JobActivityQuery {
	return NewApiUserClient(au.config).QueryJobActivities(au)
}

// QueryNotifications queries the "notifications" edge of the ApiUser entity.
func (au *ApiUser) QueryNotifications() *ChannelMessageQuery {
	return NewApiUserClient(au.config).QueryNotifications(au)
}

// QueryJobProgressHistory queries the "job_progress_history" edge of the ApiUser entity.
func (au *ApiUser) QueryJobProgressHistory() *JobProgressHistoryQuery {
	return NewApiUserClient(au.config).QueryJobProgressHistory(au)
}

// Update returns a builder for updating this ApiUser.
// Note that you need to call ApiUser.Unwrap() before calling this method if this ApiUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (au *ApiUser) Update() *ApiUserUpdateOne {
	return NewApiUserClient(au.config).UpdateOne(au)
}

// Unwrap unwraps the ApiUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (au *ApiUser) Unwrap() *ApiUser {
	_tx, ok := au.config.driver.(*txDriver)
	if !ok {
		panic("ent: ApiUser is not a transactional entity")
	}
	au.config.driver = _tx.drv
	return au
}

// String implements the fmt.Stringer.
func (au *ApiUser) String() string {
	var builder strings.Builder
	builder.WriteString("ApiUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", au.ID))
	builder.WriteString("created_at=")
	builder.WriteString(au.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(au.Username)
	builder.WriteString(", ")
	builder.WriteString("pwd_hash=")
	builder.WriteString(au.PwdHash)
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", au.Active))
	builder.WriteString(", ")
	if v := au.CbAPIURL; v != nil {
		builder.WriteString("cb_api_url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("cb_api_auth=")
	builder.WriteString(fmt.Sprintf("%v", au.CbAPIAuth))
	builder.WriteString(", ")
	builder.WriteString("cb_api_user=")
	builder.WriteString(au.CbAPIUser)
	builder.WriteString(", ")
	builder.WriteString("cb_api_pwd=")
	builder.WriteString(au.CbAPIPwd)
	builder.WriteString(", ")
	builder.WriteString("cb_api_token=")
	builder.WriteString(au.CbAPIToken)
	builder.WriteString(", ")
	builder.WriteString("cb_api_endpoints=")
	builder.WriteString(fmt.Sprintf("%v", au.CbAPIEndpoints))
	builder.WriteByte(')')
	return builder.String()
}

// NamedTokens returns the Tokens named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedTokens(name string) ([]*ApiUserToken, error) {
	if au.Edges.namedTokens == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedTokens[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedTokens(name string, edges ...*ApiUserToken) {
	if au.Edges.namedTokens == nil {
		au.Edges.namedTokens = make(map[string][]*ApiUserToken)
	}
	if len(edges) == 0 {
		au.Edges.namedTokens[name] = []*ApiUserToken{}
	} else {
		au.Edges.namedTokens[name] = append(au.Edges.namedTokens[name], edges...)
	}
}

// NamedAuditLogs returns the AuditLogs named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedAuditLogs(name string) ([]*AuditLog, error) {
	if au.Edges.namedAuditLogs == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedAuditLogs[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedAuditLogs(name string, edges ...*AuditLog) {
	if au.Edges.namedAuditLogs == nil {
		au.Edges.namedAuditLogs = make(map[string][]*AuditLog)
	}
	if len(edges) == 0 {
		au.Edges.namedAuditLogs[name] = []*AuditLog{}
	} else {
		au.Edges.namedAuditLogs[name] = append(au.Edges.namedAuditLogs[name], edges...)
	}
}

// NamedCreatedEstimates returns the CreatedEstimates named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedCreatedEstimates(name string) ([]*Estimate, error) {
	if au.Edges.namedCreatedEstimates == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedCreatedEstimates[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedCreatedEstimates(name string, edges ...*Estimate) {
	if au.Edges.namedCreatedEstimates == nil {
		au.Edges.namedCreatedEstimates = make(map[string][]*Estimate)
	}
	if len(edges) == 0 {
		au.Edges.namedCreatedEstimates[name] = []*Estimate{}
	} else {
		au.Edges.namedCreatedEstimates[name] = append(au.Edges.namedCreatedEstimates[name], edges...)
	}
}

// NamedCreatedJobs returns the CreatedJobs named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedCreatedJobs(name string) ([]*Job, error) {
	if au.Edges.namedCreatedJobs == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedCreatedJobs[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedCreatedJobs(name string, edges ...*Job) {
	if au.Edges.namedCreatedJobs == nil {
		au.Edges.namedCreatedJobs = make(map[string][]*Job)
	}
	if len(edges) == 0 {
		au.Edges.namedCreatedJobs[name] = []*Job{}
	} else {
		au.Edges.namedCreatedJobs[name] = append(au.Edges.namedCreatedJobs[name], edges...)
	}
}

// NamedCreatedPartners returns the CreatedPartners named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedCreatedPartners(name string) ([]*Partner, error) {
	if au.Edges.namedCreatedPartners == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedCreatedPartners[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedCreatedPartners(name string, edges ...*Partner) {
	if au.Edges.namedCreatedPartners == nil {
		au.Edges.namedCreatedPartners = make(map[string][]*Partner)
	}
	if len(edges) == 0 {
		au.Edges.namedCreatedPartners[name] = []*Partner{}
	} else {
		au.Edges.namedCreatedPartners[name] = append(au.Edges.namedCreatedPartners[name], edges...)
	}
}

// NamedSurveyProgress returns the SurveyProgress named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedSurveyProgress(name string) ([]*SurveyProgress, error) {
	if au.Edges.namedSurveyProgress == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedSurveyProgress[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedSurveyProgress(name string, edges ...*SurveyProgress) {
	if au.Edges.namedSurveyProgress == nil {
		au.Edges.namedSurveyProgress = make(map[string][]*SurveyProgress)
	}
	if len(edges) == 0 {
		au.Edges.namedSurveyProgress[name] = []*SurveyProgress{}
	} else {
		au.Edges.namedSurveyProgress[name] = append(au.Edges.namedSurveyProgress[name], edges...)
	}
}

// NamedEstimateActivities returns the EstimateActivities named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedEstimateActivities(name string) ([]*EstimateActivity, error) {
	if au.Edges.namedEstimateActivities == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedEstimateActivities[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedEstimateActivities(name string, edges ...*EstimateActivity) {
	if au.Edges.namedEstimateActivities == nil {
		au.Edges.namedEstimateActivities = make(map[string][]*EstimateActivity)
	}
	if len(edges) == 0 {
		au.Edges.namedEstimateActivities[name] = []*EstimateActivity{}
	} else {
		au.Edges.namedEstimateActivities[name] = append(au.Edges.namedEstimateActivities[name], edges...)
	}
}

// NamedUserActivities returns the UserActivities named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedUserActivities(name string) ([]*UserActivity, error) {
	if au.Edges.namedUserActivities == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedUserActivities[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedUserActivities(name string, edges ...*UserActivity) {
	if au.Edges.namedUserActivities == nil {
		au.Edges.namedUserActivities = make(map[string][]*UserActivity)
	}
	if len(edges) == 0 {
		au.Edges.namedUserActivities[name] = []*UserActivity{}
	} else {
		au.Edges.namedUserActivities[name] = append(au.Edges.namedUserActivities[name], edges...)
	}
}

// NamedPartnerActivities returns the PartnerActivities named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedPartnerActivities(name string) ([]*PartnerActivity, error) {
	if au.Edges.namedPartnerActivities == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedPartnerActivities[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedPartnerActivities(name string, edges ...*PartnerActivity) {
	if au.Edges.namedPartnerActivities == nil {
		au.Edges.namedPartnerActivities = make(map[string][]*PartnerActivity)
	}
	if len(edges) == 0 {
		au.Edges.namedPartnerActivities[name] = []*PartnerActivity{}
	} else {
		au.Edges.namedPartnerActivities[name] = append(au.Edges.namedPartnerActivities[name], edges...)
	}
}

// NamedJobActivities returns the JobActivities named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedJobActivities(name string) ([]*JobActivity, error) {
	if au.Edges.namedJobActivities == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedJobActivities[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedJobActivities(name string, edges ...*JobActivity) {
	if au.Edges.namedJobActivities == nil {
		au.Edges.namedJobActivities = make(map[string][]*JobActivity)
	}
	if len(edges) == 0 {
		au.Edges.namedJobActivities[name] = []*JobActivity{}
	} else {
		au.Edges.namedJobActivities[name] = append(au.Edges.namedJobActivities[name], edges...)
	}
}

// NamedNotifications returns the Notifications named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedNotifications(name string) ([]*ChannelMessage, error) {
	if au.Edges.namedNotifications == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedNotifications[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedNotifications(name string, edges ...*ChannelMessage) {
	if au.Edges.namedNotifications == nil {
		au.Edges.namedNotifications = make(map[string][]*ChannelMessage)
	}
	if len(edges) == 0 {
		au.Edges.namedNotifications[name] = []*ChannelMessage{}
	} else {
		au.Edges.namedNotifications[name] = append(au.Edges.namedNotifications[name], edges...)
	}
}

// NamedJobProgressHistory returns the JobProgressHistory named value or an error if the edge was not
// loaded in eager-loading with this name.
func (au *ApiUser) NamedJobProgressHistory(name string) ([]*JobProgressHistory, error) {
	if au.Edges.namedJobProgressHistory == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := au.Edges.namedJobProgressHistory[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (au *ApiUser) appendNamedJobProgressHistory(name string, edges ...*JobProgressHistory) {
	if au.Edges.namedJobProgressHistory == nil {
		au.Edges.namedJobProgressHistory = make(map[string][]*JobProgressHistory)
	}
	if len(edges) == 0 {
		au.Edges.namedJobProgressHistory[name] = []*JobProgressHistory{}
	} else {
		au.Edges.namedJobProgressHistory[name] = append(au.Edges.namedJobProgressHistory[name], edges...)
	}
}

// ApiUsers is a parsable slice of ApiUser.
type ApiUsers []*ApiUser
