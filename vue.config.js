//vue.config.js
module.exports = {
    //options...
    devServer: {
      proxy: {
        '^/api':{
          target: 'http://91.64.175.168:3080',
          changeOrigin: true
        },
      }
    }
}
