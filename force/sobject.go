package force

import (
	"strings"
)

// IMPORTANT! When adding fields to any of the SObject*Metadata structs, you must abide by the following:
// - Do not delete any existing fields
// - Add `omitempty` to the json tag
// - Add a field comment

// SObject is a single record/object in salesforce.
type SObject map[string]interface{}

// SObjectMetadata contains detailed metadata about a particular sObject.
//
// This fields defined in SObjectMetadata represent a cumulative/running list from all currently supported
// Salesforce API versions. Not all fields may return a value depending on the API version used for a DESCRIBE request.
//
// Reference: https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm
type SObjectMetadata struct {
	// Activateable - Reserved for future use.
	Activateable *bool `json:"activateable,omitempty"`

	// AssociateEntityType - If the object is associated with a parent object, the type of association it has to its
	// parent, such as History. Otherwise, its value is null. Available in API version 50.0 and later.
	AssociateEntityType *string `json:"associateEntityType,omitempty"`

	// AssociateParentEntity - If the object is associated with a parent object, the parent object it’s associated with.
	// Otherwise, its value is null. Available in API version 50.0 and later.
	AssociateParentEntity *string `json:"associateParentEntity,omitempty"`

	// Createable - Indicates whether the object can be created via the create() call (true) or not (false).
	Createable *bool `json:"createable,omitempty"`

	// Custom - Indicates whether the object is a custom object (true) or not (false).
	Custom *bool `json:"custom,omitempty"`

	// CustomSetting - Indicates whether the object is a custom setting object (true) or not (false).
	CustomSetting *bool `json:"customSetting,omitempty"`

	// Deletable - Indicates whether the object can be deleted via the delete() call (true) or not (false).
	Deletable *bool `json:"deletable,omitempty"`

	// DeprecatedAndHidden - Reserved for future use.
	DeprecatedAndHidden *bool `json:"deprecatedAndHidden,omitempty"`

	// FeedEnabled - Indicates whether Chatter feeds are enabled for the object (true) or not (false).
	// This property is available in API version 19.0 and later.
	FeedEnabled *bool `json:"feedEnabled,omitempty"`

	// KeyPrefix - Three-character prefix code in the object ID. Object IDs are prefixed with three-character codes
	// that specify the type of the object. For example, Account objects have a prefix of 001 and Opportunity objects
	// have a prefix of 006. Note that a key prefix can sometimes be shared by multiple objects so it does not always
	// uniquely identify an object. Use the value of this field to determine the object type of a parent in those cases
	// where the child may have more than one object type as parent (polymorphic).
	// For example, you may need to obtain the keyPrefix value for the parent of a Task or Event.
	KeyPrefix *string `json:"keyPrefix,omitempty"`

	// Label - Text for a tab or field renamed in the user interface, if applicable, or the object name, if not.
	// For example, an organization representing a medical vertical might rename Account to Patient.
	// Tabs and fields can be renamed in the Salesforce user interface. See the Salesforce online help for more information.
	Label *string `json:"label,omitempty"`

	// LabelPlural - text for an object that represents the plural version of an object name, for example, “Accounts.”
	LabelPlural *string `json:"labelPlural,omitempty"`

	// Layoutable - Indicates whether the object supports the describeLayout() call (true) or not (false).
	Layoutable *bool `json:"layoutable,omitempty"`

	// Mergeable - Indicates whether the object can be merged with other objects of its type (true) or not (false).
	// true for leads, contacts, and accounts.
	Mergeable *bool `json:"mergeable,omitempty"`

	// Name of the object. This is the same string that was passed in as the sObjectType parameter.
	Name *string `json:"name,omitempty"`

	// Queryable - Indicates whether the object can be queried via the query() call (true) or not (false).
	Queryable *bool `json:"queryable,omitempty"`

	// Replicateable - Indicates whether the object can be replicated via the getUpdated() and getDeleted()
	// calls (true) or not (false).
	Replicateable *bool `json:"replicateable,omitempty"`

	// Retrieveable - Indicates whether the object can be retrieved via the retrieve() call (true) or not (false).
	Retrieveable *bool `json:"retrieveable,omitempty"`

	// Searchable - Indicates whether the object can be searched via the search() call (true) or not (false).
	Searchable *bool `json:"searchable,omitempty"`

	// Triggerable - Indicates whether the object supports Apex triggers.
	Triggerable *bool `json:"triggerable,omitempty"`

	// Undeletable - Indicates whether an object can be undeleted using the undelete() call (true) or not (false).
	Undeletable *bool `json:"undeletable,omitempty"`

	// Updateable - Indicates whether the object can be updated via the update() call (true) or not (false).
	Updateable *bool `json:"updateable,omitempty"`

	// URLs
	URLs map[string]string `json:"urls,omitempty"`

	// UrlDetail - URL to the read-only detail page for this object. Compare with urlEdit, which is read-write.
	// Client applications can use this URL to redirect to, or access, the Salesforce user interface for standard
	// and custom objects. To provide flexibility and allow for future enhancements, returned urlDetail values are dynamic.
	// To ensure that client applications are forward compatible, it is recommended that they use this capability where possible.
	// Note that, for objects for which a stable URL is not available, this field is returned empty.
	URLDetail *string `json:"urlDetail,omitempty"`

	// URLEdit - URL to the edit page for this object. For example, the urlEdit field for the Account object returns https://yourInstance.salesforce.com/{ID}/e.
	// Substituting the {ID} field for the current object ID will return the edit page for that specific account in the Salesforce user interface.
	// Compare with urlDetail, which is read-only. Client applications can use this URL to redirect to, or access,
	// the Salesforce user interface for standard and custom objects. To provide flexibility and allow for future enhancements,
	// returned urlDetail values are dynamic. To ensure that client applications are forward compatible,
	// it is recommended that they use this capability where possible. Note that, for objects for which a stable URL is not available,
	// this field is returned empty.
	URLEdit *string `json:"urlEdit,omitempty"`

	// URLNew 	URL to the new/create page for this object. Client applications can use this URL to redirect to, or access,
	// the Salesforce user interface for standard and custom objects. To provide flexibility and allow for future enhancements,
	// returned urlNew values are dynamic. To ensure that client applications are forward compatible,
	// it is recommended that they use this capability where possible. Note that, for objects for which a stable URL is not available,
	// this field is returned empty.
	URLNew *string `json:"urlNew,omitempty"`

	// ActionOverrides - An array of action overrides. Action overrides replace the URLs specified in the urlDetail, urlEdit and urlNew fields.
	// This field is available in API version 32.0 and later.
	ActionOverrides []*SObjectActionOverrideMetadata `json:"actionOverrides,omitempty"`

	// ChildRelationships - An array of child relationships, which is the name of the sObject that has
	// a foreign key to the sObject being described.
	ChildRelationships []*SObjectChildRelationshipMetadata `json:"childRelationships,omitempty"`

	// CompactLayoutable - Indicates that the object can be used in describeCompactLayouts().
	CompactLayoutable *bool `json:"compactLayoutable,omitempty"`

	// Fields - Array of fields associated with the object. The mechanism for retrieving information from this
	// list varies among development tools.
	Fields []*SObjectFieldMetadata `json:"fields,omitempty"`

	// HasSubtypes
	HasSubtypes *bool `json:"hasSubtypes,omitempty"`

	// IDEnabled
	IDEnabled *bool `json:"idEnabled,omitempty"`

	// IsSubtype
	IsSubtype *bool `json:"isSubtype,omitempty"`

	// NamedLayoutInfos - The specific named layouts that are available for the objects other than the default layout.
	NamedLayoutInfos []*SObjectNamedLayoutInfoMetadata `json:"namedLayoutInfos,omitempty"`

	// NetworkScopeFieldName
	NetworkScopeFieldName string `json:"networkScopeFieldName,omitempty"`

	// RecordTypeInfos
	RecordTypeInfos []*SObjectRecordTypeInfoMetadata `json:"recordTypeInfos,omitempty"`

	// SearchLayoutable
	SearchLayoutable *bool `json:"searchLayoutable,omitempty"`

	// SupportedScopes is the list of supported scopes for the object. For example, Account might have supported scopes
	// of “All Accounts”, “My Accounts”, and “My Team’s Accounts”.
	SupportedScopes []*SObjectScopeInfoInfoMetadata `json:"supportedScopes,omitempty"`
}

// SObjectActionOverrideMetadata provides details about an action that replaces the default action pages for an object.
// For example, an object could be configured to replace the new/create page with a custom page.
// This type is available in API version 32.0 and later.
//
// Reference: https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm
type SObjectActionOverrideMetadata struct {
	// FormFactor represents the environment to which the action override applies.
	// For example, a Large value in this field represents the Lightning Experience desktop environment,
	// and is valid for Lightning pages and Lightning components. A Small value represents the Salesforce mobile
	// app on a phone or tablet. This field is available in API version 37.0 and later.
	FormFactor *string `json:"formFactor,omitempty"`

	// IsAvailableInTouch indicates whether the action override is available in the Salesforce mobile app (true) or not (false).
	IsAvailableInTouch *bool `json:"isAvailableInTouch,omitempty"`

	// Name of the action that overrides the default action. For example, if the new/create page was overridden with a
	// custom action, the name might be “New”.
	Name *string `json:"name,omitempty"`

	// PageID is the ID of the page for the action override.
	PageID *string `json:"pageId,omitempty"`

	// URL of the item being used for the action override, such as a Visualforce page.
	// Returns as null for Lightning page overrides.
	URL *string `json:"url,omitempty"`
}

// SObjectChildRelationshipMetadata provides details of the sObject that has a foreign key to the sObject being described.
//
// Reference: https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm
type SObjectChildRelationshipMetadata struct {
	// CascadeDelete indicates whether the child object is deleted when the parent object is deleted (true) or not (false).
	CascadeDelete *bool `json:"cascadeDelete,omitempty"`

	// ChildSObject is the name of the object on which there is a foreign key back to the parent sObject.
	ChildSObject *string `json:"childSObject,omitempty"`

	// DeprecatedAndHidden is reserved for future use.
	DeprecatedAndHidden *bool `json:"deprecatedAndHidden,omitempty"`

	// Field is the name of the field that has a foreign key back to the parent sObject.
	Field *string `json:"field,omitempty"`

	// JunctionIdListNames are the names of the lists of junction IDs associated with an object.
	// Each ID represents an object that has a relationship with the associated object.
	JunctionIdListNames []string `json:"junctionIdListNames,omitempty"`

	// JunctionReferenceTo is a collection of object names that the polymorphic keys in the junctionIdListNames
	// property can reference. You can query these object names.
	JunctionReferenceTo []string `json:"junctionReferenceTo,omitempty"`

	// RelationshipName is the name of the relationship, usually the plural of the value in childSObject.
	RelationshipName *string `json:"relationshipName,omitempty"`

	// RestrictedDelete indicates whether the parent object can’t be deleted because it is referenced by a child object (true) or not (false).
	RestrictedDelete *bool `json:"restrictedDelete,omitempty"`
}

// SObjectFieldMetadata represents a field in an API object. The array contains only the fields that the user can view,
// as defined by the user's field-level security settings.
//
// Reference: https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm
type SObjectFieldMetadata struct {
	// Aggregatable
	Aggregatable *bool `json:"aggregatable,omitempty"`

	// AutoNumber - Indicates whether this field is an autonumber field (true) or not (false). Analogous to a SQL IDENTITY type,
	// autonumber fields are read only, non-createable text fields with a maximum length of 30 characters.
	// Autonumber fields are read-only fields used to provide a unique ID that is independent of the internal object ID
	// (such as a purchase order number or invoice number). Autonumber fields are configured entirely in the Salesforce user interface.
	// The API provides access to this attribute so that client applications can determine whether a given field is an autonumber field.
	AutoNumber *bool `json:"autoNumber,omitempty"`

	// ByteLength - For variable-length fields (including binary fields), the maximum size of the field, in bytes.
	ByteLength *int `json:"byteLength,omitempty"`

	// Calculated indicates whether the field is a custom formula field (true) or not (false).
	// Note that custom formula fields are always read-only.
	Calculated *bool `json:"calculated,omitempty"`

	// CalculatedFormula
	CalculatedFormula *string `json:"calculatedFormula,omitempty"`

	// CascadeDelete
	CascadeDelete *bool `json:"cascadeDelete,omitempty"`

	// CaseSensitive indicates whether the field is case sensitive (true) or not (false).
	CaseSensitive *bool `json:"caseSensitive,omitempty"`

	// ControllerName 	The name of the field that controls the values of this picklist.
	// It only applies if type is picklist or multipicklist and dependentPicklist is true. See About Dependent Picklists.
	// The mapping of controlling field to dependent field is stored in the validFor attribute of each SObjectFieldPicklistEntryMetadata
	// for this picklist. See validFor.
	ControllerName *string `json:"controllerName,omitempty"`

	// Createable indicates whether the field can be created (true) or not (false).
	// If true, then this field value can be set in a create() call.
	Createable *bool `json:"createable,omitempty"`

	// Custom indicates whether the field is a custom field (true) or not (false).
	Custom *bool `json:"custom,omitempty"`

	// DefaultValue
	DefaultValue interface{} `json:"defaultValue,omitempty"`

	// DefaultValueFormula is the default value specified for this field if the formula is not used.
	// If no value has been specified, this field is not returned.
	DefaultValueFormula *string `json:"defaultValueFormula,omitempty"`

	// DefaultedOnCreate Indicates whether this field is defaulted when created (true) or not (false).
	// If true, then Salesforce implicitly assigns a value for this field when the object is created,
	// even if a value for this field is not passed in on the create() call. For example, in the Opportunity object,
	// the Probability field has this attribute because its value is derived from the Stage field.
	// Similarly, the Owner has this attribute on most objects because its value is derived from the current user
	// (if the Owner field is not specified).
	DefaultedOnCreate *bool `json:"defaultedOnCreate,omitempty"`

	// DependentPicklist indicates whether a picklist is a dependent picklist (true) where available values
	// depend on the chosen values from a controlling field, or not (false). See About Dependent Picklists.
	DependentPicklist *bool `json:"dependentPicklist,omitempty"`

	// DeprecatedAndHidden is reserved for future use.
	DeprecatedAndHidden *bool `json:"deprecatedAndHidden,omitempty"`

	// Digits - For fields of type integer. Maximum number of digits. The API returns an error if an integer value
	// exceeds the number of digits.
	Digits *int `json:"digits,omitempty"`

	// DisplayLocationInDecimal indicates how the geolocation values of a Location custom field appears in the user interface.
	// If true, the geolocation values appear in decimal notation. If false, the geolocation values appear as degrees,
	// minutes, and seconds.
	DisplayLocationInDecimal *bool `json:"displayLocationInDecimal,omitempty"`

	// Encrypted indicates whether this field is encrypted. This value only appears in the results of a describeSObjects()
	// call when it is true; otherwise, it is omitted from the results. This field is available in API version 31.0 and later.
	Encrypted *bool `json:"encrypted,omitempty"`

	// ExternalID
	ExternalID *bool `json:"externalId,omitempty"`

	// ExtraTypeInfo - If the field is a textarea field type, indicates if the text area is plain text (plaintextarea)
	// or rich text (richtextarea). If the field is a url field type, if this value is imageurl, the URL references
	// an image file. Available on standard fields on standard objects only,
	// for example, Account.photoUrl, Contact.photoUrl, and so on.
	ExtraTypeInfo *string `json:"extraTypeInfo,omitempty"`

	// Filterable indicates whether the field is filterable (true) or not (false).
	// If true, then this field can be specified in the WHERE clause of a query string in a query() call.
	Filterable *bool `json:"filterable,omitempty"`

	// FilteredLookupInfo 	If the field is a reference field type with a lookup filter, filteredLookupInfo contains
	// the lookup filter information for the field. If there is no lookup filter, or the filter is inactive, this field is null.
	// This field is available in API version 31.0 and later.
	FilteredLookupInfo *SObjectFieldFilteredLookupInfoMetadata `json:"filteredLookupInfo,omitempty"`

	// Groupable indicates whether the field can be included in the GROUP BY clause of a SOQL query (true) or not (false).
	// See GROUP BY in the Salesforce SOQL and SOSL Reference Guide. Available in API version 18.0 and later.
	Groupable *bool `json:"groupable,omitempty"`

	// HighScaleNumber 	Indicates whether the field stores numbers to 8 decimal places regardless of what’s specified
	// in the field details (true) or not (false). Used to handle currencies for products that cost fractions of a cent,
	// in large quantities. If high-scale unit pricing isn’t enabled in your organization, this field isn’t returned.
	// Available in API version 33.0 and later.
	HighScaleNumber *bool `json:"highScaleNumber,omitempty"`

	// HTMLFormatted indicates whether a field such as a hyperlink custom formula field has been formatted for HTML
	// and should be encoded for display in HTML (true) or not (false). Also indicates whether a field is a custom
	// formula field that has an IMAGE text function.
	HTMLFormatted *bool `json:"htmlFormatted,omitempty"`

	// IDLookup indicates whether the field can be used to specify a record in an upsert() call (true) or not (false).
	IDLookup *bool `json:"idLookup,omitempty"`

	// InlineHelpText is the text that displays in the field-level help hover text for this field.
	InlineHelpText *string `json:"inlineHelpText,omitempty"`

	// Label - Text label that is displayed next to the field in the Salesforce user interface. This label can be localized.
	Label *string `json:"label,omitempty"`

	// Length returns the maximum size of the field in Unicode characters (not bytes) or 255, whichever is less.
	// The maximum value returned by the getLength() property is 255. Available in API version 49.0 and later.
	Length *int `json:"length,omitempty"`

	// Mask
	Mask *string `json:"mask,omitempty"`

	// MaskType
	MaskType *string `json:"maskType,omitempty"`

	// Name - Field name used in API calls, such as create(), delete(), and query().
	Name *string `json:"name,omitempty"`

	// NameField indicates whether this field is a name field (true) or not (false). Used to identify the name field
	// for standard objects (such as AccountName for an Account object) and custom objects. Limited to one per object,
	// except where FirstName and LastName fields are used (such as in the Contact object).
	//
	//If a compound name is present, for example the Name field on a person account, nameField is set to true for that record.
	// If no compound name is present, FirstName and LastName have this field set to true.
	NameField *bool `json:"nameField,omitempty"`

	// NamePointing indicates whether the field's value is the Name of the parent of this object (true) or not (false).
	// Used for objects whose parents may be more than one type of object, for example a task may have an account
	// or a contact as a parent.
	NamePointing *bool `json:"namePointing,omitempty"`

	// Nillable indicates whether the field is nillable (true) or not (false). A nillable field can have empty content.
	// A non-nillable field must have a value in order for the object to be created or saved.
	Nillable *bool `json:"nillable,omitempty"`

	// Permissionable indicates whether FieldPermissions can be specified for the field (true) or not (false).
	Permissionable *bool `json:"permissionable,omitempty"`

	// PicklistValues provides the list of valid values for the picklist. Specified only if restrictedPicklist is true.
	PicklistValues []SObjectFieldPicklistEntryMetadata `json:"picklistValues,omitempty"`

	// Precision is for fields of type double. Maximum number of digits that can be stored,
	// including all numbers to the left and to the right of the decimal point (but excluding the decimal point character).
	Precision *int `json:"precision,omitempty"`

	// QueryByDistance
	QueryByDistance *bool `json:"queryByDistance,omitempty"`

	// ReferenceTargetField - Applies only to indirect lookup relationships on external objects. Name of the custom
	// field on the parent standard or custom object whose values are matched against the values of the child external
	// object's indirect lookup relationship field. This matching is done to determine which records are related to each other.
	// This field is available in API version 32.0 and later.
	ReferenceTargetField *string `json:"referenceTargetField,omitempty"`

	// ReferenceTo - For fields that refer to other objects, this array indicates the object types of the referenced objects.
	ReferenceTo []string `json:"referenceTo,omitempty"`

	// RelationshipName - The name of the relationship, if this is a master-detail relationship field.
	RelationshipName *string `json:"relationshipName,omitempty"`

	// RelationshipOrder - The type of relationship for a master-detail relationship field. Valid values are:
	//   - 0 if the field is the primary relationship
	//
	//   - 1 if the field is the secondary relationship
	RelationshipOrder int `json:"relationshipOrder,omitempty"`

	// RestrictedDelete
	RestrictedDelete *bool `json:"restrictedDelete,omitempty"`

	// RestrictedPicklist indicates whether the field is a restricted picklist (true) or not (false).
	RestrictedPicklist *bool `json:"restrictedPicklist,omitempty"`

	// Scale is for fields of type double. Number of digits to the right of the decimal point.
	// The API silently truncates any extra digits to the right of the decimal point, but it returns a fault response
	// if the number has too many digits to the left of the decimal point.
	Scale *int `json:"scale,omitempty"`

	// SoapType
	SoapType *string `json:"soapType,omitempty"`

	// Sortable indicates whether a query can sort on this field (true) or not (false).
	Sortable *bool `json:"sortable,omitempty"`

	// Type
	Type *string `json:"type,omitempty"`

	// Unique indicates whether the value must be unique true) or not false).
	Unique *bool `json:"unique,omitempty"`

	// Updateable indicates one of the following:
	// - Whether the field is updateable, (true) or not (false).  If true, then this field value can be set in an update() call.
	//
	// - If the field is in a master-detail relationship on a custom object, indicates whether the child records can be
	//reparented to different parent records (true), false otherwise.
	Updateable *bool `json:"updateable,omitempty"`

	// WriteRequiresMasterRead 	This field only applies to master-detail relationships. Indicates whether a user
	// requires read sharing access (true) or write sharing access (false) to the parent record to insert,
	// update, and delete a child record. In both cases, a user also needs Create, Edit, and Delete object permissions
	// for the child object.
	WriteRequiresMasterRead *bool `json:"writeRequiresMasterRead,omitempty"`
}

// SObjectFieldPicklistEntryMetadata represents information regarding the picklist entry.
//
// Reference: https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm
type SObjectFieldPicklistEntryMetadata struct {
	// Active indicates whether this item must be displayed (true) or not (false) in the drop-down list for the picklist field in the user interface.
	Active *bool `json:"active,omitempty"`

	// DefaultValue indicates whether this item is the default item (true) in the picklist or not (false).
	// Only one item in a picklist can be designated as the default.
	DefaultValue *bool `json:"defaultValue,omitempty"`

	// Label is the display name of this item in the picklist.
	Label *string `json:"label,omitempty"`

	// ValidFor is a set of bits where each bit indicates a controlling value for which this PicklistEntry is valid.
	// This is base64Encoded.
	ValidFor *string `json:"validFor,omitempty"`

	// Value of this item in the picklist.
	Value *string `json:"value,omitempty"`
}

// SObjectNamedLayoutInfoMetadata represents the name of the named layout for the object. Standard objects can have
// defined named layouts which are separate from the primary layout for both the profile and the record type.
//
// Reference: https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm
type SObjectNamedLayoutInfoMetadata struct {
	// Name of this layout.
	Name *string `json:"name,omitempty"`
}

// SObjectRecordTypeInfoMetadata represents the base class for the old RecordTypeMapping object.
// This object contains all of the existing fields of RecordTypeMapping except layoutId and picklistForRecordType
//
// Reference: https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm
type SObjectRecordTypeInfoMetadata struct {
	// Available indicates whether this record type is available (true) or not (false).
	// Availability is used to display a list of available record types to the user when they are creating a new record.
	Available *bool `json:"available,omitempty"`

	// DefaultRecordTypeMapping indicates whether this is the default record type mapping (true) or not (false).
	DefaultRecordTypeMapping *bool `json:"defaultRecordTypeMapping,omitempty"`

	// DeveloperName is name of this record type. Available in API versions 43.0 and later.
	DeveloperName *string `json:"developerName,omitempty"`

	// Name of this record type.
	Name *string `json:"name,omitempty"`

	// Indicates whether this is the master record type (true) or not (false).
	// The master record type is the default record type that’s used when a record has no custom record type associated with it.
	Master *bool `json:"master,omitempty"`

	// RecordTypeId is the ID of this record type.
	RecordTypeId *string `json:"recordTypeId,omitempty"`
}

// SObjectScopeInfoInfoMetadata represents the scope for an object that can be used to filter object records.
// For example, Account may have a supported ScopeInfo of “mine” (with a UI label of “My accounts”)
// which filters only Account records for the current user.
//
// Reference: https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm
type SObjectScopeInfoInfoMetadata struct {
	// Name of this scope.
	Name *string `json:"name,omitempty"`

	// Label is the UI label for this scope.
	Label *string `json:"label,omitempty"`
}

// SObjectCreateResult represents the response returned from creating an new sobject.
type SObjectCreateResult struct {
	ID      string `json:"id"`
	Success bool   `json:"success"`
}

// SObjectFieldFilteredLookupInfoMetadata contains information about the lookup filter associated with the field.
//
// Reference: https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/sforce_api_calls_describesobjects_describesobjectresult.htm
type SObjectFieldFilteredLookupInfoMetadata struct {
	// ControllingFields is an array of the field’s controlling fields when the lookup filter is dependent on the source object.
	ControllingFields []string `json:"controllingFields,omitempty"`

	// Depedent indicates whether the lookup filter is dependent upon the source object (true) or not (false).
	Dependent bool `json:"dependent,omitempty"`

	// OptionalFilter indicates whether the lookup filter is optional (true) or not (false).
	OptionalFilter bool `json:"optionalFilter,omitempty"`
}

// GetFieldNames returns a string array of field names of a SObject.
func (s *SObjectMetadata) GetFieldNames() []string {
	names := make([]string, 0)
	for _, field := range s.Fields {
		names = append(names, field.GetName())
	}
	return names
}

// GetFieldNamesString returns a single string of all field names of a SObject separated by a comma.
//
// The returned string value is often used as part of a SOQL query.
func (s *SObjectMetadata) GetFieldNamesString() string {
	return strings.Join(s.GetFieldNames(), ",")
}
