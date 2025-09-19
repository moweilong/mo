package core

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TypeConverters 定义时间类型转换器，用于 copier 的深度拷贝。
func TypeConverters() []copier.TypeConverter {
	return []copier.TypeConverter{
		{
			SrcType: time.Time{},
			DstType: &timestamppb.Timestamp{},
			Fn: func(src interface{}) (interface{}, error) {
				s, ok := src.(time.Time)
				if !ok {
					return nil, errors.New("source type not matching")
				}
				return timestamppb.New(s), nil
			},
		},
		{
			SrcType: &timestamppb.Timestamp{},
			DstType: time.Time{},
			Fn: func(src interface{}) (interface{}, error) {
				s, ok := src.(*timestamppb.Timestamp)
				if !ok {
					return nil, errors.New("source type not matching")
				}
				return s.AsTime(), nil
			},
		},
	}
}

// CopyWithConverters 深度拷贝 from 到 to，使用自定义的类型转换器.
// 自定义类型转换器用于处理 time.Time 到 timestamppb.Timestamp 的相互转换.
func CopyWithConverters(to any, from any) error {
	return copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true, Converters: TypeConverters()})
}

// Copy 深度拷贝 from 到 to，不使用自定义的类型转换器.
func Copy(to any, from any) error {
	return copier.Copy(to, from)
}
