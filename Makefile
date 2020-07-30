timestamp := $(shell date +%s)

test:
	go build
	#alfred_workflow_bundleid=hit.zhangjie.app alfred_workflow_cache=tmp alfred_workflow_data=tmp ./alfred-datetime-workflow ${timestamp}
	alfred_workflow_bundleid=hit.zhangjie.app alfred_workflow_cache=tmp alfred_workflow_data=tmp ./alfred-datetime-workflow "2020-07-31 02:26:08 CST"
	rm -rf tmp
