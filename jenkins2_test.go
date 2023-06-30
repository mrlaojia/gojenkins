package gojenkins

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
* @Author: jack.walker
* @File: jenkins2_test.go
* @CreateDate: 2023/6/29 16:55
* @ChangeDate: 2023/6/29 16:55
* @Version: 1.0.0
* @Description: 测试 jenkins.go ； 使用的 jenkins 版本为：2.406
 */

func getTestJenkins() (*Jenkins, error) {
	ctx := context.Background()
	jenkins = CreateJenkins(nil, "http://jenkins.ingress.local/", "admin", "hello123")
	return jenkins.Init(ctx)
}

func TestJenkins_CreateViewInFolder(t *testing.T) {
	ctx := context.Background()
	jenkins, err := getTestJenkins()
	assert.Nil(t, err)

	viewName := "dev_pro3"
	view1, err := jenkins.CreateViewInFolder(ctx, viewName, LIST_VIEW, "dev", "dev01")
	assert.Nil(t, err)
	assert.Equal(t, viewName, view1.GetName())
}

func TestJenkins_GetSubView(t *testing.T) {
	ctx := context.Background()
	jenkins, err := getTestJenkins()
	assert.Nil(t, err)

	viewName := "dev_pro2"
	v, err := jenkins.GetSubView(ctx, viewName, "dev", "dev01")
	assert.Nil(t, err)
	assert.Equal(t, v.GetName(), viewName)
}

func TestJenkins_GetAllSubViews(t *testing.T) {
	ctx := context.Background()
	jenkins, err := getTestJenkins()
	assert.Nil(t, err)

	vs, err := jenkins.GetAllSubViews(ctx, "dev", "dev01")
	assert.Nil(t, err)
	for _, r := range vs {
		t.Log(r.GetUrl(), r.GetName())
	}
}

func TestView_Delete(t *testing.T) {
	ctx := context.Background()
	jenkins, err := getTestJenkins()
	assert.Nil(t, err)

	viewName := "dev_pro2"
	v, err := jenkins.GetSubView(ctx, viewName, "dev", "dev01")
	assert.Nil(t, err)
	assert.Equal(t, v.GetName(), viewName)

	ok, err := v.Delete(ctx)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
}

func TestJenkins_DeleteView(t *testing.T) {
	ctx := context.Background()
	jenkins, err := getTestJenkins()
	assert.Nil(t, err)

	viewName := "wwwwwww"
	ok, err := jenkins.DeleteView(ctx, viewName)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
}
