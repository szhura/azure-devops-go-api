﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package universal

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevOps"
    "net/http"
    "net/url"
    "strconv"
)

var ResourceAreaId, _ = uuid.Parse("d397749b-f115-4027-b6dd-77a65dd10d21")

type Client struct {
    Client azureDevOps.Client
}

func NewClient(ctx context.Context, connection *azureDevOps.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API] Delete a package version from the recycle bin.
func (client *Client) DeletePackageVersionFromRecycleBin(ctx context.Context, args DeletePackageVersionFromRecycleBinArgs) error {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *args.PackageVersion

    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeletePackageVersionFromRecycleBin function
type DeletePackageVersionFromRecycleBinArgs struct {
    // (required) Name or ID of the feed.
    FeedId *string
    // (required) Name of the package.
    PackageName *string
    // (required) Version of the package.
    PackageVersion *string
}

// [Preview API] Get information about a package version in the recycle bin.
func (client *Client) GetPackageVersionMetadataFromRecycleBin(ctx context.Context, args GetPackageVersionMetadataFromRecycleBinArgs) (*UPackPackageVersionDeletionState, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *args.PackageVersion

    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UPackPackageVersionDeletionState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetPackageVersionMetadataFromRecycleBin function
type GetPackageVersionMetadataFromRecycleBinArgs struct {
    // (required) Name or ID of the feed.
    FeedId *string
    // (required) Name of the package.
    PackageName *string
    // (required) Version of the package.
    PackageVersion *string
}

// [Preview API] Restore a package version from the recycle bin to its associated feed.
func (client *Client) RestorePackageVersionFromRecycleBin(ctx context.Context, args RestorePackageVersionFromRecycleBinArgs) error {
    if args.PackageVersionDetails == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *args.PackageVersion

    body, marshalErr := json.Marshal(*args.PackageVersionDetails)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the RestorePackageVersionFromRecycleBin function
type RestorePackageVersionFromRecycleBinArgs struct {
    // (required) Set the 'Deleted' property to 'false' to restore the package.
    PackageVersionDetails *UPackRecycleBinPackageVersionDetails
    // (required) Name or ID of the feed.
    FeedId *string
    // (required) Name of the package.
    PackageName *string
    // (required) Version of the package.
    PackageVersion *string
}

// [Preview API] Delete a package version from a feed's recycle bin.
func (client *Client) DeletePackageVersion(ctx context.Context, args DeletePackageVersionArgs) (*Package, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *args.PackageVersion

    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the DeletePackageVersion function
type DeletePackageVersionArgs struct {
    // (required) Name or ID of the feed.
    FeedId *string
    // (required) Name of the package.
    PackageName *string
    // (required) Version of the package.
    PackageVersion *string
}

// [Preview API] Show information about a package version.
func (client *Client) GetPackageVersion(ctx context.Context, args GetPackageVersionArgs) (*Package, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *args.PackageVersion

    queryParams := url.Values{}
    if args.ShowDeleted != nil {
        queryParams.Add("showDeleted", strconv.FormatBool(*args.ShowDeleted))
    }
    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetPackageVersion function
type GetPackageVersionArgs struct {
    // (required) Name or ID of the feed.
    FeedId *string
    // (required) Name of the package.
    PackageName *string
    // (required) Version of the package.
    PackageVersion *string
    // (optional) True to show information for deleted versions
    ShowDeleted *bool
}

// [Preview API] Update information for a package version.
func (client *Client) UpdatePackageVersion(ctx context.Context, args UpdatePackageVersionArgs) error {
    if args.PackageVersionDetails == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *args.PackageVersion

    body, marshalErr := json.Marshal(*args.PackageVersionDetails)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the UpdatePackageVersion function
type UpdatePackageVersionArgs struct {
    // (required)
    PackageVersionDetails *PackageVersionDetails
    // (required) Name or ID of the feed.
    FeedId *string
    // (required) Name of the package.
    PackageName *string
    // (required) Version of the package.
    PackageVersion *string
}

