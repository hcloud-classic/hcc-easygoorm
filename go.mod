module hcc/easygoorm

go 1.15

replace (
	innogrid.com/hcloud-classic/hcc_errors => ../hcc_errors
	innogrid.com/hcloud-classic/model => ../model
	innogrid.com/hcloud-classic/pb => ../pb
)

require innogrid.com/hcloud-classic/model v0.0.0
