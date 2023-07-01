package gojenkins

import (
	"context"
	"errors"
	"strconv"
)

/**
* @Author: jack.walker
* @File: views2.go
* @CreateDate: 2023/7/1 15:04
* @ChangeDate：2023/7/1 15:04
* @Version：1.0.0
* @Description: 扩充 view 的功能
 */

// Delete 补充删除 view
func (v *View) Delete(ctx context.Context) (bool, error) {
	resp, err := v.Jenkins.Requester.Post(ctx, v.Base+"/doDelete", nil, nil, nil)
	if err != nil {
		return false, err
	}
	if resp.StatusCode != 200 {
		return false, errors.New(strconv.Itoa(resp.StatusCode))
	}
	return true, nil
}
