package gojenkins

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

/**
* @Author: jack.walker
* @File: jenkins2.go
* @CreateDate: 2023/6/29 16:45
* @ChangeDate: 2023/6/29 16:45
* @Version: 1.0.0
* @Description: 对原库的扩充
 */

// CreateViewInFolder 在某个文件夹下创建 view
// parents 可以指定嵌套深度，包括顶层，可以替代 j.CreateView
func (j *Jenkins) CreateViewInFolder(ctx context.Context, viewName string, viewType string, parents ...string) (*View, error) {
	return j.CreateViewWithDescInFolder(ctx, viewName, "", viewType, parents...)
}

// CreateViewWithDescInFolder 在某个文件夹下创建 view, 并且附带描述信息
// parents 可以指定嵌套深度，包括顶层，可以替代 j.CreateView
func (j *Jenkins) CreateViewWithDescInFolder(ctx context.Context, viewName, desc string, viewType string, parents ...string) (*View, error) {

	baseur := j.depJobUrl(parents...)
	view := &View{Jenkins: j, Raw: new(ViewResponse), Base: baseur + "/view/" + viewName}

	endpoint := baseur + "/createView"
	data := map[string]string{
		"name":   viewName,
		"mode":   viewType,
		"Submit": "OK",
		"json": makeJson(map[string]string{
			"name":        viewName,
			"mode":        viewType,
			"description": desc,
		}),
	}
	r, err := j.Requester.Post(ctx, endpoint, nil, view.Raw, data)

	if err != nil {
		return nil, err
	}

	if r.StatusCode == 200 {
		return j.GetSubView(ctx, viewName, parents...)
	}
	return nil, errors.New(strconv.Itoa(r.StatusCode))
}

// GetSubView 查看某个文件夹下的某个view
// parents 可以指定嵌套深度，包括顶层，替代 j.GetView
// 注意：GetView 没有对状态码判断，当view不存在时，也会判断存在
func (j *Jenkins) GetSubView(ctx context.Context, name string, parents ...string) (*View, error) {
	url := j.depJobUrl(parents...) + "/view/" + name
	view := View{Jenkins: j, Raw: new(ViewResponse), Base: url}
	status, err := view.Poll(ctx)
	if err != nil {
		return nil, err
	}
	
	if status == 200 {
		return &view, nil
	}
	return nil, errors.New(strconv.Itoa(status))
}

// GetAllSubViews 查看某个文件夹下的所有view
// parents 可以指定嵌套深度，包括顶层，可以替代 j.GetAllViews
func (j *Jenkins) GetAllSubViews(ctx context.Context, parents ...string) ([]*View, error) {
	_, err := j.Poll(ctx)
	if err != nil {
		return nil, err
	}

	jenkinsRaw := new(ExecutorResponse)
	rsp, err := j.Requester.GetJSON(ctx, j.depJobUrl(parents...), jenkinsRaw, nil)
	if j.Raw == nil || rsp.StatusCode != http.StatusOK {
		return nil, errors.New("Connection Failed, Please verify that the host and credentials are correct.")
	}

	views := make([]*View, len(jenkinsRaw.Views))
	for i, v := range jenkinsRaw.Views {
		views[i], _ = j.GetSubView(ctx, v.Name, parents...)
	}
	return views, nil
}

// DeleteView 删除 view
// parents 可以指定嵌套深度，包括顶层
func (j *Jenkins) DeleteView(ctx context.Context, viewName string, parents ...string) (bool, error) {
	view := View{
		Raw:     new(ViewResponse),
		Jenkins: j,
		Base:    j.depJobUrl(parents...) + "/view/" + viewName,
	}
	return view.Delete(ctx)
}

// depUrl 组合成 /job/x/job/x 格式的url
func (j *Jenkins) depJobUrl(parents ...string) string {
	return j.depUrl("job", parents...)
}

// depUrl 组合成 url
// tag: job, view ...
// parents: 层级关系
func (j *Jenkins) depUrl(tag string, parents ...string) string {
	var base string
	for _, p := range parents {
		base += fmt.Sprintf("/%s/%s", tag, p)
	}

	return base
}
