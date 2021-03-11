//vue.config.js
module.exports = {
    //options...
    devServer: {
      proxy: {
        '^/api':{
          target: 'http://165.227.78.216:3080',
          changeOrigin: true
        },
      }
    },
    publicPath:"./",
}
