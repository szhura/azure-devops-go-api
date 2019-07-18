﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package customerIntelligence

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
)

type Client struct {
    Client azureDevops.Client
}

func NewClient(connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// [Preview API]
// events (required)
func (client Client) PublishEvents(events *[]CustomerIntelligenceEvent) error {
    if events == nil {
        return errors.New("events is a required parameter")
    }
    body, marshalErr := json.Marshal(*events)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("b5cc35c2-ff2b-491d-a085-24b6e9f396fd")
    _, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}
