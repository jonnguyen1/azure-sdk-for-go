//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupsListByDevice.json
func ExampleBackupsClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewBackupsClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByDevicePager("Device05ForSDKTest",
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest1",
		&armstorsimple8000series.BackupsClientListByDeviceOptions{Filter: to.Ptr("createdTime%20ge%20'2017-06-22T18:30:00Z'%20and%20backupPolicyId%20eq%20'%2Fsubscriptions%2F4385cf00-2d3a-425a-832f-f4285b1c9dce%2FresourceGroups%2FResourceGroupForSDKTest%2Fproviders%2FMicrosoft.StorSimple%2Fmanagers%2FManagerForSDKTest1%2Fdevices%2FDevice05ForSDKTest%2FbackupPolicies%2FBkUpPolicy01ForSDKTest'")})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupsDelete.json
func ExampleBackupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewBackupsClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"Device05ForSDKTest",
		"880e1774-94a8-4f3e-85e6-a61e6b94a8b7",
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupsClone.json
func ExampleBackupsClient_BeginClone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewBackupsClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginClone(ctx,
		"Device05ForSDKTest",
		"880e1774-94a8-4f3e-85e6-a61e6b94a8b7",
		"7e115577-4a3b-4921-bfd4-ee5a1b9bcbb5_0000000000000000",
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest1",
		armstorsimple8000series.CloneRequest{
			BackupElement: &armstorsimple8000series.BackupElement{
				ElementID:         to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backups/880e1774-94a8-4f3e-85e6-a61e6b94a8b7/elements/7e115577-4a3b-4921-bfd4-ee5a1b9bcbb5_0000000000000000"),
				ElementName:       to.Ptr("7e115577-4a3b-4921-bfd4-ee5a1b9bcbb5_0000000000000000"),
				ElementType:       to.Ptr("managers/devices/backups/elements"),
				SizeInBytes:       to.Ptr[int64](10737418240),
				VolumeContainerID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest"),
				VolumeName:        to.Ptr("Clonedvolume1"),
				VolumeType:        to.Ptr(armstorsimple8000series.VolumeTypeTiered),
			},
			TargetAccessControlRecordIDs: []*string{
				to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACR2")},
			TargetDeviceID:   to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
			TargetVolumeName: to.Ptr("ClonedClonedvolume1"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupsRestore.json
func ExampleBackupsClient_BeginRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewBackupsClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRestore(ctx,
		"Device05ForSDKTest",
		"880e1774-94a8-4f3e-85e6-a61e6b94a8b7",
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
