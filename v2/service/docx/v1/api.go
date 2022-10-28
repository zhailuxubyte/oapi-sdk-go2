// Package docx code generated by lark suite oapi sdk gen
package docx

import (
	"context"
	"net/http"
	
	"github.com/larksuite/oapi-sdk-go/v2"
)

type DocxService struct {
	Documents *documents
	Blocks *blocks
	DocumentBlockChildren *documentBlockChildren
}

func New(app *lark.App) *DocxService {
	d := &DocxService{}
	d.Documents = &documents{app: app}
	d.Blocks = &blocks{app: app}
	d.DocumentBlockChildren = &documentBlockChildren{app: app}
	return d
}

type documents struct {
	app *lark.App
}
type blocks struct {
	app *lark.App
}
type documentBlockChildren struct {
	app *lark.App
}

func (d *documents) Create(ctx context.Context, req *DocumentCreateReq, options ...lark.RequestOptionFunc) (*DocumentCreateResp, error) {
	rawResp, err := d.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/docx/v1/documents", []lark.AccessTokenType{lark.AccessTokenTypeTenant, lark.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &DocumentCreateResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documents) Get(ctx context.Context, req *DocumentGetReq, options ...lark.RequestOptionFunc) (*DocumentGetResp, error) {
	rawResp, err := d.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/docx/v1/documents/:document_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant, lark.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &DocumentGetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documents) RawContent(ctx context.Context, req *DocumentRawContentReq, options ...lark.RequestOptionFunc) (*DocumentRawContentResp, error) {
	rawResp, err := d.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/docx/v1/documents/:document_id/raw_content", []lark.AccessTokenType{lark.AccessTokenTypeTenant, lark.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &DocumentRawContentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (b *blocks) BatchUpdate(ctx context.Context, req *DocumentBlockBatchUpdateReq, options ...lark.RequestOptionFunc) (*DocumentBlockBatchUpdateResp, error) {
	rawResp, err := b.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPatch,
		"/open-apis/docx/v1/documents/:document_id/blocks/batch_update", []lark.AccessTokenType{lark.AccessTokenTypeTenant, lark.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &DocumentBlockBatchUpdateResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (b *blocks) Get(ctx context.Context, req *DocumentBlockGetReq, options ...lark.RequestOptionFunc) (*DocumentBlockGetResp, error) {
	rawResp, err := b.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/docx/v1/documents/:document_id/blocks/:block_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant, lark.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &DocumentBlockGetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (b *blocks) List(ctx context.Context, req *DocumentBlockListReq, options ...lark.RequestOptionFunc) (*DocumentBlockListResp, error) {
	rawResp, err := b.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/docx/v1/documents/:document_id/blocks", []lark.AccessTokenType{lark.AccessTokenTypeTenant, lark.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &DocumentBlockListResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (b *blocks) Patch(ctx context.Context, req *DocumentBlockPatchReq, options ...lark.RequestOptionFunc) (*DocumentBlockPatchResp, error) {
	rawResp, err := b.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPatch,
		"/open-apis/docx/v1/documents/:document_id/blocks/:block_id", []lark.AccessTokenType{lark.AccessTokenTypeTenant, lark.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &DocumentBlockPatchResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlockChildren) BatchDelete(ctx context.Context, req *DocumentBlockChildrenBatchDeleteReq, options ...lark.RequestOptionFunc) (*DocumentBlockChildrenBatchDeleteResp, error) {
	rawResp, err := d.app.SendRequestWithAccessTokenTypes(ctx, http.MethodDelete,
		"/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children/batch_delete", []lark.AccessTokenType{lark.AccessTokenTypeTenant, lark.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &DocumentBlockChildrenBatchDeleteResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlockChildren) Create(ctx context.Context, req *DocumentBlockChildrenCreateReq, options ...lark.RequestOptionFunc) (*DocumentBlockChildrenCreateResp, error) {
	rawResp, err := d.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children", []lark.AccessTokenType{lark.AccessTokenTypeTenant, lark.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &DocumentBlockChildrenCreateResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlockChildren) Get(ctx context.Context, req *DocumentBlockChildrenGetReq, options ...lark.RequestOptionFunc) (*DocumentBlockChildrenGetResp, error) {
	rawResp, err := d.app.SendRequestWithAccessTokenTypes(ctx, http.MethodGet,
		"/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children", []lark.AccessTokenType{lark.AccessTokenTypeTenant, lark.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &DocumentBlockChildrenGetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}