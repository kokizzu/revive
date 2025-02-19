package fixtures

type RangeAllocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Range             string `json:"range,outline"`
	Data              []byte `json:"data,flow"` // MATCH /unknown option 'flow' in JSON tag/
}

type RangeAllocation struct {
	metav1.TypeMeta   `bson:",minsize,gnu"`
	metav1.ObjectMeta `bson:"metadata,omitempty"`
	Range             string `bson:"range,flow"` // MATCH /unknown option 'flow' in BSON tag/
	Data              []byte `bson:"data,inline"`
}

type RequestQueryOptions struct {
	Properties       []string `url:"properties,commmma,omitempty"` // MATCH /unknown option 'commmma' in URL tag/
	CustomProperties []string `url:"-"`
	Archived         bool     `url:"archived,myURLOption"`
}

type Fields struct {
	Field      string `datastore:",noindex,flatten,omitempty,myDatastoreOption"`
	OtherField string `datastore:",unknownOption"` // MATCH /unknown option 'unknownOption' in Datastore tag/
}

type MapStruct struct {
	Field1     string `mapstructure:",squash,reminder,omitempty,myMapstructureOption"`
	OtherField string `mapstructure:",unknownOption"` // MATCH /unknown option 'unknownOption' in Mapstructure tag/
}
