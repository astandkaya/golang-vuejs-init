const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  configureWebpack: {
    watchOptions :{
      ignored: 'node_modules',
      aggregateTimeout: 300,
      poll: 1000
    }
  }
})
