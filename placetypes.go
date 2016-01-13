package placetypes

import (
	"encoding/json"
	"errors"
)

type WOFPlacetypeName struct {
	Lang string
	Kind string
	Name string
}

type WOFPlacetypeAltNames struct {
	Lang string
	Name string
}

type WOFPlacetype struct {
	Id       int64
	Name     string
	Role     string
	Parent   []string
	AltNames []WOFPlacetypeAltNames
}

func NewPlacetypeName(name string) (*WOFPlacetypeName, error) {

	return nil, errors.New("Please write me...")
}

func NewPlacetype(placetype string) (*WOFPlacetype, error) {

	return nil, errors.New("Please write me...")
}

func IsValidPlacetype(placetype string) bool {
	return false
}

func Common() []string {
	return WithRole("common")
}

func CommonOptional() []string {
	return WithRole("common_optional")
}

func Optional() []string {
	return WithRole("optional")
}

func WithRole(role string) []string {

	places := make([]string, 0)
	return places
}

func WithRoles(roles []string) []string {

	places := make([]string, 0)
	return places
}

func IsValidRole(role string) bool {

	return false
}

func (p WOFPlacetype) Names() []*WOFPlacetypeName {

	names := make([]*WOFPlacetypeName, 0)
	return names
}

func (p WOFPlacetype) Parents() []*WOFPlacetype {

	places := make([]*WOFPlacetype, 0)
	return places
}

func (p WOFPlacetype) Ancestors() []string {

	places := make([]string, 0)
	return places
}

func (p WOFPlacetype) AncestorsWithRoles([]string) []string {

	places := make([]string, 0)
	return places
}

// please to be returning an actual thingy and not just an interface

func Spec() (interface{}, error) {

	// https://github.com/whosonfirst/whosonfirst-placetypes/blob/master/data/placetypes-spec-latest.json
	// specifically placetypes-spec-20151209.json

	// Basically this isn't going to parse in to anything useful in Go so we'll
	// need to create a separate data/spec file in whosonfirst-placetypes
	// (20160112/thisisaaronland)

	spec := []byte(`{"102312321": {"role": "optional", "name": "microhood", "parent": [102312319], "names": {}}, "102312323": {"role": "optional", "name": "macrohood", "parent": [102312317], "names": {}}, "102312325": {"role": "common_optional", "name": "venue", "parent": [102312327, 102312329, 102312331, 102312321, 102312319], "names": {}}, "102312327": {"role": "common_optional", "name": "building", "parent": [102312329, 102312331, 102312321, 102312319], "names": {}}, "102312329": {"role": "common_optional", "name": "address", "parent": [102312331, 102312321, 102312319], "names": {}}, "102312331": {"role": "common_optional", "name": "campus", "parent": [102312321, 102312319, 102312323, 102312317, 404221409], "names": {}}, "404528653": {"role": "common_optional", "name": "ocean", "parent": [102312341], "names": {}}, "102312335": {"role": "common_optional", "name": "empire", "parent": [102312309], "names": {}}, "102312341": {"role": "common_optional", "name": "planet", "parent": [], "names": {}}, "102320821": {"role": "common_optional", "name": "dependency", "parent": [102312307], "names": {}}, "136057795": {"role": "common_optional", "name": "timezone", "parent": [102312307, 102312309, 102312341], "names": {}}, "404528655": {"role": "common_optional", "name": "marinearea", "parent": [102312307, 102312309, 102312341], "names": {}}, "102371933": {"role": "optional", "name": "metroarea", "parent": [], "names": {}}, "404221409": {"role": "common_optional", "name": "localadmin", "parent": [102312313, 102312311], "names": {}}, "404221411": {"role": "optional", "name": "macroregion", "parent": [102320821, 102322043, 102312307], "names": {}}, "404221413": {"role": "optional", "name": "macrocounty", "parent": [102312311], "names": {}}, "102312307": {"role": "common", "name": "country", "parent": [102312335, 102312309], "names": {}}, "102312309": {"role": "common", "name": "continent", "parent": [102312341], "names": {}}, "102312311": {"role": "common", "name": "region", "parent": [404221411, 102320821, 102322043, 102312307], "names": {}}, "102312313": {"role": "common_optional", "name": "county", "parent": [404221413, 102312311], "names": {}}, "102322043": {"role": "common_optional", "name": "disputed", "parent": [102312307], "names": {}}, "102312317": {"role": "common", "name": "locality", "parent": [404221409, 102312313, 102312311], "names": {}}, "102312319": {"role": "common", "name": "neighbourhood", "parent": [102312323, 102312317], "names": {"eng_p": ["neighbourhood", "neighborhood"]}}}`)

	/*
	   type Debug struct {
	   	  Foo interface{}	`json:"102312321"`
	   }

	   var d Debug
	*/

	var d interface{}
	err := json.Unmarshal(spec, &d)

	if err != nil {
		return nil, err
	}

	return d, nil
}
