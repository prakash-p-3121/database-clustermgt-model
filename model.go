package database_clustermgt_model

import "time"

type EmailIDLookUpResp struct {
	UserID    string `json:"user-id"`
	ClusterID string `json:"cluster-id"`
}

type UserIDLookUpResp struct {
	EmailID   string `json:"email-id"`
	ClusterID string `json:"cluster-id"`
}

type DatabaseShard struct {
	ID           *int64    `json:"id"` //shardID
	IPAddress    *string   `json:"ip-address"`
	ClusterID    *int64    `json:"cluster-id"`
	Port         *int      `json:"port"`
	UserName     *string   `json:"user-name"`
	Password     *string   `json:"password"`
	DatabaseName *string   `json:"string"`
	StartRange   *string   `json:"start-range"`
	EndRange     *string   `json:"end-range"`
	CreatedAt    time.Time `json:"created-at"`
	UpdatedAt    time.Time `json:"updated-at"`
}

type DatabaseCluster struct {
	ID           *int64    `json:"id"` // clusterID
	TableName    *string   `json:"table-name"`
	ShardingType uint8     `json:"sharding-type"`
	CreatedAt    time.Time `json:"created-at"`
	UpdatedAt    time.Time `json:"updated-at"`
}

const (
	ShardingTypeByNumber = uint8(1)
	ShardingTypeByChar   = uint8(2)
)
