module main

go 1.12

require github.com/hashicorp/terraform v0.12.9

require (
	github.com/masterzen/winrm v0.0.0-20190308153735-1d17eaf15943
	github.com/trobert2/winrm v0.0.0-20140227124146-bafb4376f449
)

// Remove this avec terraform v0.12.8 (https://github.com/hashicorp/terraform/issues/22664)
//replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
