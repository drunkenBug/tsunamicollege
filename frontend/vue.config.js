//vue.config.js
module.exports = {
    //options...
    devServer: {
      proxy: {
        '^/api':{
          target: 'http://127.0.0.1:3080',
          changeOrigin: true
        },
      }
    },
    publicPath:"./",
}
