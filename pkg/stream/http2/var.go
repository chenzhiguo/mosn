/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package http2

import (
	"context"
	"fmt"

	"mosn.io/api"
	mosnctx "mosn.io/mosn/pkg/context"
	"mosn.io/mosn/pkg/protocol"
	"mosn.io/mosn/pkg/types"
	"mosn.io/mosn/pkg/variable"
)

var (
	headerName  = fmt.Sprintf("%s_%s", protocol.HTTP2, types.VarProtocolRequestHeader)
	headerIndex = len(headerName)
)

func init() {
	variable.RegisterPrefixVariable(headerName,
		variable.NewBasicVariable(headerName, nil, headerGetter, nil, 0))

	variable.RegisterProtocolResource(protocol.HTTP2, api.HEADER, types.VarProtocolRequestHeader)
}

func headerGetter(ctx context.Context, value *variable.IndexedValue, data interface{}) (s string, err error) {
	headers, ok := mosnctx.Get(ctx, types.ContextKeyDownStreamHeaders).(api.HeaderMap)
	if !ok {
		return variable.ValueNotFound, nil
	}
	headerKey, ok := data.(string)
	if !ok {
		return variable.ValueNotFound, nil
	}

	header, found := headers.Get(headerKey[headerIndex:])
	if !found {
		return variable.ValueNotFound, nil
	}

	return header, nil
}
