package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/blog/internal/data/model"
	"github.com/hixraid/blog/internal/response"
)

const commentIdParam = "comment_id"

func (h *Handler) createComment(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param(postIdParam))
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid param post_id")
		return
	}

	var input model.CommentInput
	if err := ctx.Bind(&input); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	commentId, err := h.service.Comment.Create(postId, input)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.NewOkResponse(ctx, model.CommentIdResponse{
		CommentId: commentId,
	})
}

func (h *Handler) commentsByPostId(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param(postIdParam))
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid param post_id")
		return
	}

	comments, err := h.service.Comment.GetAll(postId)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.NewOkResponse(ctx, comments)
}

func (h *Handler) updateCommentById(ctx *gin.Context) {
	commentId, err := strconv.Atoi(ctx.Param(commentIdParam))
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid param comment_id")
		return
	}

	var input model.CommentInput
	if err := ctx.Bind(&input); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.service.Comment.UpdateById(commentId, input)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewStatusResponse(ctx, "OK")
}

func (h *Handler) deleteCommentById(ctx *gin.Context) {
	commentId, err := strconv.Atoi(ctx.Param(commentIdParam))
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid param comment_id")
		return
	}

	err = h.service.Comment.DeleteById(commentId)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.NewStatusResponse(ctx, "OK")
}
