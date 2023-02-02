package impl

import (
	"time"
)

// func (session *session) getEmployeesDepartments(userID string, ad bool) ([]mcom.Department, error) {
// 	var user models.User
// 	var condition map[string]interface{}
// 	if ad {
// 		condition = map[string]interface{}{"account": userID}
// 	} else {
// 		condition = map[string]interface{}{"id": userID}
// 	}
// 	if err := session.db.Where(condition).Take(&user).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, mcomErr.Error{
// 				Code: mcomErr.Code_ACCOUNT_NOT_FOUND_OR_BAD_PASSWORD,
// 			}
// 		}
// 		return nil, fmt.Errorf("failed to get the department of the user: %v", err)
// 	}

// 	if user.Resigned() {
// 		return nil, mcomErr.Error{
// 			Code: mcomErr.Code_ACCOUNT_NOT_FOUND_OR_BAD_PASSWORD,
// 		}
// 	}

// 	var dep models.Department
// 	if err := session.db.Where("id = ?", user.DepartmentID).
// 		Take(&dep).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, mcomErr.Error{Code: mcomErr.Code_DEPARTMENT_NOT_FOUND}
// 		}
// 		return nil, err
// 	}

// 	var deps []models.Department
// 	if err := session.db.Where("id LIKE ?", strings.TrimRight(dep.ID, "0")+"%").
// 		Order("id").
// 		Find(&deps).Error; err != nil {
// 		return nil, err
// 	}

// 	departments := make([]mcom.Department, len(deps))
// 	for i, dep := range deps {
// 		departments[i] = mcom.Department{
// 			OID: dep.ID,
// 			ID:  dep.ID,
// 		}
// 	}
// 	return departments, nil
// }

func (session *session) createToken(user string, expiryTime time.Time) (string, error) {
	// token, uError := uuid.NewV4().String()
	// // insert token.
	// if err := session.db.Create(&models.Token{
	// 	ID:          models.EncryptedData(token),
	// 	BoundUser:   user,
	// 	ExpiryTime:  types.ToTimeNano(expiryTime),
	// 	CreatedTime: types.ToTimeNano(time.Now()),
	// 	Valid:       true,
	// 	// Info:        info,
	// }).Error; err != nil {
	// 	return "", err
	// }

	return "", nil
}
