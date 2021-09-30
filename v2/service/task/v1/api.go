// Package task code generated by lark suite oapi sdk gen
package task

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v2"
)

type TaskService struct {
	Tasks         *tasks
	Comments      *comments
	Collaborators *collaborators
	Followers     *followers
	Reminders     *reminders
}

func New(app *lark.App) *TaskService {
	t := &TaskService{}
	t.Tasks = &tasks{app: app}
	t.Comments = &comments{app: app}
	t.Collaborators = &collaborators{app: app}
	t.Followers = &followers{app: app}
	t.Reminders = &reminders{app: app}
	return t
}

type tasks struct {
	app *lark.App
}
type comments struct {
	app *lark.App
}
type collaborators struct {
	app *lark.App
}
type followers struct {
	app *lark.App
}
type reminders struct {
	app *lark.App
}

func (t *tasks) Complete(ctx context.Context, req *TaskCompleteReq, options ...lark.RequestOptionFunc) (*TaskCompleteResp, error) {
	rawResp, err := t.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/complete", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskCompleteResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *tasks) Create(ctx context.Context, req *TaskCreateReq, options ...lark.RequestOptionFunc) (*TaskCreateResp, error) {
	rawResp, err := t.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/task/v1/tasks", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskCreateResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *tasks) Delete(ctx context.Context, req *TaskDeleteReq, options ...lark.RequestOptionFunc) (*TaskDeleteResp, error) {
	rawResp, err := t.app.SendRequestWithAccessTokenTypes(ctx, http.MethodDelete,
		"/open-apis/task/v1/tasks/:task_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskDeleteResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *tasks) Get(ctx context.Context, req *TaskGetReq, options ...lark.RequestOptionFunc) (*TaskGetResp, error) {
	rawResp, err := t.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/task/v1/tasks/:task_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskGetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *tasks) Patch(ctx context.Context, req *TaskPatchReq, options ...lark.RequestOptionFunc) (*TaskPatchResp, error) {
	rawResp, err := t.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPatch,
		"/open-apis/task/v1/tasks/:task_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskPatchResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *tasks) Uncomplete(ctx context.Context, req *TaskUncompleteReq, options ...lark.RequestOptionFunc) (*TaskUncompleteResp, error) {
	rawResp, err := t.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/uncomplete", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskUncompleteResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *collaborators) Create(ctx context.Context, req *TaskCollaboratorCreateReq, options ...lark.RequestOptionFunc) (*TaskCollaboratorCreateResp, error) {
	rawResp, err := c.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/collaborators", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskCollaboratorCreateResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *collaborators) Delete(ctx context.Context, req *TaskCollaboratorDeleteReq, options ...lark.RequestOptionFunc) (*TaskCollaboratorDeleteResp, error) {
	rawResp, err := c.app.SendRequestWithAccessTokenTypes(ctx, http.MethodDelete,
		"/open-apis/task/v1/tasks/:task_id/collaborators/:collaborator_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskCollaboratorDeleteResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *collaborators) List(ctx context.Context, req *TaskCollaboratorListReq, options ...lark.RequestOptionFunc) (*TaskCollaboratorListResp, error) {
	rawResp, err := c.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/task/v1/tasks/:task_id/collaborators", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskCollaboratorListResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *comments) Create(ctx context.Context, req *TaskCommentCreateReq, options ...lark.RequestOptionFunc) (*TaskCommentCreateResp, error) {
	rawResp, err := c.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/comments", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskCommentCreateResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *comments) Delete(ctx context.Context, req *TaskCommentDeleteReq, options ...lark.RequestOptionFunc) (*TaskCommentDeleteResp, error) {
	rawResp, err := c.app.SendRequestWithAccessTokenTypes(ctx, http.MethodDelete,
		"/open-apis/task/v1/tasks/:task_id/comments/:comment_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskCommentDeleteResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *comments) Get(ctx context.Context, req *TaskCommentGetReq, options ...lark.RequestOptionFunc) (*TaskCommentGetResp, error) {
	rawResp, err := c.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/task/v1/tasks/:task_id/comments/:comment_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskCommentGetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *comments) Update(ctx context.Context, req *TaskCommentUpdateReq, options ...lark.RequestOptionFunc) (*TaskCommentUpdateResp, error) {
	rawResp, err := c.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPut,
		"/open-apis/task/v1/tasks/:task_id/comments/:comment_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskCommentUpdateResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *followers) Create(ctx context.Context, req *TaskFollowerCreateReq, options ...lark.RequestOptionFunc) (*TaskFollowerCreateResp, error) {
	rawResp, err := f.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/followers", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskFollowerCreateResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *followers) Delete(ctx context.Context, req *TaskFollowerDeleteReq, options ...lark.RequestOptionFunc) (*TaskFollowerDeleteResp, error) {
	rawResp, err := f.app.SendRequestWithAccessTokenTypes(ctx, http.MethodDelete,
		"/open-apis/task/v1/tasks/:task_id/followers/:follower_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskFollowerDeleteResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *followers) List(ctx context.Context, req *TaskFollowerListReq, options ...lark.RequestOptionFunc) (*TaskFollowerListResp, error) {
	rawResp, err := f.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/task/v1/tasks/:task_id/followers", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskFollowerListResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reminders) Create(ctx context.Context, req *TaskReminderCreateReq, options ...lark.RequestOptionFunc) (*TaskReminderCreateResp, error) {
	rawResp, err := r.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/reminders", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskReminderCreateResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reminders) Delete(ctx context.Context, req *TaskReminderDeleteReq, options ...lark.RequestOptionFunc) (*TaskReminderDeleteResp, error) {
	rawResp, err := r.app.SendRequestWithAccessTokenTypes(ctx, http.MethodDelete,
		"/open-apis/task/v1/tasks/:task_id/reminders/:reminder_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskReminderDeleteResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reminders) List(ctx context.Context, req *TaskReminderListReq, options ...lark.RequestOptionFunc) (*TaskReminderListResp, error) {
	rawResp, err := r.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/task/v1/tasks/:task_id/reminders", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &TaskReminderListResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}