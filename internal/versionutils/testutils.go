package versionutils

import (
	"context"
	"fmt"

	"github.com/elastic/terraform-provider-elasticstack/internal/clients"
	"github.com/hashicorp/go-version"
)

func CheckIfVersionIsUnsupported(minSupportedVersion *version.Version) func() (bool, error) {
	return func() (b bool, err error) {
		client, err := clients.NewAcceptanceTestingClient()
		if err != nil {
			return false, err
		}
		serverVersion, diags := client.ServerVersion(context.Background())
		if diags.HasError() {
			return false, fmt.Errorf("failed to parse the elasticsearch version %v", diags)
		}

		return serverVersion.LessThan(minSupportedVersion), nil
	}
}

func CheckIfVersionMeetsConstraints(constraints version.Constraints) func() (bool, error) {
	return func() (b bool, err error) {
		client, err := clients.NewAcceptanceTestingClient()
		if err != nil {
			return false, err
		}
		serverVersion, diags := client.ServerVersion(context.Background())
		if diags.HasError() {
			return false, fmt.Errorf("failed to parse the elasticsearch version %v", diags)
		}

		return !constraints.Check(serverVersion), nil
	}
}

func CheckIfNotServerless() func() (bool, error) {
	return func() (b bool, err error) {
		client, err := clients.NewAcceptanceTestingClient()
		if err != nil {
			return false, err
		}
		serverFlavor, diags := client.ServerFlavor(context.Background())
		if diags.HasError() {
			return false, fmt.Errorf("failed to get the elasticsearch flavor %v", diags)
		}

		return serverFlavor != clients.ServerlessFlavor, nil
	}
}
