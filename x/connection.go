/*
 * Copyright 2019 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package x

import (
	"strings"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var tlsDir = "./tls"

type TlsAuthLevel int

const (
	TlsAuthNone   TlsAuthLevel = iota // No authentication.
	TlsAuthServer                     // Client authenticates server only (similar to https).
	TlsAuthMutual                     // Client and server authenticate each other.
)

var TlsAuthLevelStr = map[string]TlsAuthLevel{
	"none":   TlsAuthNone,
	"server": TlsAuthServer,
	"mutual": TlsAuthMutual,
}

type TlsConfig struct {
	AuthLevel TlsAuthLevel
	CertDir   string
}

func AddClientTlsOptions(flag *pflag.FlagSet) {

}

func AddServerTlsOptions(flag *pflag.FlagSet) {
	flag.String("tls_auth", "none",
		"Required authentication level. One of: none, server, or mutual")
	flag.String("tls_dir", tlsDir,
		"Path to directory containing keys and certificates.")
}

func CreateConnConf(v *viper.Viper) TlsConfig {
	conf := TlsConfig{}

	return conf
}

func ConfigureConnection(v *viper.Viper) TlsConfig {
	conf := TlsConfig{}

	switch auth := strings.ToLower(v.GetString("tls_auth")); auth {
	case "none":
		glog.V(2).Info("not using TLS authentication")
	case "server":
		glog.V(2).Info("using server authentication only")
	case "mutual":
		glog.V(2).Info("using client and server authentication")
	}

	return conf
}

//func GrpcConnect(addr string, conf *ConnConf) (*grpc.ClientConn, error) {
//	callOpts := append([]grpc.CallOption{},
//		grpc.MaxCallRecvMsgSize(GrpcMaxSize),
//		grpc.MaxCallSendMsgSize(GrpcMaxSize))
//
//}

/*
func SetupConnection(host string, tlsCfg *tls.Config, useGz bool) (*grpc.ClientConn, error) {
	callOpts := append([]grpc.CallOption{},
		grpc.MaxCallRecvMsgSize(GrpcMaxSize),
		grpc.MaxCallSendMsgSize(GrpcMaxSize))

	if useGz {
		fmt.Fprintf(os.Stderr, "Using compression with %s\n", host)
		callOpts = append(callOpts, grpc.UseCompressor(gzip.Name))
	}

	dialOpts := append([]grpc.DialOption{},
		grpc.WithDefaultCallOptions(callOpts...),
		grpc.WithBlock())

	if tlsCfg != nil {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg)))
	} else {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}

	_, err := net.Dial("tcp", host)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			log.Printf("waiting for a listener to start on %s", host)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return grpc.DialContext(ctx, host, dialOpts...)
}
*/