package http

import (
	"encoding/json"
	"errors"
	"lillybox-backend/internal/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetLoginRequestID godoc
// @Title       GetLoginRequestID
// @Summary     Get login request ID
// @Description 유저가 있다면 UUID를 받아옵니다. 회원가입되지 않은 유저라면 null을 반환합니다.
// @Tags        auth
// @Accept      json
// @Param       UserAccount body GetLoginRequestIDBody true "GetLoginRequestID Body"
// @Produce     json
// @Success     200 {object} LoginResponse
// @Failure     400 {object} DefaultResponse
// @Failure     404 {object} DefaultDataResponse
// @Router      /api/v1/auth/request_id [post]
func (h *Handlers) GetLoginRequestID(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	resp, err := GetOrInsertUser(h, c.Request().Body())
	switch {
	case errors.Is(err, nil):
		return c.Status(fiber.StatusOK).JSON(LoginResponse{resp.LoginRequestID})
	case errors.Is(err, gorm.ErrRecordNotFound):
		return c.Status(fiber.ErrNotFound.Code).JSON(DefaultDataResponse{false, nil})
	default:
		return c.Status(fiber.StatusBadRequest).JSON(DefaultResponse{fiber.ErrBadRequest.Message})
	}
}

// SignUp godoc
// @Title       SignUp
// @Summary     SignUp
// @Description 유저를 insert하고 유저 정보를 가져옵니다.
// @Tags        auth
// @Accept      json
// @Param       UserInfo body SignUpBody true "SignUp Body"
// @Produce     json
// @Success     200 {object} DefaultDataResponse
// @Failure     400 {object} DefaultResponse
// @Failure     409 {object} DefaultResponse
// @Failure     500 {object} DefaultResponse
// @Router      /api/v1/auth/sign [post]
func (h *Handlers) SignUp(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	body := ParseBytesToMapObject(c.Body())
	if body != nil && KeyCheck(body, ADDR) && KeyCheck(body, NICK) {
		_, err := h.Database.InsertUser(body[ADDR], body[NICK])
		switch {
		case errors.Is(err, nil):
			break
		case errors.Is(err, gorm.ErrInvalidTransaction):
			return c.Status(fiber.ErrInternalServerError.Code).JSON(DefaultResponse{fiber.ErrInternalServerError.Message})
		case errors.Is(err, gorm.ErrInvalidData):
			return c.Status(fiber.StatusConflict).JSON(DefaultResponse{fiber.ErrConflict.Message})
		default:
			return c.Status(fiber.StatusBadRequest).JSON(DefaultResponse{fiber.ErrBadRequest.Message})
		}
		user, err := h.Database.ReadUser(body[ADDR])
		switch {
		case errors.Is(err, nil):
			return c.Status(fiber.StatusOK).JSON(DefaultDataResponse{true, user})
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.Status(fiber.StatusInternalServerError).JSON(DefaultResponse{fiber.ErrInternalServerError.Message})
		default:
			return c.Status(fiber.StatusBadRequest).JSON(DefaultResponse{fiber.ErrBadRequest.Message})
		}
	}
	return c.Status(fiber.StatusBadRequest).JSON(DefaultResponse{fiber.ErrBadRequest.Message})
}

// GetUserInfo godoc
// @Title       GetUserInfo
// @Summary     GetUserInfo
// @Description 유저의 정보를 조회합니다.
// @Tags        auth
// @Accept      json
// @Param       UserAccount body GetUserInfoBody true "GetUserInfo Body"
// @Produce     json
// @Success     200 {object} DefaultDataResponse
// @Failure     400 {object} DefaultResponse
// @Failure     500 {object} DefaultResponse
// @Router      /api/v1/auth/user [post]
func (h *Handlers) GetUserInfo(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	body := ParseBytesToMapObject(c.Body())
	if body != nil && KeyCheck(body, ADDR) {
		user, err := h.Database.ReadUser(body[ADDR])
		switch {
		case errors.Is(err, nil):
			return c.Status(fiber.StatusOK).JSON(DefaultDataResponse{true, user})
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.Status(fiber.StatusInternalServerError).JSON(DefaultResponse{fiber.ErrInternalServerError.Message})
		default:
			return c.Status(fiber.StatusBadRequest).JSON(DefaultResponse{fiber.ErrBadRequest.Message})
		}
	}
	return c.Status(fiber.StatusBadRequest).JSON(DefaultResponse{fiber.ErrBadRequest.Message})
}

// func (h *Handlers) Login(c *fiber.Ctx) error {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r)))
// 		}
// 	}()
// 	_, err := VerifySignature(h, c.Body())
// 	switch {
// 	case errors.Is(err, nil):
// 		return c.Status(fiber.StatusOK).JSON(BoolResponse{true})
// 	case errors.Is(err, fiber.ErrBadRequest):
// 		return c.Status(fiber.StatusBadRequest).JSON(DefaultResponse{fiber.ErrBadRequest.Message})
// 	default:
// 		return c.Status(fiber.StatusBadRequest).JSON(DefaultResponse{fiber.ErrBadRequest.Message})
// 	}
// }

// CheckNickname godoc
// @Title       CheckNickname
// @Summary     CheckNickname
// @Description 닉네임으로 유저를 조회합니다.
// @Tags        auth
// @Accept      json
// @Param       UserAccount body GetNicknameBody true "CheckNickname Body"
// @Produce     json
// @Success     200 {object} BoolResponse
// @Failure     400 {object} DefaultResponse
// @Failure     404 {object} BoolResponse
// @Router      /api/v1/auth/check [post]
func (h *Handlers) CheckNickname(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	body := ParseBytesToMapObject(c.Body())
	if body != nil && KeyCheck(body, NICK) {
		_, err := h.Database.ReadUserByNickname(body[NICK])
		switch {
		case errors.Is(err, nil):
			return c.Status(fiber.StatusOK).JSON(BoolResponse{true})
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.Status(fiber.StatusNotFound).JSON(BoolResponse{false})
		default:
			return c.Status(fiber.StatusBadRequest).JSON(DefaultResponse{fiber.ErrBadRequest.Message})
		}
	}
	return c.Status(fiber.StatusBadRequest).JSON(DefaultResponse{fiber.ErrBadRequest.Message})
}

// GetStreamings godoc
func (h *Handlers) GetStreamings(c *fiber.Ctx) error {
	return c.SendString("GetStreamings")
}

// CreateStreaming godoc
func (h *Handlers) CreateStreaming(c *fiber.Ctx) error {
	return c.SendString("CreateStreaming")
}

// GetStreamingByID godoc
func (h *Handlers) GetStreamingByID(c *fiber.Ctx) error {
	return c.SendString("GetStreamingByID")
}

// DeleteStreamingByID godoc
func (h *Handlers) DeleteStreamingByID(c *fiber.Ctx) error {
	return c.SendString("DeleteStreamingByID")
}

// PatchStreamingByID godoc
func (h *Handlers) PatchStreamingByID(c *fiber.Ctx) error {
	return c.SendString("PatchStreamingByID")
}

// GetOnDemands godoc
// @Title       GetOnDemands
// @Summary     GetOnDemands
// @Description 카테고리 별 전체 비디오를 조회합니다.
// @Tags        videos
// @Accept      json
// @Param       offset     query string true "Offset"
// @Param       categoryId query string true "CategoryID"
// @Produce     json
// @Success     200 {object} GetOnDemandResponse
// @Failure     400 {object} GetOnDemandResponse
// @Failure     404 {object} GetOnDemandResponse
// @Router      /api/v1/videos [get]
func (h *Handlers) GetOnDemands(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	offset := c.Query("offset", "0")
	category := c.Query("categoryId", "1")
	result, err := h.Database.ReadManyVOD(offset, category)
	switch {
	case errors.Is(err, nil):
		return c.Status(fiber.StatusOK).JSON(GetOnDemandResponse{true, ParseLillyVidMany(result)})
	case errors.Is(err, gorm.ErrRecordNotFound):
		return c.Status(fiber.StatusNotFound).JSON(GetOnDemandResponse{false, nil})
	default:
		return c.Status(fiber.StatusBadRequest).JSON(GetOnDemandResponse{false, nil})
	}
}

// GetOnDemandByID godoc
// @Title       GetOnDemandByID
// @Summary     GetOnDemandByID
// @Description 비디오 상세 페이지를 조회합니다.
// @Tags        videos
// @Accept      json
// @Param       id path string true "Video ID"
// @Produce     json
// @Success     200 {object} GetOnDemandByIDResponse
// @Failure     400 {object} DefaultDataResponse
// @Failure     404 {object} DefaultDataResponse
// @Router      /api/v1/videos/{id} [get]
func (h *Handlers) GetOnDemandByID(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	videoID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(DefaultDataResponse{false, nil})
	}
	go h.Database.UpdateVidView(uint(videoID))
	result, err := h.Database.ReadVOD(uint(videoID))
	switch {
	case errors.Is(err, nil):
		var owner = &database.Users{}
		owner, err := h.Database.ReadUserByNickname(result.Nickname)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(DefaultDataResponse{false, nil})
		}
		lillyVideo := ParseLillyVideo(result)
		lillyVideo.Account = owner.Address
		return c.Status(fiber.StatusOK).JSON(DefaultDataResponse{true, lillyVideo})
	case errors.Is(err, gorm.ErrRecordNotFound):
		return c.Status(fiber.StatusNotFound).JSON(DefaultDataResponse{false, nil})
	default:
		return c.Status(fiber.ErrBadRequest.Code).JSON(DefaultDataResponse{false, nil})
	}
}

// DeleteOnDemandByID godoc
func (h *Handlers) DeleteOnDemandByID(c *fiber.Ctx) error {
	return c.SendString("DeleteOnDemandByID")
}

// PatchOnDemandByID godoc
// @Title       PatchOnDemandByID
// @Summary     PatchOnDemandByID
// @Description 비디오를 업데이트합니다.
// @Tags        videos
// @Accept      json
// @Param       id path string true "Video ID"
// @Produce     json
// @Success     200 {object} BoolResponse
// @Failure     400 {object} DefaultDataResponse
// @Router      /api/v1/videos/{id} [patch]
func (h *Handlers) PatchOnDemandByID(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	videoID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	var dto database.UpdateVideoDto
	if err := json.Unmarshal(c.Body(), &dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	_, err = h.Database.UpdateVOD(dto, uint(videoID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	return c.Status(fiber.StatusOK).JSON(BoolResponse{true})
}

// GetComments godoc
// @Title       GetComments
// @Summary     GetComments
// @Description 해당 비디오의 전체 댓글을 조회합니다.
// @Tags        comments
// @Accept      json
// @Param       id path string true "Video ID"
// @Produce     json
// @Success     200 {object} DefaultDataResponse
// @Failure     400 {object} DefaultDataResponse
// @Router      /api/v1/videos/{id}/comments [get]
func (h *Handlers) GetComments(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	videoID := c.Params("id")
	if videoID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	id, _ := strconv.ParseUint(videoID, 0, 32)
	comment, err := h.Database.ReadComment(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	return c.Status(fiber.StatusOK).JSON(DefaultDataResponse{true, ParseLillyCommentMany(comment)})
}

// CreateComment godoc
// @Title       CreateComment
// @Summary     CreateComment
// @Description 댓글을 작성합니다.
// @Tags        comments
// @Accept      json
// @Param       id               path string                    true "Video ID"
// @Param       CreateCommentDto body database.CreateCommentDto true "Create Comment"
// @Produce     json
// @Success     200 {object} BoolResponse
// @Failure     400 {object} DefaultDataResponse
// @Router      /api/v1/videos/{id}/comments [post]
func (h *Handlers) CreateComment(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	id := c.Params("id")
	var dto database.CreateCommentDto
	if err := json.Unmarshal(c.Body(), &dto); err != nil || id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	videoID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	if _, err := h.Database.InsertComment(dto, uint(videoID)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	return c.Status(fiber.StatusOK).JSON(BoolResponse{true})
}

// PatchCommentByID godoc
// @Title       PatchCommentByID
// @Summary     PatchCommentByID
// @Description 댓글을 수정합니다.
// @Tags        comments
// @Accept      json
// @Param       id               path string                    true "Video ID"
// @Param       UpdateCommentDto body database.UpdateCommentDto true "Update Comment"
// @Produce     json
// @Success     200 {object} BoolResponse
// @Failure     400 {object} DefaultDataResponse
// @Router      /api/v1/videos/{id}/comments [patch]
func (h *Handlers) PatchCommentByID(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	id := c.Params("id")
	videoID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	var dto database.UpdateCommentDto
	if err := json.Unmarshal(c.Body(), &dto); err != nil || id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	if _, err := h.Database.UpdateComment(dto, uint(videoID)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	return c.Status(fiber.StatusOK).JSON(BoolResponse{true})

}

// DeleteCommentByID godoc
// @Title       DeleteCommentByID
// @Summary     DeleteCommentByID
// @Description 댓글을 삭제합니다.
// @Tags        comments
// @Accept      json
// @Param       id               path  string                    true "Video ID"
// @Param       comment_id       query string                    true "Comment ID"
// @Produce     json
// @Success     200 {object} BoolResponse
// @Failure     400 {object} DefaultDataResponse
// @Router      /api/v1/videos/{id}/comments [delete]
func (h *Handlers) DeleteCommentByID(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	commentID := c.Query("comment_id")
	id, _ := strconv.Atoi(commentID)
	if _, err := h.Database.DeleteComment(uint(id)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	return c.Status(fiber.StatusOK).JSON(BoolResponse{true})
}

// ReportOnDemandByID godoc
// @Title       ReportOnDemandByID
// @Summary     ReportOnDemandByID
// @Description 해당 비디오를 신고합니다.
// @Tags        report
// @Accept      json
// @Param       Nickname body GetNicknameBody true "ReportOnDemandByID Body"
// @Produce     json
// @Success     200 {object} BoolResponse
// @Failure     400 {object} DefaultDataResponse
// @Router      /api/v1/reports/videos/{id} [patch]
func (h *Handlers) ReportOnDemandByID(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	id := c.Params("id")
	videoID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	var dto database.ReportVidDto
	if err := json.Unmarshal(c.Body(), &dto); err != nil || id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	if dto.Nickname == "" {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	if _, err := h.Database.ReportVid(dto, uint(videoID)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	return c.Status(fiber.StatusOK).JSON(BoolResponse{true})
}

// ReportCommentByID godoc
// @Title       ReportCommentByID
// @Summary     ReportCommentByID
// @Description 해당 댓글을 신고합니다.
// @Tags        report
// @Accept      json
// @Param       UserNickname body GetNicknameBody true "ReportCommentByID Body"
// @Produce     json
// @Success     200 {object} BoolResponse
// @Failure     400 {object} DefaultDataResponse
// @Router      /api/v1/reports/comments/{id} [patch]
func (h *Handlers) ReportCommentByID(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			h.ErrorLogger.Error(toString("Unexpected Errors Occurred :", r))
		}
	}()
	id := c.Params("id")
	videoID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	var dto database.ReportCommentDto
	if err := json.Unmarshal(c.Body(), &dto); err != nil || id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	if dto.Nickname == "" {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	if _, err := h.Database.ReportComment(dto, uint(videoID)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(DefaultDataResponse{false, nil})
	}
	return c.Status(fiber.StatusOK).JSON(BoolResponse{true})
}
