package datatypes

type Deactivate struct {
	Status string `json:"status"`
}

type DataExport struct {
	Status string `json:"status"`
}

type AccountCheck struct {
	AccountPlan AccountCheckPlan `json:"account_plan"`
	UserCountry string           `json:"user_country"`
	Features    []string         `json:"features"`
}

type AccountCheckPlan struct {
	IsPaidSubscriptionActive       bool   `json:"is_paid_subscription_active"`
	SubscriptionPlan               string `json:"subscription_plan"`
	AccountUserRole                string `json:"account_user_role"`
	WasPaidCustomer                bool   `json:"was_paid_customer"`
	HasCustomerObject              bool   `json:"has_customer_object"`
	SubscriptionExpiresAtTimestamp any    `json:"subscription_expires_at_timestamp"`
}

const (
	// stats
	FEATURE_LOG_STATSIG_EVENTS  = "log_statsig_events"  // Enabled by default, statistical analysis report
	FEATURE_LOG_INTERCOM_EVENTS = "log_intercom_events" // Enabled by default, statistical analysis report
	// data controls
	FEATURE_DATA_DELETION_ENABLE = "data_deletion_enabled" // Enabled by default, allowing account deletion
	FEATURE_DATA_EXPORT          = "data_export_enabled"   // Enabled by default, allows data to be exported
	FEATURE_DATA_CONTROL         = "data_controls_enabled" // Enabled by default, allows users to control data
	// messages
	FEATURE_DFW_MESSAGE_FEEDBACK     = "dfw_message_feedback"                // Enabled by default
	FEATURE_DFW_INLINE_MESSAGE_REGEN = "dfw_inline_message_regen_comparison" // Enabled by default
	FEATURE_SYSTEM_MESSAGE           = "system_message"                      // Enabled by default
	// account
	FEATURE_ONFOFF_STATUE_ACCOUNT                = "oneoff_status_account"                     // If you purchase a service during a service interruption, you will be prompted for a refund
	FEATURE_SHOW_EXISTING_USER_AGE_CONFIRM_MODAL = "show_existing_user_age_confirmation_modal" // 23.05.08 Added, display age confirmation pop-up window
	// models
	FEATURE_MODEL_SWITCHER  = "model_switcher" // The model can be switched in the interface, and the Plus account is enabled by default
	FEATURE_MODEL_PREVIEWER = "model_preview"  // Remind the limit when using the preview model, the Plus account is enabled by default
	// misc
	FEATURE_DISABLE_UPGRADE_UI = "disable_upgrade_ui" // Enabled by default except for Plus accounts
	FEATURE_DISABLE_HISTORY    = "disable_history"    // Disable the session history, only the interface is reflected
	FEATURE_BUCKETED_HISTORY   = "bucketed_history"   // Enabled by default, Display history in buckets
	FEATURE_SHAREABLE_LINKS    = "shareable_links"    // 23.05.08 Added, conversation sharing feat
	// plugins
	FEATURE_DEBUG               = "debug"        // Developer permissions, debug mode
	FEATURE_PLIGIN_BROWSING     = "tools"        // Plug-in permissions
	FEATURE_PLIGIN_CODE         = "tools2"       // Plug-in permissions
	FEATURE_PLIGIN_PLUGIN       = "tools3"       // Plug-in permissions
	FEATURE_PLIGIN_PLUGIN_ADMIN = "tools3_admin" // Plug-in permissions
	FEATURE_PLIGIN_PLUGIN_DEV   = "tools3_dev"   // Plug-in permissions
)
