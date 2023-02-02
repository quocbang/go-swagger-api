package role

// func TestHasPermission(t *testing.T) {
// 	type args struct {
// 		id    kenda.FunctionOperationID
// 		roles []models.Role
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{
// 			name: "permitted",
// 			args: args{
// 				id: kenda.FunctionOperationID_GET_SERVER_STATUS,
// 				roles: []models.Role{
// 					models.Role(mcomRoles.Role_ADMINISTRATOR),
// 				},
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "not permitted",
// 			args: args{
// 				id: kenda.FunctionOperationID_GET_CONTROL_AREA_LIST,
// 				roles: []models.Role{
// 					models.Role(mcomRoles.Role_BEARER),
// 					models.Role(mcomRoles.Role_PLANNER),
// 				},
// 			},
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := HasPermission(tt.args.id, tt.args.roles)
// 			if got != tt.want {
// 				t.Errorf("HasPermission() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
