[vimflowy](https://github.com/WuTheFWasThat/vimflowy) is delightful

[vimflowy deployment docs](https://github.com/WuTheFWasThat/vimflowy/blob/master/docs/deployment.md)


### current entry-level workflow

*	ensure docker daemon running
*	launch with `make launch`
*	open http://localhost:3000 in a browser
*	can import JSON vimflowy file using the vimflowy UI
*	when finished, export JSON vimflowy file and save
*	stop with `make stop`

### TODO: better workflow

*	make it easier to version control content as text files in git without data loss
	*	run in LTE script to load JSON into fresh sqlite DB, run vimflowy, then export sqlite DB to JSON.
*	get more control over external dependencies
	*	pin container version
	*	consider vendoring upstream vimflowy repo
*	make it faster, avoid needlessly pulling of image registry

