package web

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type AppServiceHybridConnectionResourceID struct {
	ResourceGroup string
	Name          string
	AppName       string
	Namespace     string
}

func ParseAppServiceHybridConnectionID(input string) (*AppServiceHybridConnectionResourceID, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Hybrid Connection ID %q: %+v", input, err)
	}

	hybridConnection := AppServiceHybridConnectionResourceID{
		ResourceGroup: id.ResourceGroup,
	}
	if hybridConnection.Name, err = id.PopSegment("relays"); err != nil {
		return nil, err
	}

	if hybridConnection.AppName, err = id.PopSegment("sites"); err != nil {
		return nil, err
	}

	if hybridConnection.Namespace, err = id.PopSegment("hybridConnectionNamespaces"); err != nil {
		return nil, err
	}

	return &hybridConnection, nil
}
