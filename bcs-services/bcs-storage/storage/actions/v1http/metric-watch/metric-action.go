/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package metricWatch

import (
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/actions"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/actions/lib"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/operator"

	"github.com/emicklei/go-restful"
)

const (
	clusterIdTag    = "clusterId"
	resourceTypeTag = "type"
)

// Use Mongodb for storage.
const dbConfig = "metric"

var getNewTank operator.GetNewTank = lib.GetMongodbTank(dbConfig)

func WatchDynamic(req *restful.Request, resp *restful.Response) {
	request := newReqMetric(req, resp)
	request.watch()
}

func init() {
	actions.RegisterV1Action(actions.Action{Verb: "POST", Path: "/metric/watch/{clusterId}/{type}", Params: nil, Handler: WatchDynamic})
}
