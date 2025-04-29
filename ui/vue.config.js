const { defineConfig } = require('@vue/cli-service')
const packageJson = require('./package.json')

module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true
        // No pathRewrite here to keep the /api prefix
      }
    }
  },
  chainWebpack: config => {
    config.plugin('define').tap(args => {
      args[0]['process.env'].VUE_APP_VERSION = JSON.stringify(packageJson.version);
      return args;
    });
  }
})