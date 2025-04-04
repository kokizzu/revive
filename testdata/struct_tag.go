package fixtures

import "time"

type decodeAndValidateRequest struct {
	// BEAWRE : the flag of URLParam should match the const string URLParam
	URLParam    string `json:"-" path:"url_param" validate:"numeric"`
	Text        string `json:"text" validate:"max=10"`
	DefaultInt  int    `json:"defaultInt" default:"10.0"` // MATCH /field's type and default value's type mismatch/
	DefaultInt2 int    `json:"defaultInt2" default:"10"`
	// MATCH:12 /unknown option 'inline' in JSON tag/
	DefaultInt3      int             `json:"defaultInt2,inline" default:"11"` // MATCH /duplicate tag name: 'defaultInt2'/
	DefaultString    string          `json:"defaultString" default:"foo"`
	DefaultBool      bool            `json:"defaultBool" default:"trues"` // MATCH /field's type and default value's type mismatch/
	DefaultBool2     bool            `json:"defaultBool2" default:"true"`
	DefaultBool3     bool            `json:"defaultBool3" default:"false"`
	DefaultFloat     float64         `json:"defaultFloat" default:"f10.0"` // MATCH /field's type and default value's type mismatch/
	DefaultFloat2    float64         `json:"defaultFloat2" default:"10.0"`
	MandatoryStruct  mandatoryStruct `json:"mandatoryStruct" required:"trues"` // MATCH /required should be 'true' or 'false'/
	MandatoryStruct2 mandatoryStruct `json:"mandatoryStruct2" required:"true"`
	MandatoryStruct4 mandatoryStruct `json:"mandatoryStruct4" required:"false"`
	OptionalStruct   *optionalStruct `json:"optionalStruct,omitempty"`
	OptionalQuery    string          `json:"-" querystring:"queryfoo"`
	optionalQuery    string          `json:"-" querystring:"queryfoo"` // MATCH /tag on not-exported field optionalQuery/
	// No-reg test for bug https://github.com/mgechev/revive/issues/208
	Tiret       string `json:"-,"`
	BadTiret    string `json:"other,"`               // MATCH /option can not be empty in JSON tag/
	ForOmitzero string `json:"forOmitZero,omitzero"` // MATCH /unknown option 'omitzero' in JSON tag/
}

type RangeAllocation struct {
	metav1.TypeMeta   `json:",inline"` // MATCH /unknown option 'inline' in JSON tag/
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Range             string `json:"range,flow"`  // MATCH /unknown option 'flow' in JSON tag/
	Data              []byte `json:"data,inline"` // MATCH /unknown option 'inline' in JSON tag/
}

type RangeAllocation struct {
	metav1.TypeMeta   `bson:",minsize"`
	metav1.ObjectMeta `bson:"metadata,omitempty"`
	Range             string `bson:"range,flow"` // MATCH /unknown option 'flow' in BSON tag/
	Data              []byte `bson:"data,inline"`
}

type TestContextSpecificTags2 struct {
	A       int       `asn1:"explicit,tag:1"`
	B       int       `asn1:"tag:2"`
	S       string    `asn1:"tag:0,utf8"`
	Ints    []int     `asn1:"set"`
	Version int       `asn1:"optional,explicit,default:0,tag:000"` // MATCH /duplicated tag number 0/
	Time    time.Time `asn1:"explicit,tag:4,other"`                // MATCH /unknown option 'other' in ASN1 tag/
	X       int       `asn1:"explicit,tag:invalid"`                // MATCH /ASN1 tag must be a number, got 'invalid'/
}

type VirtualMachineRelocateSpecDiskLocator struct {
	DynamicData

	DiskId          int32                           `xml:"diskId,attr,cdata"`
	Datastore       ManagedObjectReference          `xml:"datastore,chardata,innerxml"`
	DiskMoveType    string                          `xml:"diskMoveType,omitempty,comment"`
	DiskBackingInfo BaseVirtualDeviceBackingInfo    `xml:"diskBackingInfo,omitempty,any"`
	Profile         []BaseVirtualMachineProfileSpec `xml:"profile,omitempty,other"` // MATCH /unknown option 'other' in XML tag/
}

type TestDuplicatedXMLTags struct {
	A int `xml:"a"`
	B int `xml:"a"` // MATCH /duplicate tag name: 'a'/
	C int `xml:"c"`
}

type TestDuplicatedBSONTags struct {
	A int `bson:"b"`
	B int `bson:"b"` // MATCH /duplicate tag name: 'b'/
	C int `bson:"c"`
}

type TestDuplicatedYAMLTags struct {
	A int `yaml:"b"`
	B int `yaml:"c"`
	C int `yaml:"c"` // MATCH /duplicate tag name: 'c'/
}

type TestDuplicatedProtobufTags struct {
	A int `protobuf:"varint,name=b"`
	B int `protobuf:"varint,name=c"`
	C int `protobuf:"varint,name=c"` // MATCH /duplicate tag name: 'c'/
}

// test case from
// sigs.k8s.io/kustomize/api/types/helmchartargs.go

type HelmChartArgs struct {
	ChartName        string                 `json:"chartName,omitempty" yaml:"chartName,omitempty"`
	ChartVersion     string                 `json:"chartVersion,omitempty" yaml:"chartVersion,omitempty"`
	ChartRepoURL     string                 `json:"chartRepoUrl,omitempty" yaml:"chartRepoUrl,omitempty"`
	ChartHome        string                 `json:"chartHome,omitempty" yaml:"chartHome,omitempty"`
	ChartRepoName    string                 `json:"chartRepoName,omitempty" yaml:"chartRepoName,omitempty"`
	HelmBin          string                 `json:"helmBin,omitempty" yaml:"helmBin,omitempty"`
	HelmHome         string                 `json:"helmHome,omitempty" yaml:"helmHome,omitempty"`
	Values           string                 `json:"values,omitempty" yaml:"values,omitempty"`
	ValuesLocal      map[string]interface{} `json:"valuesLocal,omitempty" yaml:"valuesLocal,omitempty"`
	ValuesMerge      string                 `json:"valuesMerge,omitempty" yaml:"valuesMerge,omitempty"`
	ReleaseName      string                 `json:"releaseName,omitempty" yaml:"releaseName,omitempty"`
	ReleaseNamespace string                 `json:"releaseNamespace,omitempty" yaml:"releaseNamespace,omitempty"`
	ExtraArgs        []string               `json:"extraArgs,omitempty" yaml:"extraArgs,omitempty"`
}

// Test message for holding primitive types.
type Simple struct {
	OBool                *bool    `protobuf:"varint,1,req,json=oBool"`                           // MATCH /protobuf tag lacks mandatory option 'name'/
	OInt32               *int32   `protobuf:"varint,2,opt,name=o_int32,jsonx=oInt32"`            // MATCH /unknown option 'jsonx' in protobuf tag/
	OInt32Str            *int32   `protobuf:"varint,3,rep,name=o_int32_str,name=oInt32Str"`      // MATCH /protobuf tag has duplicated option 'name'/
	OInt64               *int64   `protobuf:"varint,4,opt,json=oInt64,name=o_int64,json=oInt64"` // MATCH /protobuf tag has duplicated option 'json'/
	OSint32Str           *int32   `protobuf:"zigzag32,11,opt,name=o_sint32_str,json=oSint32Str"`
	OSint64Str           *int64   `protobuf:"zigzag64,13,opt,name=o_sint32_str,json=oSint64Str"` // MATCH /duplicate tag name: 'o_sint32_str'/
	OFloat               *float32 `protobuf:"fixed32,14,opt,name=o_float,json=oFloat"`
	ODouble              *float64 `protobuf:"fixed64,014,opt,name=o_double,json=oDouble"`      // MATCH /duplicated tag number 14/
	ODoubleStr           *float64 `protobuf:"fixed6,17,opt,name=o_double_str,json=oDoubleStr"` // MATCH /invalid protobuf tag name 'fixed6'/
	OString              *string  `protobuf:"bytes,18,opt,name=o_string,json=oString"`
	OString2             *string  `protobuf:"bytes,name=ameno"`
	OString3             *string  `protobuf:"bytes,name=ameno"` // MATCH /duplicate tag name: 'ameno'/
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type RequestQueryOption struct {
	Properties           []string `url:"properties,comma,omitempty"`
	CustomProperties     []string `url:"-"`
	Associations         []string `url:"associations,brackets,omitempty"`
	Associations2        []string `url:"associations2,semicolon,omitempty"`
	Associations3        []string `url:"associations3,space,brackets,omitempty"`
	Associations4        []string `url:"associations4,numbered,omitempty"`
	Associations5        []string `url:"associations5,space,semicolon,omitempty"` // MATCH /can not set both 'semicolon' and 'space' as delimiters in URL tag/
	PaginateAssociations bool     `url:"paginateAssociations,int,omitempty"`
	Archived             bool     `url:"archived,myURLOption"` // MATCH /unknown option 'myURLOption' in URL tag/
	IDProperty           string   `url:"idProperty,omitempty"`
}

type Fields struct {
	Field      string `datastore:",noindex,flatten,omitempty"`
	OtherField string `datastore:",unknownOption"` // MATCH /unknown option 'unknownOption' in Datastore tag/
}

type MapStruct struct {
	Field1     string `mapstructure:",squash,reminder,omitempty"`
	OtherField string `mapstructure:",unknownOption"` // MATCH /unknown option 'unknownOption' in Mapstructure tag/
}

type ValidateUser struct {
	Username    string `validate:"required,min=3,max=32"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required,min=8,max=32"`
	Biography   string `validate:"min=0,max=1000"`
	DisplayName string `validate:"displayName,min=3,max=32"` // MATCH /unknown option 'displayName' in validate tag/
	Complex     string `validate:"gt=0,dive,keys,eq=1|eq=2,endkeys,required"`
	BadComplex  string `validate:"gt=0,keys,eq=1|eq=2,endkeys,required"`              // MATCH /option 'keys' must follow a 'dive' option in validate tag/
	BadComplex2 string `validate:"gt=0,dive,eq=1|eq=2,endkeys,required"`              // MATCH /option 'endkeys' without a previous 'keys' option in validate tag/
	BadComplex3 string `validate:"gt=0,dive,keys,eq=1|eq=2,endkeys,endkeys,required"` // MATCH /option 'endkeys' without a previous 'keys' option in validate tag/
}

type TomlUser struct {
	Username string `toml:"username,omitempty"`
	Location string `toml:"location,unknown"` // MATCH /unknown option 'unknown' in TOML tag/
}

type PropertiesTags struct {
	Field int               `properties:"-"`
	Field int               `properties:"myName"`
	Field int               `properties:"myName,default=15"`
	Field int               `properties:"myName,default=sString"` // MATCH /field type and default value type mismatch in properties tag/
	Field int               `properties:",default:15"`            // MATCH /unknown option "default:15" in properties tag/
	Field int               `properties:",default=15,default=2"`  // MATCH /properties tag accepts only one default option/
	Field time.Time         `properties:"date,layout=2006-01-02"`
	Field time.Time         `properties:",layout=2006-01-02"`
	Field time.Time         `properties:"date,layout"`            // MATCH /malformed layout option for properties tag/
	Field time.Time         `properties:"date,layout=  "`         // MATCH /malformed layout option for properties tag/
	Field string            `properties:"date,layout=2006-01-02"` // MATCH /layout option is only applicable to fields of type time.Time in properties tag/
	Field []string          `properties:",default=a;b;c"`
	Field map[string]string `properties:"myName,omitempty"` // MATCH /unknown option "omitempty" in properties tag/
}
