
test:
	go build
	alfred_workflow_bundleid=hit.zhangjie.app alfred_workflow_cache=1 alfred_workflow_data=1 ./alfred-datetime-workflow $(date +%s)
