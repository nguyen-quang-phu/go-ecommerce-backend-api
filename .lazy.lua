return {
	{
		"neovim/nvim-lspconfig",
		---@class PluginLspOpts
		opts = {
			servers = {
				sqls = {
					cmd = {
						"sqls",
						"-config",
						"/home/keynold/go/src/github.com/nguyen-quang-phu/go-ecommerce-backend-api/sqls.yaml",
					},
					settings = {
						sqls = {
							connections = {
								{
									driver = "mysql",
									dataSourceName = "root:root1234@tcp(127.0.0.1:3306)/shopdevgo",
								},
							},
						},
					},
				},
			},
		},
	},
}
