// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/omec-project/util/httpwrapper"
	"github.com/omec-project/openapi"
	"github.com/omec-project/openapi/models"
	"github.com/omec-project/udr/logger"
	"github.com/omec-project/udr/producer"
)

// HTTPCreateSdmSubscriptions - Create individual sdm subscription
func HTTPCreateSdmSubscriptions(c *gin.Context) {
	var sdmSubscription models.SdmSubscription

	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.DataRepoLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&sdmSubscription, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.DataRepoLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := httpwrapper.NewRequest(c.Request, sdmSubscription)
	req.Params["ueId"] = c.Params.ByName("ueId")

	rsp := producer.HandleCreateSdmSubscriptions(req)

	for key, val := range rsp.Header {
		c.Header(key, val[0])
	}
	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.DataRepoLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

// HTTPQuerysdmsubscriptions - Retrieves the sdm subscriptions of a UE
func HTTPQuerysdmsubscriptions(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")

	rsp := producer.HandleQuerysdmsubscriptions(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.DataRepoLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
