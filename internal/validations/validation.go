package validations

import (
	"github.com/lkphuong/room-management/configs/hardcode"
	"github.com/lkphuong/room-management/internal/utils"
)

func ValidateUserInStore(store string, user utils.JwtPayload) *string {
	// #region check if user has permission to access this store
	storeIDs := user.StoreIDs
	flag := false
	for _, storeID := range storeIDs {
		if storeID == store {
			flag = true
			break
		}
	}
	if !flag && user.Code != hardcode.OPERATOR_ACCOUNT {
		errMsg := "You don't have permission to access this store"
		return &errMsg
	}

	return nil
	// #endregion
}
