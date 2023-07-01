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

func TestJenkins_CreateViewWithDescInFolder(t *testing.T) {
	ctx := context.Background()
	jenkins, err := getTestJenkins()
	assert.Nil(t, err)

	viewName := "dev_pro3"
	desc := "this is dev_pro3  view"
	view1, err := jenkins.CreateViewWithDescInFolder(ctx, viewName, desc, LIST_VIEW, "qa")
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
	assert.Equal(t, viewName, v.GetName())
}

func TestJenkins_GetAllSubViews(t *testing.T) {
	ctx := context.Background()
	jenkins, err := getTestJenkins()
	assert.Nil(t, err)

	views, err := jenkins.GetAllSubViews(ctx, "dev", "dev01")
	assert.Nil(t, err)

	assert.Equal(t, 2, len(views))
	assert.Equal(t, 1, len(views[0].Raw.Jobs))
}

func TestView_Delete(t *testing.T) {
	ctx := context.Background()
	jenkins, err := getTestJenkins()
	assert.Nil(t, err)

	viewName := "dev_pro2"
	v, err := jenkins.GetSubView(ctx, viewName, "dev", "dev01")
	assert.Nil(t, err)
	assert.Equal(t, viewName, v.GetName())

	ok, err := v.Delete(ctx)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
}
