package force

// FieldDataType represents a SObject field type.
type FieldDataType string

// FieldDataTypes represents all data types for a SObject field.
//
// Reference: https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/field_types.htm
var FieldDataTypes = struct {
	Address                    FieldDataType
	AnyType                    FieldDataType
	Calculated                 FieldDataType
	Combobox                   FieldDataType
	Currency                   FieldDataType
	DataCategoryGroupReference FieldDataType
	Email                      FieldDataType
	Encryptedstring            FieldDataType
	ID                         FieldDataType
	JunctionIdList             FieldDataType
	Location                   FieldDataType
	Masterrecord               FieldDataType
	Multipicklist              FieldDataType
	Percent                    FieldDataType
	Phone                      FieldDataType
	Picklist                   FieldDataType
	Reference                  FieldDataType
	Textarea                   FieldDataType
	URL                        FieldDataType
}{
	Address:                    "address",
	AnyType:                    "anyType",
	Calculated:                 "calculated",
	Combobox:                   "combobox",
	Currency:                   "currency",
	DataCategoryGroupReference: "DataCategoryGroupReference",
	Email:                      "email",
	Encryptedstring:            "encryptedstring",
	ID:                         "ID",
	JunctionIdList:             "JunctionIdList",
	Location:                   "location",
	Masterrecord:               "masterrecord",
	Multipicklist:              "multipicklist",
	Percent:                    "percent",
	Phone:                      "phone",
	Picklist:                   "picklist",
	Reference:                  "reference",
	Textarea:                   "textarea",
	URL:                        "url",
}

// FieldDataTypeGoDataTypeMapping represents a map of each FieldDataType's Golang data type equivalent.
//
// Values should be the string form of valid Golang data types. Reference: https://tour.golang.org/basics/11
var FieldDataTypeGoDataTypeMapping = map[FieldDataType]string{
	FieldDataTypes.Address:                    "string",
	FieldDataTypes.AnyType:                    "interface{}",
	FieldDataTypes.Calculated:                 "string",
	FieldDataTypes.Combobox:                   "string",
	FieldDataTypes.Currency:                   "json.Number",
	FieldDataTypes.DataCategoryGroupReference: "string",
	FieldDataTypes.Email:                      "string",
	FieldDataTypes.Encryptedstring:            "string",
	FieldDataTypes.ID:                         "string",
	FieldDataTypes.JunctionIdList:             "[]string",
	FieldDataTypes.Location:                   "string",
	FieldDataTypes.Masterrecord:               "string",
	FieldDataTypes.Multipicklist:              "string",
	FieldDataTypes.Percent:                    "json.Number",
	FieldDataTypes.Phone:                      "string",
	FieldDataTypes.Picklist:                   "string",
	FieldDataTypes.Reference:                  "string",
	FieldDataTypes.Textarea:                   "string",
	FieldDataTypes.URL:                        "string",
}

// ToString is a helper method to return the string of a FieldDataType.
func (f FieldDataType) ToString() string {
	return string(f)
}

// ConvertToGoDataType returns the Golang data type equivalent for a FieldDataType.
func (f FieldDataType) ConvertToGoDataType() string {
	if v, ok := FieldDataTypeGoDataTypeMapping[f]; ok {
		return v
	}
	return "string"
}
