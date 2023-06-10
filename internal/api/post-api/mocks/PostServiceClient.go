// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	ppbapi "github.com/v1tbrah/post-service/ppbapi"
)

// PostServiceClient is an autogenerated mock type for the PostServiceClient type
type PostServiceClient struct {
	mock.Mock
}

// AddHashtagToPost provides a mock function with given fields: ctx, in, opts
func (_m *PostServiceClient) AddHashtagToPost(ctx context.Context, in *ppbapi.AddHashtagToPostRequest, opts ...grpc.CallOption) (*ppbapi.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ppbapi.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.AddHashtagToPostRequest, ...grpc.CallOption) (*ppbapi.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.AddHashtagToPostRequest, ...grpc.CallOption) *ppbapi.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ppbapi.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ppbapi.AddHashtagToPostRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateHashtag provides a mock function with given fields: ctx, in, opts
func (_m *PostServiceClient) CreateHashtag(ctx context.Context, in *ppbapi.CreateHashtagRequest, opts ...grpc.CallOption) (*ppbapi.CreateHashtagResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ppbapi.CreateHashtagResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.CreateHashtagRequest, ...grpc.CallOption) (*ppbapi.CreateHashtagResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.CreateHashtagRequest, ...grpc.CallOption) *ppbapi.CreateHashtagResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ppbapi.CreateHashtagResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ppbapi.CreateHashtagRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePost provides a mock function with given fields: ctx, in, opts
func (_m *PostServiceClient) CreatePost(ctx context.Context, in *ppbapi.CreatePostRequest, opts ...grpc.CallOption) (*ppbapi.CreatePostResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ppbapi.CreatePostResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.CreatePostRequest, ...grpc.CallOption) (*ppbapi.CreatePostResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.CreatePostRequest, ...grpc.CallOption) *ppbapi.CreatePostResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ppbapi.CreatePostResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ppbapi.CreatePostRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePost provides a mock function with given fields: ctx, in, opts
func (_m *PostServiceClient) DeletePost(ctx context.Context, in *ppbapi.DeletePostRequest, opts ...grpc.CallOption) (*ppbapi.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ppbapi.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.DeletePostRequest, ...grpc.CallOption) (*ppbapi.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.DeletePostRequest, ...grpc.CallOption) *ppbapi.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ppbapi.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ppbapi.DeletePostRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHashtag provides a mock function with given fields: ctx, in, opts
func (_m *PostServiceClient) GetHashtag(ctx context.Context, in *ppbapi.GetHashtagRequest, opts ...grpc.CallOption) (*ppbapi.GetHashtagResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ppbapi.GetHashtagResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.GetHashtagRequest, ...grpc.CallOption) (*ppbapi.GetHashtagResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.GetHashtagRequest, ...grpc.CallOption) *ppbapi.GetHashtagResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ppbapi.GetHashtagResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ppbapi.GetHashtagRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPost provides a mock function with given fields: ctx, in, opts
func (_m *PostServiceClient) GetPost(ctx context.Context, in *ppbapi.GetPostRequest, opts ...grpc.CallOption) (*ppbapi.GetPostResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ppbapi.GetPostResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.GetPostRequest, ...grpc.CallOption) (*ppbapi.GetPostResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.GetPostRequest, ...grpc.CallOption) *ppbapi.GetPostResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ppbapi.GetPostResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ppbapi.GetPostRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsByHashtag provides a mock function with given fields: ctx, in, opts
func (_m *PostServiceClient) GetPostsByHashtag(ctx context.Context, in *ppbapi.GetPostsByHashtagRequest, opts ...grpc.CallOption) (*ppbapi.GetPostsByHashtagResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ppbapi.GetPostsByHashtagResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.GetPostsByHashtagRequest, ...grpc.CallOption) (*ppbapi.GetPostsByHashtagResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.GetPostsByHashtagRequest, ...grpc.CallOption) *ppbapi.GetPostsByHashtagResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ppbapi.GetPostsByHashtagResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ppbapi.GetPostsByHashtagRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsByUserID provides a mock function with given fields: ctx, in, opts
func (_m *PostServiceClient) GetPostsByUserID(ctx context.Context, in *ppbapi.GetPostsByUserIDRequest, opts ...grpc.CallOption) (*ppbapi.GetPostsByUserIDResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ppbapi.GetPostsByUserIDResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.GetPostsByUserIDRequest, ...grpc.CallOption) (*ppbapi.GetPostsByUserIDResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ppbapi.GetPostsByUserIDRequest, ...grpc.CallOption) *ppbapi.GetPostsByUserIDResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ppbapi.GetPostsByUserIDResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ppbapi.GetPostsByUserIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPostServiceClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostServiceClient creates a new instance of PostServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostServiceClient(t mockConstructorTestingTNewPostServiceClient) *PostServiceClient {
	mock := &PostServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
