package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/semeru3-go/dto"
	"github.com/purawaktra/semeru3-go/functions"
	"github.com/purawaktra/semeru3-go/utils"
	"net/http"
	"strconv"
	"time"
)

type Semeru3RequestHandler struct {
	ctrl Semeru3Controller
}

func CreateSemeru3RequestHandler(ctrl Semeru3Controller) Semeru3RequestHandler {
	return Semeru3RequestHandler{
		ctrl: ctrl,
	}
}

func (rh Semeru3RequestHandler) SelectProfile(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "SelectProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "SelectProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"UU",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "SelectProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the profile picture and check err
	utils.Debug("SelectProfile", header)
	data, err, code := rh.ctrl.SelectProfileById(bodyBytes)
	if err != nil {
		utils.Error(err, "SelectProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, data)
	return
}

func (rh Semeru3RequestHandler) InsertProfile(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "InsertProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "InsertProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"UU",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "InsertProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the profile picture and check err
	utils.Debug("InsertProfile", bodyBytes)
	err, code := rh.ctrl.InsertProfile(bodyBytes)
	if err != nil {
		utils.Error(err, "InsertProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, nil)
	return
}

func (rh Semeru3RequestHandler) UpdateProfile(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "UpdateProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "UpdateProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"UU",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "UpdateProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the profile picture and check err
	utils.Debug("UpdateProfile", header)
	err, code := rh.ctrl.UpdateProfileById(bodyBytes)
	if err != nil {
		utils.Error(err, "UpdateProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, nil)
	return
}

func (rh Semeru3RequestHandler) DeleteProfile(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "DeleteProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "DeleteProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"UU",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "DeleteProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the profile picture and check err
	utils.Debug("DeleteProfile", header)
	err, code := rh.ctrl.DeleteProfileById(bodyBytes)
	if err != nil {
		utils.Error(err, "DeleteProfile", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, nil)
	return
}
