timestamp := $(shell date +%s)

test:
	go build
	alfred_workflow_bundleid=hit.zhangjie.app alfred_workflow_cache=tmp alfred_workflow_data=tmp ./alfred-datetime-workflow ${timestamp}
	rm -rf tmp
