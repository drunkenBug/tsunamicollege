//vue.config.js
module.exports = {
    //options...
    devServer: {
      proxy: {
        '^/api':{
          // //for production:
          target: 'http://134.209.242.22:3080',

          // for local testing:
          //target:'http://localhost:3080',
          
          changeOrigin: true
        },
      }
    },
    publicPath:"./",
}
