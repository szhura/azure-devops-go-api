﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package settings

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevOps"
    "net/http"
)

type Client struct {
    Client azureDevOps.Client
}

func NewClient(ctx context.Context, connection *azureDevOps.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// [Preview API] Get all setting entries for the given user/all-users scope
func (client *Client) GetEntries(ctx context.Context, args GetEntriesArgs) (*map[string]interface{}, error) {
    routeValues := make(map[string]string)
    if args.UserScope == nil || *args.UserScope == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *args.UserScope
    if args.Key != nil && *args.Key != "" {
        routeValues["key"] = *args.Key
    }

    locationId, _ := uuid.Parse("cd006711-163d-4cd4-a597-b05bad2556ff")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue map[string]interface{}
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetEntries function
type GetEntriesArgs struct {
    // (required) User-Scope at which to get the value. Should be "me" for the current user or "host" for all users.
    UserScope *string
    // (optional) Optional key under which to filter all the entries
    Key *string
}

// [Preview API] Remove the entry or entries under the specified path
func (client *Client) RemoveEntries(ctx context.Context, args RemoveEntriesArgs) error {
    routeValues := make(map[string]string)
    if args.UserScope == nil || *args.UserScope == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *args.UserScope
    if args.Key == nil || *args.Key == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "key"} 
    }
    routeValues["key"] = *args.Key

    locationId, _ := uuid.Parse("cd006711-163d-4cd4-a597-b05bad2556ff")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the RemoveEntries function
type RemoveEntriesArgs struct {
    // (required) User-Scope at which to remove the value. Should be "me" for the current user or "host" for all users.
    UserScope *string
    // (required) Root key of the entry or entries to remove
    Key *string
}

// [Preview API] Set the specified setting entry values for the given user/all-users scope
func (client *Client) SetEntries(ctx context.Context, args SetEntriesArgs) error {
    if args.Entries == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "entries"}
    }
    routeValues := make(map[string]string)
    if args.UserScope == nil || *args.UserScope == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *args.UserScope

    body, marshalErr := json.Marshal(*args.Entries)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("cd006711-163d-4cd4-a597-b05bad2556ff")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the SetEntries function
type SetEntriesArgs struct {
    // (required) The entries to set
    Entries *map[string]interface{}
    // (required) User-Scope at which to set the values. Should be "me" for the current user or "host" for all users.
    UserScope *string
}

// [Preview API] Get all setting entries for the given named scope
func (client *Client) GetEntriesForScope(ctx context.Context, args GetEntriesForScopeArgs) (*map[string]interface{}, error) {
    routeValues := make(map[string]string)
    if args.UserScope == nil || *args.UserScope == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *args.UserScope
    if args.ScopeName == nil || *args.ScopeName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "scopeName"} 
    }
    routeValues["scopeName"] = *args.ScopeName
    if args.ScopeValue == nil || *args.ScopeValue == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "scopeValue"} 
    }
    routeValues["scopeValue"] = *args.ScopeValue
    if args.Key != nil && *args.Key != "" {
        routeValues["key"] = *args.Key
    }

    locationId, _ := uuid.Parse("4cbaafaf-e8af-4570-98d1-79ee99c56327")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue map[string]interface{}
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetEntriesForScope function
type GetEntriesForScopeArgs struct {
    // (required) User-Scope at which to get the value. Should be "me" for the current user or "host" for all users.
    UserScope *string
    // (required) Scope at which to get the setting for (e.g. "project" or "team")
    ScopeName *string
    // (required) Value of the scope (e.g. the project or team id)
    ScopeValue *string
    // (optional) Optional key under which to filter all the entries
    Key *string
}

// [Preview API] Remove the entry or entries under the specified path
func (client *Client) RemoveEntriesForScope(ctx context.Context, args RemoveEntriesForScopeArgs) error {
    routeValues := make(map[string]string)
    if args.UserScope == nil || *args.UserScope == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *args.UserScope
    if args.ScopeName == nil || *args.ScopeName == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "scopeName"} 
    }
    routeValues["scopeName"] = *args.ScopeName
    if args.ScopeValue == nil || *args.ScopeValue == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "scopeValue"} 
    }
    routeValues["scopeValue"] = *args.ScopeValue
    if args.Key == nil || *args.Key == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "key"} 
    }
    routeValues["key"] = *args.Key

    locationId, _ := uuid.Parse("4cbaafaf-e8af-4570-98d1-79ee99c56327")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the RemoveEntriesForScope function
type RemoveEntriesForScopeArgs struct {
    // (required) User-Scope at which to remove the value. Should be "me" for the current user or "host" for all users.
    UserScope *string
    // (required) Scope at which to get the setting for (e.g. "project" or "team")
    ScopeName *string
    // (required) Value of the scope (e.g. the project or team id)
    ScopeValue *string
    // (required) Root key of the entry or entries to remove
    Key *string
}

// [Preview API] Set the specified entries for the given named scope
func (client *Client) SetEntriesForScope(ctx context.Context, args SetEntriesForScopeArgs) error {
    if args.Entries == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "entries"}
    }
    routeValues := make(map[string]string)
    if args.UserScope == nil || *args.UserScope == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *args.UserScope
    if args.ScopeName == nil || *args.ScopeName == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "scopeName"} 
    }
    routeValues["scopeName"] = *args.ScopeName
    if args.ScopeValue == nil || *args.ScopeValue == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "scopeValue"} 
    }
    routeValues["scopeValue"] = *args.ScopeValue

    body, marshalErr := json.Marshal(*args.Entries)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("4cbaafaf-e8af-4570-98d1-79ee99c56327")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the SetEntriesForScope function
type SetEntriesForScopeArgs struct {
    // (required) The entries to set
    Entries *map[string]interface{}
    // (required) User-Scope at which to set the values. Should be "me" for the current user or "host" for all users.
    UserScope *string
    // (required) Scope at which to set the settings on (e.g. "project" or "team")
    ScopeName *string
    // (required) Value of the scope (e.g. the project or team id)
    ScopeValue *string
}

