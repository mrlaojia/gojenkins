package gojenkins

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
* @Author: jack.walker
* @File: views2_test.go
* @CreateDate: 2023/7/1 15:05
* @ChangeDate：2023/7/1 15:05
* @Version：1.0.0
* @Description: 使用的 jenkins 版本为：2.412
 */

func TestJenkins_DeleteView(t *testing.T) {
	ctx := context.Background()
	jenkins, err := getTestJenkins()
	assert.Nil(t, err)

	viewName := "wwwwwww"
	ok, err := jenkins.DeleteView(ctx, viewName)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
}
