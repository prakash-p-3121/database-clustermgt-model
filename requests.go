package database_clustermgt_model

import (
	"github.com/prakash-p-3121/errorlib"
	"strings"
)

type DatabaseClusterCreateReq struct {
	TableName   *string `json:"table-name"`
	ShardIDList []int64 `json:"shard-id-list"`
}

func (req *DatabaseClusterCreateReq) Validate() errorlib.AppError {
	if req.TableName == nil {
		return errorlib.NewBadReqError("table_name-nil")
	}
	*req.TableName = strings.TrimSpace(*req.TableName)
	if len(*req.TableName) == 0 {
		return errorlib.NewBadReqError("table-name-empty")
	}
	return nil
}

type DatabaseShardCreateReq struct {
	IPAddress    *string `json:"ip-address"`
	ClusterID    *int64  `json:"cluster-id"`
	Port         *int    `json:"port"`
	UserName     *string `json:"user-name"`
	Password     *string `json:"password"`
	DatabaseName *string `json:"string"`
}
