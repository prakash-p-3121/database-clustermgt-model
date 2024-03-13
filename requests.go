package database_clustermgt_model

type DatabaseClusterCreateReq struct {
	TableName   *string `json:"table-name"`
	ShardIDList []int64 `json:"shard-id-list"`
}

type DatabaseShardCreateReq struct {
	IPAddress    *string `json:"ip-address"`
	ClusterID    *int64  `json:"cluster-id"`
	Port         *int    `json:"port"`
	UserName     *string `json:"user-name"`
	Password     *string `json:"password"`
	DatabaseName *string `json:"string"`
}
