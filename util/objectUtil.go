package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type ChangeError struct {
	ErrMsg string
}

func (error *ChangeError) Error() string {
	return error.ErrMsg
}

func ToMap(data interface{}) map[string]interface{} {
	if reflect.TypeOf(data).Kind() == reflect.Map {
		resultMap := map[string]interface{}{}
		dataValue := reflect.ValueOf(data)
		for mapR := dataValue.MapRange(); mapR.Next(); {
			mapKey := mapR.Key()
			mapValue := mapR.Value()

			resultMap[mapKey.String()] = mapValue
		}
		return resultMap
	} else if reflect.TypeOf(data).Kind() == reflect.Struct {
		resultMap := map[string]interface{}{}
		jsonStr, err := json.Marshal(data)
		if err != nil {
			return resultMap
		}

		err = yaml.Unmarshal(jsonStr, &resultMap)
		if err != nil {
			return resultMap
		}
		return resultMap
	}
	return nil
}

func IsNumber(fieldKing reflect.Kind) bool {
	switch fieldKing {
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	default:
		return false
	}
}

// IsBaseType 是否是常见基本类型
func IsBaseType(fieldType reflect.Type) bool {
	fieldKind := fieldType.Kind()
	if fieldKind == reflect.Ptr {
		fieldKind = fieldType.Elem().Kind()
	}

	switch fieldKind {
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	case reflect.Bool:
		return true
	case reflect.String:
		return true
	default:
		if fieldType.String() == "time.Time" {
			return true
		}
		return false
	}
}

func ToJsonString(value interface{}) string {
	if value == nil {
		return ""
	}
	bytes, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	return string(bytes)
}

func ToString(value interface{}) string {
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%v", value)
}

func ToInt(value interface{}) int {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Int)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(int)
}

func ToInt8(value interface{}) int8 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Int8)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(int8)
}

func ToInt16(value interface{}) int16 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Int16)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(int16)
}

func ToInt32(value interface{}) int32 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Int32)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(int32)
}

func ToInt64(value interface{}) int64 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Int64)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(int64)
}

func ToUInt(value interface{}) uint {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Uint)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(uint)
}

func ToUInt8(value interface{}) uint8 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Uint8)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(uint8)
}

func ToUInt16(value interface{}) uint16 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Uint16)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(uint16)
}

func ToUInt32(value interface{}) uint32 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Uint32)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(uint32)
}

func ToUInt64(value interface{}) uint64 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Uint64)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(uint64)
}

func ToFloat32(value interface{}) float32 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Float32)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(float32)
}

func ToFloat64(value interface{}) float64 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Float64)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(float64)
}

func ToBool(value interface{}) bool {
	if value == nil {
		return false
	}
	result, err := ToValue(value, reflect.Bool)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return false
	}
	return result.(bool)
}

func ToComplex64(value interface{}) complex64 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Complex64)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(complex64)
}

func ToComplex128(value interface{}) complex128 {
	if value == nil {
		return 0
	}
	result, err := ToValue(value, reflect.Complex128)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return 0
	}
	return result.(complex128)
}

func ToValue(value interface{}, valueKind reflect.Kind) (interface{}, error) {
	if value == nil {
		return nil, nil
	}
	valueStr := ToString(value)
	return Cast(valueKind, valueStr)
}

func Cast(fieldKind reflect.Kind, valueStr string) (interface{}, error) {
	if valueStr == "nil" || valueStr == "" {
		return nil, nil
	}
	switch fieldKind {
	case reflect.Int:
		return strconv.Atoi(valueStr)
	case reflect.Int8:
		v, err := strconv.ParseInt(valueStr, 10, 8)
		if err != nil {
			return nil, err
		}
		return int8(v), nil
	case reflect.Int16:
		v, err := strconv.ParseInt(valueStr, 10, 16)
		if err != nil {
			return nil, err
		}
		return int16(v), nil
	case reflect.Int32:
		v, err := strconv.ParseInt(valueStr, 10, 32)
		if err != nil {
			return nil, err
		}
		return int32(v), nil
	case reflect.Int64:
		return strconv.ParseInt(valueStr, 10, 64)
	case reflect.Uint:
		v, err := strconv.ParseUint(valueStr, 10, 0)
		if err != nil {
			return nil, err
		}
		return uint(v), nil
	case reflect.Uint8:
		v, err := strconv.ParseUint(valueStr, 10, 8)
		if err != nil {
			return nil, err
		}
		return uint8(v), nil
	case reflect.Uint16:
		v, err := strconv.ParseUint(valueStr, 10, 16)
		if err != nil {
			return nil, err
		}
		return uint16(v), nil
	case reflect.Uint32:
		v, err := strconv.ParseUint(valueStr, 10, 32)
		if err != nil {
			return nil, err
		}
		return uint32(v), nil
	case reflect.Uint64:
		return strconv.ParseUint(valueStr, 10, 64)
	case reflect.Float32:
		v, err := strconv.ParseFloat(valueStr, 32)
		if err != nil {
			return nil, err
		}
		return float32(v), nil
	case reflect.Float64:
		return strconv.ParseFloat(valueStr, 64)
	case reflect.Complex64:
		v, err := strconv.ParseComplex(valueStr, 64)
		if err != nil {
			return nil, err
		}
		return complex64(v), nil
	case reflect.Complex128:
		return strconv.ParseComplex(valueStr, 128)
	case reflect.Bool:
		return strconv.ParseBool(valueStr)
	}
	return valueStr, nil
}

// DataToObject 其他的类型能够按照小写字母转换到对象
// 其他类型：
//  - 基本类型
//  - 结构体类型：转换后对象
//  - map类型
//  - 集合/分片类型
//  - 字符串类型：如果是json，则按照json进行转换
func DataToObject(data interface{}, targetPtrObj interface{}) error {
	if data == nil {
		fmt.Printf("data is nil")
		return nil
	}
	targetType := reflect.TypeOf(targetPtrObj)
	if targetType.Kind() != reflect.Ptr {
		fmt.Printf("targetPtrObj type is not ptr")
		return &ChangeError{ErrMsg: "targetPtrObj type is not ptr"}
	}

	switch data.(type) {
	case io.Reader:
		return ReaderToObject(data.(io.Reader), targetPtrObj)
	case string:
		return StrToObject(data.(string), targetPtrObj)
	case map[interface{}]interface{}:
		return MapToObject(data.(map[interface{}]interface{}), targetPtrObj)
	case []interface{}:
		return ArrayToObject(data.([]interface{}), targetPtrObj)
	case interface{}:
		return MapToObject(ToMap(data), targetPtrObj)
	}

	targetPtrValue := reflect.ValueOf(targetPtrObj)
	rel, err := Cast(targetPtrValue.Elem().Kind(), fmt.Sprintf("%v", data))
	if err != nil {
		return err
	}
	targetPtrValue.Elem().Set(reflect.ValueOf(rel))
	return nil
}

func ReaderToObject(reader io.Reader, targetPtrObj interface{}) error {
	if reader == nil {
		return nil
	}
	targetType := reflect.TypeOf(targetPtrObj)
	if targetType.Kind() != reflect.Ptr {
		fmt.Printf("targetPtrObj type is not ptr")
		return &ChangeError{ErrMsg: "targetPtrObj type is not ptr"}
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return StrToObject(string(data), targetPtrObj)
}

func StrToObject(contentOfJson string, targetPtrObj interface{}) error {
	if contentOfJson == "" {
		return &ChangeError{ErrMsg: "content is nil"}
	}

	targetType := reflect.TypeOf(targetPtrObj)
	if targetType.Kind() != reflect.Ptr {
		fmt.Printf("targetPtrObj type is not ptr")
		return &ChangeError{ErrMsg: "targetPtrObj type is not ptr"}
	}

	if !strings.HasPrefix(contentOfJson, "{") && !strings.HasPrefix(contentOfJson, "[") {
		targetPtrValue := reflect.ValueOf(targetPtrObj)
		rel, err := Cast(targetPtrValue.Elem().Kind(), contentOfJson)
		if err != nil {
			return err
		}
		targetPtrValue.Elem().Set(reflect.ValueOf(rel))
	}

	if strings.HasPrefix(contentOfJson, "{") && (reflect.ValueOf(targetPtrObj).Elem().Kind() == reflect.Map || reflect.ValueOf(targetPtrObj).Elem().Kind() == reflect.Struct) {
		resultMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(contentOfJson), &resultMap)
		if err != nil {
			return err
		}
		return MapToObject(resultMap, targetPtrObj)
	} else if strings.HasPrefix(contentOfJson, "[") && (reflect.ValueOf(targetPtrObj).Elem().Kind() == reflect.Slice || reflect.ValueOf(targetPtrObj).Elem().Kind() == reflect.Array) {
		var srcArray []interface{}
		err := json.Unmarshal([]byte(contentOfJson), &srcArray)
		if err != nil {
			return err
		}
		return ArrayToObject(srcArray, targetPtrObj)
	} else {
		targetPtrValue := reflect.ValueOf(targetPtrObj)
		rel, err := Cast(targetPtrValue.Elem().Kind(), contentOfJson)
		if err != nil {
			return err
		}
		targetPtrValue.Elem().Set(reflect.ValueOf(rel))
		return nil
	}
}

func ArrayToObject(dataArray interface{}, targetPtrObj interface{}) error {
	if dataArray == nil {
		return nil
	}

	if reflect.ValueOf(dataArray).Kind() != reflect.Array && reflect.ValueOf(dataArray).Kind() != reflect.Slice {
		return &ChangeError{ErrMsg: "dataArray is array type"}
	}

	targetType := reflect.TypeOf(targetPtrObj)
	if targetType.Kind() != reflect.Ptr {
		fmt.Printf("targetPtrObj type is not ptr")
		return &ChangeError{ErrMsg: "targetPtrObj type is not ptr"}
	}

	if targetType.Elem().Kind() != reflect.Slice && targetType.Elem().Kind() != reflect.Array {
		fmt.Printf("item of targetPtrObj type is not slice")
		return &ChangeError{ErrMsg: "item of targetPtrObj type is not slice"}
	}

	srcValue := reflect.ValueOf(dataArray)
	dstPtrValue := reflect.ValueOf(targetPtrObj)

	dstPrtType := reflect.TypeOf(targetPtrObj)
	dstType := dstPrtType.Elem()
	dstItemType := dstType.Elem()

	dstValue := reflect.MakeSlice(dstType, 0, 0)

	for arrayIndex := 0; arrayIndex < srcValue.Len(); arrayIndex++ {
		dataV := valueToTarget(srcValue.Index(arrayIndex), dstItemType)
		if dataV.IsValid() {
			if dataV.Kind() == reflect.Ptr {
				dstValue = reflect.Append(dstValue, dataV.Elem())
			} else {
				dstValue = reflect.Append(dstValue, dataV)
			}
		}
	}
	dstPtrValue.Elem().Set(dstValue)
	return nil
}

func MapToObject(dataMap interface{}, targetPtrObj interface{}) error {
	if dataMap == nil {
		return nil
	}
	targetType := reflect.TypeOf(targetPtrObj)
	if targetType.Kind() != reflect.Ptr {
		fmt.Printf("targetPtrObj type is not ptr")
		return &ChangeError{ErrMsg: "targetPtrObj type is not ptr"}
	}

	if targetType.Elem().Kind() != reflect.Map && targetType.Elem().Kind() != reflect.Struct {
		fmt.Printf("item of targetPtrObj type is not Map or Struct")
		return &ChangeError{ErrMsg: "item of targetPtrObj type is not slice"}
	}

	if targetType.Elem().Kind() == reflect.Map {
		srcValue := reflect.ValueOf(dataMap)
		dstValue := reflect.ValueOf(targetPtrObj)

		dstPtrType := reflect.TypeOf(targetPtrObj)
		dstType := dstPtrType.Elem()

		mapFieldValue := reflect.MakeMap(dstType)
		for mapR := srcValue.MapRange(); mapR.Next(); {
			mapKey := mapR.Key()
			mapValue := mapR.Value()

			mapKeyRealValue, err := Cast(mapFieldValue.Type().Key().Kind(), fmt.Sprintf("%v", mapKey.Interface()))
			mapValueRealValue := valueToTarget(mapValue, mapFieldValue.Type().Elem())
			if err == nil {
				if mapValueRealValue.Kind() == reflect.Ptr {
					mapFieldValue.SetMapIndex(reflect.ValueOf(mapKeyRealValue), mapValueRealValue.Elem())
				} else {
					mapFieldValue.SetMapIndex(reflect.ValueOf(mapKeyRealValue), mapValueRealValue)
				}
			}
		}
		dstValue.Elem().Set(mapFieldValue)
	} else {
		targetValue := reflect.ValueOf(targetPtrObj)
		for index, num := 0, targetType.Elem().NumField(); index < num; index++ {
			field := targetType.Elem().Field(index)
			fieldValue := targetValue.Elem().Field(index)

			doInvokeValue(reflect.ValueOf(dataMap), field, fieldValue)
		}
	}
	return nil
}

func doInvokeValue(fieldMapValue reflect.Value, field reflect.StructField, fieldValue reflect.Value) {
	// 私有字段不处理
	if IsPrivate(field.Name) {
		return
	}

	if fieldMapValue.Kind() == reflect.Ptr {
		fieldMapValue = fieldMapValue.Elem()
	}

	if v, exist := getValueFromMapValue(fieldMapValue, field.Name); exist {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		targetValue := valueToTarget(v, field.Type)
		if targetValue.IsValid() {
			if fieldValue.Kind() == reflect.Ptr {
				if targetValue.Kind() == reflect.Ptr {
					fieldValue.Elem().FieldByName(field.Name).Set(targetValue.Elem().Convert(field.Type))
				} else {
					fieldValue.Elem().FieldByName(field.Name).Set(targetValue.Convert(field.Type))
				}
			} else {
				if targetValue.Kind() == reflect.Ptr {
					fieldValue.Set(targetValue.Elem().Convert(field.Type))
				} else {
					fieldValue.Set(targetValue.Convert(field.Type))
				}
			}
		}
	}
}

func valueToTarget(srcValue reflect.Value, dstType reflect.Type) reflect.Value {
	if dstType.Kind() == reflect.Struct {
		if srcValue.Kind() == reflect.Ptr {
			srcValue = srcValue.Elem()
		}
		sourceValue := reflect.ValueOf(srcValue.Interface())
		if sourceValue.Kind() == reflect.Map || sourceValue.Kind() == reflect.Struct {
			mapFieldValue := reflect.New(dstType)
			for index, num := 0, mapFieldValue.Type().Elem().NumField(); index < num; index++ {
				field := mapFieldValue.Type().Elem().Field(index)
				fieldValue := mapFieldValue.Elem().Field(index)

				doInvokeValue(sourceValue, field, fieldValue)
			}
			return mapFieldValue
		}
	} else if dstType.Kind() == reflect.Map {
		if srcValue.Kind() == reflect.Ptr {
			srcValue = srcValue.Elem()
		}
		sourceValue := reflect.ValueOf(srcValue.Interface())
		if sourceValue.Kind() == reflect.Map {
			mapFieldValue := reflect.MakeMap(dstType)
			for mapR := sourceValue.MapRange(); mapR.Next(); {
				mapKey := mapR.Key()
				mapValue := mapR.Value()

				mapKeyRealValue, err := Cast(mapFieldValue.Type().Key().Kind(), fmt.Sprintf("%v", mapKey.Interface()))
				mapValueRealValue := valueToTarget(mapValue, mapFieldValue.Type().Elem())
				if err == nil {
					if mapValueRealValue.Kind() == reflect.Ptr {
						mapFieldValue.SetMapIndex(reflect.ValueOf(mapKeyRealValue), mapValueRealValue.Elem())
					} else {
						mapFieldValue.SetMapIndex(reflect.ValueOf(mapKeyRealValue), mapValueRealValue)
					}
				}
			}
			return mapFieldValue
		} else if sourceValue.Kind() == reflect.Struct {
			srcType := reflect.TypeOf(sourceValue)
			srcValue := reflect.ValueOf(sourceValue)
			mapFieldValue := reflect.MakeMap(dstType)

			for index, num := 0, srcType.NumField(); index < num; index++ {
				field := srcType.Field(index)
				fieldValue := srcValue.Field(index)

				mapValueRealValue := ObjectToData(fieldValue.Interface())
				mapFieldValue.SetMapIndex(reflect.ValueOf(ToLowerFirstPrefix(field.Name)), reflect.ValueOf(mapValueRealValue))

				doInvokeValue(sourceValue, field, fieldValue)
			}
			return mapFieldValue
		}
	} else if dstType.Kind() == reflect.Slice || dstType.Kind() == reflect.Array {
		if srcValue.Kind() == reflect.Ptr {
			srcValue = srcValue.Elem()
		}
		sourceValue := reflect.ValueOf(srcValue.Interface())
		if sourceValue.Kind() == reflect.Slice || sourceValue.Kind() == reflect.Array {
			arrayFieldValue := reflect.MakeSlice(dstType, 0, 0)
			for arrayIndex := 0; arrayIndex < sourceValue.Len(); arrayIndex++ {
				dataV := valueToTarget(sourceValue.Index(arrayIndex), dstType.Elem())
				if dataV.IsValid() {
					if dataV.Kind() == reflect.Ptr {
						arrayFieldValue = reflect.Append(arrayFieldValue, dataV.Elem())
					} else {
						arrayFieldValue = reflect.Append(arrayFieldValue, dataV)
					}
				}
			}
			return arrayFieldValue
		}
	} else if IsBaseType(dstType) {
		sourceValue := reflect.ValueOf(srcValue.Interface())
		if sourceValue.IsValid() && IsBaseType(sourceValue.Type()) {
			v, err := Cast(dstType.Kind(), fmt.Sprintf("%v", srcValue.Interface()))
			if err == nil {
				return reflect.ValueOf(v)
			}
		}
	} else if dstType.Kind() == reflect.Interface {
		return reflect.ValueOf(ObjectToData(srcValue.Interface()))
	} else if dstType.Kind() == reflect.Ptr {
		return srcValue
	} else {
		v, err := Cast(dstType.Kind(), fmt.Sprintf("%v", srcValue.Interface()))
		if err == nil {
			return reflect.ValueOf(v)
		}
	}
	return reflect.ValueOf(nil)
}

// ObjectToData 字段转化，其中对应字段为小写，map的话为小写
func ObjectToData(object interface{}) interface{} {
	if object == nil || reflect.ValueOf(object).Kind() == reflect.Ptr {
		return "{}"
	}

	// 只接收 map、struct、array、slice进行解析
	objKind := reflect.ValueOf(object).Kind()
	if objKind != reflect.Map && objKind != reflect.Struct && objKind != reflect.Array && objKind != reflect.Slice {
		return object
	}

	if objKind == reflect.Map {
		// Map 结构
		resultMap := map[string]interface{}{}
		objValue := reflect.ValueOf(object)
		if objValue.Len() == 0 {
			return "{}"
		}

		for mapR := objValue.MapRange(); mapR.Next(); {
			mapKey := mapR.Key()
			mapValue := mapR.Value()

			v := doObjectChange(reflect.TypeOf(mapValue.Interface()), mapValue.Interface())
			if v != nil {
				resultMap[ToLowerFirstPrefix(ToString(mapKey.Interface()))] = v
			}
		}
		return resultMap
	} else if objKind == reflect.Struct {
		// Struct 结构
		resultMap := map[string]interface{}{}
		objValue := reflect.ValueOf(object)
		objType := objValue.Type()
		for index, num := 0, objType.NumField(); index < num; index++ {
			field := objType.Field(index)
			fieldValue := objValue.Field(index)

			// 私有字段不处理
			if IsPrivate(field.Name) {
				continue
			}
			v := doObjectChange(reflect.TypeOf(fieldValue.Interface()), fieldValue.Interface())
			if v != nil {
				resultMap[ToLowerFirstPrefix(field.Name)] = v
			}
		}
		return resultMap
	} else if objKind == reflect.Array || objKind == reflect.Slice {
		// Array 结构
		var resultSlice []interface{}
		objValue := reflect.ValueOf(object)
		for index := 0; index < objValue.Len(); index++ {
			arrayItemValue := objValue.Index(index)

			v := doObjectChange(reflect.TypeOf(object).Elem(), arrayItemValue.Interface())
			if v != nil {
				resultSlice = append(resultSlice, v)
			}
		}
		return resultSlice
	}
	return nil
}

// ObjectToJson 对象转化为json，其中map对应的key为小写
func ObjectToJson(object interface{}) string {
	if object == nil || reflect.ValueOf(object).Kind() == reflect.Ptr {
		return "{}"
	}

	// 只接收 map、struct、array、slice进行解析
	objKind := reflect.ValueOf(object).Kind()
	if objKind != reflect.Map && objKind != reflect.Struct && objKind != reflect.Array && objKind != reflect.Slice {
		fmt.Printf("not support the type %v change to json", objKind.String())
		return "{}"
	}

	if objKind == reflect.Map {
		// Map 结构
		resultMap := map[string]interface{}{}
		objValue := reflect.ValueOf(object)
		if objValue.Len() == 0 {
			return "{}"
		}

		for mapR := objValue.MapRange(); mapR.Next(); {
			mapKey := mapR.Key()
			mapValue := mapR.Value()

			v := doObjectChange(reflect.TypeOf(mapValue.Interface()), mapValue.Interface())
			if v != nil {
				resultMap[ToLowerFirstPrefix(ToString(mapKey.Interface()))] = v
			}
		}
		return ToJsonString(resultMap)
	} else if objKind == reflect.Struct {
		// Struct 结构
		resultMap := map[string]interface{}{}
		objValue := reflect.ValueOf(object)
		objType := objValue.Type()
		for index, num := 0, objType.NumField(); index < num; index++ {
			field := objType.Field(index)
			fieldValue := objValue.Field(index)

			// 私有字段不处理
			if IsPrivate(field.Name) {
				continue
			}
			v := doObjectChange(reflect.TypeOf(fieldValue.Interface()), fieldValue.Interface())
			if v != nil {
				resultMap[ToLowerFirstPrefix(field.Name)] = v
			}
		}
		return ToJsonString(resultMap)
	} else if objKind == reflect.Array || objKind == reflect.Slice {
		// Array 结构
		var resultSlice []interface{}
		objValue := reflect.ValueOf(object)
		for index := 0; index < objValue.Len(); index++ {
			arrayItemValue := objValue.Index(index)

			v := doObjectChange(reflect.TypeOf(object).Elem(), arrayItemValue.Interface())
			if v != nil {
				resultSlice = append(resultSlice, v)
			}
		}
		return ToJsonString(resultSlice)
	}
	return "{}"
}

// 转换为对应类型
//
// 符号数字类型 		-> int
// 无符号类型 		-> uint
// float类型 		-> float
// complex128类型 	-> complex128
// boole类型 		-> bool
// string类型 		-> string
// 集合/分片类型 		-> [xx]；其中xx对应的类型集合中的对象再次进行转换
// 结构体 			-> 转换为map
// map 				-> 转换为map
func doObjectChange(objType reflect.Type, object interface{}) interface{} {
	if objType == nil || object == nil {
		return nil
	}
	objKind := objType.Kind()
	if objKind == reflect.Ptr {
		return nil
	}
	if objKind == reflect.Int || objKind == reflect.Int8 || objKind == reflect.Int16 || objKind == reflect.Int32 || objKind == reflect.Int64 {
		return ToInt64(object)
	} else if objKind == reflect.Uint || objKind == reflect.Uint8 || objKind == reflect.Uint16 || objKind == reflect.Uint32 || objKind == reflect.Uint64 {
		return ToUInt64(object)
	} else if objKind == reflect.Float32 || objKind == reflect.Float64 {
		return ToFloat64(object)
	} else if objKind == reflect.Complex64 {
		return ToString(object)
	} else if objKind == reflect.Complex128 {
		return ToString(object)
	} else if objKind == reflect.Bool {
		return ToBool(object)
	} else if objKind == reflect.String {
		return ToString(object)
	} else if objKind == reflect.Array || objKind == reflect.Slice {
		var resultSlice []interface{}
		objValue := reflect.ValueOf(object)
		for index := 0; index < objValue.Len(); index++ {
			arrayItemValue := objValue.Index(index)

			v := doObjectChange(reflect.TypeOf(object).Elem(), arrayItemValue.Interface())
			if v != nil {
				resultSlice = append(resultSlice, v)
			}
		}
		return resultSlice
	} else if objKind == reflect.Struct {
		resultMap := map[string]interface{}{}
		objValue := reflect.ValueOf(object)
		objType := objValue.Type()
		for index, num := 0, objType.NumField(); index < num; index++ {
			field := objType.Field(index)
			fieldValue := objValue.Field(index)

			// 私有字段不处理
			if IsPrivate(field.Name) {
				continue
			}
			v := doObjectChange(reflect.TypeOf(fieldValue.Interface()), fieldValue.Interface())
			if v != nil {
				resultMap[ToLowerFirstPrefix(field.Name)] = v
			}
		}
		return resultMap
	} else if objKind == reflect.Map {
		resultMap := map[string]interface{}{}
		objValue := reflect.ValueOf(object)
		if objValue.Len() == 0 {
			return nil
		}

		for mapR := objValue.MapRange(); mapR.Next(); {
			mapKey := mapR.Key()
			mapValue := mapR.Value()

			v := doObjectChange(reflect.TypeOf(mapValue.Interface()), mapValue.Interface())
			if v != nil {
				resultMap[ToLowerFirstPrefix(ToString(mapKey.Interface()))] = v
			}
		}
		return resultMap
	} else if objKind == reflect.Interface {
		return ObjectToData(object)
	}
	return nil
}

func getValueFromMapValue(keyValues reflect.Value, key string) (reflect.Value, bool) {
	if keyValues.Kind() == reflect.Map {
		if v1 := keyValues.MapIndex(reflect.ValueOf(key)); v1.IsValid() {
			return v1, true
		} else if v2 := keyValues.MapIndex(reflect.ValueOf(ToLowerFirstPrefix(key))); v2.IsValid() {
			return v2, true
		}
	} else if keyValues.Kind() == reflect.Struct {
		if v1 := keyValues.FieldByName(key); v1.IsValid() {
			return v1, true
		} else if v2 := keyValues.FieldByName(ToLowerFirstPrefix(key)); v2.IsValid() {
			return v2, true
		}
	}

	return reflect.ValueOf(nil), false
}

func IsPublic(s string) bool {
	return isStartUpper(s)
}

func IsPrivate(s string) bool {
	return isStartLower(s)
}

// 判断首字母是否大写
func isStartUpper(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}

// 判断首字母是否小写
func isStartLower(s string) bool {
	return unicode.IsLower([]rune(s)[0])
}

// ToLowerFirstPrefix 首字母小写
func ToLowerFirstPrefix(dataStr string) string {
	return strings.ToLower(dataStr[:1]) + dataStr[1:]
}

// ToUpperFirstPrefix 首字母大写
func ToUpperFirstPrefix(dataStr string) string {
	return strings.ToLower(dataStr[:1]) + dataStr[1:]
}

func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64Decode(dst string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(dst)
}
