// Package task code generated by lark suite oapi sdk gen
package task

import (
	"github.com/larksuite/oapi-sdk-go/v2"
)

type Href struct {
	Url   *string `json:"url,omitempty"`
	Title *string `json:"title,omitempty"`
}

type Collaborator struct {
	Id *string `json:"id,omitempty"`
}

type Comment struct {
	Content  *string `json:"content,omitempty"`
	ParentId *int64  `json:"parent_id,omitempty,string"`
	Id       *int64  `json:"id,omitempty,string"`
}

type Due struct {
	Time     *int64  `json:"time,omitempty,string"`
	Timezone *string `json:"timezone,omitempty"`
	IsAllDay *bool   `json:"is_all_day,omitempty"`
}

type Follower struct {
	Id *string `json:"id,omitempty"`
}

type Origin struct {
	PlatformI18nName *string `json:"platform_i18n_name,omitempty"`
	Href             *Href   `json:"href,omitempty"`
}

type Reminder struct {
	Id                 *int64 `json:"id,omitempty,string"`
	RelativeFireMinute *int   `json:"relative_fire_minute,omitempty"`
}

type Task struct {
	Id           *string `json:"id,omitempty"`
	Summary      *string `json:"summary,omitempty"`
	Description  *string `json:"description,omitempty"`
	CompleteTime *int64  `json:"complete_time,omitempty,string"`
	CreatorId    *string `json:"creator_id,omitempty"`
	Extra        *string `json:"extra,omitempty"`
	CreateTime   *int64  `json:"create_time,omitempty,string"`
	UpdateTime   *int64  `json:"update_time,omitempty,string"`
	Due          *Due    `json:"due,omitempty"`
	Origin       *Origin `json:"origin,omitempty"`
	CanEdit      *bool   `json:"can_edit,omitempty"`
}

type TaskCompleteReq struct {
	TaskId string `path:"task_id"`
}

type TaskCompleteResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
}

type TaskCreateReq struct {
	UserIdType *string `query:"user_id_type"`
	Task       *Task   `body:""`
}

type TaskCreateRespData struct {
	Task *Task `json:"task,omitempty"`
}

type TaskCreateResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskCreateRespData `json:"data"`
}

type TaskDeleteReq struct {
	TaskId string `path:"task_id"`
}

type TaskDeleteResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
}

type TaskGetReq struct {
	TaskId     string  `path:"task_id"`
	UserIdType *string `query:"user_id_type"`
}

type TaskGetRespData struct {
	Task *Task `json:"task,omitempty"`
}

type TaskGetResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskGetRespData `json:"data"`
}

type TaskPatchReqBody struct {
	Task         *Task    `json:"task,omitempty"`
	UpdateFields []string `json:"update_fields,omitempty"`
}

type TaskPatchReq struct {
	TaskId     string            `path:"task_id"`
	UserIdType *string           `query:"user_id_type"`
	Body       *TaskPatchReqBody `body:""`
}

type TaskPatchRespData struct {
	Task *Task `json:"task,omitempty"`
}

type TaskPatchResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskPatchRespData `json:"data"`
}

type TaskUncompleteReq struct {
	TaskId string `path:"task_id"`
}

type TaskUncompleteResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
}

type TaskCollaboratorCreateReq struct {
	TaskId       string        `path:"task_id"`
	UserIdType   *string       `query:"user_id_type"`
	Collaborator *Collaborator `body:""`
}

type TaskCollaboratorCreateRespData struct {
	Collaborator *Collaborator `json:"collaborator,omitempty"`
}

type TaskCollaboratorCreateResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskCollaboratorCreateRespData `json:"data"`
}

type TaskCollaboratorDeleteReq struct {
	TaskId         string `path:"task_id"`
	CollaboratorId string `path:"collaborator_id"`
}

type TaskCollaboratorDeleteResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
}

type TaskCollaboratorListReq struct {
	TaskId     string  `path:"task_id"`
	PageSize   *int    `query:"page_size"`
	PageToken  *string `query:"page_token"`
	UserIdType *string `query:"user_id_type"`
}

type TaskCollaboratorListRespData struct {
	Items     []*Collaborator `json:"items,omitempty"`
	PageToken *string         `json:"page_token,omitempty"`
	HasMore   *bool           `json:"has_more,omitempty"`
}

type TaskCollaboratorListResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskCollaboratorListRespData `json:"data"`
}

type TaskCommentCreateReq struct {
	TaskId  string   `path:"task_id"`
	Comment *Comment `body:""`
}

type TaskCommentCreateRespData struct {
	Comment *Comment `json:"comment,omitempty"`
}

type TaskCommentCreateResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskCommentCreateRespData `json:"data"`
}

type TaskCommentDeleteReq struct {
	TaskId    string `path:"task_id"`
	CommentId string `path:"comment_id"`
}

type TaskCommentDeleteResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
}

type TaskCommentGetReq struct {
	TaskId    string `path:"task_id"`
	CommentId string `path:"comment_id"`
}

type TaskCommentGetRespData struct {
	Comment *Comment `json:"comment,omitempty"`
}

type TaskCommentGetResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskCommentGetRespData `json:"data"`
}

type TaskCommentUpdateReqBody struct {
	Content *string `json:"content,omitempty"`
}

type TaskCommentUpdateReq struct {
	TaskId    string                    `path:"task_id"`
	CommentId string                    `path:"comment_id"`
	Body      *TaskCommentUpdateReqBody `body:""`
}

type TaskCommentUpdateRespData struct {
	Comment *Comment `json:"comment,omitempty"`
}

type TaskCommentUpdateResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskCommentUpdateRespData `json:"data"`
}

type TaskFollowerCreateReq struct {
	TaskId     string    `path:"task_id"`
	UserIdType *string   `query:"user_id_type"`
	Follower   *Follower `body:""`
}

type TaskFollowerCreateRespData struct {
	Follower *Follower `json:"follower,omitempty"`
}

type TaskFollowerCreateResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskFollowerCreateRespData `json:"data"`
}

type TaskFollowerDeleteReq struct {
	TaskId     string `path:"task_id"`
	FollowerId string `path:"follower_id"`
}

type TaskFollowerDeleteResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
}

type TaskFollowerListReq struct {
	TaskId     string  `path:"task_id"`
	PageSize   *int    `query:"page_size"`
	PageToken  *string `query:"page_token"`
	UserIdType *string `query:"user_id_type"`
}

type TaskFollowerListRespData struct {
	Items     []*Follower `json:"items,omitempty"`
	PageToken *string     `json:"page_token,omitempty"`
	HasMore   *bool       `json:"has_more,omitempty"`
}

type TaskFollowerListResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskFollowerListRespData `json:"data"`
}

type TaskReminderCreateReq struct {
	TaskId   string    `path:"task_id"`
	Reminder *Reminder `body:""`
}

type TaskReminderCreateRespData struct {
	Reminder *Reminder `json:"reminder,omitempty"`
}

type TaskReminderCreateResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskReminderCreateRespData `json:"data"`
}

type TaskReminderDeleteReq struct {
	TaskId     string `path:"task_id"`
	ReminderId string `path:"reminder_id"`
}

type TaskReminderDeleteResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
}

type TaskReminderListReq struct {
	TaskId    string  `path:"task_id"`
	PageSize  *int    `query:"page_size"`
	PageToken *string `query:"page_token"`
}

type TaskReminderListRespData struct {
	Items     []*Reminder `json:"items,omitempty"`
	PageToken *string     `json:"page_token,omitempty"`
	HasMore   *bool       `json:"has_more,omitempty"`
}

type TaskReminderListResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *TaskReminderListRespData `json:"data"`
}

type TaskUpdatedEventData struct {
	TaskId  *string `json:"task_id,omitempty"`
	ObjType *int    `json:"obj_type,omitempty"`
}

type TaskUpdatedEvent struct {
	*lark.EventV2Base
	Event *TaskUpdatedEventData `json:"event"`
}

type TaskCommentUpdatedEventData struct {
	TaskId    *string `json:"task_id,omitempty"`
	CommentId *string `json:"comment_id,omitempty"`
	ParentId  *string `json:"parent_id,omitempty"`
	ObjType   *int    `json:"obj_type,omitempty"`
}

type TaskCommentUpdatedEvent struct {
	*lark.EventV2Base
	Event *TaskCommentUpdatedEventData `json:"event"`
}