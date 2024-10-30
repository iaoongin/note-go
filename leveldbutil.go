package main

import (
	"encoding/json"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

// LevelDBWrapper 封装 LevelDB 操作
type LevelDBWrapper struct {
	db *leveldb.DB
}

type TextData struct {
	Text string
}

// NewLevelDBWrapper 创建一个新的 LevelDBWrapper
func NewLevelDBWrapper(path string) (*LevelDBWrapper, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDBWrapper{db: db}, nil
}

// Close 关闭数据库
func (l *LevelDBWrapper) Close() {
	if err := l.db.Close(); err != nil {
		log.Fatal(err)
	}
}

// Put 插入或更新键值对
func (l *LevelDBWrapper) Put(key string, value interface{}) error {
	data, err := json.Marshal(value) // 使用 JSON 序列化
	if err != nil {
		return err
	}
	return l.db.Put([]byte(key), data, nil)
}

// Get 获取指定键的值
func (l *LevelDBWrapper) Get(key string) (TextData, error) {
	data, err := l.db.Get([]byte(key), nil)
	if err != nil {
		return TextData{}, err
	}

	var value TextData                                   // 使用具体类型
	if err := json.Unmarshal(data, &value); err != nil { // 使用 JSON 反序列化
		return value, err
	}
	return value, nil
}

// Delete 删除指定键
func (l *LevelDBWrapper) Delete(key string) error {
	return l.db.Delete([]byte(key), nil)
}
