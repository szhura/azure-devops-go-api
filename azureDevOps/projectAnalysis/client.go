﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package projectAnalysis

import (
    "context"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevOps"
    "net/http"
    "net/url"
    "strconv"
    "time"
)

var ResourceAreaId, _ = uuid.Parse("7658fa33-b1bf-4580-990f-fac5896773d3")

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

// [Preview API]
func (client *Client) GetProjectLanguageAnalytics(ctx context.Context, args GetProjectLanguageAnalyticsArgs) (*ProjectLanguageAnalytics, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    locationId, _ := uuid.Parse("5b02a779-1867-433f-90b7-d23ed5e33e57")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProjectLanguageAnalytics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetProjectLanguageAnalytics function
type GetProjectLanguageAnalyticsArgs struct {
    // (required) Project ID or project name
    Project *string
}

// [Preview API]
func (client *Client) GetProjectActivityMetrics(ctx context.Context, args GetProjectActivityMetricsArgs) (*ProjectActivityMetrics, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.FromDate == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "fromDate"}
    }
    queryParams.Add("fromDate", (*args.FromDate).String())
    if args.AggregationType == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "aggregationType"}
    }
    queryParams.Add("aggregationType", string(*args.AggregationType))
    locationId, _ := uuid.Parse("e40ae584-9ea6-4f06-a7c7-6284651b466b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProjectActivityMetrics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetProjectActivityMetrics function
type GetProjectActivityMetricsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    FromDate *time.Time
    // (required)
    AggregationType *AggregationType
}

// [Preview API] Retrieves git activity metrics for repositories matching a specified criteria.
func (client *Client) GetGitRepositoriesActivityMetrics(ctx context.Context, args GetGitRepositoriesActivityMetricsArgs) (*[]RepositoryActivityMetrics, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.FromDate == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "fromDate"}
    }
    queryParams.Add("fromDate", (*args.FromDate).String())
    if args.AggregationType == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "aggregationType"}
    }
    queryParams.Add("aggregationType", string(*args.AggregationType))
    if args.Skip == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "skip"}
    }
    queryParams.Add("$skip", strconv.Itoa(*args.Skip))
    if args.Top == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "top"}
    }
    queryParams.Add("$top", strconv.Itoa(*args.Top))
    locationId, _ := uuid.Parse("df7fbbca-630a-40e3-8aa3-7a3faf66947e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []RepositoryActivityMetrics
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetGitRepositoriesActivityMetrics function
type GetGitRepositoriesActivityMetricsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Date from which, the trends are to be fetched.
    FromDate *time.Time
    // (required) Bucket size on which, trends are to be aggregated.
    AggregationType *AggregationType
    // (required) The number of repositories to ignore.
    Skip *int
    // (required) The number of repositories for which activity metrics are to be retrieved.
    Top *int
}

// [Preview API]
func (client *Client) GetRepositoryActivityMetrics(ctx context.Context, args GetRepositoryActivityMetricsArgs) (*RepositoryActivityMetrics, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RepositoryId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = (*args.RepositoryId).String()

    queryParams := url.Values{}
    if args.FromDate == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "fromDate"}
    }
    queryParams.Add("fromDate", (*args.FromDate).String())
    if args.AggregationType == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "aggregationType"}
    }
    queryParams.Add("aggregationType", string(*args.AggregationType))
    locationId, _ := uuid.Parse("df7fbbca-630a-40e3-8aa3-7a3faf66947e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue RepositoryActivityMetrics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetRepositoryActivityMetrics function
type GetRepositoryActivityMetricsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RepositoryId *uuid.UUID
    // (required)
    FromDate *time.Time
    // (required)
    AggregationType *AggregationType
}

