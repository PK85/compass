package tenant_test

import (
	"database/sql"
	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kyma-incubator/compass/components/director/pkg/tenant"
	"github.com/pkg/errors"
)

const (
	testID                           = "foo"
	createQuery                      = "INSERT INTO public.business_tenant_mappings ( id, external_name, external_tenant, parent, type, provider_name, status ) VALUES ( ?, ?, ?, ?, ?, ?, ? )"
	getByExternalIDQueryFormat       = "SELECT id, external_name, external_tenant, parent, type, provider_name, status FROM public.business_tenant_mappings WHERE external_tenant='%s'"
	updateQueryFormat                = "UPDATE public.business_tenant_mappings SET external_name = ?, parent = ?, type = ?, provider_name = ?, status = ? WHERE external_tenant='%s'"
	deleteQuery                      = "DELETE FROM public.business_tenant_mappings WHERE external_tenant = $1"
	testProviderName                 = "test-provider"
	autogeneratedProviderName        = "autogenerated"
	tenantProviderTenantIdProperty   = "tenantId"
	tenantProviderCustomerIdProperty = "customerId"
)

var (
	testError        = errors.New("test error")
	testTableColumns = []string{"id", "external_name", "external_tenant", "parent", "type", "provider_name", "status"}
	createQueryArgs  = fixTenantMappingCreateArgs(tenant.Entity{
		ID:             testID,
		Name:           testID,
		ExternalTenant: testID,
		Parent: sql.NullString{
			String: testID,
			Valid:  true,
		},
		Type:         tenant.Account,
		ProviderName: testProviderName,
		Status:       tenant.Active,
	})
)

type sqlRow struct {
	id             string
	name           string
	externalTenant string
	parent         string
	tenantType     tenant.Type
	provider       string
	status         tenant.Status
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func fixSQLRows(rows []sqlRow) *sqlmock.Rows {
	out := sqlmock.NewRows(testTableColumns)
	for _, row := range rows {
		out.AddRow(row.id, row.name, row.externalTenant, row.parent, row.tenantType, row.provider, row.status)
	}
	return out
}

func fixTenantMappingCreateArgs(ent tenant.Entity) []driver.Value {
	return []driver.Value{ent.ID, ent.Name, ent.ExternalTenant, ent.Parent, ent.Type, ent.ProviderName, ent.Status}
}
