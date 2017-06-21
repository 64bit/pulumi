// Licensed to Pulumi Corporation ("Pulumi") under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// Pulumi licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"flag"
	"fmt"

	"github.com/pkg/errors"

	"github.com/pulumi/lumi/pkg/util/cmdutil"
	"github.com/pulumi/lumi/pkg/util/rpcutil"
	"github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
	"google.golang.org/grpc"
)

// Main is the typical entrypoint for a resource provider plugin.  Using it isn't required but can cut down
// significantly on the amount of boilerplate necessary to fire up a new resource provider.
func Main(provMaker func(*HostClient) (lumirpc.ResourceProviderServer, error)) error {
	// Initialize loggers before going any further.
	cmdutil.InitLogging(false, 0, false)

	// Read the non-flags args and connect to the engine.
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return errors.New("fatal: could not connect to host RPC; missing argument")
	}
	host, err := NewHostClient(args[0])
	if err != nil {
		return errors.Errorf("fatal: could not connect to host RPC: %v", err)
	}

	// Fire up a gRPC server, letting the kernel choose a free port for us.
	port, done, err := rpcutil.Serve(0, nil, []func(*grpc.Server) error{
		func(srv *grpc.Server) error {
			prov, err := provMaker(host)
			if err != nil {
				return fmt.Errorf("failed to create resource provider: %v", err)
			}
			lumirpc.RegisterResourceProviderServer(srv, prov)
			return nil
		},
	})
	if err != nil {
		return errors.Errorf("fatal: %v", err)
	}

	// The resource provider protocol requires that we now write out the port we have chosen to listen on.
	fmt.Printf("%d\n", port)

	// Finally, wait for the server to stop serving.
	if err := <-done; err != nil {
		return errors.Errorf("fatal: %v", err)
	}

	return nil
}
