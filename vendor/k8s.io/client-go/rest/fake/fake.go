/*
Copyright 2014 The Kubernetes Authors.

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

// This is made a separate package and should only be imported by tests, because
// it imports testapi
package fake

import (
	"net/http"
	"net/url"

<<<<<<< HEAD
<<<<<<< HEAD
=======
	"k8s.io/apimachinery/pkg/apimachinery/registered"
>>>>>>> Initial dep workover
=======
>>>>>>> omg dep constraints
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
)

func CreateHTTPClient(roundTripper func(*http.Request) (*http.Response, error)) *http.Client {
	return &http.Client{
		Transport: roundTripperFunc(roundTripper),
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// RESTClient provides a fake RESTClient interface.
type RESTClient struct {
	Client               *http.Client
	NegotiatedSerializer runtime.NegotiatedSerializer
<<<<<<< HEAD
<<<<<<< HEAD
	GroupVersion         schema.GroupVersion
=======
	GroupName            string
	APIRegistry          *registered.APIRegistrationManager
>>>>>>> Initial dep workover
=======
	GroupVersion         schema.GroupVersion
>>>>>>> omg dep constraints
	VersionedAPIPath     string

	Req  *http.Request
	Resp *http.Response
	Err  error
}

func (c *RESTClient) Get() *restclient.Request {
	return c.request("GET")
}

func (c *RESTClient) Put() *restclient.Request {
	return c.request("PUT")
}

func (c *RESTClient) Patch(pt types.PatchType) *restclient.Request {
	return c.request("PATCH").SetHeader("Content-Type", string(pt))
}

func (c *RESTClient) Post() *restclient.Request {
	return c.request("POST")
}

func (c *RESTClient) Delete() *restclient.Request {
	return c.request("DELETE")
}

func (c *RESTClient) Verb(verb string) *restclient.Request {
	return c.request(verb)
}

func (c *RESTClient) APIVersion() schema.GroupVersion {
<<<<<<< HEAD
<<<<<<< HEAD
	return c.GroupVersion
=======
	return c.APIRegistry.GroupOrDie("").GroupVersion
>>>>>>> Initial dep workover
=======
	return c.GroupVersion
>>>>>>> omg dep constraints
}

func (c *RESTClient) GetRateLimiter() flowcontrol.RateLimiter {
	return nil
}

func (c *RESTClient) request(verb string) *restclient.Request {
	config := restclient.ContentConfig{
<<<<<<< HEAD
<<<<<<< HEAD
		ContentType:          runtime.ContentTypeJSON,
		GroupVersion:         &c.GroupVersion,
=======
		ContentType: runtime.ContentTypeJSON,
		// TODO this was hardcoded before, but it doesn't look right
		GroupVersion:         &c.APIRegistry.GroupOrDie("").GroupVersion,
>>>>>>> Initial dep workover
=======
		ContentType:          runtime.ContentTypeJSON,
		GroupVersion:         &c.GroupVersion,
>>>>>>> omg dep constraints
		NegotiatedSerializer: c.NegotiatedSerializer,
	}

	ns := c.NegotiatedSerializer
	info, _ := runtime.SerializerInfoForMediaType(ns.SupportedMediaTypes(), runtime.ContentTypeJSON)
	internalVersion := schema.GroupVersion{
<<<<<<< HEAD
<<<<<<< HEAD
		Group:   c.GroupVersion.Group,
		Version: runtime.APIVersionInternal,
	}
	serializers := restclient.Serializers{
		// TODO this was hardcoded before, but it doesn't look right
		Encoder: ns.EncoderForVersion(info.Serializer, c.GroupVersion),
=======
		Group:   c.APIRegistry.GroupOrDie(c.GroupName).GroupVersion.Group,
=======
		Group:   c.GroupVersion.Group,
>>>>>>> omg dep constraints
		Version: runtime.APIVersionInternal,
	}
	serializers := restclient.Serializers{
		// TODO this was hardcoded before, but it doesn't look right
<<<<<<< HEAD
		Encoder: ns.EncoderForVersion(info.Serializer, c.APIRegistry.GroupOrDie("").GroupVersion),
>>>>>>> Initial dep workover
=======
		Encoder: ns.EncoderForVersion(info.Serializer, c.GroupVersion),
>>>>>>> omg dep constraints
		Decoder: ns.DecoderToVersion(info.Serializer, internalVersion),
	}
	if info.StreamSerializer != nil {
		serializers.StreamingSerializer = info.StreamSerializer.Serializer
		serializers.Framer = info.StreamSerializer.Framer
	}
	return restclient.NewRequest(c, verb, &url.URL{Host: "localhost"}, c.VersionedAPIPath, config, serializers, nil, nil)
}

func (c *RESTClient) Do(req *http.Request) (*http.Response, error) {
	if c.Err != nil {
		return nil, c.Err
	}
	c.Req = req
	if c.Client != nil {
		return c.Client.Do(req)
	}
	return c.Resp, nil
}