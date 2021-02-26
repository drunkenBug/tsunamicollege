//vue.config.js
module.exports = {
    //options...
    devServer: {
      proxy: {
        '^/api':{
          target: 'http://localhost:3080',
          changeOrigin: true
        },
      }
    }
}
