package validate

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import "testing"

func TestConsumptionBudgetManagementGroupID(t *testing.T) {
	cases := []struct {
		Input string
		Valid bool
	}{
		{
			// empty
			Input: "",
			Valid: false,
		},

		{
			// missing ManagementGroupName
			Input: "/providers/Microsoft.Management/",
			Valid: false,
		},

		{
			// missing value for ManagementGroupName
			Input: "/providers/Microsoft.Management/managementGroups/",
			Valid: false,
		},

		{
			// missing BudgetName
			Input: "/providers/Microsoft.Management/managementGroups/12345678-1234-9876-4563-123456789012/providers/Microsoft.Consumption/",
			Valid: false,
		},

		{
			// missing value for BudgetName
			Input: "/providers/Microsoft.Management/managementGroups/12345678-1234-9876-4563-123456789012/providers/Microsoft.Consumption/budgets/",
			Valid: false,
		},

		{
			// valid
			Input: "/providers/Microsoft.Management/managementGroups/12345678-1234-9876-4563-123456789012/providers/Microsoft.Consumption/budgets/budget1",
			Valid: true,
		},

		{
			// upper-cased
			Input: "/PROVIDERS/MICROSOFT.MANAGEMENT/MANAGEMENTGROUPS/12345678-1234-9876-4563-123456789012/PROVIDERS/MICROSOFT.CONSUMPTION/BUDGETS/BUDGET1",
			Valid: false,
		},
	}
	for _, tc := range cases {
		t.Logf("[DEBUG] Testing Value %s", tc.Input)
		_, errors := ConsumptionBudgetManagementGroupID(tc.Input, "test")
		valid := len(errors) == 0

		if tc.Valid != valid {
			t.Fatalf("Expected %t but got %t", tc.Valid, valid)
		}
	}
}
