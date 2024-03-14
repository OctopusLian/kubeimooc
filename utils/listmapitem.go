package utils

import "kubeimooc/model/base"

//

func ToMap(items []base.ListMapItem) map[string]string {
	dataMap := make(map[string]string)
	for _, item := range items {
		dataMap[item.Key] = item.Value
	}
	return dataMap
}
func ToList(data map[string]string) []base.ListMapItem {
	list := make([]base.ListMapItem, 0)
	for k, v := range data {
		list = append(list, base.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return list
}
func ToListWithMapByte(data map[string][]byte) []base.ListMapItem {
	list := make([]base.ListMapItem, 0)
	for k, v := range data {
		list = append(list, base.ListMapItem{
			Key:   k,
			Value: string(v),
		})
	}
	return list
}
