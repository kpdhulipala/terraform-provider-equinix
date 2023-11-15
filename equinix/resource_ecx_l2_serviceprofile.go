package equinix

import (
	"context"
	"fmt"

	"github.com/equinix/terraform-provider-equinix/internal/config"

	"github.com/equinix/ecx-go/v2"
	"github.com/equinix/rest-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var ecxL2ServiceProfileSchemaNames = map[string]string{
	"UUID":                                "uuid",
	"State":                               "state",
	"AlertPercentage":                     "bandwidth_alert_threshold",
	"AllowCustomSpeed":                    "speed_customization_allowed",
	"AllowOverSubscription":               "oversubscription_allowed",
	"APIAvailable":                        "api_integration",
	"AuthKeyLabel":                        "authkey_label",
	"ConnectionNameLabel":                 "connection_name_label",
	"CTagLabel":                           "ctag_label",
	"Description":                         "description",
	"EnableAutoGenerateServiceKey":        "servicekey_autogenerated",
	"EquinixManagedPortAndVlan":           "equinix_managed_port_vlan",
	"IntegrationID":                       "integration_id",
	"Name":                                "name",
	"OnBandwidthThresholdNotification":    "bandwidth_threshold_notifications",
	"OnProfileApprovalRejectNotification": "profile_statuschange_notifications",
	"OnVcApprovalRejectionNotification":   "vc_statuschange_notifications",
	"OverSubscription":                    "oversubscription",
	"Private":                             "private",
	"PrivateUserEmails":                   "private_user_emails",
	"RequiredRedundancy":                  "redundancy_required",
	"SpeedFromAPI":                        "speed_from_api",
	"TagType":                             "tag_type",
	"VlanSameAsPrimary":                   "secondary_vlan_from_primary",
	"Features":                            "features",
	"Port":                                "port",
	"SpeedBand":                           "speed_band",
}

var ecxL2ServiceProfileDescriptions = map[string]string{
	"UUID":                                "Unique identifier of the service profile",
	"State":                               "Service profile provisioning status",
	"AlertPercentage":                     "Specifies the port bandwidth threshold percentage. If the bandwidth limit is met or exceeded, an alert is sent to the seller",
	"AllowCustomSpeed":                    "Boolean value that determines if customer is allowed to enter a custom connection speed",
	"AllowOverSubscription":               "Boolean value that determines if, regardless of the utilization, Equinix Fabric will continue to add connections to your links until we reach the oversubscription limit",
	"APIAvailable":                        "Boolean value that determines if API integration is enabled",
	"AuthKeyLabel":                        "Name of the authentication key label to be used by the Authentication Key service",
	"ConnectionNameLabel":                 "Custom name used for calling a connections i.e. circuit. Defaults to Connection",
	"CTagLabel":                           "C-Tag/Inner-Tag label name for the connections",
	"Description":                         "Description of the service profile",
	"EnableAutoGenerateServiceKey":        "Boolean value that indicates whether multiple connections can be created with the same authorization key",
	"EquinixManagedPortAndVlan":           "Boolean value that indicates whether the port and VLAN details are managed by Equinix",
	"IntegrationID":                       "Specifies the API integration ID that was provided to the customer during onboarding",
	"Name":                                "Name of the service profile. An alpha-numeric 50 characters string which can include only hyphens and underscores",
	"OnBandwidthThresholdNotification":    "A list of email addresses that will receive notifications about bandwidth thresholds",
	"OnProfileApprovalRejectNotification": "A list of email addresses that will receive notifications about profile status changes",
	"OnVcApprovalRejectionNotification":   "A list of email addresses that will receive notifications about connections approvals and rejections",
	"OverSubscription":                    "Oversubscription limit that will cause alerting. Default is 1x",
	"Private":                             "Boolean value that indicates whether or not this is a private profile.",
	"PrivateUserEmails":                   "A list of email addresses associated to users that will be allowed to access this service profile. Applicable for private profiles",
	"RequiredRedundancy":                  "Boolean value that determines if yourconnections will require redundancy",
	"SpeedFromAPI":                        "Boolean valuta that determines if connection speed will be derived from an API call",
	"TagType":                             "Specifies additional tagging information required by the seller profile for Dot1Q to QinQ translation",
	"VlanSameAsPrimary":                   "Indicates whether the VLAN ID of the secondary connection is the same as the primary connection",
	"Features":                            "Block of profile features configuration",
	"Port":                                "One or more definitions of ports associated with the profile",
	"SpeedBand":                           "One or more definitions of supported speed/bandwidth configurations",
}

var ecxL2ServiceProfileFeaturesSchemaNames = map[string]string{
	"CloudReach":  "allow_remote_connections",
	"TestProfile": "test_profile",
}

var ecxL2ServiceProfileFeaturesDescriptions = map[string]string{
	"CloudReach":  "Indicates whether or not connections to this profile can be created from remote metro locations",
	"TestProfile": "Indicates whether or not this profile can be used for test connections",
}

var ecxL2ServiceProfilePortSchemaNames = map[string]string{
	"ID":        "uuid",
	"MetroCode": "metro_code",
}

var ecxL2ServiceProfilePortDescriptions = map[string]string{
	"ID":        "Unique identifier of the port",
	"MetroCode": "Port location metro code",
}

var ecxL2ServiceProfileSpeedBandSchemaNames = map[string]string{
	"Speed":     "speed",
	"SpeedUnit": "speed_unit",
}

var ecxL2ServiceProfileSpeedBandDescriptions = map[string]string{
	"Speed":     "Speed/bandwidth supported by given service profile",
	"SpeedUnit": "Unit of the speed/bandwidth supported by given service profile",
}

func resourceECXL2ServiceProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceECXL2ServiceProfileCreate,
		ReadContext:   resourceECXL2ServiceProfileRead,
		UpdateContext: resourceECXL2ServiceProfileUpdate,
		DeleteContext: resourceECXL2ServiceProfileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: createECXL2ServiceProfileResourceSchema(),
	}
}

func createECXL2ServiceProfileResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		ecxL2ServiceProfileSchemaNames["UUID"]: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: ecxL2ServiceProfileDescriptions["UUID"],
		},
		ecxL2ServiceProfileSchemaNames["State"]: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: ecxL2ServiceProfileDescriptions["State"],
		},
		ecxL2ServiceProfileSchemaNames["AlertPercentage"]: {
			Type:         schema.TypeFloat,
			Optional:     true,
			ValidateFunc: validation.FloatBetween(1, 99),
			Description:  ecxL2ServiceProfileDescriptions["AlertPercentage"],
		},
		ecxL2ServiceProfileSchemaNames["AllowCustomSpeed"]: {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: ecxL2ServiceProfileDescriptions["AllowCustomSpeed"],
		},
		ecxL2ServiceProfileSchemaNames["AllowOverSubscription"]: {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: ecxL2ServiceProfileDescriptions["AllowOverSubscription"],
		},
		ecxL2ServiceProfileSchemaNames["APIAvailable"]: {
			Type:         schema.TypeBool,
			Optional:     true,
			RequiredWith: []string{ecxL2ServiceProfileSchemaNames["IntegrationID"]},
			Description:  ecxL2ServiceProfileDescriptions["IntegrationID"],
		},
		ecxL2ServiceProfileSchemaNames["AuthKeyLabel"]: {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(1, 50),
			Description:  ecxL2ServiceProfileDescriptions["AuthKeyLabel"],
		},
		ecxL2ServiceProfileSchemaNames["ConnectionNameLabel"]: {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "Connection",
			ValidateFunc: validation.StringLenBetween(1, 50),
			Description:  ecxL2ServiceProfileDescriptions["ConnectionNameLabel"],
		},
		ecxL2ServiceProfileSchemaNames["CTagLabel"]: {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(1, 50),
			Description:  ecxL2ServiceProfileDescriptions["CTagLabel"],
		},
		ecxL2ServiceProfileSchemaNames["Description"]: {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(1, 250),
			Description:  ecxL2ServiceProfileDescriptions["Description"],
		},
		ecxL2ServiceProfileSchemaNames["EnableAutoGenerateServiceKey"]: {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: ecxL2ServiceProfileDescriptions["EnableAutoGenerateServiceKey"],
		},
		ecxL2ServiceProfileSchemaNames["EquinixManagedPortAndVlan"]: {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: ecxL2ServiceProfileDescriptions["EquinixManagedPortAndVlan"],
		},
		ecxL2ServiceProfileSchemaNames["IntegrationID"]: {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(1, 50),
			Description:  ecxL2ServiceProfileDescriptions["IntegrationID"],
		},
		ecxL2ServiceProfileSchemaNames["Name"]: {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringLenBetween(1, 50),
			Description:  ecxL2ServiceProfileDescriptions["Name"],
		},
		ecxL2ServiceProfileSchemaNames["OnBandwidthThresholdNotification"]: {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: ecxL2ServiceProfileDescriptions["OnBandwidthThresholdNotification"],
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: stringIsEmailAddress(),
			},
		},
		ecxL2ServiceProfileSchemaNames["OnProfileApprovalRejectNotification"]: {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: ecxL2ServiceProfileDescriptions["OnProfileApprovalRejectNotification"],
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: stringIsEmailAddress(),
			},
		},
		ecxL2ServiceProfileSchemaNames["OnVcApprovalRejectionNotification"]: {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: ecxL2ServiceProfileDescriptions["OnVcApprovalRejectionNotification"],
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: stringIsEmailAddress(),
			},
		},
		ecxL2ServiceProfileSchemaNames["OverSubscription"]: {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1x",
			RequiredWith: []string{ecxL2ServiceProfileSchemaNames["AllowOverSubscription"]},
			Description:  ecxL2ServiceProfileDescriptions["OverSubscription"],
		},
		ecxL2ServiceProfileSchemaNames["Private"]: {
			Type:         schema.TypeBool,
			Optional:     true,
			RequiredWith: []string{ecxL2ServiceProfileSchemaNames["PrivateUserEmails"]},
			Description:  ecxL2ServiceProfileDescriptions["Private"],
		},
		ecxL2ServiceProfileSchemaNames["PrivateUserEmails"]: {
			Type:        schema.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: ecxL2ServiceProfileDescriptions["PrivateUserEmails"],
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: stringIsEmailAddress(),
			},
		},
		ecxL2ServiceProfileSchemaNames["RequiredRedundancy"]: {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: ecxL2ServiceProfileDescriptions["RequiredRedundancy"],
		},
		ecxL2ServiceProfileSchemaNames["SpeedFromAPI"]: {
			Type:         schema.TypeBool,
			Optional:     true,
			AtLeastOneOf: []string{ecxL2ServiceProfileSchemaNames["SpeedFromAPI"], ecxL2ServiceProfileSchemaNames["SpeedBand"]},
			RequiredWith: []string{ecxL2ServiceProfileSchemaNames["APIAvailable"]},
			Description:  ecxL2ServiceProfileDescriptions["SpeedFromAPI"],
		},
		ecxL2ServiceProfileSchemaNames["TagType"]: {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "CTAGED",
			ValidateFunc: validation.StringInSlice([]string{"CTAGED", "NAMED", "BOTH"}, false),
			Description:  ecxL2ServiceProfileDescriptions["TagType"],
		},
		ecxL2ServiceProfileSchemaNames["VlanSameAsPrimary"]: {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: ecxL2ServiceProfileDescriptions["VlanSameAsPrimary"],
		},
		ecxL2ServiceProfileSchemaNames["Features"]: {
			Type:        schema.TypeSet,
			Required:    true,
			MaxItems:    1,
			Description: ecxL2ServiceProfileDescriptions["Features"],
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					ecxL2ServiceProfileFeaturesSchemaNames["CloudReach"]: {
						Type:        schema.TypeBool,
						Required:    true,
						Description: ecxL2ServiceProfileFeaturesDescriptions["CloudReach"],
					},
					ecxL2ServiceProfileFeaturesSchemaNames["TestProfile"]: {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: ecxL2ServiceProfileFeaturesDescriptions["TestProfile"],
						Deprecated:  "TestProfile is no longer required and will be removed in a future release",
					},
				},
			},
		},
		ecxL2ServiceProfileSchemaNames["Port"]: {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: ecxL2ServiceProfileDescriptions["Port"],
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					ecxL2ServiceProfilePortSchemaNames["ID"]: {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
						Description:  ecxL2ServiceProfilePortDescriptions["ID"],
					},
					ecxL2ServiceProfilePortSchemaNames["MetroCode"]: {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: stringIsMetroCode(),
						Description:  ecxL2ServiceProfilePortDescriptions["MetroCode"],
					},
				},
			},
		},
		ecxL2ServiceProfileSchemaNames["SpeedBand"]: {
			Type:         schema.TypeSet,
			MinItems:     1,
			Optional:     true,
			AtLeastOneOf: []string{ecxL2ServiceProfileSchemaNames["SpeedFromAPI"], ecxL2ServiceProfileSchemaNames["SpeedBand"]},
			Description:  ecxL2ServiceProfileDescriptions["SpeedBand"],
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					ecxL2ServiceProfileSpeedBandSchemaNames["Speed"]: {
						Type:        schema.TypeInt,
						Required:    true,
						Description: ecxL2ServiceProfileSpeedBandDescriptions["Speed"],
					},
					ecxL2ServiceProfileSpeedBandSchemaNames["SpeedUnit"]: {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"MB", "GB"}, false),
						Description:  ecxL2ServiceProfileSpeedBandDescriptions["SpeedUnit"],
					},
				},
			},
		},
	}
}

func resourceECXL2ServiceProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*config.Config).Ecx
	m.(*config.Config).AddModuleToECXUserAgent(&client, d)
	var diags diag.Diagnostics
	profile := createECXL2ServiceProfile(d)
	uuid, err := client.CreateL2ServiceProfile(*profile)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(ecx.StringValue(uuid))
	diags = append(diags, resourceECXL2ServiceProfileRead(ctx, d, m)...)
	return diags
}

func resourceECXL2ServiceProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*config.Config).Ecx
	m.(*config.Config).AddModuleToECXUserAgent(&client, d)
	var diags diag.Diagnostics
	profile, err := client.GetL2ServiceProfile(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if err := updateECXL2ServiceProfileResource(profile, d); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceECXL2ServiceProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*config.Config).Ecx
	m.(*config.Config).AddModuleToECXUserAgent(&client, d)
	var diags diag.Diagnostics
	profile := createECXL2ServiceProfile(d)
	if err := client.UpdateL2ServiceProfile(*profile); err != nil {
		return diag.FromErr(err)
	}
	diags = append(diags, resourceECXL2ServiceProfileRead(ctx, d, m)...)
	return diags
}

func resourceECXL2ServiceProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*config.Config).Ecx
	m.(*config.Config).AddModuleToECXUserAgent(&client, d)
	var diags diag.Diagnostics
	if err := client.DeleteL2ServiceProfile(d.Id()); err != nil {
		restErr, ok := err.(rest.Error)
		if ok {
			// IC-PROFILE-004 =  profile does not exist
			if hasApplicationErrorCode(restErr.ApplicationErrors, "IC-PROFILE-004") {
				return diags
			}
		}
		return diag.FromErr(err)
	}
	return diags
}

func createECXL2ServiceProfile(d *schema.ResourceData) *ecx.L2ServiceProfile {
	profile := ecx.L2ServiceProfile{}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["UUID"]); ok {
		profile.UUID = ecx.String(v.(string))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["State"]); ok {
		profile.State = ecx.String(v.(string))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["AlertPercentage"]); ok {
		profile.AlertPercentage = ecx.Float64(v.(float64))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["AllowCustomSpeed"]); ok {
		profile.AllowCustomSpeed = ecx.Bool(v.(bool))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["AllowOverSubscription"]); ok {
		profile.AllowOverSubscription = ecx.Bool(v.(bool))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["APIAvailable"]); ok {
		profile.APIAvailable = ecx.Bool(v.(bool))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["AuthKeyLabel"]); ok {
		profile.AuthKeyLabel = ecx.String(v.(string))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["ConnectionNameLabel"]); ok {
		profile.ConnectionNameLabel = ecx.String(v.(string))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["CTagLabel"]); ok {
		profile.CTagLabel = ecx.String(v.(string))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["Description"]); ok {
		profile.Description = ecx.String(v.(string))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["EnableAutoGenerateServiceKey"]); ok {
		profile.EnableAutoGenerateServiceKey = ecx.Bool(v.(bool))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["EquinixManagedPortAndVlan"]); ok {
		profile.EquinixManagedPortAndVlan = ecx.Bool(v.(bool))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["IntegrationID"]); ok {
		profile.IntegrationID = ecx.String(v.(string))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["Name"]); ok {
		profile.Name = ecx.String(v.(string))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["OnBandwidthThresholdNotification"]); ok {
		profile.OnBandwidthThresholdNotification = expandSetToStringList(v.(*schema.Set))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["OnProfileApprovalRejectNotification"]); ok {
		profile.OnProfileApprovalRejectNotification = expandSetToStringList(v.(*schema.Set))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["OnVcApprovalRejectionNotification"]); ok {
		profile.OnVcApprovalRejectionNotification = expandSetToStringList(v.(*schema.Set))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["OverSubscription"]); ok {
		profile.OverSubscription = ecx.String(v.(string))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["Private"]); ok {
		profile.Private = ecx.Bool(v.(bool))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["PrivateUserEmails"]); ok {
		profile.PrivateUserEmails = expandSetToStringList(v.(*schema.Set))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["RequiredRedundancy"]); ok {
		profile.RequiredRedundancy = ecx.Bool(v.(bool))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["SpeedFromAPI"]); ok {
		profile.SpeedFromAPI = ecx.Bool(v.(bool))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["TagType"]); ok {
		profile.TagType = ecx.String(v.(string))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["VlanSameAsPrimary"]); ok {
		profile.VlanSameAsPrimary = ecx.Bool(v.(bool))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["Features"]); ok {
		featureSet := v.(*schema.Set)
		if featureSet.Len() > 0 {
			profile.Features = expandECXL2ServiceProfileFeatures(featureSet.List())
		}
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["Port"]); ok {
		profile.Ports = expandECXL2ServiceProfilePorts(v.(*schema.Set))
	}
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["SpeedBand"]); ok {
		profile.SpeedBands = expandECXL2ServiceProfileSpeedBands(v.(*schema.Set))
	}
	return &profile
}

func updateECXL2ServiceProfileResource(profile *ecx.L2ServiceProfile, d *schema.ResourceData) error {
	if err := d.Set(ecxL2ServiceProfileSchemaNames["UUID"], profile.UUID); err != nil {
		return fmt.Errorf("error reading UUID: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["State"], profile.State); err != nil {
		return fmt.Errorf("error reading State: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["AlertPercentage"], profile.AlertPercentage); err != nil {
		return fmt.Errorf("error reading AlertPercentage: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["AllowCustomSpeed"], profile.AllowCustomSpeed); err != nil {
		return fmt.Errorf("error reading AllowCustomSpeed: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["AllowOverSubscription"], profile.AllowOverSubscription); err != nil {
		return fmt.Errorf("error reading AllowOverSubscription: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["APIAvailable"], profile.APIAvailable); err != nil {
		return fmt.Errorf("error reading APIAvailable: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["AuthKeyLabel"], profile.AuthKeyLabel); err != nil {
		return fmt.Errorf("error reading AuthKeyLabel: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["ConnectionNameLabel"], profile.ConnectionNameLabel); err != nil {
		return fmt.Errorf("error reading ConnectionNameLabel: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["CTagLabel"], profile.CTagLabel); err != nil {
		return fmt.Errorf("error reading CTagLabel: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["Description"], profile.Description); err != nil {
		return fmt.Errorf("error reading Description: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["EnableAutoGenerateServiceKey"], profile.EnableAutoGenerateServiceKey); err != nil {
		return fmt.Errorf("error reading EnableAutoGenerateServiceKey: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["EquinixManagedPortAndVlan"], profile.EquinixManagedPortAndVlan); err != nil {
		return fmt.Errorf("error reading EquinixManagedPortAndVlan: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["IntegrationID"], profile.IntegrationID); err != nil {
		return fmt.Errorf("error reading IntegrationID: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["Name"], profile.Name); err != nil {
		return fmt.Errorf("error reading Name: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["OnBandwidthThresholdNotification"], profile.OnBandwidthThresholdNotification); err != nil {
		return fmt.Errorf("error reading OnBandwidthThresholdNotification: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["OnProfileApprovalRejectNotification"], profile.OnProfileApprovalRejectNotification); err != nil {
		return fmt.Errorf("error reading OnProfileApprovalRejectNotification: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["OnVcApprovalRejectionNotification"], profile.OnVcApprovalRejectionNotification); err != nil {
		return fmt.Errorf("error reading OnVcApprovalRejectionNotification: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["OverSubscription"], profile.OverSubscription); err != nil {
		return fmt.Errorf("error reading OverSubscription: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["Private"], profile.Private); err != nil {
		return fmt.Errorf("error reading Private: %s", err)
	}
	// API accepts capitalizations of the private user emails and converts it to a lowercase string
	// If API retuns same emails in lowercase we keep to suppress diff
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["PrivateUserEmails"]); ok {
		prevPrivateUserEmails := expandSetToStringList(v.(*schema.Set))
		if slicesMatchCaseInsensitive(prevPrivateUserEmails, profile.PrivateUserEmails) {
			profile.PrivateUserEmails = prevPrivateUserEmails
		}
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["PrivateUserEmails"], profile.PrivateUserEmails); err != nil {
		return fmt.Errorf("error reading PrivateUserEmails: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["RequiredRedundancy"], profile.RequiredRedundancy); err != nil {
		return fmt.Errorf("error reading RequiredRedundancy: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["SpeedFromAPI"], profile.SpeedFromAPI); err != nil {
		return fmt.Errorf("error reading SpeedFromAPI: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["TagType"], profile.TagType); err != nil {
		return fmt.Errorf("error reading TagType: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["VlanSameAsPrimary"], profile.VlanSameAsPrimary); err != nil {
		return fmt.Errorf("error reading VlanSameAsPrimary: %s", err)
	}
	var prevFeatures *ecx.L2ServiceProfileFeatures
	if v, ok := d.GetOk(ecxL2ServiceProfileSchemaNames["Features"]); ok {
		expandedFeatures := expandECXL2ServiceProfileFeatures(v.(*schema.Set).List())
		prevFeatures = &expandedFeatures
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["Features"], flattenECXL2ServiceProfileFeatures(prevFeatures, profile.Features)); err != nil {
		return fmt.Errorf("error reading Features: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["Port"], flattenECXL2ServiceProfilePorts(profile.Ports)); err != nil {
		return fmt.Errorf("error reading Port: %s", err)
	}
	if err := d.Set(ecxL2ServiceProfileSchemaNames["SpeedBand"], flattenECXL2ServiceProfileSpeedBands(profile.SpeedBands)); err != nil {
		return fmt.Errorf("error reading SpeedBand: %s", err)
	}
	return nil
}

func flattenECXL2ServiceProfileFeatures(previous *ecx.L2ServiceProfileFeatures, features ecx.L2ServiceProfileFeatures) interface{} {
	transformed := make(map[string]interface{})
	transformed[ecxL2ServiceProfileFeaturesSchemaNames["CloudReach"]] = features.CloudReach
	transformed[ecxL2ServiceProfileFeaturesSchemaNames["TestProfile"]] = features.TestProfile
	if previous != nil {
		transformed[ecxL2ServiceProfileFeaturesSchemaNames["TestProfile"]] = previous.TestProfile
	}
	return []map[string]interface{}{transformed}
}

func flattenECXL2ServiceProfilePorts(ports []ecx.L2ServiceProfilePort) interface{} {
	transformed := make([]interface{}, 0, len(ports))
	for _, port := range ports {
		transformed = append(transformed, map[string]interface{}{
			ecxL2ServiceProfilePortSchemaNames["ID"]:        port.ID,
			ecxL2ServiceProfilePortSchemaNames["MetroCode"]: port.MetroCode,
		})
	}
	return transformed
}

func flattenECXL2ServiceProfileSpeedBands(bands []ecx.L2ServiceProfileSpeedBand) interface{} {
	transformed := make([]interface{}, 0, len(bands))
	for _, band := range bands {
		transformed = append(transformed, map[string]interface{}{
			ecxL2ServiceProfileSpeedBandSchemaNames["Speed"]:     band.Speed,
			ecxL2ServiceProfileSpeedBandSchemaNames["SpeedUnit"]: band.SpeedUnit,
		})
	}
	return transformed
}

func expandECXL2ServiceProfileFeatures(features []interface{}) ecx.L2ServiceProfileFeatures {
	transformed := ecx.L2ServiceProfileFeatures{}
	if len(features) > 0 {
		feature := features[0].(map[string]interface{})
		if v, ok := feature[ecxL2ServiceProfileFeaturesSchemaNames["CloudReach"]]; ok {
			transformed.CloudReach = ecx.Bool(v.(bool))
		}
		if v, ok := feature[ecxL2ServiceProfileFeaturesSchemaNames["TestProfile"]]; ok {
			transformed.TestProfile = ecx.Bool(v.(bool))
		}
	}
	return transformed
}

func expandECXL2ServiceProfilePorts(ports *schema.Set) []ecx.L2ServiceProfilePort {
	transformed := make([]ecx.L2ServiceProfilePort, 0, ports.Len())
	for _, port := range ports.List() {
		portMap := port.(map[string]interface{})
		transformed = append(transformed, ecx.L2ServiceProfilePort{
			ID:        ecx.String(portMap[ecxL2ServiceProfilePortSchemaNames["ID"]].(string)),
			MetroCode: ecx.String(portMap[ecxL2ServiceProfilePortSchemaNames["MetroCode"]].(string)),
		})
	}
	return transformed
}

func expandECXL2ServiceProfileSpeedBands(bands *schema.Set) []ecx.L2ServiceProfileSpeedBand {
	transformed := make([]ecx.L2ServiceProfileSpeedBand, 0, bands.Len())
	for _, band := range bands.List() {
		bandMap := band.(map[string]interface{})
		transformed = append(transformed, ecx.L2ServiceProfileSpeedBand{
			Speed:     ecx.Int(bandMap[ecxL2ServiceProfileSpeedBandSchemaNames["Speed"]].(int)),
			SpeedUnit: ecx.String(bandMap[ecxL2ServiceProfileSpeedBandSchemaNames["SpeedUnit"]].(string)),
		})
	}
	return transformed
}
