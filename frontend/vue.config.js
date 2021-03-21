//vue.config.js
module.exports = {
    //options...
    devServer: {
      proxy: {
        '^/api':{
          target: 'http://142.93.169.96:3080',
          changeOrigin: true
        },
      }
    },
    publicPath:"./",
}
