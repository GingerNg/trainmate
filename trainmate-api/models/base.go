package models

// func FindAll(tblName string) []interface{} {
// 	var results []interface{}
// 	var filterMap = make(map[string]interface{})
// 	cursor := dao.NewMgo(tblName).FindMany(filterMap)
// 	for cursor.Next(context.Background()) {
// 		ud := &interface{}
// 		err := cursor.Decode(ud)
// 		results = append(results, *ud)
// 		if err != nil {
// 			fmt.Println("decode error is ", err)
// 			continue
// 		}
// 	}
// 	return results
// }
