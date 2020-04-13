const webpack = require('webpack');

//sourceMap: 映射关系 最佳实践
//{mode: 'production',devtool: 'cheap-module-source-map'}
//{mode: 'development',devtool: 'cheap-module-eval-source-map'}

const merge = require('webpack-merge');
const commonConfig = require('./webpack.common.js');
const prodConfig = {
    mode: 'production',
    // devtool: 'none',//关闭开发模式 sourceMap
    devtool: 'cheap-module-source-map',
}
module.exports = merge(commonConfig, prodConfig);