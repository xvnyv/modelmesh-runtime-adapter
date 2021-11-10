// Copyright 2021 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: ./provider.go
package s3provider

import (
	context "context"
	reflect "reflect"

	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	pullman "github.com/kserve/modelmesh-runtime-adapter/pullman"
)

// Mocks3DownloaderFactory is a mock of s3DownloaderFactory interface.
type Mocks3DownloaderFactory struct {
	ctrl     *gomock.Controller
	recorder *Mocks3DownloaderFactoryMockRecorder
}

// Mocks3DownloaderFactoryMockRecorder is the mock recorder for Mocks3DownloaderFactory.
type Mocks3DownloaderFactoryMockRecorder struct {
	mock *Mocks3DownloaderFactory
}

// NewMocks3DownloaderFactory creates a new mock instance.
func NewMocks3DownloaderFactory(ctrl *gomock.Controller) *Mocks3DownloaderFactory {
	mock := &Mocks3DownloaderFactory{ctrl: ctrl}
	mock.recorder = &Mocks3DownloaderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocks3DownloaderFactory) EXPECT() *Mocks3DownloaderFactoryMockRecorder {
	return m.recorder
}

// newDownloader mocks base method.
func (m *Mocks3DownloaderFactory) newDownloader(log logr.Logger, accessKeyID, secretAccessKey, endpoint, region, certificate string) s3Downloader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "newDownloader", log, accessKeyID, secretAccessKey, endpoint, region, certificate)
	ret0, _ := ret[0].(s3Downloader)
	return ret0
}

// newDownloader indicates an expected call of newDownloader.
func (mr *Mocks3DownloaderFactoryMockRecorder) newDownloader(log, accessKeyID, secretAccessKey, endpoint, region, certificate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "newDownloader", reflect.TypeOf((*Mocks3DownloaderFactory)(nil).newDownloader), log, accessKeyID, secretAccessKey, endpoint, region, certificate)
}

// Mocks3Downloader is a mock of s3Downloader interface.
type Mocks3Downloader struct {
	ctrl     *gomock.Controller
	recorder *Mocks3DownloaderMockRecorder
}

// Mocks3DownloaderMockRecorder is the mock recorder for Mocks3Downloader.
type Mocks3DownloaderMockRecorder struct {
	mock *Mocks3Downloader
}

// NewMocks3Downloader creates a new mock instance.
func NewMocks3Downloader(ctrl *gomock.Controller) *Mocks3Downloader {
	mock := &Mocks3Downloader{ctrl: ctrl}
	mock.recorder = &Mocks3DownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocks3Downloader) EXPECT() *Mocks3DownloaderMockRecorder {
	return m.recorder
}

// downloadBatch mocks base method.
func (m *Mocks3Downloader) downloadBatch(ctx context.Context, bucket string, targets []pullman.Target) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "downloadBatch", ctx, bucket, targets)
	ret0, _ := ret[0].(error)
	return ret0
}

// downloadBatch indicates an expected call of downloadBatch.
func (mr *Mocks3DownloaderMockRecorder) downloadBatch(ctx, bucket, targets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "downloadBatch", reflect.TypeOf((*Mocks3Downloader)(nil).downloadBatch), ctx, bucket, targets)
}

// listObjects mocks base method.
func (m *Mocks3Downloader) listObjects(bucket, prefix string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "listObjects", bucket, prefix)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// listObjects indicates an expected call of listObjects.
func (mr *Mocks3DownloaderMockRecorder) listObjects(bucket, prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "listObjects", reflect.TypeOf((*Mocks3Downloader)(nil).listObjects), bucket, prefix)
}
