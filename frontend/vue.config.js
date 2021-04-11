//vue.config.js
module.exports = {
    //options...
    devServer: {
      proxy: {
        '^/api':{
          // //for production:
          //target: 'http://142.93.169.96:3080',

          // for local testing:
          target:'http://localhost:3080',
          
          changeOrigin: true
        },
      }
    },
    publicPath:"./",
}
