package model

type ExecTime struct {
	StartTime int64 `bson:"start_time"`
	EndTime   int64 `bson:"end_time"`
}

type LogRecord struct {
	JobName string `bson:"job_name"`
	Command string `bson:"command"`
	Err     string `bson:"err"`
	Content string `bson:"content"`
	Tp      ExecTime
}

//查询实体
type FindByJobName struct {
	JobName string `bson:"job_name"`
}
