package resources_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

<<<<<<< HEAD
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
=======
	"github.com/chanzuckerberg/terraform-provider-snowflake/pkg/provider"
	"github.com/chanzuckerberg/terraform-provider-snowflake/pkg/resources"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
>>>>>>> 7dd55fa02ff8b69235d11375c3fb5f2028e5146b
	"github.com/stretchr/testify/assert"
	"github.com/viostream/terraform-provider-snowflake/pkg/provider"
	"github.com/viostream/terraform-provider-snowflake/pkg/resources"
)

func database(t *testing.T, id string, params map[string]interface{}) *schema.ResourceData {
	a := assert.New(t)
	d := schema.TestResourceDataRaw(t, resources.Database().Schema, params)
	a.NotNil(d)
	d.SetId(id)
	return d
}

func databaseGrant(t *testing.T, id string, params map[string]interface{}) *schema.ResourceData {
	a := assert.New(t)
	d := schema.TestResourceDataRaw(t, resources.DatabaseGrant().Schema, params)
	a.NotNil(d)
	d.SetId(id)
	return d
}

func fixture(name string) (string, error) {
	b, err := ioutil.ReadFile(filepath.Join("testdata", name))
	return string(b), err
}

func providers() map[string]terraform.ResourceProvider {
	p := provider.Provider()
	return map[string]terraform.ResourceProvider{
		"snowflake": p,
	}
}

func role(t *testing.T, id string, params map[string]interface{}) *schema.ResourceData {
	a := assert.New(t)
	d := schema.TestResourceDataRaw(t, resources.Role().Schema, params)
	a.NotNil(d)
	d.SetId(id)
	return d
}

func roleGrants(t *testing.T, id string, params map[string]interface{}) *schema.ResourceData {
	a := assert.New(t)
	d := schema.TestResourceDataRaw(t, resources.RoleGrants().Schema, params)
	a.NotNil(d)
	d.SetId(id)
	return d
}

func storageIntegration(t *testing.T, id string, params map[string]interface{}) *schema.ResourceData {
	a := assert.New(t)
	d := schema.TestResourceDataRaw(t, resources.StorageIntegration().Schema, params)
	a.NotNil(d)
	d.SetId(id)
	return d
}

func user(t *testing.T, id string, params map[string]interface{}) *schema.ResourceData {
	a := assert.New(t)
	d := schema.TestResourceDataRaw(t, resources.User().Schema, params)
	a.NotNil(d)
	d.SetId(id)
	return d
}

func warehouse(t *testing.T, id string, params map[string]interface{}) *schema.ResourceData {
	a := assert.New(t)
	d := schema.TestResourceDataRaw(t, resources.Warehouse().Schema, params)
	a.NotNil(d)
	d.SetId(id)
	return d
}
