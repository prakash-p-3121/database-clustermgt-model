package database_clustermgt_model

type DatabaseClusterCreateReq struct {
	ID        *int64  `json:"id"` // clusterID
	TableName *string `json:"table-name"`
	ShardSize *int64  `json:"shard-size"`
}

type DatabaseShardCreateReq struct {
	ID           *int64  `json:"id"` //shardID
	IPAddress    *string `json:"ip-address"`
	ClusterID    *int64  `json:"cluster-id"`
	Port         *int    `json:"port"`
	UserName     *string `json:"user-name"`
	Password     *string `json:"password"`
	DatabaseName *string `json:"string"`
}
