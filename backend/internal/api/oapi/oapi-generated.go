// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version (devel) DO NOT EDIT.
package oapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/discord-gophers/goapi-gen/pkg/runtime"
	openapi_types "github.com/discord-gophers/goapi-gen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Error object returned on any error.
type Error struct {
	Message string `json:"message"`
}

// FoodCategories defines model for FoodCategories.
type FoodCategories struct {
	AdditionalProperties map[string][]string `json:"-"`
}

// FoodPreferences defines model for FoodPreferences.
type FoodPreferences struct {
	Likes   []string       `json:"likes"`
	Prefers FoodCategories `json:"prefers"`
}

// FormError defines model for FormError.
type FormError struct {
	// Embedded struct due to allOf(#/components/schemas/Error)
	Error `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	FormID string `json:"form_id,omitempty"`
}

// Snowflake ID.
type ID foodtinder.ID

// Optional metadata included on login.
type LoginMetadata struct {
	// The User-Agent used for logging in.
	UserAgent string `json:"user_agent,omitempty"`
}

// Post defines model for Post.
type Post struct {
	CoverHash   string `json:"cover_hash,omitempty"`
	Description string `json:"description"`

	// Snowflake ID.
	ID     ID       `json:"id"`
	Images []string `json:"images"`

	// Location is the location where the post was made.
	Location string   `json:"location,omitempty"`
	Tags     []string `json:"tags"`
	Username string   `json:"username"`
}

// Self defines model for Self.
type Self struct {
	// Embedded struct due to allOf(#/components/schemas/User)
	User `yaml:",inline"`
	// Embedded struct due to allOf(#/components/schemas/FoodPreferences)
	FoodPreferences `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Birthday *openapi_types.Date `json:"birthday,omitempty"`
}

// Session defines model for Session.
type Session struct {
	Expiry time.Time `json:"expiry"`

	// Optional metadata included on login.
	Metadata LoginMetadata `json:"metadata"`
	Token    string        `json:"token"`
	Username string        `json:"username"`
}

// User defines model for User.
type User struct {
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Username    string `json:"username"`
}

// Error object returned on a form error.
type NotFoundError FormError

// Error object returned on any error.
type ServerError Error

// Error object returned on any error.
type UserError Error

// LoginParams defines parameters for Login.
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetNextPostsParams defines parameters for GetNextPosts.
type GetNextPostsParams struct {
	// The ID to start the pagination from, or empty to start from top.
	PrevID *ID `json:"prev_id,omitempty"`
}

// CreatePostJSONBody defines parameters for CreatePost.
type CreatePostJSONBody Post

// RegisterParams defines parameters for Register.
type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreatePostJSONRequestBody defines body for CreatePost for application/json ContentType.
type CreatePostJSONRequestBody CreatePostJSONBody

// Bind implements render.Binder.
func (CreatePostJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// Response is a common response struct for all the API calls.
// A Response object may be instantiated via functions for specific operation responses.
type Response struct {
	body        interface{}
	statusCode  int
	contentType string
}

// Render implements the render.Renderer interface. It sets the Content-Type header
// and status code based on the response definition.
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", resp.contentType)
	render.Status(r, resp.statusCode)
	return nil
}

// Status is a builder method to override the default status code for a response.
func (resp *Response) Status(statusCode int) *Response {
	resp.statusCode = statusCode
	return resp
}

// ContentType is a builder method to override the default content type for a response.
func (resp *Response) ContentType(contentType string) *Response {
	resp.contentType = contentType
	return resp
}

// MarshalJSON implements the json.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(resp.body)
}

// MarshalXML implements the xml.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(resp.body)
}

// UploadAssetJSON200Response is a constructor method for a UploadAsset response.
// A *Response is returned with the configured status code and content type from the spec.
func UploadAssetJSON200Response(body string) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// UploadAssetJSON400Response is a constructor method for a UploadAsset response.
// A *Response is returned with the configured status code and content type from the spec.
func UploadAssetJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  400,
		contentType: "application/json",
	}
}

// UploadAssetJSON413Response is a constructor method for a UploadAsset response.
// A *Response is returned with the configured status code and content type from the spec.
func UploadAssetJSON413Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  413,
		contentType: "application/json",
	}
}

// UploadAssetJSON500Response is a constructor method for a UploadAsset response.
// A *Response is returned with the configured status code and content type from the spec.
func UploadAssetJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  500,
		contentType: "application/json",
	}
}

// GetAssetJSON404Response is a constructor method for a GetAsset response.
// A *Response is returned with the configured status code and content type from the spec.
func GetAssetJSON404Response(body FormError) *Response {
	return &Response{
		body:        body,
		statusCode:  404,
		contentType: "application/json",
	}
}

// GetAssetJSON500Response is a constructor method for a GetAsset response.
// A *Response is returned with the configured status code and content type from the spec.
func GetAssetJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  500,
		contentType: "application/json",
	}
}

// ListFoodsJSON200Response is a constructor method for a ListFoods response.
// A *Response is returned with the configured status code and content type from the spec.
func ListFoodsJSON200Response(body FoodCategories) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// ListFoodsJSON500Response is a constructor method for a ListFoods response.
// A *Response is returned with the configured status code and content type from the spec.
func ListFoodsJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  500,
		contentType: "application/json",
	}
}

// LoginJSON200Response is a constructor method for a Login response.
// A *Response is returned with the configured status code and content type from the spec.
func LoginJSON200Response(body Session) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// LoginJSON401Response is a constructor method for a Login response.
// A *Response is returned with the configured status code and content type from the spec.
func LoginJSON401Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  401,
		contentType: "application/json",
	}
}

// LoginJSON500Response is a constructor method for a Login response.
// A *Response is returned with the configured status code and content type from the spec.
func LoginJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  500,
		contentType: "application/json",
	}
}

// GetNextPostsJSON200Response is a constructor method for a GetNextPosts response.
// A *Response is returned with the configured status code and content type from the spec.
func GetNextPostsJSON200Response(body []Post) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// GetNextPostsJSON400Response is a constructor method for a GetNextPosts response.
// A *Response is returned with the configured status code and content type from the spec.
func GetNextPostsJSON400Response(body FormError) *Response {
	return &Response{
		body:        body,
		statusCode:  400,
		contentType: "application/json",
	}
}

// GetNextPostsJSON500Response is a constructor method for a GetNextPosts response.
// A *Response is returned with the configured status code and content type from the spec.
func GetNextPostsJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  500,
		contentType: "application/json",
	}
}

// CreatePostJSON200Response is a constructor method for a CreatePost response.
// A *Response is returned with the configured status code and content type from the spec.
func CreatePostJSON200Response(body Post) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// CreatePostJSON400Response is a constructor method for a CreatePost response.
// A *Response is returned with the configured status code and content type from the spec.
func CreatePostJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  400,
		contentType: "application/json",
	}
}

// CreatePostJSON500Response is a constructor method for a CreatePost response.
// A *Response is returned with the configured status code and content type from the spec.
func CreatePostJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  500,
		contentType: "application/json",
	}
}

// GetLikedPostsJSON200Response is a constructor method for a GetLikedPosts response.
// A *Response is returned with the configured status code and content type from the spec.
func GetLikedPostsJSON200Response(body []Post) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// GetLikedPostsJSON500Response is a constructor method for a GetLikedPosts response.
// A *Response is returned with the configured status code and content type from the spec.
func GetLikedPostsJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  500,
		contentType: "application/json",
	}
}

// DeletePostJSON404Response is a constructor method for a DeletePost response.
// A *Response is returned with the configured status code and content type from the spec.
func DeletePostJSON404Response(body FormError) *Response {
	return &Response{
		body:        body,
		statusCode:  404,
		contentType: "application/json",
	}
}

// DeletePostJSON500Response is a constructor method for a DeletePost response.
// A *Response is returned with the configured status code and content type from the spec.
func DeletePostJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  500,
		contentType: "application/json",
	}
}

// RegisterJSON200Response is a constructor method for a Register response.
// A *Response is returned with the configured status code and content type from the spec.
func RegisterJSON200Response(body Session) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// RegisterJSON400Response is a constructor method for a Register response.
// A *Response is returned with the configured status code and content type from the spec.
func RegisterJSON400Response(body FormError) *Response {
	return &Response{
		body:        body,
		statusCode:  400,
		contentType: "application/json",
	}
}

// RegisterJSON500Response is a constructor method for a Register response.
// A *Response is returned with the configured status code and content type from the spec.
func RegisterJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  500,
		contentType: "application/json",
	}
}

// GetSelfJSON200Response is a constructor method for a GetSelf response.
// A *Response is returned with the configured status code and content type from the spec.
func GetSelfJSON200Response(body Self) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// GetSelfJSON500Response is a constructor method for a GetSelf response.
// A *Response is returned with the configured status code and content type from the spec.
func GetSelfJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  500,
		contentType: "application/json",
	}
}

// GetUserJSON200Response is a constructor method for a GetUser response.
// A *Response is returned with the configured status code and content type from the spec.
func GetUserJSON200Response(body User) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// GetUserJSON400Response is a constructor method for a GetUser response.
// A *Response is returned with the configured status code and content type from the spec.
func GetUserJSON400Response(body FormError) *Response {
	return &Response{
		body:        body,
		statusCode:  400,
		contentType: "application/json",
	}
}

// GetUserJSON404Response is a constructor method for a GetUser response.
// A *Response is returned with the configured status code and content type from the spec.
func GetUserJSON404Response(body FormError) *Response {
	return &Response{
		body:        body,
		statusCode:  404,
		contentType: "application/json",
	}
}

// GetUserJSON500Response is a constructor method for a GetUser response.
// A *Response is returned with the configured status code and content type from the spec.
func GetUserJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		statusCode:  500,
		contentType: "application/json",
	}
}

// Getter for additional properties for FoodCategories. Returns the specified
// element and whether it was found
func (a FoodCategories) Get(fieldName string) (value []string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for FoodCategories
func (a *FoodCategories) Set(fieldName string, value []string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string][]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for FoodCategories to handle AdditionalProperties
func (a *FoodCategories) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string][]string)
		for fieldName, fieldBuf := range object {
			var fieldVal []string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for FoodCategories to handle AdditionalProperties
func (a FoodCategories) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Upload an asset
	// (POST /assets)
	UploadAsset(w http.ResponseWriter, r *http.Request) *Response
	// Get the file asset by the given ID. Note that assets are not separated by type; the user must assume the type from the context that the asset ID is from.
	// (GET /assets/{id})
	GetAsset(w http.ResponseWriter, r *http.Request, id string) *Response
	// Get the list of all valid food categories and names.
	// (GET /food/list)
	ListFoods(w http.ResponseWriter, r *http.Request) *Response
	// Log in using username and password. A 401 is returned if the information is incorrect.
	// (POST /login)
	Login(w http.ResponseWriter, r *http.Request, params LoginParams) *Response
	// Get the next batch of posts
	// (GET /posts)
	GetNextPosts(w http.ResponseWriter, r *http.Request, params GetNextPostsParams) *Response
	// Create a new post
	// (POST /posts)
	CreatePost(w http.ResponseWriter, r *http.Request) *Response
	// Get the list of posts liked by the user
	// (GET /posts/liked)
	GetLikedPosts(w http.ResponseWriter, r *http.Request) *Response
	// Delete the current user's posts by ID. A 401 is returned if the user tries to delete someone else's post.
	// (DELETE /posts/{id})
	DeletePost(w http.ResponseWriter, r *http.Request, id ID) *Response
	// Register using username and password
	// (POST /register)
	Register(w http.ResponseWriter, r *http.Request, params RegisterParams) *Response
	// Get the current user
	// (GET /users/@self)
	GetSelf(w http.ResponseWriter, r *http.Request) *Response
	// Get a user by their username
	// (GET /users/{username})
	GetUser(w http.ResponseWriter, r *http.Request, username string) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	Middlewares      map[string]func(http.Handler) http.Handler
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// UploadAsset operation middleware
func (siw *ServerInterfaceWrapper) UploadAsset(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.UploadAsset(w, r)
		if resp != nil {
			render.Render(w, r, resp)
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetAsset operation middleware
func (siw *ServerInterfaceWrapper) GetAsset(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetAsset(w, r, id)
		if resp != nil {
			render.Render(w, r, resp)
		}
	})

	handler(w, r.WithContext(ctx))
}

// ListFoods operation middleware
func (siw *ServerInterfaceWrapper) ListFoods(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.ListFoods(w, r)
		if resp != nil {
			render.Render(w, r, resp)
		}
	})

	handler(w, r.WithContext(ctx))
}

// Login operation middleware
func (siw *ServerInterfaceWrapper) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params LoginParams

	// ------------- Required query parameter "username" -------------

	if err := runtime.BindQueryParameter("form", true, true, "username", r.URL.Query(), &params.Username); err != nil {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "username"})
		return
	}

	// ------------- Required query parameter "password" -------------

	if err := runtime.BindQueryParameter("form", true, true, "password", r.URL.Query(), &params.Password); err != nil {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "password"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.Login(w, r, params)
		if resp != nil {
			render.Render(w, r, resp)
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetNextPosts operation middleware
func (siw *ServerInterfaceWrapper) GetNextPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNextPostsParams

	// ------------- Optional query parameter "prev_id" -------------

	if err := runtime.BindQueryParameter("form", true, false, "prev_id", r.URL.Query(), &params.PrevID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "prev_id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetNextPosts(w, r, params)
		if resp != nil {
			render.Render(w, r, resp)
		}
	})

	handler(w, r.WithContext(ctx))
}

// CreatePost operation middleware
func (siw *ServerInterfaceWrapper) CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.CreatePost(w, r)
		if resp != nil {
			render.Render(w, r, resp)
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetLikedPosts operation middleware
func (siw *ServerInterfaceWrapper) GetLikedPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetLikedPosts(w, r)
		if resp != nil {
			render.Render(w, r, resp)
		}
	})

	handler(w, r.WithContext(ctx))
}

// DeletePost operation middleware
func (siw *ServerInterfaceWrapper) DeletePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id ID

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.DeletePost(w, r, id)
		if resp != nil {
			render.Render(w, r, resp)
		}
	})

	handler(w, r.WithContext(ctx))
}

// Register operation middleware
func (siw *ServerInterfaceWrapper) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params RegisterParams

	// ------------- Required query parameter "username" -------------

	if err := runtime.BindQueryParameter("form", true, true, "username", r.URL.Query(), &params.Username); err != nil {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "username"})
		return
	}

	// ------------- Required query parameter "password" -------------

	if err := runtime.BindQueryParameter("form", true, true, "password", r.URL.Query(), &params.Password); err != nil {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "password"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.Register(w, r, params)
		if resp != nil {
			render.Render(w, r, resp)
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetSelf operation middleware
func (siw *ServerInterfaceWrapper) GetSelf(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetSelf(w, r)
		if resp != nil {
			render.Render(w, r, resp)
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "username" -------------
	var username string

	if err := runtime.BindStyledParameter("simple", false, "username", chi.URLParam(r, "username"), &username); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "username"})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetUser(w, r, username)
		if resp != nil {
			render.Render(w, r, resp)
		}
	})

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter %s: %v", err.paramName, err.err)
}

func (err UnescapedCookieParamError) Unwrap() error { return err.err }

type UnmarshalingParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnmarshalingParamError) Error() string {
	return fmt.Sprintf("error unmarshaling parameter %s as JSON: %v", err.paramName, err.err)
}

func (err UnmarshalingParamError) Unwrap() error { return err.err }

type RequiredParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err RequiredParamError) Error() string {
	if err.err == nil {
		return fmt.Sprintf("query parameter %s is required, but not found", err.paramName)
	} else {
		return fmt.Sprintf("query parameter %s is required, but errored: %s", err.paramName, err.err)
	}
}

func (err RequiredParamError) Unwrap() error { return err.err }

type RequiredHeaderError struct {
	paramName string
}

// Error implements error.
func (err RequiredHeaderError) Error() string {
	return fmt.Sprintf("header parameter %s is required, but not found", err.paramName)
}

type InvalidParamFormatError struct {
	err       error
	paramName string
}

// Error implements error.
func (err InvalidParamFormatError) Error() string {
	return fmt.Sprintf("invalid format for parameter %s: %v", err.paramName, err.err)
}

func (err InvalidParamFormatError) Unwrap() error { return err.err }

type TooManyValuesForParamError struct {
	NumValues int
	paramName string
}

// Error implements error.
func (err TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("expected one value for %s, got %d", err.paramName, err.NumValues)
}

// ParameterName is an interface that is implemented by error types that are
// relevant to a specific parameter.
type ParameterError interface {
	error
	// ParamName is the name of the parameter that the error is referring to.
	ParamName() string
}

func (err UnescapedCookieParamError) ParamName() string  { return err.paramName }
func (err UnmarshalingParamError) ParamName() string     { return err.paramName }
func (err RequiredParamError) ParamName() string         { return err.paramName }
func (err RequiredHeaderError) ParamName() string        { return err.paramName }
func (err InvalidParamFormatError) ParamName() string    { return err.paramName }
func (err TooManyValuesForParamError) ParamName() string { return err.paramName }

type ServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      map[string]func(http.Handler) http.Handler
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

type ServerOption func(*ServerOptions)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface, opts ...ServerOption) http.Handler {
	options := &ServerOptions{
		BaseURL:     "/",
		BaseRouter:  chi.NewRouter(),
		Middlewares: make(map[string]func(http.Handler) http.Handler),
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
	}

	for _, f := range opts {
		f(options)
	}

	r := options.BaseRouter
	wrapper := ServerInterfaceWrapper{
		Handler:          si,
		Middlewares:      options.Middlewares,
		ErrorHandlerFunc: options.ErrorHandlerFunc,
	}

	r.Route(options.BaseURL, func(r chi.Router) {
		r.Post("/assets", wrapper.UploadAsset)
		r.Get("/assets/{id}", wrapper.GetAsset)
		r.Get("/food/list", wrapper.ListFoods)
		r.Post("/login", wrapper.Login)
		r.Get("/posts", wrapper.GetNextPosts)
		r.Post("/posts", wrapper.CreatePost)
		r.Get("/posts/liked", wrapper.GetLikedPosts)
		r.Delete("/posts/{id}", wrapper.DeletePost)
		r.Post("/register", wrapper.Register)
		r.Get("/users/@self", wrapper.GetSelf)
		r.Get("/users/{username}", wrapper.GetUser)

	})
	return r
}

func WithRouter(r chi.Router) ServerOption {
	return func(s *ServerOptions) {
		s.BaseRouter = r
	}
}

func WithServerBaseURL(url string) ServerOption {
	return func(s *ServerOptions) {
		s.BaseURL = url
	}
}

func WithMiddleware(key string, middleware func(http.Handler) http.Handler) ServerOption {
	return func(s *ServerOptions) {
		s.Middlewares[key] = middleware
	}
}

func WithMiddlewares(middlewares map[string]func(http.Handler) http.Handler) ServerOption {
	return func(s *ServerOptions) {
		s.Middlewares = middlewares
	}
}

func WithErrorHandler(handler func(w http.ResponseWriter, r *http.Request, err error)) ServerOption {
	return func(s *ServerOptions) {
		s.ErrorHandlerFunc = handler
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
