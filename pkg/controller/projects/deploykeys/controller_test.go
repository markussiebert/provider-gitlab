/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package deploykeys

import (
	"context"
	"net/http"
	"strconv"
	"testing"
	"time"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	gitlab "github.com/xanzy/go-gitlab"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-gitlab/apis/projects/v1alpha1"
	"github.com/crossplane-contrib/provider-gitlab/pkg/clients/projects"
	"github.com/crossplane-contrib/provider-gitlab/pkg/clients/projects/fake"
)

var (
	errorMessage           = "restult: -expected, +actual: \n%s"
	notADeployKey          resource.Managed
	testProjectID          = "testProjectId"
	testKeyID              = 123
	testKeyTitle           = "testKeyTitle"
	testKey                = "testKey"
	testCreatedAt          = time.Now()
	testExternalName       = "123"
	testGetKeyErrorMessage = "testGetKeyError"
	testCanPush            = true

	testDeployKey = gitlab.ProjectDeployKey{
		ID:        testKeyID,
		Title:     testKeyTitle,
		Key:       testKey,
		CreatedAt: &testCreatedAt,
		CanPush:   true,
	}

	testDeployKeyNoProjectID = &v1alpha1.DeployKey{}
)

type args struct {
	deployKeyService projects.DeployKeyClient
	kube             client.Client
	cr               resource.Managed
}

type deployKeyModifier func(*v1alpha1.DeployKey)

func withExternalName(id string) deployKeyModifier {
	return func(dk *v1alpha1.DeployKey) { meta.SetExternalName(dk, id) }
}

func withCanPush() deployKeyModifier {
	return func(dk *v1alpha1.DeployKey) { dk.Spec.ForProvider.CanPush = &testCanPush }
}

func withTitle() deployKeyModifier {
	return func(dk *v1alpha1.DeployKey) { dk.Spec.ForProvider.Title = testKeyTitle }
}

func withConditions(conditions ...xpv1.Condition) deployKeyModifier {
	return func(dk *v1alpha1.DeployKey) { dk.Status.ConditionedStatus.Conditions = conditions }
}

func withTestKeyRef() deployKeyModifier {
	return func(dk *v1alpha1.DeployKey) {
		dk.Spec.ForProvider.KeySecretRef.Name = "testName"
		dk.Spec.ForProvider.KeySecretRef.Namespace = "testNameSpace"
		dk.Spec.ForProvider.KeySecretRef.Key = "testKey"
	}
}

func withID() deployKeyModifier {
	return func(dk *v1alpha1.DeployKey) { dk.Status.AtProvider.ID = &testKeyID }
}

func withCreatedAt() deployKeyModifier {
	return func(dk *v1alpha1.DeployKey) { dk.Status.AtProvider.CreatedAt = &metav1.Time{Time: testCreatedAt} }
}

func buildDeployKey(modifiers ...deployKeyModifier) *v1alpha1.DeployKey {
	deployKey := &v1alpha1.DeployKey{} // why to use `&`?
	for _, modifier := range modifiers {
		modifier(deployKey)
	}

	deployKey.Spec.ForProvider.ProjectID = &testProjectID

	return deployKey
}

func TestObserve(t *testing.T) {
	type expected struct {
		dk     resource.Managed
		result managed.ExternalObservation
		err    error
	}

	testCases := map[string]struct {
		args
		expected
	}{
		"NotADeployKey": {
			args: args{
				cr: notADeployKey,
			},
			expected: expected{
				dk:  notADeployKey,
				err: errors.New(errNotDeployKey),
			},
		},
		"ProjectIDNotSet": {
			args: args{
				cr: testDeployKeyNoProjectID,
			},
			expected: expected{
				dk:     testDeployKeyNoProjectID,
				result: managed.ExternalObservation{},
				err:    errors.New(errProjectIDMissing),
			},
		},
		"NoExternalNameSet": {
			args: args{
				cr: buildDeployKey(),
			},
			expected: expected{
				dk:     buildDeployKey(),
				err:    errors.New(errIDNotAnInt),
				result: managed.ExternalObservation{},
			},
		},
		"ExternalNameNotAnInt": {
			args: args{
				cr: buildDeployKey(withExternalName("notAnInt")),
			},
			expected: expected{
				dk:     buildDeployKey(withExternalName("notAnInt")),
				result: managed.ExternalObservation{},
				err:    errors.New(errIDNotAnInt),
			},
		},
		"GetKeyClientError": {
			args: args{
				cr: buildDeployKey(withExternalName(testExternalName)),
				deployKeyService: &fake.MockClient{
					MockGetDeployKey: func(pid interface{}, deployKey int, options ...*gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
						return nil, &gitlab.Response{Response: &http.Response{StatusCode: 400}}, errors.New(testGetKeyErrorMessage)
					},
				},
			},
			expected: expected{
				dk:     buildDeployKey(withExternalName(testExternalName)),
				err:    errors.Wrap(errors.New(testGetKeyErrorMessage), errGetFail),
				result: managed.ExternalObservation{},
			},
		},
		"GetErr404": {
			args: args{
				cr: buildDeployKey(withExternalName(testExternalName)),
				deployKeyService: &fake.MockClient{
					MockGetDeployKey: func(pid interface{}, deployKey int, options ...*gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
						return nil, &gitlab.Response{Response: &http.Response{StatusCode: 404}}, errors.New("")
					},
				},
			},
			expected: expected{
				dk:     buildDeployKey(withExternalName(testExternalName)),
				result: managed.ExternalObservation{},
				err:    nil,
			},
		},
		"SuccessLateInitTrueUpToDateFalse": {
			args: args{
				cr: buildDeployKey(
					withExternalName(testExternalName),
				),
				deployKeyService: &fake.MockClient{
					MockGetDeployKey: func(pid interface{}, deployKey int, options ...*gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
						return &testDeployKey, &gitlab.Response{}, nil
					},
				},
			},
			expected: expected{
				dk: buildDeployKey(
					withExternalName(testExternalName),
					withConditions(xpv1.Available()),
					withCanPush(),
					withID(),
					withCreatedAt(),
				),
				err: nil,
				result: managed.ExternalObservation{
					ResourceExists:          true,
					ResourceUpToDate:        false,
					ResourceLateInitialized: true,
				},
			},
		},
		"SuccessLateInitFalseUpToDateTrue": {
			args: args{
				cr: buildDeployKey(
					withExternalName(testExternalName),
					withCanPush(),
					withTitle(),
				),
				deployKeyService: &fake.MockClient{
					MockGetDeployKey: func(pid interface{}, deployKey int, options ...*gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
						return &testDeployKey, &gitlab.Response{}, nil
					},
				},
			},
			expected: expected{
				dk: buildDeployKey(
					withExternalName(testExternalName),
					withCanPush(),
					withTitle(),
					withConditions(xpv1.Available()),
					withID(),
					withCreatedAt(),
				),
				err: nil,
				result: managed.ExternalObservation{
					ResourceExists:          true,
					ResourceUpToDate:        true,
					ResourceLateInitialized: false,
				},
			},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			victim := &external{kube: testCase.kube, client: testCase.deployKeyService}
			result, err := victim.Observe(context.Background(), testCase.args.cr)

			if diff := cmp.Diff(testCase.expected.err, err, test.EquateErrors()); diff != "" {
				t.Errorf(errorMessage, diff)
			}

			if diff := cmp.Diff(testCase.expected.dk, testCase.args.cr, test.EquateConditions()); diff != "" {
				t.Errorf(errorMessage, diff)
			}

			if diff := cmp.Diff(testCase.expected.result, result); diff != "" {
				t.Errorf(errorMessage, diff)
			}

		})
	}
}

func TestCreate(t *testing.T) {
	type expected struct {
		dk     resource.Managed
		result managed.ExternalCreation
		err    error
	}

	testCases := map[string]struct {
		args
		expected
	}{
		"NotADeployKey": {
			args: args{
				cr: notADeployKey,
			},
			expected: expected{
				dk:     notADeployKey,
				result: managed.ExternalCreation{},
				err:    errors.New(errNotDeployKey),
			},
		},
		"ProjectIDNotSet": {
			args: args{
				cr: testDeployKeyNoProjectID,
			},
			expected: expected{
				dk:     testDeployKeyNoProjectID,
				result: managed.ExternalCreation{},
				err:    errors.New(errProjectIDMissing),
			},
		},
		"NoKeySecretRef": {
			args: args{
				cr: buildDeployKey(),
				kube: &test.MockClient{
					MockGet: test.NewMockGetFn(errors.New("")),
				},
			},
			expected: expected{
				dk:  buildDeployKey(),
				err: errors.Wrap(errors.New(""), errKeyMissing),
			},
		},
		"FaileToAdd": {
			args: args{
				cr:   buildDeployKey(withTestKeyRef()),
				kube: &test.MockClient{MockGet: test.NewMockGetFn(nil)},
				deployKeyService: &fake.MockClient{
					MockAddDeployKey: func(pid interface{}, opt *gitlab.AddDeployKeyOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
						return nil, nil, testError()
					},
				},
			},
			expected: expected{
				dk:     buildDeployKey(withTestKeyRef()),
				err:    errors.Wrap(testError(), errCreateFail),
				result: managed.ExternalCreation{},
			},
		},
		"SuccessfullyAdd": {
			args: args{
				cr:   buildDeployKey(withTestKeyRef()),
				kube: &test.MockClient{MockGet: test.NewMockGetFn(nil)},
				deployKeyService: &fake.MockClient{
					MockAddDeployKey: func(pid interface{}, opt *gitlab.AddDeployKeyOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
						return &gitlab.ProjectDeployKey{ID: testKeyID}, nil, nil
					},
				},
			},
			expected: expected{
				err: nil,
				dk: buildDeployKey(
					withTestKeyRef(),
					withExternalName(testExternalName),
				),
				result: managed.ExternalCreation{
					ConnectionDetails: nil,
				},
			},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			victim := &external{kube: testCase.kube, client: testCase.deployKeyService}
			result, err := victim.Create(context.Background(), testCase.args.cr)

			if diff := cmp.Diff(testCase.expected.err, err, test.EquateErrors()); diff != "" {
				t.Errorf(errorMessage, diff)
			}

			if diff := cmp.Diff(testCase.expected.dk, testCase.args.cr, test.EquateConditions()); diff != "" {
				t.Errorf(errorMessage, diff)
			}

			if diff := cmp.Diff(testCase.expected.result, result); diff != "" {
				t.Errorf(errorMessage, diff)
			}

		})
	}
}

func TestUpdate(t *testing.T) {
	type expected struct {
		dk     resource.Managed
		result managed.ExternalUpdate
		err    error
	}

	testCases := map[string]struct {
		args
		expected
	}{
		"NotADeployKey": {
			args: args{
				cr: notADeployKey,
			},
			expected: expected{
				dk:  notADeployKey,
				err: errors.New(errNotDeployKey),
			},
		},
		"ProjectIDNotSet": {
			args: args{
				cr: testDeployKeyNoProjectID,
			},
			expected: expected{
				dk:     testDeployKeyNoProjectID,
				result: managed.ExternalUpdate{},
				err:    errors.New(errProjectIDMissing),
			},
		},
		"IdIsNotAnInt": {
			args: args{
				cr: buildDeployKey(withExternalName("123A")),
			},
			expected: expected{
				dk:     buildDeployKey(withExternalName("123A")),
				err:    errors.Wrap(aToIError(), errIDNotAnInt),
				result: managed.ExternalUpdate{},
			},
		},
		"FailUpdate": {
			args: args{
				cr: buildDeployKey(
					withExternalName(testExternalName),
					withID(),
					withCanPush(),
					withTitle(),
				),
				deployKeyService: &fake.MockClient{
					MockUpdateDeployKey: func(pid interface{}, deployKey int, opt *gitlab.UpdateDeployKeyOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
						return &gitlab.ProjectDeployKey{}, nil, testError()
					},
				},
			},
			expected: expected{
				dk: buildDeployKey(
					withExternalName(testExternalName),
					withID(),
					withCanPush(),
					withTitle(),
				),
				result: managed.ExternalUpdate{},
				err:    errors.Wrap(testError(), errUpdateFail),
			},
		},
		"SuccessUpdate": {
			args: args{
				cr: buildDeployKey(
					withExternalName(testExternalName),
					withID(),
					withCanPush(),
					withTitle(),
				),
				deployKeyService: &fake.MockClient{
					MockUpdateDeployKey: func(pid interface{}, deployKey int, opt *gitlab.UpdateDeployKeyOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
						return &gitlab.ProjectDeployKey{}, nil, nil
					},
				},
			},
			expected: expected{
				dk: buildDeployKey(
					withExternalName(testExternalName),
					withID(),
					withCanPush(),
					withTitle(),
				),
				result: managed.ExternalUpdate{},
				err:    nil,
			},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			victim := &external{kube: testCase.kube, client: testCase.deployKeyService}
			result, err := victim.Update(context.Background(), testCase.args.cr)

			if diff := cmp.Diff(testCase.expected.err, err, test.EquateErrors()); diff != "" {
				t.Errorf(errorMessage, diff)
			}

			if diff := cmp.Diff(testCase.expected.dk, testCase.args.cr, test.EquateConditions()); diff != "" {
				t.Errorf(errorMessage, diff)
			}

			if diff := cmp.Diff(testCase.expected.result, result); diff != "" {
				t.Errorf(errorMessage, diff)
			}

		})
	}
}

func TestDelete(t *testing.T) {
	type expected struct {
		dk  resource.Managed
		err error
	}

	testCases := map[string]struct {
		args
		expected
	}{
		"NotADeployKey": {
			args: args{
				cr: notADeployKey,
			},
			expected: expected{
				dk:  notADeployKey,
				err: errors.New(errDeleteFail),
			},
		},
		"ProjectIDNotSet": {
			args: args{
				cr: testDeployKeyNoProjectID,
			},
			expected: expected{
				dk:  testDeployKeyNoProjectID,
				err: errors.New(errProjectIDMissing),
			},
		},
		"ExternalNameNotAnInt": {
			args: args{
				cr: buildDeployKey(
					withExternalName("123A"),
				),
			},
			expected: expected{
				dk: buildDeployKey(
					withExternalName("123A"),
				),
				err: errors.Wrap(aToIError(), errIDNotAnInt),
			},
		},
		"SuccessDelete": {
			args: args{
				cr: buildDeployKey(withExternalName(testExternalName)),
				deployKeyService: &fake.MockClient{
					MockDeleteDeployKey: func(pid interface{}, deployKey int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
						return nil, nil
					},
				},
			},
			expected: expected{
				dk:  buildDeployKey(withExternalName(testExternalName)),
				err: nil,
			},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			victim := &external{kube: testCase.kube, client: testCase.deployKeyService}
			err := victim.Delete(context.Background(), testCase.args.cr)

			if diff := cmp.Diff(testCase.expected.err, err, test.EquateErrors()); diff != "" {
				t.Errorf(errorMessage, diff)
			}

			if diff := cmp.Diff(testCase.expected.dk, testCase.args.cr, test.EquateConditions()); diff != "" {
				t.Errorf(errorMessage, diff)
			}

		})
	}
}

func testError() error {
	return errors.New("error")
}

func aToIError() error {
	_, err := strconv.Atoi("123A")
	return err
}
