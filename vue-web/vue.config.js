
module.exports = {
    devServer: {
      proxy: {
        '/api': {
          target: process.env.VUE_APP_BASE_API,
          changeOrigin: true, // 允许跨域
          pathRewrite: { // 重写路径
            '^/api': ''
          }
        }
      }
    }
  }
  