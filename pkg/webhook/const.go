/*
Copyright 2023 The KubeService-Stack Authors.

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

package webhook

import (
	"errors"
)

const (
	WebhookName    = "CosignWebhook"
	WebhookEnable  = "enabled"
	WebhookDisable = "disabled"
	WebhookVersion = "v1"
)

var (
	ErrInvalidCosignCRDMoreThanOne = errors.New("invalid cosignkey CRD more than one in this namespace")
	ErrInvalidAdmissionReview      = errors.New("invalid admission review error")
	ErrInvalidAdmissionReviewObj   = errors.New("invalid admission review object error")
	ErrMissingCosignCRD            = errors.New("invalid cosignkey CRD not find in cluster")
	ErrMissingCosignCRDKeys        = errors.New("invalid cosignkey CRD key counter iz zero in cluster")
	ErrInvalidCosignVerify         = errors.New("invalid cosign verify cosignkey key error")
)
