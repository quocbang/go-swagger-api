package role

// import (
// 	"quocbang/go-swagger-api/server/swagger/models"
// 	mcomRoles "quocbang/go-swagger-api/server/swagger/models"

// 	"gitlab.kenda.com.tw/kenda/mui/server/protobuf/kenda"
// )

// // funcRoleList define function roles relationship.
// type funcRoleList map[kenda.FunctionOperationID]map[mcomRoles.Role]struct{}

// // permissionList which stores function permission roles data
// var permissionList funcRoleList

// // HasPermission will check if those user's roles have the operation id's permission.
// // it will return true if user has permission, otherwise false.
// func HasPermission(id kenda.FunctionOperationID, roles []models.Role) bool {
// 	if rpm, ok := permissionList[id]; ok {
// 		for _, userRole := range roles {
// 			if _, ok := rpm[mcomRoles.Role(userRole)]; ok {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func init() {
// 	permissionList = make(funcRoleList, len(kenda.FunctionOperationID_name))

// 	permissionList[kenda.FunctionOperationID_GET_SERVER_STATUS] = map[mcomRoles.Role]struct{}{
// 		mcomRoles.Role_ADMINISTRATOR: {},
// 	}

// 	// 材料更改 UI permission
// 	for _, p := range []kenda.FunctionOperationID{
// 		kenda.FunctionOperationID_GET_BARCODE_INFO,
// 		kenda.FunctionOperationID_UPDATE_BARCODE,
// 		kenda.FunctionOperationID_GET_UPDATE_BARCODE_STATUS_LIST,
// 		kenda.FunctionOperationID_GET_EXTEND_DAYS,
// 		kenda.FunctionOperationID_GET_CONTROL_AREA_LIST,
// 		kenda.FunctionOperationID_GET_HOLD_REASON_LIST,
// 	} {
// 		permissionList[p] = map[mcomRoles.Role]struct{}{
// 			mcomRoles.Role_ADMINISTRATOR: {},
// 			mcomRoles.Role_INSPECTOR:     {},
// 		}
// 	}

// 	// 帳號管理 UI permission
// 	for _, p := range []kenda.FunctionOperationID{
// 		kenda.FunctionOperationID_GET_ROLE_LIST,
// 		kenda.FunctionOperationID_LIST_AUTHORIZED_ACCOUNT,
// 		kenda.FunctionOperationID_LIST_UNAUTHORIZED_ACCOUNT,
// 		kenda.FunctionOperationID_CREATE_ACCOUNT,
// 		kenda.FunctionOperationID_UPDATE_ACCOUNT,
// 		kenda.FunctionOperationID_DELETE_ACCOUNT,
// 	} {
// 		permissionList[p] = map[mcomRoles.Role]struct{}{
// 			mcomRoles.Role_ADMINISTRATOR: {},
// 		}
// 	}

// 	{ // register default permission
// 		allRoles := make(map[mcomRoles.Role]struct{}, len(mcomRoles.Role_name))
// 		for r := range mcomRoles.Role_name {
// 			allRoles[mcomRoles.Role(r)] = struct{}{}
// 		}

// 		for p := range kenda.FunctionOperationID_name {
// 			perm := kenda.FunctionOperationID(p)

// 			if _, ok := permissionList[perm]; ok {
// 				continue
// 			}

// 			permissionList[perm] = allRoles
// 		}
// 	}
// }
